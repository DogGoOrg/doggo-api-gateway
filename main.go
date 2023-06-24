package main

// "github.com/DogGoOrg/doggo-api-gw/pb"

type server struct {
	// pb.UnimplementedGatewayServer
}

func runRest() {
	// ctx := context.Background()
	// ctx, cancel := context.WithCancel(ctx)
	// defer cancel()

	// mux := runtime.NewServeMux()
	// opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	// err := pb.RegisterGatewayHandlerFromEndpoint(ctx, mux, "localhost:12201", opts)
	// if err != nil {
	// 	panic(err)
	// }
	// log.Printf("server listening at 8081")
	// if err := http.ListenAndServe(":8081", mux); err != nil {
	// 	panic(err)
	// }
}

func runGrpc() {
	// lis, err := net.Listen("tcp", ":12201")
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
	// s := grpc.NewServer()
	// pb.RegisterGatewayServer(s, &server{})
	// log.Printf("server listening at %v", lis.Addr())
	// if err := s.Serve(lis); err != nil {
	// 	panic(err)
	// }
}

func main() {
	go runRest()
	runGrpc()
}
