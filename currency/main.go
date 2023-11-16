package main

import (
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc-go/protos/currency/protos"
	"net"
	"os"

	"grpc-go/server"
)

func main() {
	log := hclog.Default()

	// create an instance of the Currency server
	cs := server.NewCurrency(log)

	// create a new gRPC server, use WithInsecure to allow http connections
	gs := grpc.NewServer()

	protos.RegisterCurrencyServer(gs, cs)

	//TODO: This is for dev environment only
	//Help grpcurl command to list down the methods/services
	reflection.Register(gs)

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to create listener", "error", err)
		os.Exit(1)
	}
	gs.Serve(l)
}
