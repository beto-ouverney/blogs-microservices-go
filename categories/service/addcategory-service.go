package service

import (
	"github.com/beto-ouverney/blogs-microservices/categories/entity"
	"github.com/beto-ouverney/blogs-microservices/categories/errors"
)

func (c *CategoryService) Add(category *entity.Category) (*entity.Category, *errors.CustomError) {
	categoryExist, err := c.Model.GetByName(category.Name)
	if categoryExist != nil {
		return nil, &errors.CustomError{Code: errors.ECONFLICT, Op: "categorycategory.AddCategory", Err: err}
	}

	newCategory, err := c.Model.Add(category)
	if err == nil {
		return newCategory, nil
	}

	return nil, err
}
