package auth

import (
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"

	authJWT "microgo/internal/auth/auth_jwt"
)

// type response struct {
// 	token  string `json:"token"`
// 	code   int16  `json:"code"`
// 	status bool   `json:"status"`
// }

func Login(w http.ResponseWriter, r *http.Request) {
	api_key := os.Getenv("API_KEY")

	if r.Header["Access"] != nil {
		if r.Header["Access"][0] != api_key {
			fmt.Println("no match:" + r.Header["Access"][0] + " > " + api_key)
			return
		} else {
			kid := uuid.New().String() // from DB
			token, err := authJWT.CreateToken(kid)
			if err != nil {
				fmt.Printf("Error creating JWT")

				return
			}

			fmt.Fprintf(w, token)
		}
	}
}
