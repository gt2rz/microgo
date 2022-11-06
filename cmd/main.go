package main

import (
	"fmt"
	"net/http"
	"os"

	"microgo/internal/auth"
	"microgo/internal/utils"

	authJWT "microgo/internal/auth/auth_jwt"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Super secret area")
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Refresh super secret area")
}

func main() {

	utils.LoadEnvs("../.env")

	PORT := os.Getenv("APP_PORT")

	http.HandleFunc("/api/v0.1/", authJWT.ValidateToken(Home))
	http.HandleFunc("/api/v0.1/login", auth.Login)
	// http.HandleFunc("/api/v1/refresh", auth.Refresh)

	fmt.Println("Listen and Server at :3500")
	http.ListenAndServe(":"+PORT, nil)
}
