package dates

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
// TODO: why won't these import to the routes file
func GetAllDates(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "my first api")
}

func AddNewDate(c *gin.Context) {

}

func RemoveDate(c *gin.Context) {

}