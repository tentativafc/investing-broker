package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/tentativafc/investing-broker/user-rest/dto"
	"github.com/tentativafc/investing-broker/user-rest/repo"
	"github.com/tentativafc/investing-broker/user-rest/service"
)

var ur repo.UserRepository = repo.NewUserRepository()
var us service.UserService = service.NewUserService(ur)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user = us.CreateUser(user)
	json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var login dto.LoginData
	_ = json.NewDecoder(r.Body).Decode(&login)

	var loginResponse, err = us.Login(login)
	if err != nil {
		w.WriteHeader(401)
	} else {
		json.NewEncoder(w).Encode(&loginResponse)
	}
}

func RecoverLogin(w http.ResponseWriter, r *http.Request) {

	var recoverLoginData dto.RecoverLoginData
	_ = json.NewDecoder(r.Body).Decode(&recoverLoginData)

	var recoverLoginDataResponse, err = us.RecoverLogin(recoverLoginData)
	if err != nil {
		w.WriteHeader(404)
	} else {
		json.NewEncoder(w).Encode(&recoverLoginDataResponse)
		w.WriteHeader(201)
	}
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var userId = params["id"]
	authorization := r.Header.Get("Authorization")
	userResponse, err := us.GetuserById(authorization, userId)
	if err != nil {
		w.WriteHeader(404)
	} else {
		json.NewEncoder(w).Encode(&userResponse)
	}
}

func HandleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/", Home).Methods("GET")
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users/login", Login).Methods("POST")
	router.HandleFunc("/users/recover", RecoverLogin).Methods("POST")
	router.HandleFunc("/users/{id}", GetUserById).Methods("GET")

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":8081", handler))
}

func main() {
	HandleRequests()
}
