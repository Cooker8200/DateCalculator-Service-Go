package routes

import (
	aws "DateCalculator-Service-Go/data/aws"
	mon "DateCalculator-Service-Go/data/mongo"
	"context"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBHandlerFunc func(*gin.Context, *mongo.Client)

func configureMongo() *mongo.Client {
	var mongoClient *mongo.Client

	envs, err := godotenv.Read(".env")
	if err != nil {
			log.Fatal("Error loading .env file", err)
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(envs["mongo_url"]).SetServerAPIOptions(serverAPI)
	// opts := options.Client().ApplyURI("mongodb://127.0.0.1:27017/")

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

func InitRouter() {
	router := gin.Default()
	mongoClient := configureMongo()

	router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:4200", "http://localhost:3000"},
    AllowMethods:     []string{"GET", "PUT", "DELETE", "OPTIONS"},
    AllowHeaders:     []string{"Origin", "Access-Control-Allow-Origin"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
    AllowOriginFunc: func(origin string) bool {
      return origin == "http://localhost:4200"
    },
    MaxAge: 12 * time.Hour,
  }))

	router.GET("/dates/aws", aws.GetAllDates)
	router.PUT("/dates/aws", aws.AddNewDate)
	router.DELETE("/dates/aws", aws.RemoveDate)

	router.GET("/dates/mongo", func(c *gin.Context) {
		mon.GetAllDates(c, mongoClient)
	})
	router.PUT("/dates/mongo", func(c *gin.Context) {
		mon.AddNewDate(c, mongoClient)
	})
	router.DELETE("/dates/mongo", func(c *gin.Context) {
		mon.RemoveDate(c, mongoClient)
	})
	router.DELETE("/dates/mongo/wipe", func(c *gin.Context) {
		mon.WipeDatabase(c, mongoClient)
	})

	router.Run("localhost:3001")
}
