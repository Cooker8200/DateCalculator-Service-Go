package routes

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
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
	// os.Setenv("AWS_PROFILE", "go")
	// cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("go"))
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("", "", "")))
	if err != nil {
		log.Fatalf("Unable to load SDK config: %v\n", err)
	}

	svc := dynamodb.NewFromConfig(cfg)

	resp, err := svc.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String("dates"),
	})
	if err != nil {
		log.Printf("Failed to scan: %v\n", err)
	} else {
		log.Printf("results %v\n", resp)
		c.IndentedJSON(http.StatusOK, resp.Items)
	}
}

func AddNewDate(c *gin.Context) {

}

func RemoveDate(c *gin.Context) {

}