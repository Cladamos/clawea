package main

import (
	"clawea/ui"
	"fmt"
	"os"

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
	width     int
	height    int
	keys      keyMap
	help      help.Model
	apiErrMsg string
	weather   weatherMsg
	loading   bool
}

var keys = keyMap{
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q, ctrl+c", "quit"),
	),
}

func initialModel() *model {
	return &model{
		keys:    keys,
		help:    help.New(),
		loading: true,
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

	weatherStats := ui.WeatherStats.Render(fmt.Sprintf(
		"Weather:       %s\n"+
			"Temperature:   %.0f°C\n"+
			"Feels Like:    %.0f°C\n"+
			"Humidity:      %d%%\n"+
			"Wind Speed:    %.0f km/h\n",
		currDecodedWeather.Label,
		m.weather.Current.Temperature,
		m.weather.Current.ApparentTemperature,
		m.weather.Current.Humidity,
		m.weather.Current.WindSpeed,
	))
	currDay := ui.Box.Width(m.width - 4).Render(lipgloss.JoinHorizontal(lipgloss.Top, ui.WeatherIcon.Render(currDecodedWeather.Icon), weatherStats))
	return lipgloss.JoinVertical(lipgloss.Top, countryText, currDay, helpView)
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
