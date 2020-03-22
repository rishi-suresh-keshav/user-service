package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

type WelcomeResponse struct {
	ServiceName string
}

func Home(w http.ResponseWriter, r *http.Request) {
	b := 0
	println("Address of b: " , &b)
	testFunction(&b)

	log.Println("Reached Welcome endpoint")
	response, _ := json.Marshal(WelcomeResponse{ServiceName: "Welcome to User Service!"})
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func testFunction(b *int) {
	println(b)
	println(&b)
	println(*b)
}