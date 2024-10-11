package kucoin

import (
	"coinstrove/consts"
	"coinstrove/internal/core/domain"
	"coinstrove/internal/core/ports"
	"coinstrove/internal/core/services"
)

type newKucoinService struct {
	priceRepo        ports.PriceRepository
	broadcastHandler ports.BroadCastHandler
	data             domain.Response
	publisher        ports.Publisher
}

func NewKucoinService(priceRepo ports.PriceRepository, broadcaster ports.BroadCastHandler, publisher ports.Publisher) ports.PriceService {
	return &newKucoinService{
		priceRepo:        priceRepo,
		broadcastHandler: broadcaster,
		publisher:        publisher,
	}
}

func (kucoin *newKucoinService) GetThePrice() {
	kucoin.data = kucoin.priceRepo.Get(consts.KUCOIN)
	kucoin.BroadCast()
	kucoin.WriteToQue()
}

func (kucoin *newKucoinService) BroadCast() {
	kucoin.broadcastHandler.BroadCast(kucoin.data)
	services.Rates = append(services.Rates, kucoin.data)

}

func (kucoin *newKucoinService) WriteToQue() {
	kucoin.publisher.Publish(kucoin.data)
}
