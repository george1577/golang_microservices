package controllers

import (
	"encoding/json"
	"fmt"
	"george_custom/golang_microservice/introduction/mvc/services"
	"george_custom/golang_microservice/introduction/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Someone just made a call...")

	// get a user_id field in the url
	// controller is only handling the request, if the userID is valid, will pass to the service
	userIDParam := request.URL.Query().Get("user_id")
	userID, err := strconv.ParseInt(userIDParam, 10, 64)
	if err != nil {
		userErr := &utils.ApplicationError{
			Message:    "user_id must be an integer",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(userErr)
		writer.WriteHeader(userErr.StatusCode)
		writer.Write(jsonValue)
		// Just return the bad request to the client

		// writer.Write([]byte("user_id must be an integer"))
		return
	}

	// send the userID to the service, then service will call the model(domain) to extract the userID from the database
	user, appliErr := services.GetUser(userID)
	if appliErr != nil {
		// log.Fatalf("something went wrong... %v\n", err)
		// fmt.Printf("User with %v was not found...\n", userID)

		// WriteHeader will write the status code of the request, ex: 404=not found, 400=bad request...etc.
		jsonValue, _ := json.Marshal(appliErr)
		writer.WriteHeader(appliErr.StatusCode)
		writer.Write(jsonValue)
		// writer.Write([]byte(appliErr.Message))
		return
	}

	// return user to client
	jsonValue, _ := json.Marshal(user)
	writer.Write(jsonValue)

}
