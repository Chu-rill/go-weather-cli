package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch   int64   `json:"time_epoch"`
				TempC       float64 `json:"temp_c"`
				Condition   struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	q := "Abuja"
	if len(os.Args) >= 2 {
		q = os.Args[1]
	}

	// Fetch weather data
	res, err := http.Get("https://api.weatherapi.com/v1/forecast.json?key=fd43b9fafb65445e9b261906250403&q=" + q + "&days=1")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	// Parse JSON response
	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	// Extract relevant data
	location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour
	fmt.Printf("📍 %s, %s: 🌡️ %.1f°C, ☁️ %s\n", location.Name, location.Country, current.TempC, current.Condition.Text)

	// Loop through hourly forecast
	fmt.Println("📅 Hourly Forecast:")
	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)

		// Skip past hours
		if date.Before(time.Now()) {
			continue
		}

		message := fmt.Sprintf("🕒 %s - 🌡️ %.1f°C, ☔ %.0f%%, ☁️ %s",
			date.Format("15:04"),
			hour.TempC,
			hour.ChanceOfRain,
			hour.Condition.Text,
		)

		// Colorize output based on rain chance
		if hour.ChanceOfRain < 40 {
			fmt.Println(message)
		} else {
			color.Red(message) // Requires "github.com/fatih/color"
		}
	}
}
