package main

import (
	"github.com/gorilla/mux"
	"github.com/rishi-suresh-keshav/user-service/lib/repository"
	"github.com/rishi-suresh-keshav/user-service/lib/routes"
	"log"
	"net/http"
)

//var db *gorm.DB

func main() {
	router := mux.NewRouter()
	route.RegisterAllRoutes(router)
	http.Handle("/", router)

	initDBConnection()

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Something went wrong listening to 8080. ", err)
	}
	log.Println("Server is listening at port 8080")
}

func initDBConnection() {
	//config.Connect()
	//db = config.GetDB()

	repository.InitDatabaseConnect()
}
