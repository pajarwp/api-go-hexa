package jwt

import (
	"api-go-hexa/business/user/model"
	"api-go-hexa/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateJWTToken(u *model.UserLoginModel) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 8).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.GetConfigs().SecretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}
