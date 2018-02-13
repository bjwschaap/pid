package main

import (
	"github.com/wcharczuk/go-chart"
	"bytes"
	"fmt"
	"io/ioutil"
)

func makeChart(x, y, ys []float64) error {
	graph := chart.Chart{
		XAxis: chart.XAxis{
			ValueFormatter: func(v interface{}) string {
				if vf, isFloat := v.(float64); isFloat {
					return fmt.Sprintf("%0.f", vf)
				}
				return ""
			},
			Style:     chart.StyleShow(),
		},
		YAxis: chart.YAxis{
			Style:     chart.StyleShow(),
			ValueFormatter: func(v interface{}) string {
				if vf, isFloat := v.(float64); isFloat {
					return fmt.Sprintf("%0.f", vf)
				}
				return ""
			},
		},
		YAxisSecondary: chart.YAxis{
			Style:     chart.StyleShow(),
			ValueFormatter: func(v interface{}) string {
				if vf, isFloat := v.(float64); isFloat {
					return fmt.Sprintf("%0.f", vf)
				}
				return ""
			},
		},
		Background: chart.Style{
			Padding: chart.Box{
				Top:  20,
				Left: 20,
			},
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Name: "Input",
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.ColorBlue,
					StrokeWidth: 0.75,
					FillColor:   chart.ColorBlue.WithAlpha(40),
				},
				XValues: x,
				YValues: y,
			},
			chart.ContinuousSeries{
				Name: "Output",
				YAxis: chart.YAxisSecondary,
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.ColorRed,
					StrokeWidth: 0.75,
					FillColor:   chart.ColorRed.WithAlpha(15),
				},
				XValues: x,
				YValues: ys,
			},
		},
	}
	graph.Elements = []chart.Renderable{
		chart.LegendThin(&graph),
	}
	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("chart.png", buffer.Bytes(), 0644)
	return err
}
