package dates

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

// TODO: why won't these import to the routes file
func GetAllDates(c *gin.Context) {
	svc := configureAWS()

	resp, err := svc.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String("dates"),
	})
	if err != nil {
		log.Printf("Failed to scan: %v\n", err)
	} else {
		c.IndentedJSON(http.StatusOK, resp.Items)
	}
}

func AddNewDate(c *gin.Context) {

}

func RemoveDate(c *gin.Context) {

}

func configureAWS() (*dynamodb.Client) {
	envs, err := godotenv.Read(".env")

	if err != nil {
			log.Fatal("Error loading .env file")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(envs["aws_access_key_id"], envs["aws_secret_access_key"], "")))
	if err != nil {
		log.Fatalf("Unable to load SDK config: %v\n", err)
	}

	svc := dynamodb.NewFromConfig(cfg)
	return svc
}