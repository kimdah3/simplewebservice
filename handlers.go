package main

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"strconv"
)

//ConvertCurrency responds to request "convert?amount={amount}&currency={currency}",
//then caluclates and converts values, then returns json or xml depending on headers specification.
func ConvertCurrency(w http.ResponseWriter, r *http.Request) {
	//Get values from query params, and the content-type header
	amount, _ := strconv.Atoi(r.FormValue("amount"))
	currency := r.FormValue("currency")
	ctype := r.Header.Get("Content-Type")
	status := http.StatusOK //HTTP status: OK for default

	//Gets the requested values
	currencies := GetCurrencies("http://api.fixer.io/latest?base=" + currency)

	//if values are empty, the requested currency is not available
	if currencies.Base == "" {
		status = http.StatusBadRequest
	}

	//calculates and converts the request into our formatting
	convertedCurrs := ConvertCurrencies(currencies, amount)

	//if values are empty and the request was ok, something went wrong internally in converting
	if convertedCurrs.Currency == "" && status != http.StatusBadRequest {
		status = http.StatusInternalServerError
	}

	//sets the responsetype
	if status == http.StatusOK {
		if ctype == "application/xml" {
			w.Header().Set("Content-Type", "application/xml; charset=UTF-8")
			if err := xml.NewEncoder(w).Encode(convertedCurrs); err != nil {
				panic(err)
			}
		} else {
			//default json
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			if err := json.NewEncoder(w).Encode(convertedCurrs); err != nil {
				panic(err)
			}
		}
	}

	//writes the respons
	w.WriteHeader(status)

}
