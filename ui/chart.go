package ui

import (
	"fmt"

	"github.com/NimbleMarkets/ntcharts/canvas"
	"github.com/NimbleMarkets/ntcharts/linechart"
	"github.com/charmbracelet/lipgloss"
)

func DrawChart(width int, height int, data []float64, chartType string) string {
	var maxData, minData = data[0], data[0]
	for _, v := range data {
		if v < minData {
			minData = v
		}
		if v > maxData {
			maxData = v
		}
	}
	chart := linechart.New(
		width,
		height,
		0,       // minX
		24,      // maxX
		minData, // minY
		maxData, // maxY
		linechart.WithXLabelFormatter(func(_ int, v float64) string {
			currHour := int(v)
			// Add 24 hour label without overlapping the last label
			if currHour == 23 {
				return "24"
			}
			// Add labels every 4 hours
			if currHour%4 == 0 {
				return fmt.Sprintf("%02d", currHour)
			}
			return ""
		}),
	)
	for i := 0; i < len(data)-1; i++ {
		p1 := canvas.Float64Point{X: float64(i), Y: data[i]}
		p2 := canvas.Float64Point{X: float64(i + 1), Y: data[i+1]}
		if chartType == "temp" {
			chart.DrawBrailleLineWithStyle(p1, p2, tempChartColorStyle)
		}
		if chartType == "precipitation" {
			chart.DrawBrailleLineWithStyle(p1, p2, precipitationChartColorStyle)
		}
	}

	chart.DrawXYAxisAndLabel()

	var chartString string
	if chartType == "temp" {
		legend := tempChartLegendStyle.Render("Daily Temperature")
		chartString = lipgloss.JoinVertical(lipgloss.Top, legend, paddingChartStyle.Render(chart.View()))
	}
	if chartType == "precipitation" {
		chartString = paddingChartStyle.Render(chart.View())
	}
	return chartString
}
