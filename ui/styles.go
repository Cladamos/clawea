package ui

import "github.com/charmbracelet/lipgloss"

var (
	Box          = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1).Margin(1).MarginTop(0).Height(8)
	CountryText  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("10")).Align(lipgloss.Center).MarginLeft(2).MarginTop(1)
	WeatherIcon  = lipgloss.NewStyle().Border(lipgloss.NormalBorder(), false, true, false, false).BorderForeground(lipgloss.Color("244")).PaddingRight(2).Height(6)
	WeatherStats = lipgloss.NewStyle().PaddingLeft(2)

	// Icon Color Styles
	yellowStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("226")).Render
	lightBlueStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("153")).Render
	blueStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("33")).Render
)
