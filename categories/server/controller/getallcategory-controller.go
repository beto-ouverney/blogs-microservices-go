package controller

import (
	"github.com/beto-ouverney/blogs-microservices/categories/proto/pb"
	"github.com/beto-ouverney/blogs-microservices/categories/server/service"
)

func (s *Server) GetAllCategories(input *pb.GetAllCategoriesRequest, stream pb.CategoriesService_GetAllCategoriesServer) error {
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
		stream.Send(res)
	}
	return nil
}
