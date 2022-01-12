package main

import (
	"context"
	"crypto/tls"
	"log"

	ia "github.com/oridisboss/investAPI"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

const (
	address     = "invest-public-api.tinkoff.ru:443"
	defaultName = "Robotlatform"
)

var Instruments ia.InstrumentsServiceClient
var MarketData ia.MarketDataServiceClient

var ctx context.Context
var conn *grpc.ClientConn
var md metadata.MD

func GRPCConnect() {
	ctx = context.Background()

	var err error
	conn, err = grpc.Dial(address,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect to grpc server: %v", err)
	}

	md = metadata.New(map[string]string{"Authorization": "Bearer " + GetSetting("token")})
	ctx = metadata.NewOutgoingContext(ctx, md)

	Instruments = ia.NewInstrumentsServiceClient(conn)
	MarketData = ia.NewMarketDataServiceClient(conn)
}
