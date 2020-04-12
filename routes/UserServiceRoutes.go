package route

import (
	"github.com/gorilla/mux"
	"github.com/rishi-suresh-keshav/user-service/handlers"
)

var RegisterAllRoutes = func(router *mux.Router, handler *handlers.UserHandler) {
	router.HandleFunc("/user-service/user", handler.CreateUser).Methods("POST")
	router.HandleFunc("/user-service/user/{id}", handler.GetUserByID).Methods("GET")
}
