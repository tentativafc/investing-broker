package util

import "testing"

func TestCreateAndDecodeToken(t *testing.T) {

	var userId = "myuserid"

	token, err := CreateToken(userId)

	if err != nil {
		t.Errorf("Error is not expected but got %v", err)
	}

	userIdDecoded, err := GetUserIdFromToken(token)

	if err != nil {
		t.Errorf("Error is not expected but got %v", err)
	}

	if userId != userIdDecoded {
		t.Errorf("Expected %v but got %v", userId, userIdDecoded)

	}

}

func TestDecodeInvalidToken(t *testing.T) {

	token := "token123"

	_, err := GetUserIdFromToken(token)

	if err == nil {
		t.Error("Error is expected")
	}
}
