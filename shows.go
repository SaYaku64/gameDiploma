package main

import (
	"github.com/gin-gonic/gin"
)

// Shows start page
func showIndexPage(c *gin.Context) {
	render(c, gin.H{
		"title":  "Home Page",
		"busket": BusketSlice}, "index.html")
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
