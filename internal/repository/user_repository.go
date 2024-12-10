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
