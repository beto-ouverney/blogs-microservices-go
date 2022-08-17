package main

import (
	"github.com/beto-ouverney/blogs-microservices/categories/proto/pb"
	"github.com/beto-ouverney/blogs-microservices/categories/server/controller"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {

	lis, err := net.Listen("tcp", os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	cat := controller.Server{}

	grpcServer := grpc.NewServer()

	pb.RegisterCategoriesServiceServer(grpcServer, &cat)

	log.Println("Listening on Port:" + os.Getenv("PORT"))

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
