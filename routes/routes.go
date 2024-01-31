package routes

import (
	aws "DateCalculator-Service-Go/data/aws"
	mongo "DateCalculator-Service-Go/data/mongo"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()

	router.GET("/dates/aws", aws.GetAllDates)
	router.PUT("/dates/aws", aws.AddNewDate)
	router.DELETE("/dates/aws", aws.RemoveDate)

	router.GET("/dates/mongo", mongo.GetAllDates)
	router.PUT("/dates/mongo", mongo.AddNewDate)
	router.DELETE("/dates/mongo", mongo.RemoveDate)
	router.DELETE("/dates/mongo/wipe", mongo.WipeDatabase)

	router.Run("localhost:3001")
}
