package usermodel

import (
	"github.com/beto-ouverney/blogs-microservices/user/server/entity"
	"github.com/beto-ouverney/blogs-microservices/user/server/errors"

	"github.com/jmoiron/sqlx"
)

type IUserModel interface {
	GetByEmail(email string) (*entity.User, *errors.CustomError)
	Add(user *entity.User) (*entity.User, *errors.CustomError)
	GetAll() ([]entity.UserWithoutPassword, *errors.CustomError)
	GetByID(id int64) (*entity.UserWithoutPassword, *errors.CustomError)
	Delete(id int64) *errors.CustomError
}

type modelSqlx struct {
	sqlx *sqlx.DB
}

func NewSqlxModel(sqlx *sqlx.DB) IUserModel {
	return &modelSqlx{sqlx}
}
