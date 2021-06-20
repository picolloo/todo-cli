package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const PostgresDriver = "postgres"
const User = "docker"
const Host = "todo-cli-db"
const Password = "docker"
const DbName = "todo-cli"

func NewDBConnection() *sql.DB {
	dataSourceName := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		Host, 5432, User, Password, DbName)

	db, err := sql.Open(PostgresDriver, dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
