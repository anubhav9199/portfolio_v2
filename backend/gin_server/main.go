package main

import (
	"gin_server/configs"
	"gin_server/routes"

	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    //run database
    configs.ConnectDB()

    //routes
    routes.Route(router)

    router.Run("localhost:6000")
}

