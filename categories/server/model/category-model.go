package model

import (
	"github.com/beto-ouverney/blogs-microservices/categories/server/config"
	"github.com/beto-ouverney/blogs-microservices/categories/server/entity"
	"github.com/beto-ouverney/blogs-microservices/categories/server/errors"
	"github.com/jmoiron/sqlx"
)

type ICategoryModel interface {
	Add(name string) (*entity.Category, *errors.CustomError)
	GetAll() (*[]entity.Category, *errors.CustomError)
	GetByName(name string) (*entity.Category, *errors.CustomError)
}

type modelSqlx struct {
	sqlx *sqlx.DB
}

func New() ICategoryModel {
	return &modelSqlx{config.GetSqlx()}
}
