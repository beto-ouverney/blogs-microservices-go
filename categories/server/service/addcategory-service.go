package service

import (
	"errors"
	"github.com/beto-ouverney/blogs-microservices/categories/server/entity"
	. "github.com/beto-ouverney/blogs-microservices/categories/server/errors"
)

func (c *CategoryService) Add(name string) (*entity.Category, *CustomError) {

	if name == "" {
		return nil, &CustomError{Code: EINVALID, Op: "category-service.Add", Err: errors.New(ErrorResponse["missingFields"].Message)}
	}
	categoryExist, err := c.Model.GetByName(name)

	if categoryExist != nil {
		return nil, &CustomError{Code: ECONFLICT, Op: "category-service.AddCategory", Err: errors.New("category already exists")}
	}

	newCategory, err := c.Model.Add(name)
	if err == nil {
		return newCategory, nil
	}

	return nil, err
}
