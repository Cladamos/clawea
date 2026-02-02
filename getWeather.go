package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	tea "github.com/charmbracelet/bubbletea"
)

type location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type daily struct {
	WeatherCodes []int     `json:"weather_code"`
	MinTemps     []float64 `json:"temperature_2m_min"`
	MaxTemps     []float64 `json:"temperature_2m_max"`
}

type weatherMsg struct {
	Daily daily
}

type apiErrorMsg struct {
	message string
}

func getLocation() (lat float64, lon float64, e error) {
	const url = "http://ip-api.com/json/?fields=status,message,lat,lon"
	res, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return 0, 0, err
	}
	var location location
	if err := json.NewDecoder(res.Body).Decode(&location); err != nil {
		return 0, 0, err
	}
	return location.Lat, location.Lon, nil
}

func GetWeather() tea.Cmd {
	lat, lon, ipApiErr := getLocation()
	baseURL := "https://api.open-meteo.com/v1/forecast"
	params := url.Values{}
	params.Add("latitude", fmt.Sprintf("%f", lat))
	params.Add("longitude", fmt.Sprintf("%f", lon))
	params.Add("daily", "weather_code,temperature_2m_max,temperature_2m_min")
	params.Add("forecast_days", "4")

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	return func() tea.Msg {
		if ipApiErr != nil {
			return apiErrorMsg{message: "Request failed: " + ipApiErr.Error()}
		}

		fmt.Println(fullURL)
		res, err := http.Get(fullURL)
		if err != nil {
			return apiErrorMsg{message: "Request failed: " + err.Error()}
		}
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			return apiErrorMsg{message: "Error: Received status code " + res.Status}
		}
		var weather weatherMsg
		if err := json.NewDecoder(res.Body).Decode(&weather); err != nil {
			return apiErrorMsg{message: "Failed to parse response: " + err.Error()}
		}
		return weatherMsg{Daily: weather.Daily}
	}
}
