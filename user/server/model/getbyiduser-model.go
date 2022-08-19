package usermodel

import (
	"context"

	"github.com/beto-ouverney/blogs-microservices/user/server/entity"
	"github.com/beto-ouverney/blogs-microservices/user/server/errors"
)

func (model *modelSqlx) GetByID(id int64) (*entity.UserWithoutPassword, *errors.CustomError) {
	var user entity.UserWithoutPassword

	err := model.sqlx.GetContext(context.Background(), &user, `
		SELECT 
			id,
			displayName,
			email,
			image
		FROM Users 
		WHERE id=?
	`, id)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return nil, &errors.CustomError{Code: errors.EINTERNAL, Op: "usermodel.GetByEmail", Err: err}
	}
	return &user, nil

}
