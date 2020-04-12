package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rishi-suresh-keshav/user-service/config"
	"github.com/rishi-suresh-keshav/user-service/models"
	"github.com/rishi-suresh-keshav/user-service/repository"
	"github.com/rishi-suresh-keshav/user-service/repository/user_repo"
	"log"
	"net/http"
	"strconv"
)

func NewUserHandler(db *config.DB) *UserHandler {
	return &UserHandler{
		userRepo: user_repo.NewMysqlUserRepo(db.SQL),
	}
}

type UserHandler struct {
	userRepo repository.UserRepo
}

func (userHandler *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//id := vars["id"]
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 0, 0)
	print("id", id)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	user, err := userHandler.userRepo.GetUserById(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}
	log.Println("User Details: ", user)
	respondwithJSON(w, http.StatusOK, user)
}

func(UserHandler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	newUser := models.User{}
	json.NewDecoder(r.Body).Decode(&newUser)

	newId, err := UserHandler.userRepo.CreateUser(r.Context(), &newUser)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Something wrong happened")
	}

	respondwithJSON(w, http.StatusOK, newId)
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}