package mongo

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Date struct {
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required"`
	Date string `json:"date" binding:"required"`
}

func configureMongo() *mongo.Client {
	var mongoClient *mongo.Client

	envs, err := godotenv.Read(".env")
	if err != nil {
			log.Fatal("Error loading .env file")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(envs["mongo_url"]).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Print("Could not connect to mongo.", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Print("Could not ping mongo.", err)
	}
	mongoClient = client

	return mongoClient
}

func GetAllDates(c *gin.Context) {
	mongo := configureMongo()

	dates, err := mongo.Database("dateCalculator").Collection("dates").Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Print("Failed to find all dates")
		c.IndentedJSON(http.StatusInternalServerError, "Failed to find all dates. Mongo may not be running")
	}

	var parsedDates []bson.M
	if err = dates.All(context.TODO(), &parsedDates); err != nil {
		log.Print("Error parsing results")
		c.IndentedJSON(http.StatusInternalServerError, "Failed to parse found dates")
	}

	c.IndentedJSON(http.StatusOK, parsedDates)
}

func AddNewDate(c *gin.Context) {
	mongo := configureMongo()

	var newDate Date

	if err := (c.ShouldBindJSON(&newDate)); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to bind request body")
	}

	result, err := mongo.Database("dateCalculator").Collection("dates").InsertOne(context.TODO(), newDate)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to add new date")
	}

	c.IndentedJSON(http.StatusOK, result)
}

func RemoveDate(c *gin.Context) {
	mongo := configureMongo()

	var dateToRemove Date

	if err := (c.ShouldBindJSON(&dateToRemove)); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to bind request body")
	}

	result, err := mongo.Database("dateCalculator").Collection("dates").DeleteOne(context.TODO(), bson.D{{Key: "name", Value: dateToRemove.Name}})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to add remove date")
	}

	c.IndentedJSON(http.StatusOK, result)
}

func WipeDatabase(c *gin.Context) {
	mongo := configureMongo()

	if err := mongo.Database("dateCalculator").Drop(context.TODO()); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Did not drop the db")
	}

	if err := mongo.Database("dateCalculator").CreateCollection(context.TODO(), "dates"); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Error creating fresh database")
	}

	c.IndentedJSON(http.StatusOK, "Fresh database ready to go!")
}
