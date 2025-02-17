package service

import (
	"people-matching-api/model"
	"people-matching-api/repository"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (userService *UserService) GetUser() ([]model.UserResponse, error) {
	user, err := userService.UserRepository.GetUser()
	if err != nil {
		return nil, err
	}

	userResponses := model.UsersToResponses(*user)

	return userResponses, nil
}

func (userService *UserService) RegisterUser(userRequest model.UserRequest) error {
	user := model.User{
		FacebookID: userRequest.FacebookID,
		Name:       userRequest.Name,
		Age:        userRequest.Age,
		Friends:    userRequest.Friends,
	}

	err := userService.UserRepository.RegisterUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (userService *UserService) MatchUser(age int, friends string) (*model.MatchUserResponse, error) {
	matchUser, err := userService.UserRepository.MatchUser(age, friends)
	if err != nil {
		return nil, err
	}

	result := model.MatchUserResponse{
		IsMatch: matchUser,
	}

	return &result, nil
}
