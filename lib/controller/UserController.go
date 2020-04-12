package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rishi-suresh-keshav/user-service/lib/repository"
	"github.com/rishi-suresh-keshav/user-service/lib/parser"
	"log"
	"net/http"
	"strconv"
)

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	ID, err:= strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	user := repository.GetUserByID(ID)
	log.Println("User Details: " , user)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	newUser := repository.User{}
	parser.ParseRequest(r, &newUser)
	newUser.CreateUser()
}
