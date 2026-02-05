package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	Box             = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1).Margin(1).MarginTop(0).Height(9)
	CurrDayBox      = Box.BorderForeground(lipgloss.Color("10"))
	UpComingDaysBox = Box.BorderForeground(lipgloss.Color("12"))

	TitleText    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("10")).Align(lipgloss.Center).MarginLeft(2).MarginTop(1)
	CountryText  = TitleText.Foreground(lipgloss.Color("10"))
	UpcomingText = TitleText.Foreground(lipgloss.Color("12")).MarginTop(0)

	Divider         = lipgloss.NewStyle().MarginLeft(1).MarginRight(1).Height(9)
	CurrDivider     = Divider.Foreground(lipgloss.Color("10")).Render(strings.Repeat("│\n", 8) + "│")
	UpComingDivider = Divider.Foreground(lipgloss.Color("12")).Render(strings.Repeat("│\n", 8) + "│")

	WeatherIcon  = lipgloss.NewStyle().Height(5).Width(15)
	WeatherStats = lipgloss.NewStyle().Padding(0, 2).Height(5).BorderForeground(lipgloss.Color("10"))

	// Icon Color Styles
	yellowStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("226")).Render
	lightBlueStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("153")).Render
	blueStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("33")).Render

	// Chart Styles
	tempChartStyle       = lipgloss.NewStyle().PaddingLeft(2)
	tempChartLegendStyle = lipgloss.NewStyle().PaddingLeft(2).PaddingBottom(1).Foreground(lipgloss.Color("226"))
	tempChartColorStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("226"))
)
