package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"

	calcpb "github.com/gRPC-basics/calc_pb"
	errorpb "github.com/gRPC-basics/error_pb"
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
	errClient := errorpb.NewMiscServiceClient(conn)
	doUnary(client)
	doServerStreaming(client)
	doClientStreaming(client)
	doBiDiStreaming(client)
	doDivideErrorHandling(errClient)
	doPingWithDeadline(errClient, 5*time.Second)
	doPingWithDeadline(errClient, 2*time.Second)
}

func doUnary(c calcpb.CalculatorServiceClient) {
	fmt.Println(`*********************`)
	fmt.Println(`Unary`)
	fmt.Println(`*********************`)
	req := &calcpb.SumRequest{
		Number1: 20,
		Number2: 10,
	}

	resp, err := c.DoSum(context.Background(), req)
	if err != nil {
		fmt.Println("Error in DoSum ", err)
	}
	log.Printf("Output of DoSum: %v", resp.GetSum())

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
	log.Printf("output of Multiply: %v", mul.GetResp())

}

func doServerStreaming(c calcpb.CalculatorServiceClient) {
	fmt.Println(`*********************`)
	fmt.Println(`Server Streaming`)
	fmt.Println(`*********************`)
	req := &calcpb.GigaByteRequest{
		GbValue: "2",
	}

	stream, err := c.GetByteValues(context.Background(), req)
	if err != nil {
		fmt.Printf("Server Stream Error %v\n", err)
	}

	for {
		val, err := stream.Recv()
		if err == io.EOF {
			fmt.Println(`Calculation complete for Server Streaming`)
			break
		}
		if err != nil {
			fmt.Printf("Stream Error %v\n", err)
		}
		fmt.Println(`Receiving Values from Server:  `, val.GetResp())
	}

}

func doClientStreaming(c calcpb.CalculatorServiceClient) {
	fmt.Println(`*********************`)
	fmt.Println(`Client Streaming`)
	fmt.Println(`*********************`)

	stream, err := c.AddAllNumbers(context.Background())
	if err != nil {
		fmt.Printf("Client Stream Error %v\n", err)
	}

	numbers := make([]float32, 0)
	for i := 0; i < 5; i++ {
		numbers = append(numbers, float32(i))
	}
	for _, number := range numbers {
		stream.Send(&calcpb.AddSumRequest{
			StreamNo: number,
		})
		fmt.Println("Req: Stream Value: ", number)
		time.Sleep(1 * time.Second)
	}
	val, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Printf("Stream Error %v\n", err)
	}
	fmt.Println(`Calculated Sum from Server:  `, val.GetSum())

}

func doBiDiStreaming(c calcpb.CalculatorServiceClient) {
	fmt.Println(`*********************`)
	fmt.Println(`Client Streaming`)
	fmt.Println(`*********************`)

	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		fmt.Printf("BiDi Stream Error %v\n", err)
	}

	// create wait channel
	waitCh := make(chan struct{})

	// send go routines
	go func() {
		numbers := []int32{2, 7, 7, 4, 9, 19, 5}
		for _, number := range numbers {
			fmt.Println("Sending Value: ", number)
			stream.Send(&calcpb.FindMaximumRequest{
				Number: number,
			})
			time.Sleep(500 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	// receive go routines
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				fmt.Printf("Error in Receiving Values %v \n", err)
				break
			}

			fmt.Println("Received New Maximum Value: ", resp.GetMax())

		}
		close(waitCh)
	}()

	<-waitCh

}

func doDivideErrorHandling(client errorpb.MiscServiceClient) {
	fmt.Println(`*********************`)
	fmt.Println(`Divide by Zero : Error Handling`)
	fmt.Println(`*********************`)

	resp, err := client.DivideWithError(context.Background(), &errorpb.DivideRequest{
		Number1: 10,
		Number:  0,
	})
	if err != nil {

		status, ok := status.FromError(err)
		if ok {
			fmt.Println("Error message from server: ", status.Message())
			fmt.Println("Error code from server: ", status.Code())
		} else {
			fmt.Printf("Error in DivideErrorHandling: %v \n", err)
		}
		return
	}

	fmt.Println("Output of DivideErrorHandling: ", resp.GetAnswer())

}

func doPingWithDeadline(client errorpb.MiscServiceClient, timeout time.Duration) {
	fmt.Println(`*********************`)
	fmt.Printf("Ping With Deadline : Error Handling with timeout: %v\n", timeout)
	fmt.Println(`*********************`)

	req := &errorpb.PingRequest{
		Word: "ping.....",
	}

	ctx, cancelFun := context.WithTimeout(context.Background(), timeout)
	defer cancelFun()

	resp, err := client.Ping(ctx, req)
	if err != nil {
		status, ok := status.FromError(err)
		if ok {
			if status.Code() == codes.DeadlineExceeded {
				fmt.Println("Message: ", status.Message())
			} else {
				fmt.Println("Error message from server: ", status.Message())
				fmt.Println("Error code from server: ", status.Code())
			}
		} else {
			fmt.Printf("Error in DivideErrorHandling: %v \n", err)
		}
		return
	}
	fmt.Println("Output of doPingWithDeadline: ", resp.GetResponse())
}
