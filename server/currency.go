package server

import (
	"context"
	"log"

	"github.com/hashicorp/go-hclog"
	"github.com/shakirck/grpctest/data"
	protos "github.com/shakirck/grpctest/protos/currency"
)

type Currency struct {
	l hclog.Logger
	protos.UnimplementedCurrencyServer
}

func NewCurrency(l hclog.Logger) *Currency {

	return &Currency{l, protos.UnimplementedCurrencyServer{}}
}
func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.l.Info("Handle GetRate", "base ", rr.GetBase(), "destination ", rr.GetDestination())
	tr, err := data.NewRates(hclog.Default())
	if err != nil {
		return nil, err
	}
	log.Default().Printf("%#v", tr.Rates["RUB"])
	return &protos.RateResponse{Rate: float32(rr.GetBase()) * float32(tr.Rates[rr.GetDestination().String()])}, nil
}
