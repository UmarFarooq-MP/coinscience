package services

import (
	"coinstrove/internal/core/domain"
)

type LowestRate struct {
	ExchangeName string
	// Currency consts.CURRENCY
	Rate string
}

// Binance , ETH , Rate

// create a an arary to store all the rates
// map from Exchange -> Exchange Response object
var Rates map[string]domain.Response

func GetRatesCache() map[string]domain.Response {
	if Rates == nil {
		Rates = make(map[string]domain.Response)
	}
	return Rates
}

// map from Currency -> LowestRate
var cheapestRates map[string]LowestRate

func GetCheapestRatesCache() map[string]LowestRate {
	if cheapestRates == nil {
		cheapestRates = make(map[string]LowestRate)
	}
	return cheapestRates
}

// func returnTheLowestValues(values []float32) float32 {
// }

// Binance.  Currencies
func UpdateCheapestRates() {
	tempVals := map[string]LowestRate{}

	for exchangeName, exchangeResponse := range Rates {
		for _, currency := range exchangeResponse.Data.Currencies {

			// if the currency is not in the map
			if _, ok := tempVals[currency.Name]; !ok {
				tempVals[currency.Name] = LowestRate{
					ExchangeName: exchangeName,
					Rate:         currency.Price,
				}
			}

			// if the currency is in the map
			if currency.Price > tempVals[currency.Name].Rate {
				tempVals[currency.Name] = LowestRate{
					ExchangeName: exchangeName,
					Rate:         currency.Price,
				}
			}
		}
	}

	cheapestRates = tempVals
}
