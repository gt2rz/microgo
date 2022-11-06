package utils

import (
	"fmt"
	"log"
	"github.com/joho/godotenv"

)

func LoadEnvs(file string){

	err := godotenv.Load(file)
	if err != nil {
		log.Fatal("Error loading envs file")
	}

	fmt.Println("Envs file loaded")
}