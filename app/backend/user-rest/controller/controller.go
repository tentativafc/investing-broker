package controller

import (
	"encoding/json"
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

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u dto.User
	json.NewDecoder(r.Body).Decode(&u)
	ur, err := us.CreateUser(u)
	if err != nil {
		HandleError(err, w)
	} else {
		json.NewEncoder(w).Encode(&ur)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var u dto.UserUpdate
	json.NewDecoder(r.Body).Decode(&u)
	params := mux.Vars(r)
	var uId = params["id"]
	authorization := r.Header.Get("Authorization")
	u.ID = uId
	u, err := us.UpdateUser(u, authorization)
	if err != nil {
		HandleError(err, w)
	} else {
		json.NewEncoder(w).Encode(&u)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var l dto.LoginData
	_ = json.NewDecoder(r.Body).Decode(&l)
	ur, err := us.Login(l)
	if err != nil {
		HandleError(err, w)
	} else {
		json.NewEncoder(w).Encode(&ur)
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
	var code int
	switch err.(type) {
	case *errorUR.NotFoundError:
		code = err.(*errorUR.NotFoundError).Code()
	case *errorUR.AuthError:
		code = err.(*errorUR.AuthError).Code()
	default:
		code = http.StatusInternalServerError
	}
	er := dto.ErrorResponse{Code: code, Message: err.Error()}
	// Return error response
	NewError(w, er)
}

func NewError(w http.ResponseWriter, er dto.ErrorResponse) {
	w.WriteHeader(er.Code)
	json.NewEncoder(w).Encode(er)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
}

func HandleRequests() {

	router := mux.NewRouter()
	router.HandleFunc("/users/login", Login).Methods("POST")
	router.HandleFunc("/users/{id}", GetUserById).Methods("GET")
	router.HandleFunc("/users/recover", RecoverLogin).Methods("POST")
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":8082", handler))
}
