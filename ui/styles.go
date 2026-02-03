package ui

import "github.com/charmbracelet/lipgloss"

var (
	Box            = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1).Margin(1).Height(8)
	CountryText    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("10")).Align(lipgloss.Center).Margin(2).MarginBottom(0)
	yellowStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("226")).Render
	lightBlueStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("153")).Render
	blueStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("33")).Render
)
