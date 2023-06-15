package users

import (
	"Weather/internal/models"
	"database/sql"
)

type UserRepo interface {
	AddUser(cookieValue string) error
	GetUser(cookieValue string) (models.User, error)
}

type userRepo struct {
	db *sql.DB
}

func Repo(db *sql.DB) UserRepo {
	return userRepo{db: db}
}

func (r userRepo) AddUser(cookieValue string) error {
	_, err := r.db.Exec("insert into users(id) values (?)", cookieValue)

	return err
}

func (r userRepo) GetUser(cookieValue string) (models.User, error) {
	row := r.db.QueryRow("select * from users where id = ?", cookieValue)
	if row.Err() != nil {
		return models.User{}, row.Err()
	}

	var user models.User
	err := row.Scan(&user.ID, &user.Pined.Name, &user.Pined.Lat, &user.Pined.Lon, &user.Pined.Status)
	if row.Err() != nil {
		return models.User{}, err
	}

	return user, err
}
