package main

import (
	"encoding/json"
	"net/http"

	tea "github.com/charmbracelet/bubbletea"
)

type locationMsg struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type apiErrorMsg struct {
	message string
}

func getLocation() tea.Cmd {
	const url = "http://ip-api.com/json/?fields=status,message,lat,lon"
	return func() tea.Msg {
		res, err := http.Get(url)
		if err != nil {
			return apiErrorMsg{message: "Request failed: " + err.Error()}
		}
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			return apiErrorMsg{message: "Error: Received status code " + res.Status}
		}
		var location locationMsg
		if err := json.NewDecoder(res.Body).Decode(&location); err != nil {
			return apiErrorMsg{message: "Failed to parse response: " + err.Error()}
		}
		return locationMsg{Lat: location.Lat, Lon: location.Lon}
	}
}
