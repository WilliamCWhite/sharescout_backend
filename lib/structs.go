package lib

import (
	"time"

	"github.com/shopspring/decimal"
)

type RequestInterval struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}


type ApiPoint struct {
	Value decimal.Decimal `json:"value"`
	Timestamp int64 `json:"time"`
}

type ResponsePoint struct {
	Timestamp int64 `json:"timestamp"`
	Price float64 `json:"price"`
	PercentGrowth float64 `json:"percentGrowth"`
	ThousandIn float64 `json:"thousandIn"`
	PercentReturns float64 `json:"percentReturns"`
}

type DividendPoint struct {
	Amount decimal.Decimal `json:"amount"`
	Timestamp int64 `json:"timestamp"`
}
