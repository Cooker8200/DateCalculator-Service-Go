package routes

import (
	dates "DateCalculator-Service-Go/data"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()

	router.GET("/dates", dates.GetAllDates)
	router.PUT("/dates", dates.AddNewDate)
	router.DELETE("/dates", dates.RemoveDate)

	router.Run("localhost:3001")
}
