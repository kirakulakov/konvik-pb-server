package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/kirakulakov/konvik-pb-server"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type konvikServer struct {
	pb.UnimplementedKonvikServer
}

func (s *konvikServer) Ping(ctx context.Context, _ *pb.PingRequest) (*pb.PingResponse, error) {

	response := pb.PingResponse{
		Timestamp: timestamppb.Now(),
	}
	return &response, nil
}

func newServer() *konvikServer {
	s := &konvikServer{}
	return s
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		log.Fatal(fmt.Sprintf("Error while creating listener:%v", err))
	}

	srv := grpc.NewServer()
	pb.RegisterKonvikServer(srv, newServer())

	log.Println("Server started!")
	if err := srv.Serve(lis); err != nil {
		log.Fatal(fmt.Sprintf("Error while starting server:%v", err))
	}

}
