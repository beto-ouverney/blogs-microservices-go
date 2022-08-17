package service

import (
	"github.com/beto-ouverney/blogs-microservices/categories/server/entity"
	"github.com/beto-ouverney/blogs-microservices/categories/server/errors"
)

func (u *CategoryService) GetAll() (*[]entity.Category, *errors.CustomError) {
	categories, err := u.Model.GetAll()
	if err == nil {
		return categories, nil
	}

	return nil, err
}
