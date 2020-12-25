package service

import (
	"time"

	"github.com/google/uuid"

	"strings"

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

func (us UserService) CreateUser(u dto.User) (dto.User, error) {
	u.ID = uuid.New().String()
	userDb := repo.UserDB{ID: u.ID, Firstname: u.Firstname, Lastname: u.Lastname, Email: u.Email, Password: u.Password, CreatedAt: time.Now()}
	_, err := us.ur.CreateUser(userDb)
	if err != nil {
		return dto.User{}, errorUR.NewGenericError(err.Error())

	}
	return u, nil
}

func (us UserService) UpdateUser(u dto.UserUpdate, authorization string) (dto.UserUpdate, error) {

	if !strings.HasPrefix(authorization, "Bearer ") {
		err := errorUR.NewAuthError("Token not found")
		return dto.UserUpdate{}, err
	}

	token := util.GetSubstringAfter(authorization, "Bearer ")

	userIdJwt, err := util.GetUserIdFromToken(token)

	if err != nil {
		err = errorUR.NewAuthError("Token expired or invalid")
		return dto.UserUpdate{}, err
	}

	if u.ID != userIdJwt {
		err = errorUR.NewAuthError("Invalid credentials")
		return dto.UserUpdate{}, err
	}

	userDb := repo.UserDB{ID: u.ID, Firstname: u.Firstname, Lastname: u.Lastname, Email: u.Email, UpdatedAt: time.Now()}
	_, err = us.ur.UpdateUser(userDb)
	if err != nil {
		return dto.UserUpdate{}, errorUR.NewGenericError(err.Error())

	}
	return u, nil
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
	return lr, nil
}

func (us UserService) RecoverLogin(recover dto.RecoverLoginData) (dto.RecoverLoginDataResponse, error) {
	var rl dto.RecoverLoginDataResponse
	userDb, err := us.ur.FindByEmail(recover.Email)
	if err != nil {
		return rl, errorUR.NewNotFountError("User not found")
	}

	tempPassword := uuid.New().String()
	r, err := us.ur.CreateRecoverPassword(userDb, uuid.New(), tempPassword)
	if err != nil {
		err = errorUR.NewAuthError("Error creating recover password")
		return dto.RecoverLoginDataResponse{}, err
	}
	rl = dto.RecoverLoginDataResponse{ID: r.ID, Email: userDb.Email}
	return rl, nil
}

func (us UserService) GetuserById(authorization string, userId string) (dto.UserResponse, error) {

	if !strings.HasPrefix(authorization, "Bearer ") {
		err := errorUR.NewAuthError("Token not found")
		return dto.UserResponse{}, err
	}

	token := util.GetSubstringAfter(authorization, "Bearer ")

	userIdJwt, err := util.GetUserIdFromToken(token)

	if err != nil {
		err = errorUR.NewAuthError("Token expired or invalid")
		return dto.UserResponse{}, err
	}

	if userId != userIdJwt {
		err = errorUR.NewAuthError("Invalid credentials")
		return dto.UserResponse{}, err

	}

	userDb, err := us.ur.FindById(userId)
	if err != nil {
		err = errorUR.NewNotFountError("User not found")
		return dto.UserResponse{}, err

	}
	return dto.UserResponse{ID: userDb.ID, Firstname: userDb.Firstname, Lastname: userDb.Lastname, Email: userDb.Email}, nil
}
