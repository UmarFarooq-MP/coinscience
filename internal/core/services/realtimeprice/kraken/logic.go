package kraken

import (
	"coinstrove/consts"
	"coinstrove/internal/core/domain"
	"coinstrove/internal/core/ports"
	"coinstrove/internal/core/services"
)

type newKrakenService struct {
	priceRepo        ports.PriceRepository
	broadcastHandler ports.BroadCastHandler
	data             domain.Response
	publisher        ports.Publisher
}

func NewKrakenService(priceRepo ports.PriceRepository, broadcaster ports.BroadCastHandler, publisher ports.Publisher) ports.PriceService {
	return &newKrakenService{
		priceRepo:        priceRepo,
		broadcastHandler: broadcaster,
		publisher:        publisher,
	}
}

func (kraken *newKrakenService) GetThePrice() {
	kraken.data = kraken.priceRepo.Get(consts.KRAKEN)
	kraken.BroadCast()
	kraken.WriteToQue()
}

func (kraken *newKrakenService) BroadCast() {
	kraken.broadcastHandler.BroadCast(kraken.data)
	services.Rates = append(services.Rates, kraken.data)

}

func (kraken *newKrakenService) WriteToQue() {
	kraken.publisher.Publish(kraken.data)
}
