package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	env := `C:\Users\suryatejam\Documents\jwt gen\app.env`
	err := godotenv.Load(env)

	if err != nil {
		log.Println("Error while reading env variables")
	}
}
