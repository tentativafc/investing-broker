package service

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/tentativafc/investing-broker/app/backend/sts-service/util"
	"github.com/tentativafc/investing-broker/app/backend/user-service/dto"
	errorUR "github.com/tentativafc/investing-broker/app/backend/user-service/error"
	"github.com/tentativafc/investing-broker/app/backend/user-service/repo"
	"github.com/tentativafc/investing-broker/app/backend/user-service/stspb"
)

type UserService struct {
	ur repo.UserRepository
	sc stspb.StsClient
}

func (us UserService) CreateUser(u dto.User) (*dto.UserResponse, error) {
	u.ID = uuid.New().String()
	userDb := repo.User{ID: u.ID, Firstname: u.Firstname, Lastname: u.Lastname, Email: u.Email, Password: u.Password, CreatedAt: time.Now()}
	_, err := us.ur.CreateUser(userDb)
	if err != nil {
		return nil, errorUR.NewGenericError(err.Error(), err)

	}
	cc, err := us.sc.GenerateClientCredentials(context.Background(), &stspb.GenerateClientCredentialsRequest{ClientName: u.Email})
	if err != nil {
		return nil, errorUR.NewGenericError("Error to generating client credentials", err)
	}

	tr, err := us.sc.GenerateToken(context.Background(), &stspb.TokenRequest{ClientId: cc.ClientId, ClientSecret: cc.ClientSecret})
	if err != nil {
		return nil, errorUR.NewAuthError("Error to generating jwt", err)
	}
	return &dto.UserResponse{Token: tr.Token, ID: userDb.ID, Firstname: userDb.Firstname, Lastname: userDb.Lastname, Email: userDb.Email}, nil
}

func (us UserService) UpdateUser(u dto.UserUpdate, authorization string) (*dto.UserUpdate, error) {

	if !strings.HasPrefix(authorization, "Bearer ") {
		err := errorUR.NewAuthError("Token not found", nil)
		return nil, err
	}

	token := util.GetSubstringAfter(authorization, "Bearer ")

	vtr, err := us.sc.ValidateToken(context.Background(), &stspb.ValidateTokenRequest{Token: token})

	if err != nil {
		err = errorUR.NewAuthError("Token expired or invalid", err)
		return nil, err
	}

	if u.Email != vtr.ClientName {
		err = errorUR.NewAuthError("Invalid credentials", nil)
		return nil, err
	}

	userDb := repo.User{ID: u.ID, Firstname: u.Firstname, Lastname: u.Lastname, Email: u.Email, UpdatedAt: time.Now()}
	_, err = us.ur.UpdateUser(userDb)
	if err != nil {
		return nil, errorUR.NewGenericError(err.Error(), err)

	}
	return &u, nil
}

func (us UserService) Login(l dto.LoginData) (*dto.UserResponse, error) {

	err := l.Validate()
	if err != nil {
		return nil, err
	}

	userDb, err := us.ur.FindByEmail(l.Email)
	if err != nil {
		return nil, errorUR.NewNotFoundError("User not found")
	}

	ccr, err := us.ur.FindClientCredentialsByClientName(l.Email)
	if ccr == nil || err != nil {
		return nil, errorUR.NewBadRequestError("Error to find client credentials.", err)
	}

	tr, err := us.sc.GenerateToken(context.Background(), &stspb.TokenRequest{ClientId: ccr.ClientId, ClientSecret: ccr.ClientSecret})
	if err != nil {
		return nil, errorUR.NewAuthError("Error to generating jwt", err)
	}
	return &dto.UserResponse{Token: tr.Token, ID: userDb.ID, Firstname: userDb.Firstname, Lastname: userDb.Lastname, Email: userDb.Email}, nil
}

func (us UserService) RecoverLogin(recover dto.RecoverLoginData) (*dto.RecoverLoginDataResponse, error) {
	userDb, err := us.ur.FindByEmail(recover.Email)
	if err != nil {
		return nil, errorUR.NewNotFoundError("User not found")
	}

	tempPassword := uuid.New().String()
	r, err := us.ur.CreateRecoverPassword(userDb, uuid.New(), tempPassword)
	if err != nil {
		return nil, errorUR.NewAuthError("Error creating recovery password", err)
	}
	return &dto.RecoverLoginDataResponse{ID: r.ID, Email: userDb.Email}, nil

}

func (us UserService) GetuserById(authorization string, userId string) (*dto.UserResponse, error) {

	if !strings.HasPrefix(authorization, "Bearer ") {
		err := errorUR.NewAuthError("Token not found", nil)
		return nil, err
	}

	token := util.GetSubstringAfter(authorization, "Bearer ")

	vtr, err := us.sc.ValidateToken(context.Background(), &stspb.ValidateTokenRequest{Token: token})

	if err != nil {
		err = errorUR.NewAuthError("Token expired or invalid", err)
		return nil, err
	}

	userDb, err := us.ur.FindById(userId)
	if err != nil {
		err = errorUR.NewNotFoundError("User not found")
		return nil, err
	}

	if userDb.Email != vtr.ClientName {
		err = errorUR.NewAuthError("Invalid credentials", nil)
		return nil, err
	}

	return &dto.UserResponse{ID: userDb.ID, Firstname: userDb.Firstname, Lastname: userDb.Lastname, Email: userDb.Email, Token: token}, nil
}

func NewUserService(ur repo.UserRepository, sc stspb.StsClient) UserService {
	us := UserService{ur: ur, sc: sc}
	return us
}
