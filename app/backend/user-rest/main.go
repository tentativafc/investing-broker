package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	guuid "github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/dgrijalva/jwt-go"
)

var db *gorm.DB

type User struct {
	ID        string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
}

type UserResponse struct {
	ID        string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
}

type LoginData struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type RecoverLoginData struct {
	Email string `json:"email,omitempty"`
}

type LoginResponse struct {
	Token    string        `json:"auth_token,omitempty"`
	UserData *UserResponse `json:"user,omitempty"`
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

type RecoverLoginDB struct {
	ID                string `gorm:"primarykey"`
	UserID            string
	User              UserDB `gorm:"foreignKey:UserID"`
	TemporaryPassword string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (RecoverLoginDB) TableName() string {
	return "recover_password"
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

func Login(w http.ResponseWriter, r *http.Request) {

	var login LoginData
	_ = json.NewDecoder(r.Body).Decode(&login)

	var userDb UserDB
	if err := db.Where("email = ?", login.Email).First(&userDb).Error; err != nil {
		w.WriteHeader(401)
	} else {
		userResponse := UserResponse{ID: userDb.ID, Firstname: userDb.Firstname, Lastname: userDb.Lastname, Email: userDb.Email}
		// TODO error handler
		token, _ := CreateToken(userResponse.ID)
		loginResponse := LoginResponse{Token: token, UserData: &userResponse}
		json.NewEncoder(w).Encode(&loginResponse)
	}
}

func RecoverLogin(w http.ResponseWriter, r *http.Request) {

	var recover RecoverLoginData
	_ = json.NewDecoder(r.Body).Decode(&recover)

	var userDb UserDB
	if err := db.Where("email = ?", recover.Email).First(&userDb).Error; err != nil {
		w.WriteHeader(404)
	} else {
		ID := guuid.New().String()
		tempPassword := guuid.New().String()
		recoverLoginDB := RecoverLoginDB{ID: ID, UserID: userDb.ID, User: userDb, TemporaryPassword: tempPassword, CreatedAt: time.Now()}
		db.Create(&recoverLoginDB)
		w.WriteHeader(201)
	}
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
	router.HandleFunc("/users/login", Login).Methods("POST")
	router.HandleFunc("/users/recover", RecoverLogin).Methods("POST")
	router.HandleFunc("/users/{id}", GetUserById).Methods("GET")
	router.HandleFunc("/users", GetUsersByEmail).Queries("email", "{email}").Methods("GET")

	handler := cors.AllowAll().Handler(router)

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
	db.AutoMigrate(&RecoverLoginDB{})
}

func CreateToken(userId string) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "mamaandtito") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func main() {
	dbConfig()
	handleRequests()
}
