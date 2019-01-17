package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	graphmerge "github.com/Codehardt/go-graphmerge"
	chart "github.com/wcharczuk/go-chart"
)

func getTicks(min int, max int) []chart.Tick {
	res := []chart.Tick{}
	for ; min <= max; min++ {
		res = append(res, chart.Tick{float64(min), fmt.Sprint(min)})
	}
	return res
}

func main() {
	// Graph 1
	x1 := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	y1 := []float64{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}
	// Graph 2
	x2 := []float64{0, 2, 4, 6, 8, 10}
	y2 := []float64{0, 1, 2, 3, 4, 5}
	// Graph 3
	x3 := []float64{11, 15, 18, 22}
	y3 := []float64{-1, -2, -3, -4}
	// Merged Graph
	x, y, err := graphmerge.GraphMerge(x1, y1, x2, y2, x3, y3)
	if err != nil {
		panic(err)
	}
	graph := chart.Chart{
		XAxis: chart.XAxis{
			Style: chart.StyleShow(),
			Ticks: getTicks(0, 22),
		},
		YAxis: chart.YAxis{
			Style: chart.StyleShow(),
			Ticks: getTicks(-4, 23),
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Name:    "Graph 1",
				XValues: x1,
				YValues: y1,
				Style: chart.Style{
					Show:     true,
					DotWidth: 3,
				},
			},
			chart.ContinuousSeries{
				Name:    "Graph 2",
				XValues: x2,
				YValues: y2,
				Style: chart.Style{
					Show:     true,
					DotWidth: 3,
				},
			},
			chart.ContinuousSeries{
				Name:    "Graph 3",
				XValues: x3,
				YValues: y3,
				Style: chart.Style{
					Show:     true,
					DotWidth: 3,
				},
			},
			chart.ContinuousSeries{
				Name:    "Merged",
				XValues: x,
				YValues: y,
				Style: chart.Style{
					Show:     true,
					DotWidth: 3,
				},
			},
		},
	}
	graph.Elements = []chart.Renderable{
		chart.Legend(&graph),
	}
	var b bytes.Buffer
	if err := graph.Render(chart.PNG, &b); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile("example.png", b.Bytes(), 0644); err != nil {
		panic(err)
	}
}
