package main

import (
	"fmt"
	pb "grpcServer/datafiles"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port      = ":50051"
	noOfSteps = 3
)

type server struct {
	pb.UnimplementedMoneyTransactionServer
}

func (s *server) MakeTransaction(in *pb.TransactionRequest, stream pb.MoneyTransaction_MakeTransactionServer) error {
	log.Printf("Got request for money Transfer...")
	log.Printf("Amount: %f, from A/c:%s, to A/c:%s", in.Amount, in.From, in.To)

	for i := 0; i < noOfSteps; i++ {
		time.Sleep(time.Second * 3)
		if err := stream.Send(&pb.TransactionResponse{
			Status:      "good",
			Step:        int32(i),
			Description: fmt.Sprintf("Description of step %d", int32(i)),
		}); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, "status", err)
		}
	}
	log.Printf("Successfully transfered amount %v from %v to %v", in.Amount, in.From, in.To)
	return nil
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
