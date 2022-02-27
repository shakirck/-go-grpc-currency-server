package main

import (
	"net"

	"github.com/hashicorp/go-hclog"
	protos "github.com/shakirck/grpctest/protos/currency"
	"github.com/shakirck/grpctest/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()
	gs := grpc.NewServer()
	cs := server.NewCurrency(log)
	protos.RegisterCurrencyServer(gs, cs)
	reflection.Register(gs)
	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("osmething wronggo", err)
	}

	gs.Serve(l)

}
