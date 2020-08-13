package domain

import (
	"fmt"
	"george_custom/golang_microservice/introduction/mvc/utils"
	"net/http"
)

// define an interface that has GetUser() method implemented
type userServiceInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

// mimicking the database
var (
	// take the pointer of the User, *User can be nil but not User object itself, this is more convenient if User is not found in database
	users = map[int64]*User{
		123: {ID: 123, FirstName: "George", LastName: "Chen", Email: "george1577@yahoo.com.tw"},
		165: {ID: 165, FirstName: "Roger", LastName: "Federer", Email: "roger_federer@yahoo.com.tw"},
		67:  {ID: 67, FirstName: "Rafael", LastName: "Nadal", Email: "rnadal@yahoo.com.tw"},
		88:  {ID: 88, FirstName: "Novak", LastName: "Djokovic", Email: "novakD@gmail.com"},
	}

	UserDao userServiceInterface
)

func init() {
	UserDao = &userDao{}
}

func (u *userDao) GetUser(userID int64) (*User, *utils.ApplicationError) {
	// extract the user data, if not found, will be nil
	fmt.Println("Accessing the database....")
	user := users[userID]
	if user != nil {
		return user, nil
	}
	// if userID is not found
	applicationErr := &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
	return nil, applicationErr

}
