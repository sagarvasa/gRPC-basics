package main

import (
	"context"
	"fmt"
	"log"

	calcpb "github.com/gRPC-basics/calc_pb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client Main Function")

	conn, err := grpc.Dial("localhost:50001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to dial %v", err)
	}

	defer conn.Close()

	client := calcpb.NewCalculatorServiceClient(conn)
	doUnary(client)
}

func doUnary(c calcpb.CalculatorServiceClient) {
	req := &calcpb.SumRequest{
		Number1: 20,
		Number2: 10,
	}

	resp, err := c.DoSum(context.Background(), req)
	if err != nil {
		fmt.Println("Error in DoSum ", err)
	}
	log.Printf("Output of DoSum: %v", resp.Sum)

	/* Position of field in struct can be modified */
	DiffResp, err := c.CalcDiff(context.Background(), &calcpb.DiffRequest{
		Number2: 10,
		Number1: 20,
	})
	if err != nil {
		fmt.Println("Error in CalcDiff ", err)
	}
	log.Printf("output of calcDiff: %v", DiffResp.Difference)

	mul, err := c.Multiply(context.Background(), &calcpb.MultiplierRequest{
		Number3: 20,
		Number2: 10,
		Number1: 10,
	})
	if err != nil {
		fmt.Println("Error in Multiply ", err)
	}
	log.Printf("output of Multiply: %v", mul.Resp)

}
