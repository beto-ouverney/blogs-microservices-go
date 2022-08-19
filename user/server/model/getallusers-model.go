package usermodel

import (
	"context"

	"github.com/beto-ouverney/blogs-microservices/user/server/entity"
	"github.com/beto-ouverney/blogs-microservices/user/server/errors"
)

func (model *modelSqlx) GetAll() ([]entity.UserWithoutPassword, *errors.CustomError) {
	var users []entity.UserWithoutPassword

	err := model.sqlx.SelectContext(context.Background(), &users, "SELECT id, displayName, email, image FROM Users")

	if err != nil {
		return nil, &errors.CustomError{Code: errors.EINTERNAL, Op: "usermodel.GetAllUsers", Err: err}
	}
	return users, nil
}
