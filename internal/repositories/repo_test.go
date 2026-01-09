package repositories

import (
	"address-book-server-v3/internal/common/types"
	// "address-book-server-v3/internal/models"
	"address-book-server-v3/test"
	"fmt"
	"log"
	"testing"

	wgconfig "bitbucket.org/vayana/walt-gorm.go/config"
	wgormconfig "bitbucket.org/vayana/walt-gorm.go/config"
	"github.com/golang-migrate/migrate/v4"
	sql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/suite"

	"gorm.io/gorm"
)

type RepositoryTestSuite struct {
	suite.Suite
	suite *test.TestingSuite
	db    *gorm.DB
}

func (suite *RepositoryTestSuite) SetupSuite() {
	// fmt.Println("=======================> Starting test suite setup")

	// Init application test container
	suite.suite = test.NewTestingSuite(nil, nil)

	dbConfig, err := wgconfig.NewDBConfig(
		types.DB_HOSTNAME,
		types.DB_PORT,
		types.DB_USERNAME,
		types.DB_PASSWORD,
		types.DB_NAME,
		types.DB_TYPE,
	).Get()
	// fmt.Println("=================================> DB config: ", dbConfig)
	if err != nil {
		log.Fatalf("Could not connect to MySQL database: %v", err)

	}

	// fmt.Println("=================================> Before ConnectToDatabase(): ")
	db, err := wgormconfig.ConnectToDatabase(dbConfig).Get()
	// fmt.Println("=================================> After ConnectToDatabase(): ")

	if err != nil {
		log.Fatalf("Could not connect to MySQL database: %v", err)
	}

	suite.db = db
	// fmt.Println("=======================> DB configured")

	// Run migrations
	if err := suite.migrate(); err != nil {
		log.Fatalf("Could not migrate database: %v", err)
	}

	// Seed minimal required data
	// suite.db.Create(&models.ErrorMessage{
	// 	Code: "test",
	// 	One:  "test",
	// })

}

func (suite *RepositoryTestSuite) SetupTest() {
	suite.db.Exec("DELETE FROM users")
	suite.db.Exec("DELETE FROM addresses")
	// suite.db.Exec("DELETE FROM error_messages WHERE code='test' AND one='test'")
}

func (suite *RepositoryTestSuite) TearDownSuite() {
	sqlDB, err := suite.db.DB()
	if err != nil {
		log.Fatalf("Could not get DB: %v", err)
	}
	sqlDB.Close()
}

func (suite *RepositoryTestSuite) migrate() error {
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

func TestRepositorySuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}
