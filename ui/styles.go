package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	Box                   = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1).Margin(0, 1).Height(9)
	CurrDayBox            = Box.BorderForeground(lipgloss.Color("10"))
	UpComingDaysBox       = Box.BorderForeground(lipgloss.Color("12"))
	DailyWeatherStatsBox  = Box.BorderForeground(lipgloss.Color("11"))
	DailyPrecipitationBox = Box.BorderForeground(lipgloss.Color("14"))

	HelpView = lipgloss.NewStyle().Align(lipgloss.Center).Height(1)

	TitleText         = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("10")).Align(lipgloss.Center).MarginLeft(2).MarginTop(1)
	CountryText       = TitleText.Foreground(lipgloss.Color("10"))
	UpcomingText      = TitleText.Foreground(lipgloss.Color("12"))
	DailyWeatherText  = TitleText.Foreground(lipgloss.Color("11"))
	PrecipitationText = TitleText.Foreground(lipgloss.Color("14"))
	LoadingText       = TitleText.Foreground(lipgloss.Color("10"))
	TooSmallText      = TitleText.Foreground(lipgloss.Color("10"))

	Divider             = lipgloss.NewStyle().MarginLeft(1).MarginRight(1).Height(9)
	CurrDayDivider      = Divider.Foreground(lipgloss.Color("10")).Render(strings.Repeat("│\n", 8) + "│")
	UpComingDayDivider  = Divider.Foreground(lipgloss.Color("12")).Render(strings.Repeat("│\n", 8) + "│")
	DailyWeatherDivider = Divider.Foreground(lipgloss.Color("11")).Render(strings.Repeat("│\n", 8) + "│")

	WeatherIcon  = lipgloss.NewStyle().Height(5).Width(15)
	WeatherStats = lipgloss.NewStyle().Padding(0, 2).Height(5).BorderForeground(lipgloss.Color("10"))

	// Icon Color Styles
	yellowStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("226")).Render
	lightBlueStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("153")).Render
	blueStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("33")).Render

	// Chart Styles
	paddingChartStyle = lipgloss.NewStyle().Padding(0, 2)
	legendStyle       = lipgloss.NewStyle().Padding(0, 2).PaddingBottom(1)

	tempChartLegendStyle = legendStyle.Foreground(lipgloss.Color("226"))
	tempChartColorStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("226"))

	precipitationChartColorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("14"))
)
