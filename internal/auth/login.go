package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"

	authJWT "microgo/internal/auth/jwt"
)

type data struct {
	Token     string `json:"token"`
	ExpiredAt int64  `json:"expiredAt`
}

type response struct {
	Code   int16 `json:"code"`
	Status bool  `json:"status"`
	Data   data  `json:"data"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	api_key := os.Getenv("API_KEY")

	if r.Header["Access"] != nil {
		if r.Header["Access"][0] != api_key {
			fmt.Println("no match:" + r.Header["Access"][0] + " > " + api_key)
			return
		} else {
			kid := uuid.New().String() // from DB

			tokenData, err := authJWT.CreateToken(kid)

			if err != nil {
				fmt.Printf("Error creating JWT")
				return
			}

			r := &response{
				Code: http.StatusOK,
				Data: data{
					Token:     tokenData.Token,
					ExpiredAt: tokenData.ExpiredAt,
				},
				Status: true,
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(r)

		}
	}
}
