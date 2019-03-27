package controllers

import (
	m "github.com/uonagent/ssp-api/models"
	c "github.com/uonagent/ssp-api/config"
	"errors"
)

// StorageController contains access methods
type StorageController interface {
	GetCalculators() ([]*m.Calculator, error)
	GetCalculator(id string) (*m.Calculator, error)
}

// NewStorageController creates new StorageController using Config 
func NewStorageController(config *c.Config) (*StorageController, error) {
	var e error
	e = nil
	var sc StorageController
	switch config.StorageParams["db-name"] {
	case "mock-db":
		sc = NewMockDB(config)
	default:
		sc = nil
		e = errors.New("Unknown database type: " + config.StorageParams["db-name"])
	}
	return &sc, e
}