package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/raufhm/calculator_go_grpc/calculatorpb/v1"
	"google.golang.org/grpc"
)

type Server struct {
	calculatorpb.UnimplementedCalculatorServiceServer
}

func (server *Server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Receive Sum RPC: %v", req)
	firstNumber := req.FirstNumber
	secondNumber := req.SecondNumber
	result := firstNumber + secondNumber
	res := &calculatorpb.SumResponse{
		SumResult: result,
	}

	return res, nil
}

func main() {
	fmt.Print("Calculator Server\n")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}

}
