package repository

import (
	"Weather/internal/repository/users"
	"database/sql"
)

type Repo struct {
	UserSrv users.UserRepo
}

func New(db *sql.DB) *Repo {
	return &Repo{
		UserSrv: users.Repo(db),
	}
}
