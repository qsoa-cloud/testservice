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

	// Prepare gRPC client
	var client pb.TestClient
	service.OnInit(func() error {
		conn, err := qgrpc.Dial("qcloud://testservice/")
		if err != nil {
			log.Fatalf("Cannot dial grpc: %v", err)
		}

		client = pb.NewTestClient(conn)

		return nil
	})

	// Provides HTTP service
	service.HandleHttp("/", &http.Handler{Client: client})

	// Run service
	service.Run()
}
