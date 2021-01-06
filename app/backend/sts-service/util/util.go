package util

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(clientId string) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["client_id"] = clientId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(ACCESS_SECRET))
	if err != nil {
		return "", err
	}
	return token, nil
}

func GetUserIdFromToken(accessToken string) (string, error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(ACCESS_SECRET), nil
	})

	var userId string

	if err != nil {
		return userId, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId = claims["user_id"].(string)
	}
	return userId, err
}

func GetSubstringAfter(value string, a string) string {
	// Get substring after a string.
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:]
}
