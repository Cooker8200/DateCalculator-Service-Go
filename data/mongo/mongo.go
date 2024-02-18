package mongo

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Date struct {
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required"`
	Date string `json:"date" binding:"required"`
}


func GetAllDates(c *gin.Context, mongo *mongo.Client) {
	log.Print("MONGO:::", mongo)
	dates, err := mongo.Database("dateCalculator").Collection("dates").Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Print("Failed to find all dates")
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"body": "Failed to find all dates. Mongo may not be running"})
		return
	}

	var parsedDates []bson.M
	if err = dates.All(context.TODO(), &parsedDates); err != nil {
		log.Print("Error parsing results")
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"body": "Failed to parse found dates"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"body": parsedDates})
}

func getSpecificDate(mongoClient *mongo.Client, id interface{}) []primitive.M {
	date, err := mongoClient.Database("dateCalculator").Collection("dates").Find(context.TODO(), bson.M{"_id": id})
	if err != nil {
		log.Print("Failed to find dates: ", id)
	}

	var parsedDate []bson.M
	if err = date.All(context.TODO(), &parsedDate); err != nil {
		log.Print("Error parsing result")
	}

	return parsedDate
}

func AddNewDate(c *gin.Context, mongo *mongo.Client) {
	var newDate Date

	if err := (c.ShouldBindJSON(&newDate)); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to bind request body")
		return
	}

	result, err := mongo.Database("dateCalculator").Collection("dates").InsertOne(context.TODO(), newDate)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"body": "Failed to add new date"})
		return
	}

	insertedDate := getSpecificDate(mongo, result.InsertedID)

	c.IndentedJSON(http.StatusOK, gin.H{"body": insertedDate})
}

func RemoveDate(c *gin.Context, mongo *mongo.Client) {
	var dateToRemove Date

	if err := (c.ShouldBindJSON(&dateToRemove)); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"body": "Failed to bind request body"})
		return
	}

	result, err := mongo.Database("dateCalculator").Collection("dates").DeleteOne(context.TODO(), bson.D{{Key: "name", Value: dateToRemove.Name}})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"body": "Failed to add remove date"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"body": result})
}

func WipeDatabase(c *gin.Context, mongo *mongo.Client) {
	if err := mongo.Database("dateCalculator").Drop(context.TODO()); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"body": "Did not drop the db"})
		return
	}

	if err := mongo.Database("dateCalculator").CreateCollection(context.TODO(), "dates"); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"body": "Error creating fresh database"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"body": "Fresh database ready to go!"})
}
