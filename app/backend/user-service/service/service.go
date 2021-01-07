package service

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"

	"strings"

	"github.com/tentativafc/investing-broker/user-service/dto"
	errorUR "github.com/tentativafc/investing-broker/user-service/error"
	"github.com/tentativafc/investing-broker/user-service/repo"
	"github.com/tentativafc/investing-broker/user-service/stspb"
	"github.com/tentativafc/investing-broker/user-service/util"
	"google.golang.org/grpc"
)

type UserService struct {
	ur repo.UserRepository
	sc stspb.StsClient
}

func NewUserService(ur repo.UserRepository) UserService {

	cc, err := grpc.Dial("sts_server:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatal("Could not connect: %v", err)
	}

	// defer cc.Close()

	c := stspb.NewStsClient(cc)

	us := UserService{ur: ur, sc: c}

	return us
}

func (us UserService) CreateUser(u dto.User) (dto.UserResponse, error) {
	u.ID = uuid.New().String()
	userDb := repo.UserDB{ID: u.ID, Firstname: u.Firstname, Lastname: u.Lastname, Email: u.Email, Password: u.Password, CreatedAt: time.Now()}
	_, err := us.ur.CreateUser(userDb)
	var ur dto.UserResponse
	if err != nil {
		return ur, errorUR.NewGenericError(err.Error())

	}

	tr, err := us.sc.GenerateToken(context.Background(), &stspb.TokenRequest{ClientId: userDb.ID})
	if err != nil {
		return ur, errorUR.NewAuthError("Error to generating jwt")
	}
	ur = dto.UserResponse{Token: tr.Token, ID: userDb.ID, Firstname: userDb.Firstname, Lastname: userDb.Lastname, Email: userDb.Email}
	return ur, nil
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

func (us UserService) Login(l dto.LoginData) (dto.UserResponse, error) {

	var ur dto.UserResponse
	err := l.Validate()
	if err != nil {
		return ur, err
	}

	userDb, err := us.ur.FindByEmail(l.Email)
	if err != nil {
		return ur, errorUR.NewNotFoundError("User not found")
	}
	tr, err := us.sc.GenerateToken(context.Background(), &stspb.TokenRequest{ClientId: userDb.ID})
	if err != nil {
		return ur, errorUR.NewAuthError("Error to generating jwt")
	}
	ur = dto.UserResponse{Token: tr.Token, ID: userDb.ID, Firstname: userDb.Firstname, Lastname: userDb.Lastname, Email: userDb.Email}
	return ur, nil
}

func (us UserService) RecoverLogin(recover dto.RecoverLoginData) (dto.RecoverLoginDataResponse, error) {
	var rl dto.RecoverLoginDataResponse
	userDb, err := us.ur.FindByEmail(recover.Email)
	if err != nil {
		return rl, errorUR.NewNotFoundError("User not found")
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
		err = errorUR.NewNotFoundError("User not found")
		return dto.UserResponse{}, err
	}
	return dto.UserResponse{ID: userDb.ID, Firstname: userDb.Firstname, Lastname: userDb.Lastname, Email: userDb.Email, Token: token}, nil
}
