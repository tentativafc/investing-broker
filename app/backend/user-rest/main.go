package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/tentativafc/investing-broker/user-rest/dto"
	errorUR "github.com/tentativafc/investing-broker/user-rest/error"
	"github.com/tentativafc/investing-broker/user-rest/repo"
	"github.com/tentativafc/investing-broker/user-rest/service"
)

var ur repo.UserRepository = repo.NewUserRepository()
var us service.UserService = service.NewUserService(ur)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u dto.User
	_ = json.NewDecoder(r.Body).Decode(&u)
	u, err := us.CreateUser(u)
	if err != nil {
		HandleError(err, w)
	} else {
		json.NewEncoder(w).Encode(&u)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var l dto.LoginData
	_ = json.NewDecoder(r.Body).Decode(&l)

	lr, err := us.Login(l)
	if err != nil {
		HandleError(err, w)
	} else {
		json.NewEncoder(w).Encode(&lr)
	}
}

func RecoverLogin(w http.ResponseWriter, r *http.Request) {

	var recoverLoginData dto.RecoverLoginData
	_ = json.NewDecoder(r.Body).Decode(&recoverLoginData)

	rl, err := us.RecoverLogin(recoverLoginData)
	if err != nil {
		HandleError(err, w)
	} else {
		json.NewEncoder(w).Encode(&rl)
		w.WriteHeader(201)
	}
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var uId = params["id"]
	authorization := r.Header.Get("Authorization")
	u, err := us.GetuserById(authorization, uId)
	if err != nil {
		HandleError(err, w)
	} else {
		json.NewEncoder(w).Encode(&u)
	}
}

func HandleError(err error, w http.ResponseWriter) {
	switch err.(type) {
	case *errorUR.NotFoundError:
		w.WriteHeader(404)
	case *errorUR.AuthError:
		w.WriteHeader(401)
	default:
		w.WriteHeader(500)
	}
	json.NewEncoder(w).Encode(dto.ErrorResponse{Msg: err.Error()})

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
