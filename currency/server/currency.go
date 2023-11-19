package server

import (
	"context"
	"github.com/hashicorp/go-hclog"
	"grpc-go/data"
	"grpc-go/protos/currency/protos"
)

type Currency struct {
	log   hclog.Logger
	rates *data.ExchangeRates
}

func NewCurrency(l hclog.Logger, r *data.ExchangeRates) *Currency {
	return &Currency{l, r}
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle request for GetRate", "base", rr.GetBase(), "dest", rr.GetDestination())
	r, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())

	if err != nil {
		c.log.Error("Error in fetching rates", "Error:", err)
		return nil, err
	}

	return &protos.RateResponse{Rate: r}, nil
}

//func (c *Currency) MustEmbedUnimplementedCurrencyServer() {
//	//TODO implement me
//	panic("implement me")
//}
