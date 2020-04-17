package grpc_client

import (
	"fmt"

	"bitbucket.org/alien_soft/courier_service/config"
	pbf "bitbucket.org/alien_soft/courier_service/genproto/fare_service"
	"google.golang.org/grpc"
)

type GrpcClientI interface {
	FareService() pbf.FareServiceClient
}

type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

func New(cfg config.Config) (*GrpcClient, error) {
	connFare, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.FareServiceHost, cfg.FareServicePort),
		grpc.WithInsecure())

	if err != nil {
		return nil, fmt.Errorf("user service dial host: %s port: %d",
			cfg.FareServiceHost, cfg.FareServicePort)
	}

	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"fare_service": pbf.NewFareServiceClient(connFare),
		},
	}, nil
}

func (g *GrpcClient) FareService() pbf.FareServiceClient {
	return g.connections["fare_service"].(pbf.FareServiceClient)
}
