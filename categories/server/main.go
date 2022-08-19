package main

import (
	"github.com/beto-ouverney/blogs-microservices/categories/proto/pb"
	"github.com/beto-ouverney/blogs-microservices/categories/server/config"
	"github.com/beto-ouverney/blogs-microservices/categories/server/controller"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	lis, err := net.Listen("tcp", config.PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	cat := controller.Server{}

	grpcServer := grpc.NewServer()

	pb.RegisterCategoriesServiceServer(grpcServer, &cat)

	log.Println("Listening on Port" + config.PORT)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
