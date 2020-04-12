package route

import (
	"github.com/gorilla/mux"
	"github.com/rishi-suresh-keshav/user-service/lib/controller"
)

var RegisterAllRoutes = func(router *mux.Router) {
	router.HandleFunc("/", controller.Home).Methods("GET")
	router.HandleFunc("/user-service/user", controller.CreateUser).Methods("POST")
	router.HandleFunc("/user-service/user/{id}", controller.GetUserByID).Methods("GET")
}