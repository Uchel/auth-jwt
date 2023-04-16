package jwt_auth

import (
	"database/sql"
	"log"
)

type UserRepo interface {
	GetByUsername(username string) (string, string)
}

type userRepo struct {
	db *sql.DB
}

func (u *userRepo) GetByUsername(username string) (string, string) {
	var user User

	query := "SELECT username, password FROM users WHERE username=$1"

	row := u.db.QueryRow(query, username)

	if err := row.Scan(&user.Username, &user.Password); err != nil {
		log.Println(err)
	}

	if user.Username == "" {
		return "user not found", "password uncorrect"
	}

	return user.Username, user.Password
}
func NewUserRepo(db *sql.DB) UserRepo {
	repo := new(userRepo)

	repo.db = db

	return repo
}
