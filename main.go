package main

import (
	"jual-beli-motor/repository"
	"jual-beli-motor/routes"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(500)
	}
	repository.InitDB()

	routes.Routes()
}
