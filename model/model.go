package model

import (
	"github.com/cladamos/clawea/pages"
	"github.com/cladamos/clawea/ui"
	"github.com/cladamos/clawea/weather"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type keyMap struct {
	Quit key.Binding
	Next key.Binding
	Prev key.Binding
}

type model struct {
	width          int
	height         int
	keys           keyMap
	apiErrMsg      weather.ApiErrorMsg
	weather        weather.WeatherMsg
	currDayWeather weather.CurrDayWeatherMsg
	loadingSpinner spinner.Model
	paginator      paginator.Model
	loading        bool
	tempLoading    bool
	page           int
}

var keys = keyMap{
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q, ctrl+c", "quit"),
	),
	Next: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("right, l", "next"),
	),
	Prev: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("left, h", "prev"),
	),
}

func InitialModel() *model {
	p := paginator.New()
	p.Type = paginator.Dots
	p.PerPage = 1
	p.ActiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "235", Dark: "252"}).Render("•")
	p.InactiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "250", Dark: "238"}).Render("•")
	p.SetTotalPages(2)

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Tick()

	return &model{
		keys:           keys,
		loading:        true,
		tempLoading:    true,
		loadingSpinner: s,
		paginator:      p,
		page:           0,
	}
}
func (m *model) Init() tea.Cmd {
	return tea.Batch(m.loadingSpinner.Tick, weather.GetWeather(), weather.TickEveryHour())
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
		case key.Matches(msg, m.keys.Next):
			if m.page < 1 {
				m.page++
				m.paginator.Page++
			}
			return m, nil
		case key.Matches(msg, m.keys.Prev):
			if m.page > 0 {
				m.page--
				m.paginator.Page--
			}
			return m, nil
		}
	case weather.HourlyTickMsg:
		m.loading = true
		return m, tea.Batch(weather.GetWeather(), weather.TickEveryHour())
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.loadingSpinner, cmd = m.loadingSpinner.Update(msg)
		return m, cmd
	case weather.ApiErrorMsg:
		m.apiErrMsg = msg
		return m, nil
	case weather.WeatherMsg:
		m.weather = msg
		m.loading = false
		return m, weather.GetCurrDayWeather(m.weather.Location.Lat, m.weather.Location.Lon)
	case weather.CurrDayWeatherMsg:
		m.currDayWeather = msg
		m.tempLoading = false
		return m, nil
	}

	return m, nil
}

/*
	I really messed up the code while trying to make responsive layout
	I tried to make it understandable with comments but if you don't get it, it counts on me :D
*/

func (m *model) View() string {
	loadingText := lipgloss.JoinHorizontal(lipgloss.Center, m.loadingSpinner.View(), ui.LoadingText.Render(" Loading..."))
	tooSmallText := lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, ui.TooSmallText.Render("Terminal is too small"))

	if m.loading {
		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, loadingText)
	}

	// Terminal is too small to show anything
	if m.height < 14 || m.width < 30 || (m.height < 17 && m.width < 50) {
		return tooSmallText
	}

	currPage := ""
	switch m.page {
	case 0:
		currPage = pages.Overview(m.weather, m.currDayWeather.Hourly.Temperatures, m.width, m.height, m.tempLoading, m.loading, loadingText)
	case 1:
		currPage = pages.DailyStast(m.currDayWeather, m.width, m.height)
	}
	return lipgloss.JoinVertical(lipgloss.Center, currPage, m.paginator.View())
}
