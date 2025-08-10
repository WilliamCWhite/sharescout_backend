package lib

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shopspring/decimal"
)

// data.chart.result[0].events.dividends;

type DividendEvent struct {
	Amount float64 `json:"amount"`
	Date int64 `json:"date"`
}

type DividendEvents struct {
	Dividends map[string]DividendEvent `json:"dividends"`
}

type Result struct {
	Events DividendEvents `json:"events"`
}

type Chart struct {
	Result []Result `json:"result"`
}

type ChartResponse struct {
	Chart Chart `json:"chart"`
}

type DividendPoint struct {
	Amount decimal.Decimal `json:"amount"`
	Timestamp int64 `json:"timestamp"`
}

// INFO: Make sure the request interval is sufficiently large
func GetDividends(ticker string, reqInterval RequestInterval) ([]DividendPoint, error) {
	url := fmt.Sprintf("https://query2.finance.yahoo.com/v8/finance/chart/%s?period1=%d&period2=%d&interval=1mo&events=div", 
		ticker, reqInterval.StartDate.Unix(), reqInterval.EndDate.Unix(),
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error creating dividend request: %w", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0 Safari/537.36")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error getting dividends: %w", err)
	}
	defer resp.Body.Close()

	var data ChartResponse
	err = json.NewDecoder(resp.Body).Decode(&data);
	if err != nil {
		return nil, fmt.Errorf("Error decoding dividends: %w", err)
	}

	if len(data.Chart.Result) > 0 {
		dividends := data.Chart.Result[0].Events.Dividends

		dividendPoints := make([]DividendPoint, len(dividends))
		i := 0
		for _, div := range dividends {
			dividendPoints[i] = DividendPoint{
				Timestamp: div.Date,
				Amount: decimal.NewFromFloat(div.Amount),
			}
			i++
		}
		return dividendPoints, nil
	}

	return nil, fmt.Errorf("No results found")
}

