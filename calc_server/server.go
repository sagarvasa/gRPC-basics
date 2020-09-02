package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"strconv"
	"time"

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

func (*server) GetByteValues(req *calcpb.GigaByteRequest, stream calcpb.CalculatorService_GetByteValuesServer) error {
	fmt.Println("********************")
	fmt.Println("In GetByteValues Method")
	fmt.Println("********************")
	gbValue, _ := strconv.ParseFloat(req.GbValue, 64)
	gbValue = gbValue * (math.Pow(2, 30))
	multiplier := 1.0

	for gbValue > 1 {

		gbValue = gbValue / (math.Pow(2, 5))
		multiplier = multiplier * (math.Pow(2, 5))

		err := stream.Send(&calcpb.GigaByteResponse{
			Resp: (strconv.FormatFloat(gbValue, 'f', 2, 64) + " Multiply by --> " + strconv.FormatFloat(multiplier, 'f', 2, 64)),
		})
		if err != nil {
			fmt.Printf("Error in Conversion: %v\n", err)
		}
		time.Sleep(1 * time.Second)
	}

	return nil
}

func (*server) AddAllNumbers(stream calcpb.CalculatorService_AddAllNumbersServer) error {
	fmt.Println("********************")
	fmt.Println("In AddAllNumbers Method: Client Streaming")
	fmt.Println("********************")
	var sum float32 = 0.0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&calcpb.AddSumResponse{
				Sum: sum,
			})
		}
		if err != nil {
			fmt.Printf("Error in adding all numbers: %v\n", err)
		}
		sum += req.GetStreamNo()
	}

}

func (*server) FindMaximum(stream calcpb.CalculatorService_FindMaximumServer) error {
	fmt.Println("********************")
	fmt.Println("In FindMaximum Method: BiDirectional Streaming")
	fmt.Println("********************")

	maximum := int32(0)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			fmt.Printf("Error in recieving stream: %v\n", err)
			return nil
		}

		number := req.GetNumber()
		if number > maximum {
			maximum = number
			sendErr := stream.Send(&calcpb.FindMaximumResponse{
				Max: maximum,
			})

			if sendErr != nil {
				fmt.Printf("Error in Sending stream: %v\n", sendErr)
				return sendErr
			}
		}
		time.Sleep(500 * time.Millisecond)
	}

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
