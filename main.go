// main.go

package main

import (
	"context"
	logger "logger"

	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var router *gin.Engine

type obj map[string]interface{}

// Localization - var that shows user localization
var Localization string

// Client for MongoDB
var Client = connDB()

// Collection - connect to collection from DB
var Collection = Client.Database(mongoDatabase).Collection(mongoCollUsers)

func main() {

	AllUsersMap.FillAllUsers()
	FillBuildingInfoMap()

	gin.SetMode(gin.ReleaseMode)

	router = gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.Static("/static", "./static")

	initializeRoutes()

	err := router.Run()
	if err != nil {
		logger.Warning.Println(err)
	}

	disconnDB(Client)
}

func render(c *gin.Context, templateName string, data gin.H) {

	data["logged"] = IsLoggedIn(c)

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, AddStandartGin(c, data))
	}
}

func connDB() mongo.Client {
	// Creating DB Client
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoConn))
	if err != nil {
		logger.Warning.Println(err)
	}

	// Connect
	err = client.Connect(context.TODO())
	if err != nil {
		logger.Warning.Println(err)
	}

	// Checking connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		logger.Warning.Println(err)
	}

	return *client
}

func disconnDB(client mongo.Client) {
	// Disconnect
	err := client.Disconnect(context.TODO())
	if err != nil {
		logger.Warning.Println(err)
	}
}
