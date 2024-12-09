package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrNameIsRequired        = errors.New("name is required")
	ErrUserIDIsRequired      = errors.New("user id is required")
	ErrInvalidIDParsed       = errors.New("invalid id parsed")
	ErrDescriptionIsRequired = errors.New("description is required")
)

type Meal struct {
	ID          string    `db:"id"`
	UserID      string    `db:"user_id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Date        time.Time `db:"date"`
	IsDiet      bool      `db:"is_diet"`
}

func NewMeal(userID, name, description string, date time.Time, IsDiet bool) (*Meal, error) {
	if date.IsZero() {
		date = time.Now()
	}

	meal := &Meal{
		ID:          uuid.New().String(),
		UserID:      userID,
		Name:        name,
		Description: description,
		Date:        date,
		IsDiet:      IsDiet,
	}

	err := meal.Validate()
	if err != nil {
		return nil, err
	}

	return meal, nil
}

func (m *Meal) Validate() error {
	if m.UserID == "" {
		return ErrUserIDIsRequired
	}

	_, err := uuid.Parse(m.UserID)
	if err != nil {
		return ErrInvalidIDParsed
	}

	if m.Name == "" {
		return ErrNameIsRequired
	}

	if m.Description == "" {
		return ErrDescriptionIsRequired
	}

	return nil
}
