package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/gowithvikash/grpc_with_go/server_streaming_api/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.GreetServiceServer
}

var (
	network = "tcp"
	address = "0.0.0.0:50051"
)

func main() {
	lis, err := net.Listen(network, address)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
func (s *Server) Greet_Many_Times(in *pb.GreetRequest, stream pb.GreetService_Greet_Many_TimesServer) error {
	fmt.Println("___ Greet_Many_Times Function Was Invoked At Client___")
	for i := 0; i < 5; i++ {
		if err := stream.Send(&pb.GreetResponse{Result: fmt.Sprintf("Hello , %s !\n", in.Name)}); err != nil {
			log.Fatal(err)
		}
	}
	return nil

}
