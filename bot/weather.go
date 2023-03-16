package bot

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/RahimjonovMuhammadUmar/weather_bot/config"
)

func getCityTemperature(cityName string, cfg config.Config) *float64 {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", cityName, cfg.OpenWeatherApiKey)
	response, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to send request to api.openweathermap.org '%s'", cityName)
		return nil
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		var data map[string]interface{}
		err := json.NewDecoder(response.Body).Decode(&data)
		if err != nil {
			log.Printf("Failed to decode response to data '%s'", cityName)
			return nil
		}
		if temperature, ok := data["main"].(map[string]interface{})["temp"].(float64); ok {
			return &temperature
		} else {
			log.Printf("Failed to parse temperature data from response for city '%s': %v", cityName, data)
			return nil
		}
	} else {
		log.Printf("OpenWeatherMap API returned status code %d for city '%s'", response.StatusCode, cityName)
		return nil
	}
}
