package repositories

import (
	"bitbucket.org/vayana/walt-go/logger"
	"gorm.io/gorm"
)


type RepoContext struct {
	db *gorm.DB
	log *logger.Logger
}

func NewRepoContext(db *gorm.DB, log *logger.Logger) *RepoContext {
	return &RepoContext{
		db:  db,
		log: log,
	}
}