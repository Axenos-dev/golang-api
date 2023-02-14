package upload_picture

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type Response struct {
	Message string
}

func Upload_Picture(data UploadPicture, token_string string) (Response, bool) {
	user_id := decode_token(token_string)

	if user_id == "" {
		return Response{Message: "invalid token"}, false
	}

	path_to_img := "./images/" + uuid.New().String() + ".jpg"

	successful_download := Download_Image(data.Url, path_to_img)
	if !successful_download {
		return Response{"Failed to download image"}, false
	}

	insert_into_db(user_id, data.Url, "/")

	return Response{Message: "image loaded"}, true
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
