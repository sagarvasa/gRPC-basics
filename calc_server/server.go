package main

import (
	"context"
	"fmt"
	"log"
	"net"

	calcpb "github.com/gRPC-basics/calc_pb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) DoSum(ctx context.Context, req *calcpb.SumRequest) (*calcpb.SumResponse, error) {
	var firstNumber int32 = req.Number1
	var secondNumber int32 = req.Number2
	var sum int32 = firstNumber + secondNumber

	res := &calcpb.SumResponse{
		Sum: int64(sum),
	}

	return res, nil

}

func (*server) CalcDiff(ctx context.Context, req *calcpb.DiffRequest) (*calcpb.DiffResponse, error) {
	firstNumber := req.Number1
	secondNumber := req.Number2

	res := &calcpb.DiffResponse{
		Difference: firstNumber - secondNumber,
	}

	return res, nil

}

func (*server) Multiply(ctx context.Context, req *calcpb.MultiplierRequest) (*calcpb.MultiplierResponse, error) {
	firstNumber := req.Number1
	secondNumber := req.Number2
	thirdNumber := req.Number3

	res := &calcpb.MultiplierResponse{
		Resp: firstNumber * secondNumber * thirdNumber,
	}

	return res, nil

}

func main() {
	fmt.Println("Starting Server")

	listener, err := net.Listen("tcp", "0.0.0.0:50001")
	if err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}

	s := grpc.NewServer()
	calcpb.RegisterCalculatorServiceServer(s, &server{})
	if err = s.Serve(listener); err != nil {
		log.Fatalf("Failed to Serve: %v", err)
	}

}
