package main

import (
	"context"
	pb "grpcServer/datafiles"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedMoneyTransactionServer
}

func (s *server) MakeTransaction(ctx context.Context, in *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	log.Printf("Got request for money Transfer...")
	log.Printf("Amount: %f, from A/c:%s, to A/c:%s", in.Amount, in.From, in.To)

	return &pb.TransactionResponse{
		Confirmation: true,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMoneyTransactionServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to server: %v", err)
	}
}
