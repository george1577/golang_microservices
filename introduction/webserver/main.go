package main

import (
	"fmt"
	"log"
	"net/http"
)

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Someone just made a call")
	message := []byte("Hello World!\n")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Currently serving now.....")
	http.HandleFunc("/hello", viewHandler)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
