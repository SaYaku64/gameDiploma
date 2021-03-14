package main

import (
	"github.com/gin-gonic/gin"
)

// Shows start page
func showIndexPage(c *gin.Context) {
	render(c, gin.H{
		"title":   "Home Page",
		"version": generateSessionToken(),
		"tiles0":  TilesSlice[0],
		"tiles1":  TilesSlice[1],
		"tiles2":  TilesSlice[2],
		"tiles3":  TilesSlice[3],
		"tiles4":  TilesSlice[4],
		"tiles5":  TilesSlice[5],
		"tiles6":  TilesSlice[6],
		"tiles7":  TilesSlice[7],
		"tiles8":  TilesSlice[8],
		"tiles9":  TilesSlice[9],
	}, "index.html")
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
