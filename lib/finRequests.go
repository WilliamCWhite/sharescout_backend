package lib

import (
	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
)

func GetServerDataPoints(ticker string, reqInterval RequestInterval) []ServerDataPoint {
	interval := DetermineInterval(reqInterval.StartDate, reqInterval.EndDate)

	params := &chart.Params{
		Symbol: ticker,
		Start: datetime.New(&reqInterval.StartDate),
		End: datetime.New(&reqInterval.EndDate),
		Interval: interval,
	}

	iter := chart.Get(params)

	jsonRes := []ServerDataPoint{}
	for iter.Next() {
		bar := iter.Bar()
		rBar := ServerDataPoint{
			Value: bar.Close,
			Timestamp: int64(bar.Timestamp), //bar.Timestamp happens to give the exact data format the frontend wants
		}
		jsonRes = append(jsonRes, rBar)
	}
	return jsonRes
}

