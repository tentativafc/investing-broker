package service

import (
	"time"

	"github.com/google/uuid"

	"github.com/tentativafc/investing-broker/user-rest/dto"
	errorUR "github.com/tentativafc/investing-broker/user-rest/error"
	"github.com/tentativafc/investing-broker/user-rest/repo"
	"github.com/tentativafc/investing-broker/user-rest/util"
)

type UserService struct {
	ur repo.UserRepository
}

func NewUserService(ur repo.UserRepository) UserService {
	us := UserService{ur: ur}
	return us
}

func (us UserService) CreateUser(u dto.User) dto.User {
	u.ID = uuid.New().String()
	userDb := repo.UserDB{ID: u.ID, Firstname: u.Firstname, Lastname: u.Lastname, Email: u.Email, Password: u.Password, CreatedAt: time.Now()}
	us.ur.CreateUser(userDb)
	return u
}

func (us UserService) Login(login dto.LoginData) (dto.LoginResponse, error) {
	var lr dto.LoginResponse
	userDb, err := us.ur.FindByEmail(login.Email)
	if err != nil {
		return lr, errorUR.NewNotFountError("User not found")
	}

	authToken, err := util.CreateToken(userDb.ID)
	if err != nil {
		return lr, errorUR.NewAuthError("Error to generating jwt")
	}
	lr = dto.LoginResponse{Token: authToken, UserData: dto.UserResponse{ID: userDb.ID, Firstname: userDb.Firstname, Lastname: userDb.Lastname, Email: userDb.Email}}
	return lr, err
}

func (us UserService) RecoverLogin(recover dto.RecoverLoginData) (dto.RecoverLoginDataResponse, error) {
	var rl dto.RecoverLoginDataResponse
	userDb, err := us.ur.FindByEmail(recover.Email)
	if err != nil {
		return rl, errorUR.NewNotFountError("User not found")
	}

	tempPassword := uuid.New().String()
	r := us.ur.CreateRecoverPassword(userDb, uuid.New(), tempPassword)
	rl = dto.RecoverLoginDataResponse{ID: r.ID, Email: userDb.Email}
	return rl, err
}

func (us UserService) GetuserById(token string, userId string) (u dto.UserResponse, err error) {

	userIdJwt, err := util.GetUserIdFromToken(token)

	if err != nil {
		err = errorUR.NewAuthError("Token expired or invalid")
		return
	}

	if userId != userIdJwt {
		err = errorUR.NewAuthError("Invalid credentials")
		return
	}

	userDb, err := us.ur.FindById(userId)
	if err != nil {
		err = errorUR.NewNotFountError("User not found")
		return
	}

	u = dto.UserResponse{ID: userDb.ID, Firstname: userDb.Firstname, Lastname: userDb.Lastname, Email: userDb.Email}
	return
}
