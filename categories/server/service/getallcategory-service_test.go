package service_test

import (
	mocks "github.com/beto-ouverney/blogs-microservices/categories/mocks/server/model"
	"github.com/beto-ouverney/blogs-microservices/categories/server/entity"
	"github.com/beto-ouverney/blogs-microservices/categories/server/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategoryServiceGetAllCategories(t *testing.T) {
	assertions := assert.New(t)
	t.Log("Testing if GetAll returns all categories")
	allCategories := &[]entity.Category{
		{
			ID:   1,
			Name: "Category 1",
		},
		{
			ID:   2,
			Name: "Category 2",
		},
	}

	m := new(mocks.ICategoryModel)
	m.On("GetAll").Return(allCategories, nil)

	s := service.CategoryService{Model: m}

	res, err := s.GetAll()

	assertions.Nil(err, "Error should be nil")
	assertions.Equal(res, allCategories, "Add should return new category")

}
