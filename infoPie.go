package main

import (
	"math/rand"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var (
	answerVariantsFirst   = []string{getLoc("firstQ1", Phrases), getLoc("firstQ2", Phrases), getLoc("firstQ3", Phrases), getLoc("firstQ4", Phrases), getLoc("firstQ5", Phrases), getLoc("firstQ6", Phrases)}
	answerVariantsSecond  = []string{getLoc("secondQ1", Phrases), getLoc("secondQ2", Phrases), getLoc("secondQ3", Phrases), getLoc("secondQ4", Phrases)}
	answerVariantsSeventh = []string{getLoc("seventhQ1", Phrases), getLoc("seventhQ2", Phrases), getLoc("seventhQ3", Phrases), getLoc("seventhQ4", Phrases), getLoc("seventhQ5", Phrases)}
)

func generatePieItems(vars []string) []opts.PieData {
	items := make([]opts.PieData, 0)
	for i := 0; i < len(vars); i++ {
		items = append(items, opts.PieData{Name: vars[i], Value: rand.Intn(100)})
	}
	return items
}

func (as *AllSurveys) ClearPie() {
	os.Remove("templates/pie.html")
	os.Create("templates/pie.html")
	router.LoadHTMLGlob("templates/*")
}

func (as *AllSurveys) GenPie(str string, vars []string) (*charts.Pie, *charts.Pie) {
	///////////
	pieBasic := charts.NewPie()
	pieBasic.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:      getLoc(str, Phrases) + getLoc("pieBasic", Phrases),
			TitleStyle: &opts.TextStyle{FontStyle: "bold"},
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "infographic",
		}),
	)

	pieItems := generatePieItems(vars)

	pieBasic.AddSeries("pie", pieItems).
		SetSeriesOptions(charts.WithLabelOpts(
			opts.Label{
				Show:      true,
				Formatter: "{b}: {c}",
			}),
		)
	///////////
	pieRose := charts.NewPie()
	pieRose.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:      getLoc(str, Phrases) + getLoc("pieRose", Phrases),
			TitleStyle: &opts.TextStyle{FontStyle: "bold"},
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "infographic",
		}),
	)

	pieRose.AddSeries("pie", pieItems).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:      true,
				Formatter: "{b}: {c}",
			}),
			charts.WithPieChartOpts(opts.PieChart{
				Radius:   []string{"25%", "80%"},
				RoseType: "area",
			}),
		)
	return pieBasic, pieRose
}

func (as *AllSurveys) RenderPie(pies []string) {
	f, _ := os.Create("templates/pie.html")
	for _, pieStr := range pies {
		switch pieStr {
		case "firstQ":
			pieBasic, pieRose := as.GenPie(pieStr, answerVariantsFirst)
			pieBasic.Render(f)
			pieRose.Render(f)
		case "secondQ":
			pieBasic, pieRose := as.GenPie(pieStr, answerVariantsSecond)
			pieBasic.Render(f)
			pieRose.Render(f)
		case "seventhQ":
			pieBasic, pieRose := as.GenPie(pieStr, answerVariantsSeventh)
			pieBasic.Render(f)
			pieRose.Render(f)
		}
	}
	router.LoadHTMLGlob("templates/*")
}
