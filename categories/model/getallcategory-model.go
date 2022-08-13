package model

import (
	"context"
	"github.com/beto-ouverney/blogs-microservices/categories/entity"
	"github.com/beto-ouverney/blogs-microservices/categories/errors"
)

func (model *modelSqlx) GetAll() (*[]entity.Category, *errors.CustomError) {

	var categories []entity.Category

	err := model.sqlx.SelectContext(context.Background(), &categories, `SELECT id AS "categories.id",name AS "categories.name" FROM Categories`)

	if err != nil {
		return nil, &errors.CustomError{Code: errors.EINTERNAL, Op: "categorymodel.GetAllCategories", Err: err}
	}

	defer model.sqlx.Close()

	return &categories, nil
}
