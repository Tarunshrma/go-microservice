package server

import (
	"context"
	"github.com/hashicorp/go-hclog"
	"grpc-go/data"
	"grpc-go/protos/currency/protos"
	"io"
	"time"
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

func (c *Currency) SubscribeRates(src protos.Currency_SubscribeRatesServer) error {

	// handle client messages
	go func() {
		for {
			rr, err := src.Recv()
			if err == io.EOF {
				c.log.Info("Client has closed connection")
				break
			}

			// any other error means the transport between the server and client is unavailable
			if err != nil {
				c.log.Error("Unable to read from client", "error", err)
				break
			}

			c.log.Info("Handle client request", "request_base", rr.GetBase(), "request_dest", rr.GetDestination())
		}
	}()

	// handle server responses
	// we block here to keep the connection open
	for {
		err := src.Send(&protos.RateResponse{Rate: 19.2})
		if err != nil {
			return err
		}

		time.Sleep(5 * time.Second)
	}
}

//func (c *Currency) MustEmbedUnimplementedCurrencyServer() {
//	//TODO implement me
//	panic("implement me")
//}
