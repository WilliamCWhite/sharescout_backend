package lib

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func GenerateResponsePoints(apiPoints []ApiPoint, dividends []DividendPoint) ([]ResponsePoint, error) {
	if len(apiPoints) == 0 {
		return nil, fmt.Errorf("apiPoints is empty")
	}


	responsePoints := make([]ResponsePoint, len(apiPoints))

	initialTimestamp := apiPoints[0].Timestamp
	initialPrice := apiPoints[0].Value
	initialBalance := decimal.NewFromFloat(1000)
	hundred := decimal.NewFromFloat(100)
	shares := initialBalance.Div(initialPrice)
	j := 0 // used to iterate through dividends
	for i := 0; i < len(apiPoints); {
		// Conditions to add transaction
		if len(dividends) > 0 && j < len(dividends) {
			// Skip dividends that come before any api point
			if (dividends[j].Timestamp < initialTimestamp) {
				j++
				continue
			}

			// For any dividend before the upcoming point, buy shares with
			// the dividend at the upcoming closing price
			if (dividends[j].Timestamp < apiPoints[i].Timestamp) {
				// add dividendAmount/Price shares
				shares = shares.Add( dividends[j].Amount.Mul(shares).Div(apiPoints[i].Value) )
				j++
			}
		}

		current_price := apiPoints[i].Value
		price, _ := current_price.Float64()
		// (p - p0) / p0 * 100
		percentGrowth, _ := (current_price.Sub(initialPrice)).Div(initialPrice).Mul(hundred).Float64()
		thousandInDecimal := shares.Mul(current_price)
		thousandIn, _ := thousandInDecimal.Float64()
		// (p - p0) / p0 * 100
		percentReturns, _ := (thousandInDecimal.Sub(initialBalance)).Div(initialBalance).Mul(hundred).Float64()

		responsePoints[i] = ResponsePoint{
			Timestamp: apiPoints[i].Timestamp,
			Price: price,
			PercentGrowth: percentGrowth,
			ThousandIn: thousandIn,
			PercentReturns: percentReturns,
		}
		i++
	}

	return responsePoints, nil
}
