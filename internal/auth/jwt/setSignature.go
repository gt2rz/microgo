package jwt

import (
	"os"
)

func SetSignature(sign string) []byte {
	secret := os.Getenv("ACCESS_SECRET")

	return []byte(secret + sign)
}
