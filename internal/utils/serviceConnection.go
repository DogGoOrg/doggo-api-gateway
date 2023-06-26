package utils

import (
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcConnectionController struct{}

func (p *GrpcConnectionController) ConnGrpc(envParamName string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(os.Getenv(envParamName), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
		return nil, err
	}

	return conn, nil
}
