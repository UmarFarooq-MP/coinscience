package bitpay

import (
	"coinstrove/consts"
	"coinstrove/internal/core/domain"
	"coinstrove/internal/core/ports"
	"coinstrove/internal/core/services"
)

type newBitPayService struct {
	priceRepo        ports.PriceRepository
	broadcastHandler ports.BroadCastHandler
	data             domain.Response
	publisher        ports.Publisher
}

func NewBitPayService(priceRepo ports.PriceRepository, broadcaster ports.BroadCastHandler, publisher ports.Publisher) ports.PriceService {
	return &newBitPayService{
		priceRepo:        priceRepo,
		broadcastHandler: broadcaster,
		publisher:        publisher,
	}
}

func (bitpay *newBitPayService) GetThePrice() {
	bitpay.data = bitpay.priceRepo.Get(consts.BITPAY)
	bitpay.BroadCast()
	bitpay.WriteToQue()
}

func (bitpay *newBitPayService) BroadCast() {
	bitpay.broadcastHandler.BroadCast(bitpay.data)
	services.Rates = append(services.Rates, bitpay.data)

}

func (bitpay *newBitPayService) WriteToQue() {
	bitpay.publisher.Publish(bitpay.data)
}
