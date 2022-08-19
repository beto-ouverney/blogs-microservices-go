package service

import (
	"github.com/beto-ouverney/blogs-microservices/categories/server/entity"
	"github.com/beto-ouverney/blogs-microservices/categories/server/errors"
	"github.com/beto-ouverney/blogs-microservices/categories/server/model"
)

type ICategoryService interface {
	GetAll() (*[]entity.Category, *errors.CustomError)
	Add(name string) (*entity.Category, *errors.CustomError)
}

type CategoryService struct {
	Model model.ICategoryModel
}

func New() *CategoryService {
	return &CategoryService{
		Model: model.New(),
	}
}
