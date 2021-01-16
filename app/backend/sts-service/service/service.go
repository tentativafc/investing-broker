package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/tentativafc/investing-broker/app/backend/sts-service/dto"
	"github.com/tentativafc/investing-broker/app/backend/sts-service/repo"

	errorSts "github.com/tentativafc/investing-broker/app/backend/sts-service/error"
)

type StsService struct {
	ccr repo.ClientCredentialsRepository
}

func (s StsService) GenerateClientCredentials(ccr dto.ClientCredentialsRequest) (*dto.ClientCredentials, error) {

	err := ccr.Validate()
	if err != nil {
		return nil, errorSts.NewBadRequestError("Error to validate request", err)
	}
	cc, err := repo.NewClientCredentialsRepository().FindByClientName(ccr.ClientName)
	if err != nil {
		return nil, errorSts.NewGenericError("Error to find client credentials by Client Name.", err)
	}
	if cc != nil {
		return nil, errorSts.NewBadRequestError("Client credentials already exists with this Client Name.", nil)
	}
	cc = &repo.ClientCredentials{ClientName: ccr.ClientName, ClientId: uuid.New().String(), ClientSecret: uuid.New().String(), CreatedAt: time.Now()}
	_, err = s.ccr.CreateClientCredentials(cc)
	if err != nil {
		return nil, errorSts.NewGenericError("Error to create client credentials.", err)
	}
	return &dto.ClientCredentials{ClientName: cc.ClientName, ClientId: cc.ClientId, ClientSecret: cc.ClientSecret}, nil
}

func (s StsService) GenerateToken(tr dto.TokenRequest) (string, error) {

	err := tr.Validate()
	if err != nil {
		return "", errorSts.NewBadRequestError(err.Error(), err)
	}

	cr, err := s.ccr.FindByClientId(tr.ClientId)
	if cr == nil || err != nil {
		return "", errorSts.NewGenericError("Error to find client credentials.", err)
	}
	if cr.ClientSecret != tr.ClientSecret {
		return "", errorSts.NewBadRequestError("Invalid credentials.", nil)
	}
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["client_id"] = cr.ClientId
	atClaims["client_name"] = cr.ClientName
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	at.Header["client_id"] = cr.ClientId
	token, err := at.SignedString([]byte(cr.ClientSecret))
	if err != nil {
		return "", errorSts.NewGenericError("Error to generate token.", err)
	}

	return token, nil
}

func (s StsService) ValidateToken(req dto.ValidateTokenRequest) (*dto.ValidateTokenResponse, error) {

	err := req.Validate()
	if err != nil {
		return nil, errorSts.NewBadRequestError(err.Error(), err)
	}

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errorSts.NewGenericError(fmt.Sprint("Unexpected signing method: %v", token.Header["alg"]), nil)
		}

		if _, ok := token.Header["client_id"]; !ok {
			return nil, errorSts.NewGenericError("Invalid token.", nil)
		}

		cr, err := s.ccr.FindByClientId(token.Header["client_id"].(string))
		if err != nil {
			return cr, errorSts.NewGenericError("Error to find client credentials.", err)
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(cr.ClientSecret), nil
	})

	if err != nil {
		return nil, err
	}

	var vtr dto.ValidateTokenResponse

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		vtr = dto.ValidateTokenResponse{Token: req.Token, ClientId: claims["client_id"].(string), ClientName: claims["client_name"].(string)}
	}
	return &vtr, err
}

func NewStsService() StsService {
	return StsService{ccr: repo.NewClientCredentialsRepository()}
}
