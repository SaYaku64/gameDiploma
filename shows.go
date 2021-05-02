package main

import (
	"github.com/gin-gonic/gin"
)

// AddStandartGin - adds standart gin values to specialized
func AddStandartGin(c *gin.Context, temp gin.H) gin.H {
	for k, v := range GetStandartGin(c) {
		temp[k] = v
	}
	return temp
}

// GetStandartGin - used for adding standart gin values to all pages
func GetStandartGin(c *gin.Context) gin.H {
	standard := gin.H{
		"version": generateSessionToken(),
		"loc":     Localization,
		"main":    getLoc("main", GinLabels),
		"back":    getLoc("back", GinLabels),
	}
	if IsLoggedIn(c) {
		standard["logout"] = getLoc("logout", GinLabels)
	} else {
		standard["signIn"] = getLoc("signIn", GinLabels)
		standard["reg"] = getLoc("reg", GinLabels)
		standard["close"] = getLoc("close", GinLabels)
		standard["email"] = getLoc("email", GinLabels)
		standard["login"] = getLoc("login", GinLabels)
		standard["password"] = getLoc("password", GinLabels)
		standard["remember"] = getLoc("remember", GinLabels)
	}
	return standard
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
		title := building[locIndex].Name
		description := building[locIndex].Description

		render(c, "territory.html", gin.H{
			"title":       title,
			"description": description,
		})
	} else {
		errorTitle := getLoc("errorTitle", Errors)
		errorText := getLoc("territoryNotFound", Errors)

		render(c, "unfound.html", gin.H{
			"title":       errorTitle,
			"description": errorText,
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
