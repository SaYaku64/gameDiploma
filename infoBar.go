package main

import (
	"math/rand"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// var Question1Strs []string = []string{"abit", "bak1", "bak3", "mah", "no", "tutor"}

// // generate random data for bar chart
// func (as *AllSurveys) GenerateBarItems(strs []string) []opts.BarData {

// 	items := make([]opts.BarData, 0)
// 	temp := make(map[string]int)
// 	for _, s := range as.Cache {
// 		temp[s.Question1] = temp[s.Question1] + 1
// 	}
// 	for _, str := range strs {
// 		items = append(items, opts.BarData{Value: temp[str]})
// 	}

// 	return items
// }

// // func (as *AllSurveys) CountDifferentValue(arr []string, cr string) int {
// // 	dict := make(map[string]int)
// // 	for _, str := range arr {
// // 		dict[str] = dict[str] + 1
// // 	}
// // 	return dict[cr]
// // }
// func (as *AllSurveys) GenerateBarChart(w http.ResponseWriter, _ *http.Request) {
// 	// create a new line instance
// 	bar := charts.NewBar()
// 	// set some global options like Title/Legend/ToolTip or anything else
// 	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
// 		Title: getLoc("firstQ", Phrases),
// 		// Subtitle: "It's extremely easy to use, right?",
// 	}))

// 	// Put data into instance
// 	bar.SetXAxis(Question1Strs).
// 		AddSeries(getLoc("firstQ1", Phrases), allSurveys.GenerateBarItems(Question1Strs))
// 	bar.Render(w)
// }

// var (
// 	itemCnt = 7
// 	weeks   = []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
// )

// func generateBarItems() []opts.BarData {
// 	items := make([]opts.BarData, 0)
// 	for i := 0; i < itemCnt; i++ {
// 		items = append(items, opts.BarData{Value: rand.Intn(300)})
// 	}
// 	return items
// }

// func (as *AllSurveys) ClearBar() {
// 	os.Remove("templates/bar.html")
// 	os.Create("templates/bar.html")
// 	router.LoadHTMLGlob("templates/*")
// }

// func (as *AllSurveys) GenBar() {
// 	bar := charts.NewBar()
// 	bar.SetGlobalOptions(
// 		charts.WithTitleOpts(opts.Title{Title: "label options"}),
// 	)
// 	bar.SetXAxis(weeks).
// 		AddSeries("Category A", generateBarItems()).
// 		AddSeries("Category B", generateBarItems()).
// 		SetSeriesOptions(
// 			charts.WithLabelOpts(opts.Label{
// 				Show:     true,
// 				Position: "top",
// 			}),
// 		)

// 	f, _ := os.Create("templates/bar.html")
// 	bar.Render(f)
// 	router.LoadHTMLGlob("templates/*")
// }

var (
	infoFirstQAns  = []string{"Площа перед корпусом №1", "Студентський сквер", "Кулиничі біля корпусу №3", "Стадіон", "Паркова зона"}
	infoSecondQAns = []string{"Алея сакури", "Кінотеатр", "Торгово-розважальний центр", "Пустка", "Ашан"}
	firstQAns      = []string{getLoc("firstQ1", Phrases), getLoc("firstQ2", Phrases), getLoc("firstQ3", Phrases), getLoc("firstQ4", Phrases), getLoc("firstQ5", Phrases), getLoc("firstQ6", Phrases)}
	secondQAns     = []string{getLoc("secondQ1", Phrases), getLoc("secondQ2", Phrases), getLoc("secondQ3", Phrases), getLoc("secondQ4", Phrases)}
	fourthQAns     = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	fifthQAns      = []string{getLoc("fifthQ1", Phrases), getLoc("fifthQ2", Phrases), getLoc("fifthQ3", Phrases)}
	sixthQAns      = []string{getLoc("sixthQ1", Phrases), getLoc("sixthQ2", Phrases), getLoc("sixthQ3", Phrases), getLoc("sixthQ4", Phrases), getLoc("sixthQ5", Phrases)}
	seventhQAns    = []string{getLoc("seventhQ1", Phrases), getLoc("seventhQ2", Phrases), getLoc("seventhQ3", Phrases), getLoc("seventhQ4", Phrases), getLoc("seventhQ5", Phrases)}
)

func (as *AllSurveys) ClearBar() {
	os.Remove("templates/bar.html")
	os.Create("templates/bar.html")
	router.LoadHTMLGlob("templates/*")
}

func generateBarItems(vars []string) []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < len(vars); i++ {
		items = append(items, opts.BarData{Value: rand.Intn(300)})
	}
	return items
}

func (as *AllSurveys) GenBar(str string, vars []string) *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:      getLoc(str, Phrases) + getLoc("barBasic", Phrases),
			TitleStyle: &opts.TextStyle{FontStyle: "bold"},
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "infographic",
		}),
	)

	barItems := generateBarItems(vars)

	bar.SetXAxis(vars).
		AddSeries("", barItems).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:     true,
				Position: "top",
			}),
		)
	return bar
}

func (as *AllSurveys) RenderBar(bars []string) {
	f, _ := os.Create("templates/bar.html")
	for _, barStr := range bars {
		switch barStr {
		case "infoFirstQ":
			barBasic := as.GenBar(barStr, infoFirstQAns)
			barBasic.Render(f)
		case "infoSecondQ":
			barBasic := as.GenBar(barStr, infoSecondQAns)
			barBasic.Render(f)
		case "firstQ":
			barBasic := as.GenBar(barStr, firstQAns)
			barBasic.Render(f)
		case "secondQ":
			barBasic := as.GenBar(barStr, secondQAns)
			barBasic.Render(f)
		case "thirdQ":
			barBasic := as.GenBar(barStr, fourthQAns)
			barBasic.Render(f)
		case "fourthQ":
			barBasic := as.GenBar(barStr, fourthQAns)
			barBasic.Render(f)
		case "fifthQ":
			barBasic := as.GenBar(barStr, fifthQAns)
			barBasic.Render(f)
		case "sixthQ":
			barBasic := as.GenBar(barStr, sixthQAns)
			barBasic.Render(f)
		case "seventhQ":
			barBasic := as.GenBar(barStr, seventhQAns)
			barBasic.Render(f)
		}
	}
	router.LoadHTMLGlob("templates/*")
}
