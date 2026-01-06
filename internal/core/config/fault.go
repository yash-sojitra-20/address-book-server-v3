package config

import (
	"address-book-server-v3/internal/models"

	"bitbucket.org/vayana/walt-go/logger"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func NewFaultWrapper(db *gorm.DB, logger *logger.Logger) (*i18n.Bundle, error) {
	errorMessages, err := GetErrorMessages(db)
	if err != nil {
		return nil, err
	}
	fb := i18n.NewBundle(language.English)
	msgs := make([]*i18n.Message, len(errorMessages))

	for idx, f := range errorMessages {
		msg := &i18n.Message{
			ID:    f.Code,
			One:   f.One,
			Other: f.Other,
		}
		msgs[idx] = msg
	}
	fb.AddMessages(language.English, msgs...) // #nosec G104
	return fb, nil
}

func GetErrorMessages(db *gorm.DB) ([]models.ErrorMessage, error) {
	var errorMessages []models.ErrorMessage
	if err := db.Find(&errorMessages).Error; err != nil {
		return nil, err
	}
	return errorMessages, nil
}
