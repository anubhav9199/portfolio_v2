package controllers

import (
	"gin_server/configs"
	"gin_server/models"
	"gin_server/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)


var dataCollection *mongo.Collection = configs.GetCollection(configs.DB, "Services")
var validate = validator.New()


func GetPortfolioServices() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var services []models.Services
		defer cancel()

		results, err := dataCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		defer results.Close(ctx)
		for results.Next(ctx) {
			var service models.Services
			if err = results.Decode(&service); err != nil {
				c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			services = append(services, service)
		}

		c.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": services}})

	}
}
