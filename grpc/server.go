//go:generate protoc -I pb --go_out=plugins=grpc:pb pb/service.proto
package grpc

import (
	"context"
	"log"

	"google.golang.org/grpc/status"

	"testservice/grpc/pb"
)

type Server struct{}

func (Server) Sum(ctx context.Context, req *pb.SumReq) (*pb.SumResp, error) {
	return &pb.SumResp{Sum: req.N1 + req.N2}, nil
}

func (Server) PingPong(server pb.Test_PingPongServer) error {
	for {
		msg, err := server.Recv()
		if err != nil {
			if err.Error() != "EOF" {
				log.Printf("Recv err: %v, code: %s", err, status.Code(err))
			}
			break
		}

		if err := server.Send(&pb.Pong{Text: msg.Text}); err != nil {
			log.Printf("Send err: %v, code: %s", err, status.Code(err))
			break
		}
	}

	return nil
}
