package app

import (
	"fmt"
	"george_custom/golang_microservice/introduction/mvc/controllers"
	"log"
	"net/http"
)

func StartApp() {
	fmt.Println("Start serving....")
	http.HandleFunc("/users", controllers.GetUser)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
