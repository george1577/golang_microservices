package services

import (
	"fmt"
	"george_custom/golang_microservice/introduction/mvc/domain"
	"george_custom/golang_microservice/introduction/mvc/utils"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	userDaoMock     usersDaoMock
	getUserFunction func(int64) (*domain.User, *utils.ApplicationError)
)

func init() {
	// domain.UserDao is an interface
	domain.UserDao = &usersDaoMock{}
	// share the same interface defined in user_dao
	// need to assign a pointer since GetUser() method was declared with pointer reciever func (u *usersDaoMock) GetUser()...
	// since domain.UserDao variable takes the value of usersDaoMock, when calling GetUser(), it will call the one
	// implemented here instead of GetUser() in user_dao.go, thus we can avoid access the database directly
}

type usersDaoMock struct{}

// copy from user_dao, need to mimic GetUser() behavior
// use another getUserFunction() to provide flexibility of how we want the function to return based on different cases
func (u *usersDaoMock) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userID)
}

// this test is for testing in service package only, isolated from domain package
// we do not want to call the domain database itself
func TestGetUserNotFoundInDatabase(t *testing.T) {
	// customize the return value for getUserFunction
	getUserFunction = func(userID int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("user %v was not found", userID),
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}

	// mock domain.UserDao layer
	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 was not found", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	// customize the return value for getUserFunction
	getUserFunction = func(userID int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{ID: 123, FirstName: "George", LastName: "Chen", Email: "george1577@yahoo.com.tw"}, nil
	}

	user, err := UserService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.ID)
	assert.EqualValues(t, "George", user.FirstName)
	assert.EqualValues(t, "Chen", user.LastName)
	assert.EqualValues(t, "george1577@yahoo.com.tw", user.Email)
}

// no need to test the userID not integer case, since when request reaches the user service, it should be integer, the filter of integer happens
// in user_controller

/*
   artifact and mocking layer
1. During the user_service testing, we need to call GetUser() and return domain.UserDao.GetUser(userID), which will further call
   GetUser() in the domain package(which touch the database), this is not the behavior we want since we're testing the service package
   only we should isolate it.

2. Thus we need to mock the domain.UserDao layer in the users_service_test.go, the method to do it is using an interface, thus we can
   declare variable separately in both user_dao.go and users_service_test.go, by doing so, we can have separate behaviors.

3. We first declare an userServiceInterface, which requires GetUser() method in user_dao and declare a public variable UserService
   so other files can also access it(first letter capital), since we can't call GetUser() directly using interface variable, we need
   to assign a value from another type, which could be any type that satisfy the interface requirement(in other words, have GetUser()
   method implemented). This is why we have UserDao = &userDao{} in user_dao and domain.UserDao = &usersDaoMock{}, both userDao
   and usersDaoMock type has GetUser() method implemented, the reason we need a pointer & here is because we use pointer receiver
   when implementing GetUser() method. We could put the value assignment in func init() so it will be executed first

4. In users_service_test.go, in GetUser() method we return another function(which also returns (*domain.User, *utils.ApplicationError) as
   we require), this provides a flexibility when we want to test different cases, otherwise if we would return a fixed value, there is no
   way that we can change it for different cases

5. After finishing 1-4, during the test, users_service_test.go will call GetUser() function in user_service, which will then
   call domain.UserDao.GetUser(), since we have initialize the domain.UserDao value by domain.UserDao = &usersDaoMock{}, the GetUser()
   we call here will be from the one we implemented for usersDaoMock type, not the one in user_dao, thus avoiding accessing the database
   directly, without domain.UserDao = &usersDaoMock{}, it will call the one implemented for userDao type in user_dao
*/
