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
	svc := configureMongo()

	dates, err := svc.Database("dateCalculator").Collection("dates").Find(context.TODO(), bson.D{{}})
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

func AddNewDate() {

}

func RemoveDate() {

}

func PopulateDatabase() {

}

func WipeDatabase() {

}