package controllers

import (
	m "github.com/uonagent/ssp-api/models"
)

// StorageController contains access methods
type StorageController interface {
	GetCalculators() []*m.Calculator
	GetCalculator(id string) *m.Calculator
	GetExecutor(id string) m.Executor
	GetFormatter(id string) m.Formatter
}
