package services

import (
	"george_custom/golang_microservice/introduction/mvc/domain"
	"george_custom/golang_microservice/introduction/mvc/utils"
)

// UserService type and use struct as the underlying type
type userService struct{}

var (
	UserService userService
)

// GetUser is the method defined under UserService type
func (s *userService) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userID) // return user, error from this function

}
