package controllers

import (
	"errors"
	"strconv"

	c "github.com/uonagent/ssp-api/config"
	m "github.com/uonagent/ssp-api/models"
)

// MockDB implements fake db
type MockDB struct {
	CalculatorsRepo []*m.Calculator
}

// NewMockDB creates new MockDB instance
func NewMockDB(config *c.Config) *MockDB {
	calcs := make([]*m.Calculator, 1, 1)

	calcMatricesStyles := make([]*m.MatrixStyle, 4)
	calcMatricesStyles[0] = &m.MatrixStyle{
		MinXSize: 1,
		MinYSize: 1,
		MaxXSize: -1,
		MaxYSize: -1,
		InitialXSize: 3,
		InitialYSize: 3,
		Datatype: "realnumber",
		IsInput: true,
	}
	calcMatricesStyles[1] = &m.MatrixStyle{
		MinXSize: 1,
		MinYSize: 1,
		MaxXSize: 1,
		MaxYSize: -1,
		InitialXSize: 1,
		InitialYSize: 3,
		Datatype: "text",
		IsInput: true,
	}
	calcMatricesStyles[2] = &m.MatrixStyle{
		InitialXSize: 1,
		InitialYSize: 1,
		IsInput: false,
		Datatype: "money-BYN",
	}
	calcMatricesStyles[3] = &m.MatrixStyle{
		InitialXSize: 2,
		InitialYSize: 2,
		IsInput: false,
		Datatype: "money-BYN",
	}
	calcs[0] = &m.Calculator{
		ID: "1",
		Title: "test",
		Locale: "by-BY",
		Syles: calcMatricesStyles,
		ExecutorID: "1",
	}
	return &MockDB{
		CalculatorsRepo: calcs,
	}
}

// GetCalculators GetCalculators
func (m *MockDB) GetCalculators() ([]*m.Calculator, error) {
	return m.CalculatorsRepo, nil
}

// GetCalculator GetCalculator
func (m *MockDB) GetCalculator(id string) (*m.Calculator, error) {
	index, e := strconv.Atoi(id)
	if e != nil {
		return nil, e
	}
	if index < 0 || index > len(m.CalculatorsRepo) {
		return nil, errors.New("Index ot of range")
	}
	return m.CalculatorsRepo[index], nil
}
