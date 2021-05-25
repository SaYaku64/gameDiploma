package main

import (
	"logger"
	"net/http"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/types"

	"github.com/gin-gonic/gin"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var Question1Strs []string = []string{"abit", "bak1", "bak3", "mah", "no", "tutor"}

// generate random data for bar chart
func (as *AllSurveys) GenerateBarItems(strs []string) []opts.BarData {

	items := make([]opts.BarData, 0)
	temp := make(map[string]int)
	for _, s := range as.Cache {
		temp[s.Question1] = temp[s.Question1] + 1
	}
	for _, str := range strs {
		items = append(items, opts.BarData{Value: temp[str]})
	}

	return items
}

// func (as *AllSurveys) CountDifferentValue(arr []string, cr string) int {
// 	dict := make(map[string]int)
// 	for _, str := range arr {
// 		dict[str] = dict[str] + 1
// 	}
// 	return dict[cr]
// }
func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	items = append(items, opts.LineData{Value: 0})
	items = append(items, opts.LineData{Value: 0})
	items = append(items, opts.LineData{Value: 1})
	items = append(items, opts.LineData{Value: 0})
	items = append(items, opts.LineData{Value: 2})
	items = append(items, opts.LineData{Value: 2})
	items = append(items, opts.LineData{Value: 2})
	items = append(items, opts.LineData{Value: 4})
	items = append(items, opts.LineData{Value: 3})
	items = append(items, opts.LineData{Value: 7})
	items = append(items, opts.LineData{Value: 10})
	return items
}

func httpserver(w http.ResponseWriter, _ *http.Request) {
	// create a new line instance
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title: "Оцінка території університету",
		}))

	// Put data into instance
	line.SetXAxis([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}).
		AddSeries("Частота", generateLineItems()).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	line.Render(w)
}

func (as *AllSurveys) GenerateBarChart(w http.ResponseWriter, _ *http.Request) {
	// create a new line instance
	bar := charts.NewBar()
	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: getLoc("firstQ", Phrases),
		// Subtitle: "It's extremely easy to use, right?",
	}))

	// Put data into instance
	bar.SetXAxis(Question1Strs).
		AddSeries(getLoc("firstQ1", Phrases), SurveyMap.GenerateBarItems(Question1Strs))
	bar.Render(w)
}

// SurveyMap - cache of all surveys
var SurveyMap = AllSurveys{
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

	sur := Survey{uint64(len(SurveyMap.Cache) + 1), q1, q2, q3, q4, q5, q6}

	if err := SurveyMap.AddNewSurvey(sur); err != nil {
		logger.Error.Println("infographics.go -> surveyComplete -> AddNewSurvey: err =", err)
		c.JSON(http.StatusOK, gin.H{
			"error":   true,
			"message": getLoc("internalError", Errors),
		})
		return
	}
	user, ok := GetCurrentUser(c)
	if !ok {
		logger.Error.Println("infographics.go -> surveyComplete -> GetCurrentUser: !ok")
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
