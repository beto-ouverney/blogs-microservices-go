package model

import (
	"context"
	"github.com/beto-ouverney/blogs-microservices/categories/server/entity"
	"github.com/beto-ouverney/blogs-microservices/categories/server/errors"
)

func (model *modelSqlx) Add(name string) (*entity.Category, *errors.CustomError) {
	result, err := model.sqlx.ExecContext(context.Background(), "INSERT INTO Categories (name) VALUES (?)", name)
	if err != nil {

		return nil, &errors.CustomError{Code: errors.EINTERNAL, Op: "categorymodel.AddCategory", Err: err}
	}

	defer model.sqlx.Close()

	newCategory := &entity.Category{
		Name: name,
	}
	newCategory.ID, err = result.LastInsertId()
	if err != nil {
		return nil, &errors.CustomError{Code: errors.EINVALID, Op: "usermodel.AddUser", Err: err}
	}

	return newCategory, nil
}
