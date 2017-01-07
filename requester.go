package main

import (
	"encoding/json"
	"log"
	"net/http"
)

//GetCurrencies uses a client and sends a request, then decoding to specified type Currencies
func GetCurrencies(url string) Currencies {

	// Build  request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	// create a Client
	client := &http.Client{}

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}

	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record Currencies

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	return record
}
