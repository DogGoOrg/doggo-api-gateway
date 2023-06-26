package server

import (
	"log"

	"github.com/DogGoOrg/doggo-api-gateway/proto/proto_services/Account"
	"github.com/DogGoOrg/doggo-api-gateway/proto/proto_services/Tracker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceConnections struct {
	AccountService *Account.AccountClient
	TrackerService *Tracker.TrackerClient
}

const (
	accountServiceAddr = "localhost:50051"
	trackerServiceAddr = "localhost:50052"
)

func dialServiceConnection(addr string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		conn.Close()
		return nil, err
	}

	return conn, nil
}

func EstablishConnection() *ServiceConnections {
	connList := new(ServiceConnections)

	//account service
	acc, err := dialServiceConnection(accountServiceAddr)
	if err != nil {
		panic(err)
	}

	accountServiceClient := Account.NewAccountClient(acc)
	connList.AccountService = &accountServiceClient

	//tracker service
	track, err := dialServiceConnection(trackerServiceAddr)
	if err != nil {
		panic(err)
	}

	trackerServiceClient := Tracker.NewTrackerClient(track)
	connList.TrackerService = &trackerServiceClient

	return connList
}
