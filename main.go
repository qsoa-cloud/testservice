package main

import (
	"database/sql"
	"log"

	"gopkg.qsoa.cloud/service"
	"gopkg.qsoa.cloud/service/qgrpc"
	"gopkg.qsoa.cloud/service/qhttp"
	_ "gopkg.qsoa.cloud/service/qmysql"

	"testservice/grpc"
	"testservice/grpc/pb"
	"testservice/http"
)

func main() {
	// Prepare gRpc client
	conn, err := qgrpc.Dial("qcloud://" + service.GetService() + "/")
	if err != nil {
		log.Fatalf("Cannot dial grpc: %v", err)
	}
	defer conn.Close()

	grpcClient := pb.NewTestClient(conn)

	// Prepare mysql connection
	db, err := sql.Open("qmysql", "example_db")
	if err != nil {
		log.Fatalf("Cannot open mysql database: %v", err)
	}
	defer db.Close()

	// Provide HTTP service
	qhttp.Handle("/", http.New(grpcClient, db))

	// Provide gRPC service
	pb.RegisterTestServer(qgrpc.GetServer(), grpc.Server{})

	// Run service
	service.Run()
}
