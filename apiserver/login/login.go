package login

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Response struct {
	Token string
}

func Login(user User) (Response, bool) {
	user_db := check_in_db(user)
	if user_db.Id == "" {
		return Response{
			Token: "",
		}, false
	}

	if !user.compare_passwords(user_db) {
		return Response{
			Token: "",
		}, false
	}

	token, err := GenerateJWT(user_db.Id)
	if err != nil {
		return Response{
			Token: "",
		}, false
	}

	return Response{
		Token: token,
	}, true
}

func GenerateJWT(userID string) (string, error) {
	// Create a new token object
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 12).Unix()

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte("very_secret_key"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
