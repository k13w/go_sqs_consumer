package repository

import "database/sql"

type UserRepositoryPsql struct {
	DB *sql.DB
}

func CreateUser(db *sql.DB) *UserRepositoryPsql {
	return &UserRepositoryPsql{DB: db}
}
