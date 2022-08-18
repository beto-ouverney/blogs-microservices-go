package service

import (
	"github.com/beto-ouverney/blogs-microservices/categories/server/entity"
	"github.com/beto-ouverney/blogs-microservices/categories/server/errors"
)

func (c *CategoryService) Add(name string) (*entity.Category, *errors.CustomError) {
	categoryExist, err := c.Model.GetByName(name)

	if categoryExist != nil {
		return nil, &errors.CustomError{Code: errors.ECONFLICT, Op: "categorycategory.AddCategory", Err: err}
	}

	newC := New()

	newCategory, err := newC.Model.Add(name)
	if err == nil {
		return newCategory, nil
	}

	return nil, err
}
