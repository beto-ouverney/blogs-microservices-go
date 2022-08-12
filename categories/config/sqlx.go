package config

import (
	"github.com/jmoiron/sqlx"
	"os"
)

func GetSqlx() *sqlx.DB {
	return sqlx.MustConnect("mysql", os.Getenv("DB_CONNECTION"))
}
