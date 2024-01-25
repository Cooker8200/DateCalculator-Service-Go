package routes

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName string
}

type Date struct {
	dateName string
	dateType string
	date string
}

func InitRouter() {
	router := gin.Default()

	router.GET("/dates", GetAllDates)
	router.PUT("/dates", AddNewDate)
	router.DELETE("/dates", RemoveDate)

	router.Run("localhost:3001")
}

func configureAWS() (*dynamodb.Client) {
	envs, err := godotenv.Read(".env")

	if err != nil {
			log.Fatal("Error loading .env file")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(envs["aws_access_key_id"], envs["aws_secret_access_key"], "")))
	if err != nil {
		log.Fatalln("Unable to load SDK config: ", err)
	}

	svc := dynamodb.NewFromConfig(cfg)
	return svc
}

func GetAllDates(c *gin.Context) {
	svc := configureAWS()

	resp, err := svc.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String("dates"),
	})
	if err != nil {
		log.Println("Failed to scan: ", err)
		c.IndentedJSON(http.StatusInternalServerError, "Something went wrong.")
	} else {
		c.IndentedJSON(http.StatusOK, resp.Items)
	}
}

func AddNewDate(c *gin.Context) {
	// newDate := c.Request.Body
	log.Println("request body: ", c)
	var newDate Date
	
	svc := configureAWS()

	if err := c.BindJSON(&newDate); err != nil {
		return
	}
	log.Print("new date: ", newDate)
	resp, err := svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("dates"),
		Item: map[string]types.AttributeValue{
			"name": &types.AttributeValueMemberS{Value: "test"},
			"type": &types.AttributeValueMemberS{Value: "test"},
			"date": &types.AttributeValueMemberS{Value: "test"},
		},
	})
	if err != nil {
		log.Println("Failed to add new date: ", err)
		c.IndentedJSON(http.StatusInternalServerError, "Failed to add new date.")
	} else {
		c.IndentedJSON(http.StatusOK, resp.ResultMetadata)
	}
}

func RemoveDate(c *gin.Context) {
	// key := c.Request
	// var key = "keyToRemove"

	svc := configureAWS()

	_, err := svc.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String("dates"),
		Key: map[string]types.AttributeValue{
			"name": &types.AttributeValueMemberS{Value: "test"},
		},
	})
	if err != nil {
		log.Println("Failed to delete item: ", err)
		c.IndentedJSON(http.StatusInternalServerError, "Failed to remove date")
	} else {
		c.IndentedJSON(http.StatusOK, "Date removed")
	}
}