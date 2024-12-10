package usecase

import (
	"errors"
	"fmt"
	"savioafs/daily-diet-app-go/internal/dto"
	"savioafs/daily-diet-app-go/internal/entity"
	"savioafs/daily-diet-app-go/internal/repository"
)

type MealUsecase struct {
	repository repository.MealStorer
}

func NewMealUseCase(repo repository.MealStorer) MealUsecase {
	return MealUsecase{repository: repo}
}

func (u *MealUsecase) Create(meal *entity.Meal) (*entity.Meal, error) {
	mealID, err := u.repository.Create(meal)
	if err != nil {
		return nil, err
	}

	meal.ID = mealID

	return meal, nil
}

func (u *MealUsecase) FindMealByID(id string) (*entity.Meal, error) {
	if id == "" {
		return nil, errors.New("meal id cannot empty")
	}

	meal, err := u.repository.GetMealByID(id)
	if err != nil {
		return nil, err
	}

	if meal == nil {
		return nil, errors.New("meal not found")
	}

	return meal, nil
}

func (u *MealUsecase) GetAllMealsByUser(userID string) ([]entity.Meal, error) {
	if userID == "" {
		return nil, errors.New("user id cannot empty")
	}

	meals, err := u.repository.GetAllMealsByUser(userID)
	if err != nil {
		return nil, err
	}

	if meals == nil {
		return nil, errors.New("meals by id not found")
	}

	return meals, nil
}

func (u *MealUsecase) GetMealsUserByStatus(userID string, status bool) ([]entity.Meal, error) {
	if userID == "" {
		return nil, errors.New("user id cannot empty")
	}

	var mealsWithStatusSelected []entity.Meal

	meals, err := u.repository.GetAllMealsByUser(userID)
	if err != nil {
		return nil, err
	}

	if meals == nil {
		return nil, errors.New("meals by user not found")
	}

	for _, meal := range meals {
		if meal.IsDiet == status {
			mealsWithStatusSelected = append(mealsWithStatusSelected, meal)
		}
	}

	if mealsWithStatusSelected == nil {
		return nil, errors.New("meals by user with current status not found")
	}

	return mealsWithStatusSelected, nil

}

func (u *MealUsecase) MetricsMealsByUser(userID string) (dto.MetricsOutputDTO, error) {
	if userID == "" {
		return dto.MetricsOutputDTO{}, errors.New("user id cannot empty")
	}

	meals, err := u.repository.GetAllMealsByUser(userID)
	if err != nil {
		return dto.MetricsOutputDTO{}, err
	}

	if meals == nil {
		return dto.MetricsOutputDTO{}, errors.New("meals by user not found")
	}

	var (
		dietMealsCount, nonDietMealsCount int
	)

	for _, meal := range meals {
		if meal.IsDiet {
			dietMealsCount++
		} else {
			nonDietMealsCount++
		}
	}

	totalMeals := (len(meals))
	dietPercent := float32(dietMealsCount) / float32(totalMeals) * 100
	nonDietPercent := float32(nonDietMealsCount) / float32(totalMeals) * 100

	metricsOutput := dto.MetricsOutputDTO{
		TotalMeals:        totalMeals,
		TotalMealsDiet:    dietMealsCount,
		TotalMealsNonDiet: nonDietMealsCount,
		DietPercent:       float64(dietPercent),
		NonDietPercent:    float64(nonDietPercent),
	}

	return metricsOutput, nil

}

func (u *MealUsecase) UpdateMeal(id string, meal *entity.Meal) error {
	if id == "" {
		return errors.New(" id cannot empty")
	}

	if meal == nil {
		return errors.New("meal details cannot be nil")
	}

	err := u.repository.UpdateMeal(id, meal)
	if err != nil {
		return err
	}

	return nil
}

func (u *MealUsecase) DeleteMeal(id string) error {
	if id == "" {
		return errors.New("meal id cannot empty")
	}

	err := u.repository.DeleteMeal(id)
	if err != nil {
		return fmt.Errorf("failed to delete meal: %w", err)
	}

	return nil
}
