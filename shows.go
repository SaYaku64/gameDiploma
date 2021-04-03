package main

import (
	"github.com/gin-gonic/gin"
)

// AddStandartGin - adds standart gin values to specialized
func AddStandartGin(temp gin.H) gin.H {
	for k, v := range GetStandartGin() {
		temp[k] = v
	}
	return temp
}

// GetStandartGin - used for adding standart gin values to all pages
func GetStandartGin() gin.H {
	return gin.H{
		"version": generateSessionToken(),
		"loc":     Localization,
	}
}

// Shows start page
func showIndexPage(c *gin.Context) {
	render(c, "index.html", gin.H{
		"title": "NULESandbox",
	})
}

func showObjectPage(c *gin.Context) {
	name := c.Params.ByName("name")

	var locIndex = getLocIndex()

	building, ok := BuildingInfo[name]
	if ok {
		back := Buttons["back"][locIndex]
		title := building[locIndex].Name
		description := building[locIndex].Description

		render(c, "territory.html", gin.H{
			"title":       title,
			"description": description,
			"back":        back,
		})
	} else {
		errorText := Errors["territoryNotFound"][locIndex]
		errorTitle := Errors["errorTitle"][locIndex]
		main := Buttons["main"][locIndex]

		render(c, "unfound.html", gin.H{
			"title":       errorTitle,
			"description": errorText,
			"main":        main,
		})
		return
	}
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
