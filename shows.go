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
		"version": generateSessionToken(),
		"loc":     Localization,
		"main":    getLoc("main", GinLabels),
		"back":    getLoc("back", GinLabels),
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
		SurveyMap.GenerateBarChart(c.Writer, c.Request)
		httpserver(c.Writer, c.Request)
	}
	c.Redirect(http.StatusTemporaryRedirect, "/survey")
}

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
