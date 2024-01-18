package routes

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
)

type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName string
}

func InitRouter() {
	router := gin.Default()

	router.GET("/dates", GetAllDates)

	router.Run("localhost:3001")
}

// data functions - to be moved to separate file when I figure it out
func GetAllDates(c *gin.Context) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("Unable to load SDK config")
	}

	svc := dynamodb.NewFromConfig(cfg)

	resp, err := svc.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String("dates"),
	})
	if err != nil {
		log.Printf("Failed to scan", err)
	}

	c.IndentedJSON(http.StatusOK, resp)
}

func AddNewDate(c *gin.Context) {

}

func RemoveDate(c *gin.Context) {

}