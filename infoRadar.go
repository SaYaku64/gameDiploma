package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"os"

// 	"github.com/go-echarts/go-echarts/v2/charts"
// 	"github.com/go-echarts/go-echarts/v2/opts"
// )

// var (
// 	wcDataFirst = map[string]interface{}{
// 		getLoc("firstQ1", Phrases): rand.Intn(10000-5000+1) + 5000,
// 		getLoc("firstQ2", Phrases): rand.Intn(10000-5000+1) + 5000,
// 		getLoc("firstQ3", Phrases): rand.Intn(10000-5000+1) + 5000,
// 		getLoc("firstQ4", Phrases): rand.Intn(10000-5000+1) + 5000,
// 		getLoc("firstQ5", Phrases): rand.Intn(10000-5000+1) + 5000,
// 		getLoc("firstQ6", Phrases): rand.Intn(10000-5000+1) + 5000,
// 	}

// 	wcDataSecond = map[string]interface{}{
// 		getLoc("secondQ1", Phrases): rand.Intn(10000-5000+1) + 5000,
// 		getLoc("secondQ2", Phrases): rand.Intn(10000-5000+1) + 5000,
// 		getLoc("secondQ3", Phrases): rand.Intn(10000-5000+1) + 5000,
// 		getLoc("secondQ4", Phrases): rand.Intn(10000-5000+1) + 5000,
// 	}

// 	wcDataSeventh = map[string]interface{}{
// 		getLoc("seventhQ1", Phrases): rand.Intn(10000-5000+1) + 5000,
// 		getLoc("seventhQ2", Phrases): rand.Intn(10000-5000+1) + 5000,
// 		getLoc("seventhQ3", Phrases): rand.Intn(10000-5000+1) + 5000,
// 		getLoc("seventhQ4", Phrases): rand.Intn(10000-5000+1) + 5000,
// 		getLoc("seventhQ5", Phrases): rand.Intn(10000-5000+1) + 5000,
// 	}
// )

// func generateWCData(data map[string]interface{}) (items []opts.RadarData) {
// 	items = make([]opts.RadarData, 0)
// 	for k, v := range data {
// 		items = append(items, opts.RadarData{Name: k, Value: v})
// 	}
// 	fmt.Println(items)
// 	return
// }

// func (as *AllSurveys) ClearRadar() {
// 	os.Remove("templates/radar.html")
// 	os.Create("templates/radar.html")
// 	router.LoadHTMLGlob("templates/*")
// }

// func (as *AllSurveys) GenRadar(str string, vars map[string]interface{}) *charts.Radar {
// 	wc := charts.NewRadar()
// 	wc.SetGlobalOptions(
// 		charts.WithTitleOpts(opts.Title{
// 			Title:      getLoc(str, Phrases) + getLoc("radar", Phrases),
// 			TitleStyle: &opts.TextStyle{FontStyle: "bold"},
// 		}),
// 		charts.WithInitializationOpts(opts.Initialization{
// 			Theme: "infographic",
// 		}),
// 	)

// 	wc.AddSeries("radar", generateWCData(vars)).
// 		SetSeriesOptions(
// 			charts.WithWorldCloudChartOpts(
// 				opts.RadarChart{
// 					SizeRange: []float32{18, 80},
// 					Shape:     "cardioid",
// 				}),
// 		)
// 	return wc
// }

// func (as *AllSurveys) RenderRadar(radars []string) {
// 	f, _ := os.Create("templates/radar.html")
// 	for _, radarStr := range radars {
// 		switch radarStr {
// 		case "infoFirstQ":
// 			radar := as.GenRadar(radarStr, wcDataInfoFirst)
// 			radar.Render(f)
// 		case "infoSecondQ":
// 			radar := as.GenRadar(radarStr, wcDataInfoSecond)
// 			radar.Render(f)
// 		case "firstQ":
// 			radar := as.GenRadar(radarStr, wcDataFirst)
// 			radar.Render(f)
// 		case "secondQ":
// 			radar := as.GenRadar(radarStr, wcDataSecond)
// 			radar.Render(f)
// 		case "seventhQ":
// 			radar := as.GenRadar(radarStr, wcDataSeventh)
// 			radar.Render(f)
// 		}
// 	}
// 	router.LoadHTMLGlob("templates/*")
// }
