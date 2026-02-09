package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type location struct {
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Country string  `json:"country"`
	Region  string  `json:"regionName"`
}

type daily struct {
	WeatherCodes []int     `json:"weather_code"`
	MinTemps     []float64 `json:"temperature_2m_min"`
	MaxTemps     []float64 `json:"temperature_2m_max"`
	Dates        []string  `json:"time"`
}

type current struct {
	WeatherCode         int     `json:"weather_code"`
	Temperature         float64 `json:"temperature_2m"`
	ApparentTemperature float64 `json:"apparent_temperature"`
	Humidity            int     `json:"relative_humidity_2m"`
	WindSpeed           float64 `json:"wind_speed_10m"`
	IsDay               int     `json:"is_day"`
	Date                string  `json:"time"`
}

type WeatherMsg struct {
	Daily    daily   `json:"daily"`
	Current  current `json:"current"`
	Location location
}

type hourly struct {
	WeatherCodes               []int     `json:"weather_code"`
	PrecipitationProbabilities []float64 `json:"precipitation_probability"`
	Temperatures               []float64 `json:"temperature_2m"`
}

type CurrDayWeatherMsg struct {
	Hourly hourly `json:"hourly"`
}

type ApiErrorMsg struct {
	message string
}
type HourlyTickMsg struct {
}

func getLocation() (loc location, e error) {
	const url = "http://ip-api.com/json/?fields=status,message,country,regionName,lat,lon"
	res, err := http.Get(url)
	if err != nil {
		return location{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return location{}, err
	}
	var currLocation location
	if err := json.NewDecoder(res.Body).Decode(&currLocation); err != nil {
		return location{}, err
	}
	return currLocation, nil
}

func GetWeather() tea.Cmd {
	return func() tea.Msg {
		location, ipApiErr := getLocation()
		baseURL := "https://api.open-meteo.com/v1/forecast"
		params := url.Values{}
		params.Add("latitude", fmt.Sprintf("%f", location.Lat))
		params.Add("longitude", fmt.Sprintf("%f", location.Lon))
		params.Add("daily", "weather_code,temperature_2m_max,temperature_2m_min")
		params.Add("current", "weather_code,temperature_2m,relative_humidity_2m,is_day,apparent_temperature,wind_speed_10m")
		params.Add("forecast_days", "6")

		fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
		if ipApiErr != nil {
			return ApiErrorMsg{message: "Request failed: " + ipApiErr.Error()}
		}

		res, err := http.Get(fullURL)
		if err != nil {
			return ApiErrorMsg{message: "Request failed: " + err.Error()}
		}
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			return ApiErrorMsg{message: "Error: Received status code " + res.Status}
		}
		var weather WeatherMsg
		if err := json.NewDecoder(res.Body).Decode(&weather); err != nil {
			return ApiErrorMsg{message: "Failed to parse response: " + err.Error()}
		}
		return WeatherMsg{Daily: weather.Daily, Current: weather.Current, Location: location}
	}
}

func GetCurrDayWeather(lat, lon float64) tea.Cmd {
	return func() tea.Msg {
		baseURL := "https://api.open-meteo.com/v1/forecast"
		params := url.Values{}
		params.Add("latitude", fmt.Sprintf("%f", lat))
		params.Add("longitude", fmt.Sprintf("%f", lon))
		params.Add("hourly", "temperature_2m,weather_code,precipitation_probability")
		params.Add("timezone", "auto")
		params.Add("forecast_days", "1")

		fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

		res, err := http.Get(fullURL)
		if err != nil {
			return ApiErrorMsg{message: "Request failed: " + err.Error()}
		}
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			return ApiErrorMsg{message: "Error: Received status code " + res.Status}
		}
		var currDayWeather CurrDayWeatherMsg
		if err := json.NewDecoder(res.Body).Decode(&currDayWeather); err != nil {
			return ApiErrorMsg{message: "Failed to parse response: " + err.Error()}
		}
		return CurrDayWeatherMsg{Hourly: currDayWeather.Hourly}
	}
}

func TickEveryHour() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(time.Hour)
		return HourlyTickMsg{}
	}
}
