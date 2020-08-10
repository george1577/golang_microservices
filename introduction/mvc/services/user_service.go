package services

import (
	"george_custom/golang_microservice/introduction/mvc/domain"
	"george_custom/golang_microservice/introduction/mvc/utils"
)

func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID) // return user, error from this function

}
