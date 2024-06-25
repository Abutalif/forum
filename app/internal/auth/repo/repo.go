package repo

import (
	"database/sql"
	"fmt"
	"forum/app/entities"
	"forum/app/internal/auth"
)

type authRepo struct {
	db *sql.DB
}

func NewAuthRepo() auth.AuthRepo {
	return &authRepo{}
}

func (a *authRepo) GetUser(name string) entities.User {
	var user entities.User
	query := fmt.Sprintf(`SELECT name, password FROM users WHERE name=%v;`, name)

	row := a.db.QueryRow(query)
	err := row.Scan(&user.Name, &user.Password)
	if err == sql.ErrNoRows {
		return entities.User{}
	}
	// TODO: handle other errors
	return user
}
