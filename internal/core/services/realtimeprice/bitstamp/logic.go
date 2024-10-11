package bitstamp

import (
	"coinstrove/consts"
	"coinstrove/internal/core/domain"
	"coinstrove/internal/core/ports"
	"coinstrove/internal/core/services"
)

type newBitstampService struct {
	priceRepo        ports.PriceRepository
	broadcastHandler ports.BroadCastHandler
	data             domain.Response
	publisher        ports.Publisher
}

func NewBitstampService(priceRepo ports.PriceRepository, broadcaster ports.BroadCastHandler, publisher ports.Publisher) ports.PriceService {
	return &newBitstampService{
		priceRepo:        priceRepo,
		broadcastHandler: broadcaster,
		publisher:        publisher,
	}
}

func (bitstamp *newBitstampService) GetThePrice() {
	bitstamp.data = bitstamp.priceRepo.Get(consts.BITSTAMP)
	bitstamp.BroadCast()
	bitstamp.WriteToQue()
}

func (bitstamp *newBitstampService) BroadCast() {
	bitstamp.broadcastHandler.BroadCast(bitstamp.data)
	services.Rates = append(services.Rates, bitstamp.data)

}

func (bitstamp *newBitstampService) WriteToQue() {
	bitstamp.publisher.Publish(bitstamp.data)
}
