package main

import (
	"github.com/gin-gonic/gin"
)

// Shows start page
func showIndexPage(c *gin.Context) {
	render(c, "index.html", gin.H{
		"title":   "NULESandbox",
		"version": generateSessionToken(),
	})
}

func showObjectPage(c *gin.Context) {
	name := c.Params.ByName("name")

	var object ObjectInfoType
	var ok bool
	var back string
	var main string
	var errorTitle string
	var errorText string

	if Localization == "UA" {
		object, ok = ObjectInfo[name]
		if ok {
			back = Buttons["back"]
		} else {
			errorText = Errors["objectNotFound"]
			errorTitle = Errors["errorTitle"]
			main = Buttons["main"]
		}
	} else {
		object, ok = ObjectInfoEN[name]
		if ok {
			back = ButtonsEN["back"]
		} else {
			errorText = ErrorsEN["objectNotFound"]
			errorTitle = ErrorsEN["errorTitle"]
			main = ButtonsEN["main"]
		}
	}

	if !ok {
		render(c, "unfound.html", gin.H{
			"title":       errorTitle,
			"description": errorText,
			"main":        main,
		})
		return
	}

	render(c, "object.html", gin.H{
		"title":       object.Name,
		"description": object.Description,
		"back":        back,
	})
}

// // Shows Conversation page (user articles)
// func showConversationPage(c *gin.Context) {
// 	articles := getArticleFromDB()

// 	// Call the render function with the name of the template to render
// 	render(c, gin.H{
// 		"title":   "Conversation",
// 		"payload": articles,
// 		"busket":  BusketSlice}, "conversation.html")
// }
