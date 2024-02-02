package routes

import (
	aws "DateCalculator-Service-Go/data/aws"
	mongo "DateCalculator-Service-Go/data/mongo"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:4200", "http://localhost:3000"},
    AllowMethods:     []string{"GET", "PUT", "DELETE", "OPTIONS"},
    AllowHeaders:     []string{"Origin", "Access-Control-Allow-Origin"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
    AllowOriginFunc: func(origin string) bool {
      return origin == "http://localhost:4200"
    },
    MaxAge: 12 * time.Hour,
  }))

	router.GET("/dates/aws", aws.GetAllDates)
	router.PUT("/dates/aws", aws.AddNewDate)
	router.DELETE("/dates/aws", aws.RemoveDate)

	router.GET("/dates/mongo", mongo.GetAllDates)
	router.PUT("/dates/mongo", mongo.AddNewDate)
	router.DELETE("/dates/mongo", mongo.RemoveDate)
	router.DELETE("/dates/mongo/wipe", mongo.WipeDatabase)

	router.Run("localhost:3001")
}
