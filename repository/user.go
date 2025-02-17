package repository

import (
	"people-matching-api/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (userRepository *UserRepository) GetUser() (*[]model.User, error) {
	var users []model.User
	err := userRepository.db.Model(&model.User{}).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (userRepository *UserRepository) RegisterUser(user model.User) error {
	err := userRepository.db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (userRepository *UserRepository) MatchUser(age int, friends string) (bool, error) {
	var user model.User

	err := userRepository.db.Where("age = ? AND friends = ?", age, friends).First(&user).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
