package service_test

import (
	"errors"
	"testing"

	mocks "github.com/beto-ouverney/blogs-microservices/categories/mocks/server/model"
	"github.com/beto-ouverney/blogs-microservices/categories/server/entity"
	. "github.com/beto-ouverney/blogs-microservices/categories/server/errors"
	"github.com/beto-ouverney/blogs-microservices/categories/server/service"
	"github.com/stretchr/testify/assert"
)

func TestAddCategoryCanAddNewCategory(t *testing.T) {
	t.Log("Testing if AddCategory can add a new category")
	assertions := assert.New(t)

	name := "Go Lang"

	newCategory := &entity.Category{
		ID:   2,
		Name: "Go Lang",
	}

	m := new(mocks.ICategoryModel)
	m.On("GetByName", name).Return(nil, nil)
	m.On("Add", name).Return(&entity.Category{ID: 2, Name: "Go Lang"}, nil)

	s := service.CategoryService{Model: m}

	res, err := s.Add(name)

	assertions.Nil(err, "Error should be nil")
	assertions.Equal(res, newCategory, "Add should return new category")

}

func TestAddCategoryNotAddNewCategoryWithoutNameField(t *testing.T) {
	t.Log("Testing if AddCategory can`t add a new category without name field")
	assertions := assert.New(t)

	name := ""

	m := new(mocks.ICategoryModel)

	s := service.CategoryService{Model: m}

	res, err := s.Add(name)

	assertions.Nil(res, "new category response should be nil")
	assertions.Equal(err, &CustomError{Code: EINVALID, Op: "category-service.Add", Err: errors.New(ErrorResponse["missingFields"].Message)}, "Add should return error")

}

func TestAddCategoryNotAddCategoryAlreadtExists(t *testing.T) {
	t.Log("Testing if AddCategory can`t add a new category if it already exists")
	assertions := assert.New(t)

	name := "Go Lang"

	category := &entity.Category{
		ID:   2,
		Name: "Go Lang",
	}

	m := new(mocks.ICategoryModel)
	m.On("GetByName", name).Return(category, nil)

	s := service.CategoryService{Model: m}

	res, err := s.Add(name)

	assertions.Nil(res, "new category response should be nil")
	assertions.Equal(err, &CustomError{Code: ECONFLICT, Op: "category-service.AddCategory", Err: errors.New("category already exists")}, "Add should return error")

}
