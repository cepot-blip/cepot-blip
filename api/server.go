package api

import (
	"fmt"
	"log"
	"os"

	"github.com/cepot-blip/fullstack/api/controllers"
	"github.com/cepot-blip/fullstack/api/seed"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error dalam mendapatkan file env, not comming through %v", err)
	} else {
		fmt.Println("berhasil mendapatkan env")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	seed.Load(server.DB)

	server.Run(":9000")

}
