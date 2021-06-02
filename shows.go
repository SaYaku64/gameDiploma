package main

import (
	"net/http"

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
		"version":        generateSessionToken(),
		"loc":            Localization,
		"main":           getLoc("main", GinLabels),
		"back":           getLoc("back", GinLabels),
		"updateCharts":   getLoc("updateCharts", GinLabels),
		"filterCharts":   getLoc("filterCharts", GinLabels),
		"bar":            getLoc("bar", GinLabels),
		"pie":            getLoc("pie", GinLabels),
		"word":           getLoc("word", GinLabels),
		"radar":          getLoc("radar", GinLabels),
		"tech":           getLoc("tech", GinLabels),
		"clear":          getLoc("clear", GinLabels),
		"help":           getLoc("help", GinLabels),
		"infoFirstQ":     getLoc("infoFirstQ", Phrases),
		"infoSecondQ":    getLoc("infoSecondQ", Phrases),
		"resultFirstQ":   getLoc("resultFirstQ", Phrases),
		"resultSecondQ":  getLoc("resultSecondQ", Phrases),
		"resultThirdQ":   getLoc("resultThirdQ", Phrases),
		"resultFourthQ":  getLoc("resultFourthQ", Phrases),
		"resultFifthQ":   getLoc("resultFifthQ", Phrases),
		"resultSixthQ":   getLoc("resultSixthQ", Phrases),
		"resultSeventhQ": getLoc("resultSeventhQ", Phrases),
		"helpTitle":      getLoc("helpTitle", Phrases),
		"help1":          getLoc("help1", Phrases),
		"help2":          getLoc("help2", Phrases),
		"help3":          getLoc("help3", Phrases),
		"help4":          getLoc("help4", Phrases),
		"help5":          getLoc("help5", Phrases),
		"help6":          getLoc("help6", Phrases),
		"help7":          getLoc("help7", Phrases),
		"help8":          getLoc("help8", Phrases),
		"help9":          getLoc("help9", Phrases),
	}
	if IsLoggedIn(c) {
		standard["logout"] = getLoc("logout", GinLabels)
		standard["infographics"] = getLoc("infographics", GinLabels)
		return standard
	}
	standard["signIn"] = getLoc("signIn", GinLabels)
	standard["reg"] = getLoc("reg", GinLabels)
	standard["close"] = getLoc("close", GinLabels)
	standard["email"] = getLoc("email", GinLabels)
	standard["login"] = getLoc("login", GinLabels)
	standard["password"] = getLoc("password", GinLabels)
	standard["remember"] = getLoc("remember", GinLabels)

	return standard
}

// Shows start page
func showIndexPage(c *gin.Context) {
	render(c, "index.html", gin.H{
		"title": "NULESandbox",
	})
}

// Shows infographics page
func showInfographicsPage(c *gin.Context) {
	if IsAsked(c) {
		render(c, "infographics.html", gin.H{
			"title": getLoc("infographicsTitle", Phrases),
		})
	}
	c.Redirect(http.StatusTemporaryRedirect, "/survey")
}

// // Shows infographics page
// func showBarPage(c *gin.Context) {
// 	if IsAsked(c) {
// 		render(c, "infographics.html", gin.H{
// 			"title": getLoc("infographicsTitle", Phrases),
// 		})
// 		allSurveys.GenerateBarChart(c.Writer, c.Request)
// 		httpserver(c.Writer, c.Request)
// 	}
// 	c.Redirect(http.StatusTemporaryRedirect, "/survey")
// }

// Shows survey page
func showSurveyPage(c *gin.Context) {
	if IsAsked(c) {
		render(c, "survey.html", gin.H{
			"title":     getLoc("surveyTitle", Phrases),
			"surThanks": getLoc("surThanks", Phrases),
		})
		return
	}
	render(c, "survey.html", gin.H{
		"choose":    getLoc("choose", GinLabels),
		"other":     getLoc("other", GinLabels),
		"send":      getLoc("send", GinLabels),
		"title":     getLoc("surveyTitle", Phrases),
		"surPhrase": getLoc("surPhrase", Phrases),
		"surConf":   getLoc("surConf", Phrases),
		"firstQ":    getLoc("firstQ", Phrases),
		"firstQ1":   getLoc("firstQ1", Phrases),
		"firstQ2":   getLoc("firstQ2", Phrases),
		"firstQ3":   getLoc("firstQ3", Phrases),
		"firstQ4":   getLoc("firstQ4", Phrases),
		"firstQ5":   getLoc("firstQ5", Phrases),
		"firstQ6":   getLoc("firstQ6", Phrases),
		"secondQ":   getLoc("secondQ", Phrases),
		"secondQ1":  getLoc("secondQ1", Phrases),
		"secondQ2":  getLoc("secondQ2", Phrases),
		"secondQ3":  getLoc("secondQ3", Phrases),
		"secondQ4":  getLoc("secondQ4", Phrases),
		"thirdQ":    getLoc("thirdQ", Phrases),
		"fourthQ":   getLoc("fourthQ", Phrases),
		"fifthQ":    getLoc("fifthQ", Phrases),
		"fifthQ1":   getLoc("fifthQ1", Phrases),
		"fifthQ2":   getLoc("fifthQ2", Phrases),
		"fifthQ3":   getLoc("fifthQ3", Phrases),
		"sixthQ":    getLoc("sixthQ", Phrases),
		"sixthQ1":   getLoc("sixthQ1", Phrases),
		"sixthQ2":   getLoc("sixthQ2", Phrases),
		"sixthQ3":   getLoc("sixthQ3", Phrases),
		"sixthQ4":   getLoc("sixthQ4", Phrases),
		"sixthQ5":   getLoc("sixthQ5", Phrases),
		"seventhQ":  getLoc("seventhQ", Phrases),
		"seventhQ1": getLoc("seventhQ1", Phrases),
		"seventhQ2": getLoc("seventhQ2", Phrases),
		"seventhQ3": getLoc("seventhQ3", Phrases),
		"seventhQ4": getLoc("seventhQ4", Phrases),
		"seventhQ5": getLoc("seventhQ5", Phrases),
	})
}

func showObjectPage(c *gin.Context) {
	tid := c.Params.ByName("name")
	bid := tid
	var locIndex = getLocIndex()

	if IsLoggedIn(c) {
		userBid := GetUserBuildingByTID(c, tid)
		if userBid != "" {
			bid = userBid
		}
	}

	buildingIT, ok := BuildingCache[bid]
	if ok {
		title := buildingIT.Info[locIndex].Name
		description := buildingIT.Info[locIndex].Description

		content := gin.H{
			"title":       title,
			"description": description,
			"details":     getLoc("details", GinLabels),
			"history":     getLoc("history", GinLabels),
			"photo":       getLoc("photo", GinLabels),
			"ddetails":    getLoc(bid, Details),
			"hhistory":    getLoc(bid, History),
			"pphoto":      bid,
		}
		allowed := GetAllowedBuildings(tid)
		if len(allowed) > 0 && IsLoggedIn(c) {
			allowedRender := []AllowedBuildings{}

			for _, build := range allowed {
				allowedRender = append(allowedRender, AllowedBuildings{build.BID, build.Info[locIndex].Name, tid})
			}

			content["allowed"] = allowedRender
			content["changeTerr"] = getLoc("changeTerr", GinLabels)
		}

		render(c, "territory.html", content)
		return
	}

	errorTitle := getLoc("errorTitle", Errors)
	errorText := getLoc("territoryNotFound", Errors)
	render(c, "unfound.html", gin.H{
		"title":       errorTitle,
		"description": errorText,
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
