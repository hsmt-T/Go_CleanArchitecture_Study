package main

import (
	"go_cleanArchitecture_study/infrastructure/router"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(".envが存在しません")
	}

	router.StartEcho()
	// router.StartGin()
}