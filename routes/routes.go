package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
func InitRouter() {
	router := gin.Default()

	router.GET("/dates", GetAllDates)

	router.Run("localhost:3001")
}

// data functions - to be moved to separate file when I figure it out
func GetAllDates(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "my first api")
}

func AddNewDate(c *gin.Context) {

}

func RemoveDate(c *gin.Context) {

}