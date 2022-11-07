package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type data struct {
	Token     string `json:"token`
	ExpiredAt int64  `json:"expiredAt"`
}

func CreateToken(kid string) (data, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	// signature := authJWT.setSignature(kid)
	signature := SetSignature("")

	now := time.Now()
	exp := now.Add(time.Hour).Unix()

	claims["kid"] = kid
	claims["iat"] = now.Unix()
	claims["exp"] = exp
	claims["scope"] = "test"
	claims["iss"] = "localhost:3500"     //Signer
	claims["aud"] = "localhost:3500/api" //Audience or clients

	tokenStr, err := token.SignedString([]byte(signature))

	if err != nil {
		fmt.Println(err.Error())
		d := data{
			Token:     "",
			ExpiredAt: 0,
		}
		return d, err
	}

	d := data{
		Token:     tokenStr,
		ExpiredAt: exp,
	}

	return d, nil
}
