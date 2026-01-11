package services

import (
	"address-book-server-v3/internal/common/types"
	"address-book-server-v3/test"
	"fmt"
	"log"
	"testing"

	wgconfig "bitbucket.org/vayana/walt-gorm.go/config"
	wgormconfig "bitbucket.org/vayana/walt-gorm.go/config"
	sql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type ServiceTestSuit struct {
	suite.Suite
	suite test.TestingSuite
	db    *gorm.DB
}

func (suite *ServiceTestSuit) SetupSuite() {
	// Init application test container
	suite.suite = *test.NewTestingSuite(nil, nil)

	dbConfig, err := wgconfig.NewDBConfig(
		types.DB_HOSTNAME,
		types.DB_PORT,
		types.DB_USERNAME,
		types.DB_PASSWORD,
		types.DB_NAME,
		types.DB_TYPE,
	).Get()
	if err != nil {
		log.Fatalf("Could not connect to MySQL database: %v", err)
	}

	db, err := wgormconfig.ConnectToDatabase(dbConfig).Get()
	if err != nil {
		log.Fatalf("Could not connect to MySQL database: %v", err)
	}

	suite.db = db

	if err := suite.migrate(); err != nil {
		log.Fatalf("Could not migrate database: %v", err)
	}

	// Seed minimal required data
	// suite.db.Create(&models.ErrorMessage{
	// 	Code: "test",
	// 	One:  "test",
	// })

}

func (suite *ServiceTestSuit) SetupTest() {
	suite.db.Exec("DELETE FROM users")
	suite.db.Exec("DELETE FROM addresses")
	// suite.db.Exec("DELETE FROM error_messages WHERE code='test' AND one='test'")
}

func (suite *ServiceTestSuit) TearDownSuite() {
	sqlDB, err := suite.db.DB()
	if err != nil {
		log.Fatalf("Could not get DB: %v", err)
	}
	sqlDB.Close()
}

func (suite *ServiceTestSuit) migrate() error {
	sqlDB, err := suite.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	driver, err := sql.WithInstance(sqlDB, &sql.Config{})
	if err != nil {
		return fmt.Errorf("failed to create mysql driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://../../database/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to init migrate: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration failed: %w", err)
	}

	return nil
}

func TestServiceSuit(t *testing.T) {
	suite.Run(t, new(ServiceTestSuit))
}
