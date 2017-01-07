package main

import "math"

//ConvertCurrencies do some calculations and returns a type CurrenciesConverted with values
func ConvertCurrencies(curr Currencies, amount int) CurrenciesConverted {
	f := float64(amount)
	converted := CurrenciesConverted{}
	converted.Amount = amount
	converted.Currency = curr.Base

	converted.Converted.AUD = toFixed(curr.Rates.AUD*f, 2)
	converted.Converted.BGN = toFixed(curr.Rates.BGN*f, 2)
	converted.Converted.BRL = toFixed(curr.Rates.BRL*f, 2)
	converted.Converted.CAD = toFixed(curr.Rates.CAD*f, 2)
	converted.Converted.CHF = toFixed(curr.Rates.CHF*f, 2)
	converted.Converted.CNY = toFixed(curr.Rates.CNY*f, 2)
	converted.Converted.CZK = toFixed(curr.Rates.CZK*f, 2)
	converted.Converted.DKK = toFixed(curr.Rates.DKK*f, 2)
	converted.Converted.GBP = toFixed(curr.Rates.GBP*f, 2)
	converted.Converted.HKD = toFixed(curr.Rates.HKD*f, 2)
	converted.Converted.HRK = toFixed(curr.Rates.HRK*f, 2)
	converted.Converted.HUF = toFixed(curr.Rates.HUF*f, 2)
	converted.Converted.IDR = toFixed(curr.Rates.IDR*f, 2)
	converted.Converted.ILS = toFixed(curr.Rates.ILS*f, 2)
	converted.Converted.INR = toFixed(curr.Rates.INR*f, 2)
	converted.Converted.JPY = toFixed(curr.Rates.JPY*f, 2)
	converted.Converted.KRW = toFixed(curr.Rates.KRW*f, 2)
	converted.Converted.MXN = toFixed(curr.Rates.MXN*f, 2)
	converted.Converted.MYR = toFixed(curr.Rates.MYR*f, 2)
	converted.Converted.NOK = toFixed(curr.Rates.NOK*f, 2)
	converted.Converted.NZD = toFixed(curr.Rates.NZD*f, 2)
	converted.Converted.PHP = toFixed(curr.Rates.PHP*f, 2)
	converted.Converted.PLN = toFixed(curr.Rates.PLN*f, 2)
	converted.Converted.RON = toFixed(curr.Rates.RON*f, 2)
	converted.Converted.RUB = toFixed(curr.Rates.RUB*f, 2)
	converted.Converted.SGD = toFixed(curr.Rates.SGD*f, 2)
	converted.Converted.THB = toFixed(curr.Rates.THB*f, 2)
	converted.Converted.TRY = toFixed(curr.Rates.TRY*f, 2)
	converted.Converted.USD = toFixed(curr.Rates.USD*f, 2)
	converted.Converted.ZAR = toFixed(curr.Rates.ZAR*f, 2)
	converted.Converted.EUR = toFixed(curr.Rates.EUR*f, 2)

	return converted
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
