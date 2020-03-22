package route

import (
	"controller"
	"github.com/gorilla/mux"
)

var RegisterAllRoutes = func(router *mux.Router) {
	router.HandleFunc("/", controller.Home).Methods("GET")
	router.HandleFunc("/user-service/user", controller.CreateUser).Methods("POST")
	router.HandleFunc("/user-service/user/{id}", controller.GetUserByID).Methods("GET")
}