package apirest

import (
	"net/http"

	"microgo/internal/auth"

	M "microgo/internal/auth/middleware"

	C "microgo/internal/controllers"
)

func initializeAuthRoutes() {

	http.HandleFunc("/api/v0.1/", M.ValidateToken(C.Home))
	http.HandleFunc("/api/v0.1/auth/login", auth.Login)
}
