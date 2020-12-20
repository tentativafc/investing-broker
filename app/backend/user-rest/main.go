package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	guuid "github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var db *gorm.DB

type User struct {
	ID        string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
}

type UserDB struct {
	ID        string `gorm:"primarykey"`
	Firstname string
	Lastname  string
	Email     string `gorm:"uniqueIndex"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (UserDB) TableName() string {
	return "user"
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = guuid.New().String()

	userDb := UserDB{ID: user.ID, Firstname: user.Firstname, Lastname: user.Lastname, Email: user.Email, Password: user.Password, CreatedAt: time.Now()}
	db.Create(&userDb)

	json.NewEncoder(w).Encode(user)
}

func GetUsersByEmail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var email = params["email"]
	var userDb UserDB
	db.Where("email = ?", email).First(&userDb)
	user := User{ID: userDb.ID, Firstname: userDb.Firstname, Lastname: userDb.Lastname, Email: userDb.Email}
	json.NewEncoder(w).Encode(&user)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var userId = params["id"]

	var userDb UserDB
	if err := db.First(&userDb, &userId).Error; err != nil {
		w.WriteHeader(404)
	} else {
		fmt.Printf("Not Null")
		user := User{ID: userDb.ID, Firstname: userDb.Firstname, Lastname: userDb.Lastname, Email: userDb.Email}
		json.NewEncoder(w).Encode(&user)
	}
}

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/", Home).Methods("GET")
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", GetUserById).Methods("GET")
	router.HandleFunc("/users", GetUsersByEmail).Queries("email", "{email}").Methods("GET")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8081", handler))
}

func dbConfig() {
	db_configs := "host=localhost user=postgres password=123456 dbname=postgres port=5432"
	var err error
	db, err = gorm.Open(postgres.Open(db_configs), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&UserDB{})
}

func main() {
	dbConfig()
	handleRequests()
}
