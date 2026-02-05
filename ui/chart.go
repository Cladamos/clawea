package ui

import (
	"fmt"

	"github.com/NimbleMarkets/ntcharts/canvas"
	"github.com/NimbleMarkets/ntcharts/linechart"
	"github.com/charmbracelet/lipgloss"
)

func DrawChart(width int, height int, temps []float64) string {
	var maxTemp, minTemp = temps[0], temps[0]
	for _, v := range temps {
		if v < minTemp {
			minTemp = v
		}
		if v > maxTemp {
			maxTemp = v
		}
	}
	chart := linechart.New(
		width,
		height,
		0,       // minX
		24,      // maxX
		minTemp, // minY
		maxTemp, // maxY
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
	for i := 0; i < len(temps)-1; i++ {
		p1 := canvas.Float64Point{X: float64(i), Y: temps[i]}
		p2 := canvas.Float64Point{X: float64(i + 1), Y: temps[i+1]}
		chart.DrawBrailleLineWithStyle(p1, p2, tempChartColorStyle)
	}

	chart.DrawXYAxisAndLabel()
	legend := tempChartLegendStyle.Render("Daily Temperature")
	chartString := lipgloss.JoinVertical(lipgloss.Top, legend, tempChartStyle.Render(chart.View()))
	return chartString
}
