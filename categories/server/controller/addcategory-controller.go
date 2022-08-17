package controller

import (
	"github.com/beto-ouverney/blogs-microservices/categories/proto/pb"
	serviceCategory "github.com/beto-ouverney/blogs-microservices/categories/server/service"
	"io"
)

func (s *Server) Add(stream pb.CategoriesService_AddCategoryServer) error {

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		service := serviceCategory.New()

		newCategory, err := service.Add(req.GetName())

		if err != nil {
			return err
		}
		stream.Send(&pb.AddCategoryResponse{Id: newCategory.ID, Name: newCategory.Name})

	}
}
