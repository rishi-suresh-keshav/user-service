package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rishi-suresh-keshav/user-service/config"
	"github.com/rishi-suresh-keshav/user-service/handlers"
	"github.com/rishi-suresh-keshav/user-service/routes"
	"net/http"
)

const (
	dbName = "user"
	dbPass = ""
	dbHost = "docker.for.mac.localhost" // Weird Docker issue for local https://stackoverflow.com/questions/52504318/unable-to-connect-to-mysql-server-with-go-and-docker-dial-tcp-127-0-0-13306
	dbPort = "3306"
)

func main() {
	//dbName := os.Getenv("DB_NAME")
	//dbPass := os.Getenv("DB_PASS")
	//dbHost := os.Getenv("DB_HOST")
	//dbPort := os.Getenv("DB_PORT")

	println("this is db", dbName, dbHost, dbPass, dbPort)

	connection, err := config.ConnectSQL(dbHost, dbPort, "root", dbPass, dbName)
	if err != nil {
		fmt.Println(err)
		//os.Exit(-1)
	}

	uHandler := handlers.NewUserHandler(connection)
	router := mux.NewRouter()
	route.RegisterAllRoutes(router, uHandler)
	http.Handle("/", router)

	fmt.Println("Server listen at :8080")
	http.ListenAndServe(":8080", router)
}