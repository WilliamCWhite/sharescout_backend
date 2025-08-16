package lib

import (
	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
)

// Uses piquette/finance-go to get the closing prices and timestamps for a ticker across an interval
func GetApiPoints(ticker string, reqInterval RequestInterval) []ApiPoint {
	interval := DetermineInterval(reqInterval.StartDate, reqInterval.EndDate)

	params := &chart.Params{
		Symbol: ticker,
		Start: datetime.New(&reqInterval.StartDate),
		End: datetime.New(&reqInterval.EndDate),
		Interval: interval,
	}

	iter := chart.Get(params)

	jsonRes := []ApiPoint{}
	for iter.Next() {
		bar := iter.Bar()
		rBar := ApiPoint{
			Value: bar.Close,
			Timestamp: int64(bar.Timestamp), //bar.Timestamp happens to give the exact data format the frontend wants
		}
		jsonRes = append(jsonRes, rBar)
	}
	return jsonRes
}
