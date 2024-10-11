package gate_io

import (
	"coinstrove/consts"
	"coinstrove/internal/core/domain"
	"coinstrove/internal/core/ports"
	"coinstrove/internal/core/services"
)

type newGateIOService struct {
	priceRepo        ports.PriceRepository
	broadcastHandler ports.BroadCastHandler
	data             domain.Response
	publisher        ports.Publisher
}

func NewGateIOService(priceRepo ports.PriceRepository, broadcaster ports.BroadCastHandler, publisher ports.Publisher) ports.PriceService {
	return &newGateIOService{
		priceRepo:        priceRepo,
		broadcastHandler: broadcaster,
		publisher:        publisher,
	}
}

func (gateio *newGateIOService) GetThePrice() {
	gateio.data = gateio.priceRepo.Get(consts.GATEIO)
	gateio.BroadCast()
	gateio.WriteToQue()
}

func (gateio *newGateIOService) BroadCast() {
	gateio.broadcastHandler.BroadCast(gateio.data)
	services.Rates = append(services.Rates, gateio.data)

}

func (gateio *newGateIOService) WriteToQue() {
	gateio.publisher.Publish(gateio.data)
}
