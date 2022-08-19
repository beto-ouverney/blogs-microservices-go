package usermodel

import (
	"context"

	"github.com/beto-ouverney/blogs-microservices/user/server/errors"

	"github.com/beto-ouverney/blogs-microservices/user/server/entity"
)

func (model *modelSqlx) Add(user *entity.User) (*entity.User, *errors.CustomError) {

	result, err := model.sqlx.ExecContext(context.Background(), `
		INSERT INTO Users (
			displayName,
			email,
			password,
			image
		) VALUES (?, ?, ?, ?)
	`, user.DisplayName, user.Email, user.Password, user.Image)
	if err != nil {
		return nil, &errors.CustomError{Code: errors.EINVALID, Op: "usermodel.AddUser", Err: err}
	}
	user.ID, err = result.LastInsertId()
	if err != nil {
		return nil, &errors.CustomError{Code: errors.EINVALID, Op: "usermodel.AddUser", Err: err}
	}
	return user, nil
}
