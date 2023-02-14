package get_pictures

import (
	"github.com/dgrijalva/jwt-go"
)

type Response struct {
	Message string
	Images  []Picture
}

func Get_Pictures(token_string string) (Response, bool) {
	user_id := decode_token(token_string)

	if user_id == "" {
		return Response{"Invalid token", []Picture{}}, false
	}

	results := Select_From_DB(user_id)

	return Response{
		"Images found",
		results,
	}, true
}

func decode_token(token_string string) string {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(token_string, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("very_secret_key"), nil
	})

	if err != nil || !token.Valid {
		return ""
	}

	return claims["user_id"].(string)
}
