package model

import (
	"github.com/beto-ouverney/blogs-microservices/categories/config"
	"github.com/beto-ouverney/blogs-microservices/categories/entity"
	"github.com/beto-ouverney/blogs-microservices/categories/errors"
	"github.com/jmoiron/sqlx"
)

type ICategoryModel interface {
	Add(category *entity.Category) (*entity.Category, *errors.CustomError)
	GetAll() (*[]entity.Category, *errors.CustomError)
	GetByName(name string) (*entity.Category, *errors.CustomError)
}

type modelSqlx struct {
	sqlx *sqlx.DB
}

func New() ICategoryModel {
	return &modelSqlx{config.GetSqlx()}
}
