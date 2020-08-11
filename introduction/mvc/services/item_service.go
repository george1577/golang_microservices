package services

import (
	"george_custom/golang_microservice/introduction/mvc/domain"
	"george_custom/golang_microservice/introduction/mvc/utils"
)

type itemService struct{}

var (
	ItemService itemService
)

// define method for itemService type
func (s *itemService) GetItem(itemID string) (*domain.Item, utils.ApplicationError) {
	return nil, utils.ApplicationError{}
}
