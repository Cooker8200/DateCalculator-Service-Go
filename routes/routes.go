package routes

import (
	"DateCalculator-Service-Go/dates"

	"github.com/gin-gonic/gin"
)
func InitRouter() {
	router := gin.Default()

	router.GET("/dates", GetAllDates)

	router.Run("localhost:3001")
}
