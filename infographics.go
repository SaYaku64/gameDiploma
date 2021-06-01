package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// allSurveys - cache of all surveys
var allSurveys = AllSurveys{
	Cache: make(map[uint64]Survey),
}

// FillAllSurveys - fills map of all surveys
func (as *AllSurveys) FillAllSurveys() {
	surveys := GetAllSurveys()
	for _, survey := range surveys {
		as.RLock()
		as.Cache[survey.ID] = survey
		as.RUnlock()
	}
}

// UpdateSurveyInMap - updates user in map of users
func (as *AllSurveys) UpdateSurveyInMap(sur Survey) {
	as.RLock()
	as.Cache[sur.ID] = sur
	as.RUnlock()
}

func surveyComplete(c *gin.Context) {
	q1 := c.PostForm("selY")
	q2 := c.PostForm("selU")

	q3 := 5
	q3, _ = strconv.Atoi(c.PostForm("rangeCum"))

	q4 := 5
	q4, _ = strconv.Atoi(c.PostForm("rangeCorp"))

	chk1, _ := strconv.ParseBool(c.PostForm("chk1"))
	chk2, _ := strconv.ParseBool(c.PostForm("chk2"))
	chk3, _ := strconv.ParseBool(c.PostForm("chk3"))
	q5 := [3]bool{chk1, chk2, chk3}

	chk4, _ := strconv.ParseBool(c.PostForm("chk4"))
	chk5, _ := strconv.ParseBool(c.PostForm("chk5"))
	chk6, _ := strconv.ParseBool(c.PostForm("chk6"))
	chk7, _ := strconv.ParseBool(c.PostForm("chk7"))
	chk8, _ := strconv.ParseBool(c.PostForm("chk8"))
	q6 := [5]bool{chk4, chk5, chk6, chk7, chk8}

	q7 := c.PostForm("selD")

	sur := Survey{uint64(len(allSurveys.Cache) + 1), q1, q2, q3, q4, q5, q6, q7}

	if err := allSurveys.AddNewSurvey(sur); err != nil {
		Error.Println("infographics.go -> surveyComplete -> AddNewSurvey: err =", err)
		c.JSON(http.StatusOK, gin.H{
			"error":   true,
			"message": getLoc("internalError", Errors),
		})
		return
	}
	user, ok := GetCurrentUser(c)
	if !ok {
		Error.Println("infographics.go -> surveyComplete -> GetCurrentUser: !ok")
		c.JSON(http.StatusOK, gin.H{
			"error":   true,
			"message": getLoc("unauth", Errors),
		})
		return
	}
	AllUsersMap.UpdateAskedInDB(user, true)
	c.Set("asked", true)

	c.Redirect(http.StatusFound, "/survey")
}

func clearInfo(c *gin.Context) {
	allSurveys.ClearPie()
	allSurveys.ClearBar()
	c.JSON(http.StatusOK, gin.H{
		"message": getLoc("chartClear", Errors),
	})
}

func infographicsShow(c *gin.Context) {
	radio := c.PostForm("radio")

	changeCheck, _ := strconv.ParseBool(c.PostForm("changeCheck"))
	placedCheck, _ := strconv.ParseBool(c.PostForm("placedCheck"))
	check1, _ := strconv.ParseBool(c.PostForm("check1"))
	check2, _ := strconv.ParseBool(c.PostForm("check2"))
	check3, _ := strconv.ParseBool(c.PostForm("check3"))
	check4, _ := strconv.ParseBool(c.PostForm("check4"))
	check5, _ := strconv.ParseBool(c.PostForm("check5"))
	check6, _ := strconv.ParseBool(c.PostForm("check6"))
	check7, _ := strconv.ParseBool(c.PostForm("check7"))

	var checkedStrings []string

	if changeCheck {
		checkedStrings = append(checkedStrings, "infoFirstQ")
	}
	if placedCheck {
		checkedStrings = append(checkedStrings, "infoSecondQ")
	}
	if check1 {
		checkedStrings = append(checkedStrings, "firstQ")
	}
	if check2 {
		checkedStrings = append(checkedStrings, "secondQ")
	}
	if check3 {
		checkedStrings = append(checkedStrings, "thirdQ")
	}
	if check4 {
		checkedStrings = append(checkedStrings, "fourthQ")
	}
	if check5 {
		checkedStrings = append(checkedStrings, "fifthQ")
	}
	if check6 {
		checkedStrings = append(checkedStrings, "sixthQ")
	}
	if check7 {
		checkedStrings = append(checkedStrings, "seventhQ")
	}

	if !changeCheck && !placedCheck && !check1 && !check2 && !check3 && !check4 && !check5 && !check6 && !check7 {
		allSurveys.ClearPie()
		// allSurveys.ClearWord()
		allSurveys.ClearBar()
		c.JSON(http.StatusOK, gin.H{
			"message": getLoc("chartClear", Errors),
		})
		return
	}

	switch radio {
	case "barRadio":
		allSurveys.ClearPie()
		allSurveys.RenderBar(checkedStrings)
	case "pieRadio":
		// allSurveys.ClearWord()
		allSurveys.ClearBar()
		if !check1 && !check2 && !check7 {
			allSurveys.ClearPie()
			c.JSON(http.StatusOK, gin.H{
				"message": getLoc("chartClear", Errors),
			})
			return
		}
		allSurveys.RenderPie(checkedStrings)
	// case "radarRadio":
	// 	allSurveys.ClearPie()
	// 	allSurveys.ClearBar()
	// 	if !check1 && !check2 && !check5 && !check6 && !check7 {
	// 		// allSurveys.ClearWord()
	// 		c.JSON(http.StatusOK, gin.H{
	// 			"message": getLoc("chartClear", Errors),
	// 		})
	// 		return
	// 	}
	// 	allSurveys.RenderWord(checkedStrings)
	default:
		Error.Println("infographics.go -> infographicsShow: Unknown radio")
		c.JSON(http.StatusOK, gin.H{
			"error":   true,
			"message": getLoc("internalError", Errors),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": getLoc("chartOk", Errors),
	})
}
