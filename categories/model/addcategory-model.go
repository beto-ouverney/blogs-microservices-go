package model

import (
	"context"
	"github.com/beto-ouverney/blogs-microservices/categories/entity"
	"github.com/beto-ouverney/blogs-microservices/categories/errors"
)

func (model *modelSqlx) Add(category *entity.Category) (*entity.Category, *errors.CustomError) {
	result, err := model.sqlx.ExecContext(context.Background(), "INSERT INTO Categories (name) VALUES (?)", category.Name)
	if err != nil {

		return nil, &errors.CustomError{Code: errors.EINTERNAL, Op: "categorymodel.AddCategory", Err: err}
	}

	defer model.sqlx.Close()

	category.ID, err = result.LastInsertId()
	if err != nil {
		return nil, &errors.CustomError{Code: errors.EINVALID, Op: "usermodel.AddUser", Err: err}
	}

	return category, nil
}
