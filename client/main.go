package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/gowithvikash/grpc_with_go/server_streaming_api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)
	do_Greet_Many_Times(c)

}

func do_Greet_Many_Times(c pb.GreetServiceClient) {
	fmt.Println("\n_______________ Action Number : 01 _______________")
	fmt.Println("_____  do_Greet_Many_Times Function Was Invoked At Client  ____")
	stream, err := c.Greet_Many_Times(context.Background(), &pb.GreetRequest{Name: "Vikash Parashar"})
	if err != nil {
		log.Fatal(err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("___ do_Greet_Everyone_Result: %v\n", res.Result)
	}

}
