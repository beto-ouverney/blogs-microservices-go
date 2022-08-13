package service

import (
	"github.com/beto-ouverney/blogs-microservices/categories/entity"
	"github.com/beto-ouverney/blogs-microservices/categories/errors"
	"github.com/beto-ouverney/blogs-microservices/categories/model"
)

type ICategoryService interface {
	GetAll() (*[]entity.Category, *errors.CustomError)
	Add(category *entity.Category) (*entity.Category, *errors.CustomError)
}

type CategoryService struct {
	Model model.ICategoryModel
}

func New() *CategoryService {
	return &CategoryService{
		Model: model.New(),
	}
}
