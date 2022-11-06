package auth_jwt

import (
	"os"
)

func setSignature(sign string) []byte {
	secret := os.Getenv("ACCESS_SECRET")

	return []byte(secret + sign)
}
