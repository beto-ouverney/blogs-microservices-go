package service_test

import (
	"github.com/beto-ouverney/blogs-microservices/categories/server/entity"
	"github.com/beto-ouverney/blogs-microservices/categories/server/service/mock"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestCategoryService_GetAll(t *testing.T) {
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

	controller := gomock.NewController(t)

	m := mock.NewMockICategoryService(controller)
	m.EXPECT().GetAll().Return(allCategories, nil).AnyTimes()

}
