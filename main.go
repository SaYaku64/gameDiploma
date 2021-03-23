// main.go

package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var router *gin.Engine

// Localization - var that shows user localization
var Localization string

// Client for MongoDB
var Client = connDB()

func main() {

	gin.SetMode(gin.ReleaseMode)

	router = gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.Static("/static", "./static")

	initializeRoutes()

	Localization = "UA" // change with cookie

	err := router.Run()
	if err != nil {
		log.Println(err)
	}

	disconnDB(Client)
}

func render(c *gin.Context, templateName string, data gin.H) {
	loggedInInterface, ok := c.Get("is_logged_in")
	if ok != true {
		log.Println("Error in getting 'is_logged_in' parameter")
		return
	}
	data["is_logged_in"], ok = loggedInInterface.(bool)
	if ok != true {
		log.Println("LoggedInInterface isn't bool")
		return
	}

	// adminnedInterface, ok := c.Get("adminned")
	// if ok != true {
	// 	log.Println("Error in getting 'adminned' parameter")
	// 	return
	// }
	// data["adminned"], ok = adminnedInterface.(bool)
	// if ok != true {
	// 	log.Println("adminnedInterface isn't bool")
	// 	return
	// }

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}

func connDB() mongo.Client {
	// Creating DB Client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Println(err)
	}

	// Connect
	err = client.Connect(context.TODO())
	if err != nil {
		log.Println(err)
	}

	// Checking connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println(err)
	}

	return *client
}

func disconnDB(client mongo.Client) {
	// Disconnect
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Println(err)
	}
}
