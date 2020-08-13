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
// to avoid calling domain.UserDao itself, we need to mimic domain.UserDao.Getuser() behavior in users_service_test.go
func (s *userService) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userID) // return user, error from this function

}
