package controller

import (
	"github.com/beto-ouverney/blogs-microservices/categories/proto/pb"
	serviceCategory "github.com/beto-ouverney/blogs-microservices/categories/server/service"
	"io"
)

func (s *Server) AddCategory(stream pb.CategoriesService_AddCategoryServer) error {

	service := serviceCategory.New()

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		newCategory, errC := service.Add(req.GetName())

		if errC != nil {
			return errC
		}

		err = stream.Send(&pb.AddCategoryResponse{Id: newCategory.ID, Name: newCategory.Name})
		if err != nil {
			return err
		}
	}
}
