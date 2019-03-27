package models

// Calculator datatype container
type Calculator struct {
	ID         string            `json:"id"`
	Title      string            `json:"title"`
	Locale     string            `json:"locale"`
	Syles      []*MatrixStyle    `json:"styles"`
	ExecutorID string            `json:"executorId"`
	Params     map[string]string `json:"params"`
	//DefaultValues [][][]*DataCell   `json:"matrices"`
}

// Executor perfoms calculation
type Executor = func(in [][][]*DataCell) [][][]*DataCell

// DataCell datatype containing data from cell
type DataCell struct {
	Data      string `json:"data"`
	IsDefined bool   `json:"isDefined"`
}

// MatrixStyle contains style parameters for matrix
type MatrixStyle struct {
	MinXSize     int               `json:"minXSize"`
	MinYSize     int               `json:"minYSize"`
	MaxXSize     int               `json:"maxXSize"`
	MaxYSize     int               `json:"maxYSize"`
	InitialXSize int               `json:"initialXSize"`
	InitialYSize int               `json:"initialYSize"`
	Datatype     string            `json:"datatype"`
	IsInput      bool              `json:"isInput"`
	StyleParams  map[string]string `json:"styleParams"`
}

// CalculationRequest /calculate POST
type CalculationRequest struct {
	InputMatrices [][][]*DataCell `json:"inputMatrix"`
	ExecutorID    string          `json:"executorId"`
}

// FormatRequest /format-data POST
type FormatRequest struct {
	InputString string `json:"inputString"`
	Locale      string `json:"locale"`
	ConverterID string `json:"converterId"`
}

// FormatRespond /format-data POST
type FormatRespond struct {
	OutputString string `json:"outputString"`
	HasErrors    bool   `json:"hasErrors"`
	ErrorMessage string `json:"errorMessage"`
}
