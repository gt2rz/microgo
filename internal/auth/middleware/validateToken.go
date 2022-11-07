package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"

	authJWT "microgo/internal/auth/jwt"
)

func ValidateToken(next func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)

				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("Not authorized"))
				}

				signature := authJWT.SetSignature("")

				return signature, nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Not authorized: " + err.Error()))
			}

			if token.Valid {

				claims, ok := token.Claims.(jwt.MapClaims)
				if ok {
					iss := claims["iss"].(string)
					// aud := claims["aud"].(string)

					if iss != "localhost:3500" {
						w.WriteHeader(http.StatusUnauthorized)
						w.Write([]byte("Not authorized: iss not valid"))

					}
					// else if aud != "localhost:3500/api" {
					// 	w.WriteHeader(http.StatusUnauthorized)
					// 	w.Write([]byte("Not authorized: aud not valid"))
					// }

					// } else {

					next(w, r)

				}
				// }

			}

		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Not authorized"))
		}
	})
}
