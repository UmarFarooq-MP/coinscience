package coinbase

import (
	"coinstrove/consts"
	"coinstrove/internal/core/domain"
	"coinstrove/internal/core/ports"
	"coinstrove/internal/core/services"
)

type newCoinBaseService struct {
	priceRepo        ports.PriceRepository
	broadcastHandler ports.BroadCastHandler
	data             domain.Response
	publisher        ports.Publisher
}

func NewCoinBaseService(priceRepo ports.PriceRepository, broadcaster ports.BroadCastHandler, publisher ports.Publisher) ports.PriceService {
	return &newCoinBaseService{
		priceRepo:        priceRepo,
		broadcastHandler: broadcaster,
		publisher:        publisher,
	}
}

func (coinbase *newCoinBaseService) GetThePrice() {
	coinbase.data = coinbase.priceRepo.Get(consts.COINBASE)
	coinbase.BroadCast()
	coinbase.WriteToQue()
}

func (coinbase *newCoinBaseService) BroadCast() {
	coinbase.broadcastHandler.BroadCast(coinbase.data)
	services.Rates = append(services.Rates, coinbase.data)

}

func (coinbase *newCoinBaseService) WriteToQue() {
	coinbase.publisher.Publish(coinbase.data)
}
