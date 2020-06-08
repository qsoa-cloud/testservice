package main

import (
	"log"

	"gopkg.qsoa.cloud/service"
	qgrpc "gopkg.qsoa.cloud/service/grpc"

	"testservice/grpc"
	"testservice/grpc/pb"
	"testservice/http"
)

func main() {
	// Provides gRPC service
	pb.RegisterTestServer(service.GetGrpcServer(), &grpc.Server{})

	// Provides HTTP service
	httpHandler := &http.Handler{}
	service.HandleHttp("/", httpHandler)

	service.OnInit(func() error {
		// Prepare gRPC client

		conn, err := qgrpc.Dial("qcloud://testservice/")
		if err != nil {
			log.Fatalf("Cannot dial grpc: %v", err)
		}

		client := pb.NewTestClient(conn)
		httpHandler.Client = client

		return nil
	})

	// Run service
	service.Run()
}
