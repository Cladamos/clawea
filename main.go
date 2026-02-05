package main

import (
	"clawea/ui"
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type keyMap struct {
	Quit key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Quit},
	}
}

type model struct {
	width       int
	height      int
	keys        keyMap
	help        help.Model
	apiErrMsg   string
	weather     weatherMsg
	temps       tempMsg
	loading     bool
	tempLoading bool
}

var keys = keyMap{
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q, ctrl+c", "quit"),
	),
}

func initialModel() *model {
	return &model{
		keys:        keys,
		help:        help.New(),
		loading:     true,
		tempLoading: true,
	}
}
func (m *model) Init() tea.Cmd {
	return getWeather()
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	case apiErrorMsg:
		m.apiErrMsg = msg.message
		return m, nil
	case weatherMsg:
		m.weather = msg
		m.loading = false
		return m, getTempData(m.weather.Location.Lat, m.weather.Location.Lon)
	case tempMsg:
		m.temps = msg
		m.tempLoading = false
		return m, nil
	}
	return m, nil
}

func (m *model) View() string {
	if m.loading {
		return "Loading..."
	}
	helpView := m.help.View(m.keys)

	// -4 width comes from margin (2) and lipgloss add extra (2) characters to width
	currDecodedWeather := ui.WeatherCodeDecoder(m.weather.Current.WeatherCode, m.weather.Current.IsDay == 0)

	countryText := ui.CountryText.Render(m.weather.Location.Region + ", " + m.weather.Location.Country)
	weatherIcon := ui.WeatherIcon.Render(currDecodedWeather.Icon)
	weatherStats := ui.WeatherStats.Render(fmt.Sprintf(
		"Weather:       %s\n"+
			"Temperature:   %.0f°C\n"+
			"Feels Like:    %.0f°C\n"+
			"Humidity:      %d%%\n"+
			"Wind Speed:    %.0f km/h",
		currDecodedWeather.Label,
		m.weather.Current.Temperature,
		m.weather.Current.ApparentTemperature,
		m.weather.Current.Humidity,
		m.weather.Current.WindSpeed,
	))

	var currDayChart string
	if m.tempLoading {
		currDayChart = "Loading..."
	} else {
		if m.width > 80 {
			currDayChart = ui.DrawChart(m.width-60, 7, m.temps.Hourly.Temperatures)
		} else {
			currDayChart = ""
		}
	}
	currDayBox := ui.CurrDayBox.Width(m.width - 4).Render(lipgloss.JoinHorizontal(lipgloss.Center, weatherIcon, ui.CurrDivider, weatherStats, ui.CurrDivider, currDayChart))

	if m.width < 60 {
		currDayBox = ui.CurrDayBox.Width(m.width - 4).Render(lipgloss.Place(m.width-4, 9, lipgloss.Center, lipgloss.Center,
			lipgloss.JoinVertical(lipgloss.Center, weatherIcon, weatherStats)))
	}

	upComingText := ui.UpcomingText.Render("Upcoming Days")

	var upComingDays []string
	// Each card has 18 width with dividers so 5 cards = 90, 3 cards= 54
	maxCards := (m.width - 4) / 18
	if maxCards > 5 {
		maxCards = 5
	}

	// We are skipping first day because it is the current day
	for i := 1; i <= maxCards; i++ {
		date, _ := time.Parse("2006-01-02", m.weather.Daily.Dates[i])
		upComingDays = append(upComingDays, lipgloss.JoinVertical(lipgloss.Center, date.Format("Mon 02"),
			ui.WeatherIcon.Render(ui.WeatherCodeDecoder(m.weather.Daily.WeatherCodes[i], false).Icon),
			fmt.Sprintf("%.0f°C  %.0f°C", m.weather.Daily.MaxTemps[i], m.weather.Daily.MinTemps[i])))
		if i != maxCards {
			upComingDays = append(upComingDays, ui.UpComingDivider)
		}
	}
	//TODO: Make upcoming days box responsive
	upComingDaysRow := lipgloss.JoinHorizontal(lipgloss.Top, upComingDays...)
	var UpcomingDaysBox = ui.UpComingDaysBox.Width(m.width - 4).Render(lipgloss.Place(m.width-4, 9, lipgloss.Center, lipgloss.Center, upComingDaysRow))

	return lipgloss.JoinVertical(lipgloss.Top, countryText, currDayBox, upComingText, UpcomingDaysBox, helpView)
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
