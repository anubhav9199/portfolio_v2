package main

import (
	"gin_server/configs"
	"gin_server/routes"

	"github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
)

func main() {
    log := logrus.New()
    log.SetLevel(logrus.DebugLevel)
    
    router := gin.Default()

    // adding logs
    router.Use(func(c *gin.Context) {
        log.WithFields(logrus.Fields{
          "method": c.Request.Method,
          "path":   c.Request.URL.Path,
        }).Info("Received request")
        c.Next()
      })

    //run database
    configs.ConnectDB()

    //routes
    routes.Route(router)

    router.Run("localhost:8002")
}

