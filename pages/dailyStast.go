package pages

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/cladamos/clawea/ui"
	"github.com/cladamos/clawea/weather"
)

type hourlyWeatherCard struct {
	hour   string
	icon   string
	styled string
}

func createWeatherCards(currDayWeather weather.CurrDayWeatherMsg) []hourlyWeatherCard {
	var cards []hourlyWeatherCard
	var isDay bool

	currHour := time.Now().Hour()
	for i := currHour; i < 24; i++ {
		// I hardcoded 06.00-18.00 as day time
		if i <= 5 || i >= 17 {
			isDay = false
		} else {
			isDay = true
		}
		hour := fmt.Sprintf("%02d:00", i)
		icon := ui.WeatherIcon.Render(ui.WeatherCodeDecoder(currDayWeather.Hourly.WeatherCodes[i], isDay).Icon)
		currCard := hourlyWeatherCard{
			hour:   hour,
			icon:   icon,
			styled: lipgloss.JoinVertical(lipgloss.Center, hour, icon),
		}
		if len(cards) == 0 {
			cards = append(cards, currCard)
		} else {
			lastCard := cards[len(cards)-1]
			// If the last card has same weather code and it's the last hour, merge it
			if (lastCard.icon == currCard.icon) && (i == 23) {
				lastCard.hour = lastCard.hour + " - " + fmt.Sprintf("%02d:00", i)
				lastCard.styled = lipgloss.JoinVertical(lipgloss.Center, lastCard.hour, lastCard.icon)
				cards[len(cards)-1] = lastCard
			}
			if lastCard.icon != currCard.icon {
				// If the last card not one hour ago, start new card
				if lastCard.hour[:2] != fmt.Sprintf("%02d", i-1) {
					// Merge hours that share same weather code
					lastCard.hour = lastCard.hour + " - " + fmt.Sprintf("%02d:00", i-1)
					lastCard.styled = lipgloss.JoinVertical(lipgloss.Center, lastCard.hour, lastCard.icon)
					cards[len(cards)-1] = lastCard
				}
				cards = append(cards, currCard)
			}
		}
	}
	return cards
}

func DailyStast(currDayWeather weather.CurrDayWeatherMsg, width int, height int) string {

	cardDatas := createWeatherCards(currDayWeather)
	// Each card has 18 width with dividers so 5 cards = 90, 3 cards= 54
	maxCards := (width - 4) / 18
	if maxCards > 5 {
		maxCards = 5
	}
	if maxCards > len(cardDatas) {
		maxCards = len(cardDatas)
	}
	var cards []string
	for i := 0; i < maxCards; i++ {
		cards = append(cards, "\n"+cardDatas[i].styled)
		if i != maxCards-1 {
			cards = append(cards, ui.DailyWeatherDivider)
		}
	}
	weatherStatsRow := lipgloss.JoinHorizontal(lipgloss.Top, cards...)
	weatherStatsText := ui.DailyWeatherText.Render("Daily Weather")
	weatherStatsBox := ui.DailyWeatherStatsBox.Width(width - 4).Render(lipgloss.Place(width-4, 9, lipgloss.Center, lipgloss.Center, weatherStatsRow))

	precipitationText := ui.PrecipitationText.Render("Daily Precipitation Probabilities (%)")
	precipitationChart := ui.DrawChart(width-16, 9, currDayWeather.Hourly.PrecipitationProbabilities, "precipitation")
	precipitationInside := ui.DailyPrecipitationBox.Width(width - 4).Render(lipgloss.Place(width-6, 9, lipgloss.Center, lipgloss.Center, precipitationChart))
	precipitationBox := lipgloss.JoinVertical(lipgloss.Top, precipitationText, precipitationInside)

	return lipgloss.JoinVertical(lipgloss.Top, weatherStatsText, weatherStatsBox, precipitationBox)
}
