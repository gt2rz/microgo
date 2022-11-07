package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func RefreshToken(kid string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	// signature := authJWT.setSignature(kid)
	signature := SetSignature("")

	now := time.Now()

	claims["kid"] = kid
	claims["iat"] = now.Unix()
	claims["exp"] = now.Add(time.Hour).Unix()
	claims["scope"] = "test"
	claims["iss"] = "localhost:3500"     //Signer
	claims["aud"] = "localhost:3500/api" //Audience or clients

	tokenStr, err := token.SignedString([]byte(signature))

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return tokenStr, nil
}
