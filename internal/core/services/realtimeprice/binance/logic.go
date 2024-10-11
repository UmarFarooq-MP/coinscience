package binance

import (
	"coinstrove/consts"
	"coinstrove/internal/core/domain"
	"coinstrove/internal/core/ports"
	"coinstrove/internal/core/services"
)

type newBinanceService struct {
	priceRepo        ports.PriceRepository
	data             domain.Response
	broadcastHandler ports.BroadCastHandler
	publisher        ports.Publisher
}

func NewBinanceService(priceRepo ports.PriceRepository, broadcaster ports.BroadCastHandler, publisher ports.Publisher) ports.PriceService {
	return &newBinanceService{
		priceRepo:        priceRepo,
		broadcastHandler: broadcaster,
		publisher:        publisher,
	}
}

func (binance *newBinanceService) GetThePrice() {
	binance.data = binance.priceRepo.Get(consts.BINANCE)
	binance.BroadCast()
	binance.WriteToQue()
}

func (binance *newBinanceService) BroadCast() {
	binance.broadcastHandler.BroadCast(binance.data)
	rates := services.GetRatesCache()
	rates[binance.data.Data.ExchangeName] = binance.data
}

func (binance *newBinanceService) WriteToQue() {
	binance.publisher.Publish(binance.data)
}
