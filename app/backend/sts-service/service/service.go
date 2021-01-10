package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/tentativafc/investing-broker/app/backend/sts-service/dto"
	"github.com/tentativafc/investing-broker/app/backend/sts-service/repo"
)

type StsService struct {
	ccr repo.ClientCredentialsRepository
}

func (s StsService) CreateClientCredentials(ccr dto.ClientCredentialsRequest) (dto.ClientCredentials, error) {

	err := ccr.Validate()
	if err != nil {
		return dto.ClientCredentials{}, err
	}

	cc := repo.ClientCredentials{ClientName: ccr.ClientName, ClientId: uuid.New().String(), ClientSecret: uuid.New().String(), CreatedAt: time.Now()}
	_, err = s.ccr.CreateClientCredentials(cc)
	if err != nil {
		return dto.ClientCredentials{}, err
	}
	return dto.ClientCredentials{ClientName: cc.ClientName, ClientId: cc.ClientId, ClientSecret: cc.ClientSecret}, nil
}

func (s StsService) CreateToken(clientId string, clientSecret string) (string, error) {

	cr, err := s.ccr.FindByClientId(clientId)
	if err != nil {
		return "", errors.Wrap(err, "Error to find client credentials.")
	}
	if cr.ClientSecret != clientSecret {
		return "", errors.Wrap(err, "Invalid credentials.")
	}
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["client_id"] = cr.ClientId
	atClaims["client_name"] = cr.ClientName
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(cr.ClientSecret))
	if err != nil {
		return "", errors.Wrap(err, "Error to generate token.")
	}
	return token, nil
}

func NewStsService() StsService {
	return StsService{ccr: repo.NewClientCredentialsRepository()}
}

// func (s StsService) GetUserIdFromToken(accessToken string) (string, error) {

// 	cr, err := s.ur.FindByClientId(clientId)
// 	if err != nil {
// 		return cr, errors.Wrap(err, "Error to find client credentials.")
// 	}

// 	// Parse takes the token string and a function for looking up the key. The latter is especially
// 	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
// 	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
// 	// to the callback, providing flexibility.
// 	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
// 		// Don't forget to validate the alg is what you expect:
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}

// 		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
// 		return []byte(ACCESS_SECRET), nil
// 	})

// 	var userId string

// 	if err != nil {
// 		return userId, err
// 	}

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		userId = claims["user_id"].(string)
// 	}
// 	return userId, err
// }
