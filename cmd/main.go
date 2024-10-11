package main

import (
	"log"
	"net/http"
	"time"

	"coinstrove/api/websocket"
	"coinstrove/internal/core/ports"
	"coinstrove/internal/core/publisher"
	"coinstrove/internal/core/services/realtimeprice/binance"
	"coinstrove/internal/core/services/realtimeprice/bitfinex"
	"coinstrove/internal/core/services/realtimeprice/bitstamp"
	"coinstrove/internal/core/services/realtimeprice/coinbase"
	gate_io "coinstrove/internal/core/services/realtimeprice/gate.io"
	"coinstrove/internal/core/services/realtimeprice/huobi"
	"coinstrove/internal/core/services/realtimeprice/kraken"
	"coinstrove/internal/core/services/realtimeprice/kucoin"
	"coinstrove/internal/core/services/realtimeprice/okx"
	"coinstrove/repositories/apirepository"
)

func getPriceAfterFiveSeconds(priceService []ports.PriceService) {
	ticker := time.NewTicker(7 * time.Second)
	for range ticker.C {
		log.Printf("Starting To Fetch Latest Price")
		for _, value := range priceService {
			go value.GetThePrice()
		}
		log.Printf("Fetch Price competed")
	}
}

func startServer() {
	log.Println("Starting Server on port")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func main() {
	apiRepo := apirepository.NewAPIRepository()
	handler := websocket.NewHandler()
	broadCastManager := websocket.NewBroadcastManager(handler)

	quePublisher, err := publisher.NewRabbitMQPublisher("amqp://guest:guest@queue:5672/")
	if err != nil {
		log.Printf("Error while Initiating Rabbit MQ Connection with message %v", err)
	}
	quePublisher.Init()
	defer quePublisher.Close()

	websocket.NewRouter(handler)
	go startServer()

	priceService := []ports.PriceService{
		binance.NewBinanceService(apiRepo, broadCastManager, quePublisher),
		gate_io.NewGateIOService(apiRepo, broadCastManager, quePublisher),
		kraken.NewKrakenService(apiRepo, broadCastManager, quePublisher),
		coinbase.NewCoinBaseService(apiRepo, broadCastManager, quePublisher),
		// bitpay.NewBitPayService(apiRepo, broadCastManager, quePublisher),
		bitfinex.NewBitfinexService(apiRepo, broadCastManager, quePublisher),
		bitstamp.NewBitstampService(apiRepo, broadCastManager, quePublisher),
		huobi.NewHuobiService(apiRepo, broadCastManager, quePublisher),
		kucoin.NewKucoinService(apiRepo, broadCastManager, quePublisher),
		okx.NewOkxService(apiRepo, broadCastManager, quePublisher),
	}

	// WebSocket Endpoints
	go getPriceAfterFiveSeconds(priceService)

	select {}
}
