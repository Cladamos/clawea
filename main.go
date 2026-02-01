package main

import (
	"clawea/ui"
	"fmt"
	"os"
	"strconv"

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
	weather   string
	apiErrMsg string
	lat       float64
	lon       float64
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
		weather: "Loading...",
	}
}
func (m *model) Init() tea.Cmd {
	return getLocation()
}

type weatherMsg struct{ weather string }

func checkWeather() tea.Cmd {
	return func() tea.Msg {
		return weatherMsg{weather: "Sunny"}
	}
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
	case locationMsg:
		m.lat = msg.Lat
		m.lon = msg.Lon
		return m, nil
	}
	return m, nil
}

func (m *model) View() string {
	helpView := m.help.View(m.keys)
	// -4 width comes from margin (2) and lipgloss add extra (2) characters to width
	box1 := ui.Box.Width(m.width - 4).Render(ui.Thunderstorm)
	return lipgloss.JoinVertical(lipgloss.Top, box1, helpView, m.apiErrMsg, strconv.FormatFloat(m.lat, 'f', 6, 64), strconv.FormatFloat(m.lon, 'f', 6, 64))
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
