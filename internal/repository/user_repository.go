package repository

import (
	"database/sql"
	"savioafs/daily-diet-app-go/internal/entity"
)

type UserRepositoryPG struct {
	DB *sql.DB
}

func NewUserRepositoryPG(db *sql.DB) *UserRepositoryPG {
	return &UserRepositoryPG{DB: db}
}

func (r *UserRepositoryPG) Create(user *entity.User) error {
	stmt, err := r.DB.Prepare("INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	err = stmt.QueryRow(user.ID, user.Name, user.Email, user.Password).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepositoryPG) FindByEmail(email string) (*entity.User, error) {
	stmt, err := r.DB.Prepare("SELECT id, name, email, password FROM users WHERE email = $1")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var user entity.User

	err = stmt.QueryRow(email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
