package main

import (
	"fmt"
	"net/http"
	"os"

	"microgo/internal/auth"
	"microgo/internal/utils"

	M "microgo/internal/auth/middleware"

	C "microgo/internal/controllers"
)

func main() {

	utils.LoadEnvs("../.env")

	PORT := os.Getenv("APP_PORT")

	http.HandleFunc("/api/v0.1/", M.ValidateToken(C.Home))
	http.HandleFunc("/api/v0.1/login", auth.Login)
	// http.HandleFunc("/api/v1/refresh", auth.Refresh)

	fmt.Println("Listen and Server at :3500")
	http.ListenAndServe(":"+PORT, nil)
}
