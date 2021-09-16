package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	sdk "agones.dev/agones/sdks/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "test/helloworld"
)

type controller struct{}

func NewController() pb.GreeterServer {
	return controller{}
}

func (controller) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	name := req.GetName()
	return &pb.HelloReply{Message: fmt.Sprintf("Hello %s!", name)}, nil
}

func main() {
	port := flag.String("port", "7654", "The port to listen to traffic on")
	flag.Parse()
	if ep := os.Getenv("PORT"); ep != "" {
		port = &ep
	}
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", *port))
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Printf("Listening :%s", *port)

	s, err := sdk.NewSDK()
	if err != nil {
		log.Fatalf(err.Error())
	}

	if err = s.Ready(); err != nil {
		log.Fatalf(err.Error())
	}

	go func() {
		for range time.Tick(time.Second) {
			if err = s.Health(); err != nil {
				log.Fatalf(err.Error())
			}
		}
	}()

	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, NewController())

	reflection.Register(grpcServer)
	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalf(err.Error())
	}
}
