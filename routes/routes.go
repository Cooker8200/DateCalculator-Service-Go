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
	"github.com/joho/godotenv"
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
	var envs map[string]string
	envs, err := godotenv.Read(".env")

	if err != nil {
			log.Fatal("Error loading .env file")
	}

	aws_access_key_id := envs["aws_access_key_id"]
	aws_secret_access_key := envs["aws_secret_access_key"]

	// os.Setenv("AWS_PROFILE", "go")
	// cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("go"))
	// DO NOT COMMIT HARD CODED VALUES
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(aws_access_key_id, aws_secret_access_key, "")))
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