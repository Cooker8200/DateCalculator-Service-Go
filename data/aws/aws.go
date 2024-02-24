package aws

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

type Date struct {
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required"`
	Date string `json:"date" binding:"required"`
}

func configureAWS() *dynamodb.Client {
	envs, err := godotenv.Read(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(envs["aws_access_key_id"], envs["aws_secret_access_key"], "")),
	)
	if err != nil {
		log.Fatalln("Unable to load SDK config: ", err)
	}

	dynamo := dynamodb.NewFromConfig(cfg)
	return dynamo
}

func GetAllDates(c *gin.Context) {
	dynamo := configureAWS()

	resp, err := dynamo.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String("dates"),
	})
	if err != nil {
		log.Println("Failed to scan: ", err)
		errMessage := "Failed to scan dates ::: " + err.Error()
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"body": errMessage})
		return
	}

	var dates []Date
	err = attributevalue.UnmarshalListOfMaps(resp.Items, &dates)
	if err != nil {
		log.Print("Failed to do unmarshal data")
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"body": dates})
}

func AddNewDate(c *gin.Context) {
	dynamo := configureAWS()

	var dateToAdd Date

	if err := c.ShouldBindJSON(&dateToAdd); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"body": "Failed to bind request body"})
		return
	}

	_, err := dynamo.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("dates"),
		Item: map[string]types.AttributeValue{
			"name": &types.AttributeValueMemberS{Value: dateToAdd.Name},
			"type": &types.AttributeValueMemberS{Value: dateToAdd.Type},
			"date": &types.AttributeValueMemberS{Value: dateToAdd.Date},
		},
	})
	if err != nil {
		log.Println("Failed to add new date: ", err)
		errMessage := "Failed to add new date ::: " + err.Error()
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"body": errMessage})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"body": "Date added"})
}

func RemoveDate(c *gin.Context) {
	dynamo := configureAWS()

	var dateToRemove Date

	if err := c.ShouldBindJSON(&dateToRemove); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"body": "Failed to bind request body"})
		return
	}

	_, err := dynamo.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String("dates"),
		Key: map[string]types.AttributeValue{
			"name": &types.AttributeValueMemberS{Value: dateToRemove.Name},
		},
	})
	if err != nil {
		log.Println("Failed to delete item: ", err)
		errMessage := "Failed to remove date ::: " + err.Error()
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"body": errMessage})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"body": "Date removed"})
}
