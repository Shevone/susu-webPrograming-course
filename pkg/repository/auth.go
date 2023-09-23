package repository

import (
	"database/sql"
	"fmt"
	"web-programing-susu/pkg/models"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	sqlText := fmt.Sprintf("INSERT INTO" + " " + usersTable + " (login, password, name, email, admin) values ($1,$2,$3,$4,$5) RETURNING id")
	row := r.db.QueryRow(sqlText, user.Login, user.Password, user.Name, user.EMail, user.Role)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
