package pages

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/cladamos/clawea/ui"
	"github.com/cladamos/clawea/weather"
)

func Overview(weather weather.WeatherMsg, temps []float64, width int, height int,
	isVertical bool, isOneBoxLayout bool, tempLoading bool, loading bool, loadingText string) string {

	// -4 width comes from margin (2) and lipgloss add extra (2) characters to width
	currDecodedWeather := ui.WeatherCodeDecoder(weather.Current.WeatherCode, weather.Current.IsDay == 0)
	currDate, _ := time.Parse("2006-01-02T15:04", weather.Current.Date)
	rawCountryText := weather.Location.Region + ", " + weather.Location.Country + " | " + currDate.Format("Mon 02 15:04")
	countryText := ui.CountryText.Render(rawCountryText)
	weatherIcon := ui.WeatherIcon.Render(currDecodedWeather.Icon)
	weatherStats := ui.WeatherStats.Render(fmt.Sprintf(
		"Weather:       %s\n"+
			"Temperature:   %.0f°C\n"+
			"Feels Like:    %.0f°C\n"+
			"Humidity:      %d%%\n"+
			"Wind Speed:    %.0f km/h",
		currDecodedWeather.Label,
		weather.Current.Temperature,
		weather.Current.ApparentTemperature,
		weather.Current.Humidity,
		weather.Current.WindSpeed,
	))

	// Show currDayTempChart if terminal is wide enough
	var currDayTempChart string
	if tempLoading {
		currDayTempChart = loadingText
	} else {
		if width > 75 {
			currDayTempChart = lipgloss.JoinHorizontal(lipgloss.Center, ui.CurrDayDivider, ui.DrawChart(width-60, 7, temps, "temp"))
		} else {
			currDayTempChart = ""
		}
	}
	currDayBoxInside := ui.CurrDayBox.Width(width - 4).Render(lipgloss.JoinHorizontal(lipgloss.Center, weatherIcon, ui.CurrDayDivider, weatherStats, currDayTempChart))

	// Vertical Layout //
	// If terminal is not wide enough, we will show boxes vertically
	// Remove countryText margin
	currDayBox := lipgloss.JoinVertical(lipgloss.Top, countryText, currDayBoxInside)
	if isVertical {
		countryText = ui.CountryText.MarginLeft(0).Render(rawCountryText)
		currDayBoxInside = ui.CurrDayBox.Width(width - 4).Render(lipgloss.Place(width-4, 9, lipgloss.Center, lipgloss.Center,
			lipgloss.JoinVertical(lipgloss.Center, weatherIcon, weatherStats)))
		currDayBox = lipgloss.JoinVertical(lipgloss.Center, countryText, currDayBoxInside)
	}

	// If terminal really really short we will show only the current day box without countryText
	if height < 16 {
		currDayBox = lipgloss.JoinVertical(lipgloss.Center, currDayBoxInside)
		return currDayBox
	}

	upComingDaysBox := ""
	// If terminal tall enough we will show upcoming days
	if !isOneBoxLayout {
		upComingText := ui.UpcomingText.Render("Upcoming Days")
		var upComingDays []string
		// Each card has 18 width with dividers so 5 cards = 90, 3 cards= 54
		maxCards := (width - 4) / 18
		if maxCards > 5 {
			maxCards = 5
		}

		// We are skipping first day because it is the current day
		for i := 1; i <= maxCards; i++ {
			date, _ := time.Parse("2006-01-02", weather.Daily.Dates[i])
			upComingDays = append(upComingDays, lipgloss.JoinVertical(lipgloss.Center, date.Format("Mon 02"),
				ui.WeatherIcon.Render(ui.WeatherCodeDecoder(weather.Daily.WeatherCodes[i], false).Icon),
				fmt.Sprintf("%.0f°C  %.0f°C", weather.Daily.MaxTemps[i], weather.Daily.MinTemps[i])))
			if i != maxCards {
				upComingDays = append(upComingDays, ui.UpComingDayDivider)
			}
		}
		upComingDaysRow := lipgloss.JoinHorizontal(lipgloss.Top, upComingDays...)
		upComingDaysBoxInside := ui.UpComingDaysBox.Width(width - 4).Render(lipgloss.Place(width-4, 9, lipgloss.Center, lipgloss.Center, upComingDaysRow))
		upComingDaysBox = lipgloss.JoinVertical(lipgloss.Top, upComingText, upComingDaysBoxInside)

		// Vertical Layout //
		// If terminal is not wide enough, we will show only the upcoming days box inside
		if isVertical {
			upComingDaysBox = upComingDaysBoxInside
		}
	}
	return lipgloss.JoinVertical(lipgloss.Top, currDayBox, upComingDaysBox)
}
