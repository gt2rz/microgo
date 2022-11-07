package main

import (
	"fmt"
	"net/http"
	"os"

	"microgo/internal/utils"

	rest "microgo/internal/routes/apirest"
)

func main() {

	utils.LoadEnvs("../.env")

	PORT := os.Getenv("APP_PORT")

	rest.InitializeApiRestRoutes()

	fmt.Println("Listen and Server at :3500")
	http.ListenAndServe(":"+PORT, nil)
}
