package usermodel

import (
	"context"

	"github.com/beto-ouverney/blogs-microservices/user/server/entity"
	"github.com/beto-ouverney/blogs-microservices/user/server/errors"
)

func (model *modelSqlx) GetByEmail(email string) (*entity.User, *errors.CustomError) {

	var user entity.User

	err := model.sqlx.GetContext(context.Background(), &user, `
		SELECT 
			id,
			displayName,
			email,
			password,
			image
		FROM Users 
		WHERE email=?
	`, email)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return nil, &errors.CustomError{Code: errors.EINTERNAL, Op: "usermodel.GetByEmail", Err: err}
	}
	return &user, nil

}
