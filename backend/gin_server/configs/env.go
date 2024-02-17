package configs

import (
	"os"
)


func EnvMongoURI() string {
	_, cerr := os.Getwd()
	if cerr != nil {
		panic(cerr)
	}
	return os.Getenv("MONGO_URI")
}
