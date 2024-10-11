package services

import (
	"coinstrove/consts"
	"coinstrove/internal/core/domain"
	"flag"
)

type LowestRate struct {
	ExchangeName consts.EXCHANGE
	//Currency consts.CURRENCY
	Rate float32
}

// Binance , ETH , Rate

// create a an arary to store all the rates
var Rates map[string]domain.Response

func GetRatesCache() map[string]domain.Response {
	if Rates == nil {
		Rates = make(map[string]domain.Response)
	}
	return Rates
}

var cheapestRates map[consts.CURRENCY]float32

func GetCheapestRatesCache() map[consts.CURRENCY]float32 {
	if cheapestRates == nil {
		cheapestRates = make(map[consts.CURRENCY]float32)
	}
	return cheapestRates
}

func returnTheLowestValues(values []float32) float32 {

}

// Binance.  Currencies
func UpdateLowestRatesCache(coinName string) {
	lowestValue := 0
	exchangeName := ""
	for key,value := range Rates {
		Rates[coinName].Data.
	}
}
