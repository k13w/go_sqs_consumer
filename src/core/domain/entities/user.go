package entities

import "database/sql"

type User struct {
	Username string
	Level    int
}

type UserEntity struct {
	DbConnect *sql.DB
}

func (user UserEntity) CreateUser(username string, level int) error {
	createUserSql := `INSERT INTO "user" (username, level, id) VALUES ($1, $2, DEFAULT)`

	_, err := user.DbConnect.Exec(createUserSql, username, level)
	if err != nil {
		return err
	}

	return nil
}
