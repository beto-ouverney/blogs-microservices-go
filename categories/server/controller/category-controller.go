package controller

import (
	"github.com/beto-ouverney/blogs-microservices/categories/proto/pb"
)

type Server struct {
	pb.UnimplementedCategoriesServiceServer
}
