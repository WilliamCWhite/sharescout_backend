package lib

import (
	"time"

	"github.com/shopspring/decimal"
)

type RequestInterval struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

// Contains float value and UNIX timestamp / 1000
type DataPoint struct {
	Value float64 `json:"value"`
	Timestamp int64 `json:"time"`
}

type ServerDataPoint struct {
	Value decimal.Decimal `json:"value"`
	Timestamp int64 `json:"time"`
}

type Transaction struct {
	Timestamp int64 `json:"time"`
	Price decimal.Decimal `json:"price"`
	Shares decimal.Decimal `json:"shares"`
	Ticker string `json:"ticker"`
}

type HoldingPoint struct {
	Timestamp int64 `json:"time"`
	Price decimal.Decimal `json:"price"`
	Shares decimal.Decimal `json:"shares"`
	Input decimal.Decimal `json:"input"` // negative input is output (withdraws)
	Value decimal.Decimal `json:"value"`
}

type TestHoldingPoint struct {
	Timestamp time.Time `json:"time"`
	Price decimal.Decimal `json:"price"`
	Shares decimal.Decimal `json:"shares"`
	Input decimal.Decimal `json:"input"` // negative input is output (withdraws)
	Value decimal.Decimal `json:"value"`
}
