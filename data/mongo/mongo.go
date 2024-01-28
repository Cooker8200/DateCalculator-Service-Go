package mongo

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func configureMongo() *mongo.Client{
	var mongoClient *mongo.Client

	envs, err := godotenv.Read(".env")
	if err != nil {
			log.Fatal("Error loading .env file")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(envs["mongo_url"]).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	mongoClient = client

	return mongoClient
	// clientOptions := options.Client().ApplyURI(envs["mongo_url"])
	// client, err := mongo.Connect(context.TODO(), clientOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = client.Ping(context.TODO(), nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func GetAllDates(c *gin.Context) {
	svc := configureMongo()
log.Print("made it here")
	dates := svc.Database("dates")

	log.Print(dates)
	
	c.IndentedJSON(http.StatusOK, "did something, but still working on it")
}

func AddNewDate() {

}

func RemoveDate() {

}

func PopulateDatabase() {

}

func WipeDatabase() {

}