package model

import (
	"context"
	"github.com/beto-ouverney/blogs-microservices/categories/server/entity"
	"github.com/beto-ouverney/blogs-microservices/categories/server/errors"
)

func (m *modelSqlx) GetByName(name string) (*entity.Category, *errors.CustomError) {
	var category entity.Category

	err := m.sqlx.GetContext(context.Background(), &category, `SELECT id AS "categories.id" , name AS "categories.name" FROM Categories WHERE name = ?`, name)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return nil, &errors.CustomError{Code: errors.EINTERNAL, Op: "categorymodel.GetByNameCategory", Err: err}
	}
	return &category, nil
}
