package configs

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)


func EnvMongoURI() string {
	wd,cerr := os.Getwd()
	if cerr != nil {
		panic(cerr)
	}
	parent := filepath.Dir(wd)
	_grand_parent := filepath.Dir(parent)
	grand_parent := _grand_parent + "/.env"
	
	err := godotenv.Load(grand_parent)
	if err != nil {
		log.Fatal("Error loading environment file")
	}

	return os.Getenv("MONGO_URI")
}
