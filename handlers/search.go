package handlers

import (
	"fmt"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

type Quote struct {
	Symbol string `json:"symbol"`
	Shortname string `json:"shortname"`
	Type string `json:"typeDisp"`
	Sector string `json:"sectorDisp"`
}

type SearchResponse struct {
	Quotes []Quote `json:"quotes"`
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Println("SearchHandler can only receive get requests")
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	input := vars["input"]
	encodedInput := url.QueryEscape(input)
	
	fmt.Printf("Processing search with input %v\n", encodedInput)


	url := fmt.Sprintf("https://query2.finance.yahoo.com/v1/finance/search?q=%s", encodedInput)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error creating autocomplete request: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0 Safari/537.36")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error getting autocomplete response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	
	var result SearchResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Printf("Error decoding json: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var sendData []Quote

	j := 0
	for i := 0; i < len(result.Quotes); i++ {
		if (j >= 5) {
			break;
		}
		if (result.Quotes[i].Type == "Option") {
			continue
		} else {
			sendData = append(sendData, result.Quotes[i])
			j++
		}

	}


	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(sendData)
	if err != nil {
		fmt.Printf("Error encoding json: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
