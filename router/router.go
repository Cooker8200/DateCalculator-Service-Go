package router

import (
	"log"
	"net/http"
	"DateCalculator-Service-Go/data"

	"github.com/gin-gonic/gin"
)

func InitRouter()  {
	router := gin.Default()
	router.GET("/dates", getAllDates)

	log.Println("API running port 3001")
	http.ListenAndServe(":3001", router)
}
