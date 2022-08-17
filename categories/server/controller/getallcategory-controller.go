package controller

import (
	"github.com/beto-ouverney/blogs-microservices/categories/proto/pb"
	"github.com/beto-ouverney/blogs-microservices/categories/server/errors"
	"github.com/beto-ouverney/blogs-microservices/categories/server/service"
)

func (s *Server) GetAll(stream pb.CategoriesService_GetAllCategoriesServer) *errors.CustomError {
	services := service.New()
	categories, err := services.GetAll()
	if err != nil {
		return err
	}

	for _, category := range *categories {

		res := &pb.GetAllCategoriesResponse{
			Id:   category.ID,
			Name: category.Name,
		}
		stream(res)
	}
	return nil
}
