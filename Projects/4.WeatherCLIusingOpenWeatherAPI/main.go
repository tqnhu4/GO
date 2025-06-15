package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time" // Import time package for timeout
)

// WeatherAPIKey is your OpenWeatherMap API key.
// IMPORTANT: Replace "YOUR_API_KEY" with your actual key.
const WeatherAPIKey = "YOUR_API_KEY"

const BaseURL = "http://api.openweathermap.org/data/2.5/weather"

// Define structs to unmarshal the JSON response from OpenWeatherMap
// You can find the full structure in OpenWeatherMap API documentation
// For simplicity, we'll only unmarshal the fields we need.

type Weather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type Main struct {
	Temp     float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
	Pressure int     `json:"pressure"`
	Humidity int     `json:"humidity"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}

type Sys struct {
	Country string `json:"country"`
	Sunrise int64   `json:"sunrise"`
	Sunset  int64   `json:"sunset"`
}

type WeatherResponse struct {
	Weather []Weather `json:"weather"`
	Main    Main    `json:"main"`
	Wind    Wind    `json:"wind"`
	Name    string  `json:"name"` // City name
	Sys     Sys     `json:"sys"`
	Cod     int     `json:"cod"` // Status code from API (e.g., 200 for success)
	Message string  `json:"message"` // Error message if Cod is not 200
}

func main() {
	if WeatherAPIKey == "YOUR_API_KEY" || WeatherAPIKey == "" {
		fmt.Println("Error: Please replace 'YOUR_API_KEY' with your actual OpenWeatherMap API key in main.go")
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <city_name>")
		fmt.Println("Example: go run main.go London")
		fmt.Println("         go run main.go 'New York'") // Use quotes for city names with spaces
		os.Exit(1)
	}

	cityName := strings.Join(os.Args[1:], " ") // Join arguments to handle city names with spaces

	weatherData, err := fetchWeatherData(cityName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching weather data: %v\n", err)
		os.Exit(1)
	}

	if weatherData.Cod != 200 {
		fmt.Fprintf(os.Stderr, "Error from API (%d): %s\n", weatherData.Cod, weatherData.Message)
		os.Exit(1)
	}

	printWeatherData(weatherData)
}

// fetchWeatherData makes an HTTP GET request to OpenWeatherMap API
func fetchWeatherData(city string) (WeatherResponse, error) {
	// Construct the URL with parameters
	// units=metric for Celsius, imperial for Fahrenheit, default is Kelvin
	// appid is your API key
	queryParams := url.Values{}
	queryParams.Set("q", city)
	queryParams.Set("units", "metric") // Get temperature in Celsius
	queryParams.Set("appid", WeatherAPIKey)

	fullURL := fmt.Sprintf("%s?%s", BaseURL, queryParams.Encode())

	// Create an HTTP client with a timeout
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(fullURL)
	if err != nil {
		return WeatherResponse{}, fmt.Errorf("failed to make HTTP request: %w", err)
	}
	defer resp.Body.Close() // Ensure the response body is closed

	if resp.StatusCode != http.StatusOK {
		// Attempt to read the error message from the API response
		bodyBytes, _ := io.ReadAll(resp.Body)
		var apiError WeatherResponse
		json.Unmarshal(bodyBytes, &apiError)
		if apiError.Message != "" {
			return WeatherResponse{}, fmt.Errorf("API responded with status %d: %s", resp.StatusCode, apiError.Message)
		}
		return WeatherResponse{}, fmt.Errorf("API responded with status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return WeatherResponse{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var weatherResponse WeatherResponse
	err = json.Unmarshal(body, &weatherResponse)
	if err != nil {
		return WeatherResponse{}, fmt.Errorf("failed to unmarshal JSON response: %w", err)
	}

	return weatherResponse, nil
}

// printWeatherData formats and prints the weather information
func printWeatherData(data WeatherResponse) {
	fmt.Printf("Weather in %s, %s:\n", data.Name, data.Sys.Country)
	fmt.Printf("------------------------------\n")
	if len(data.Weather) > 0 {
		fmt.Printf("Condition: %s (%s)\n", data.Weather[0].Main, data.Weather[0].Description)
	}
	fmt.Printf("Temperature: %.1f°C (Feels like: %.1f°C)\n", data.Main.Temp, data.Main.FeelsLike)
	fmt.Printf("Min/Max Temp: %.1f°C / %.1f°C\n", data.Main.TempMin, data.Main.TempMax)
	fmt.Printf("Humidity: %d%%\n", data.Main.Humidity)
	fmt.Printf("Pressure: %d hPa\n", data.Main.Pressure)
	fmt.Printf("Wind Speed: %.1f m/s\n", data.Wind.Speed)

	sunriseTime := time.Unix(data.Sys.Sunrise, 0)
	sunsetTime := time.Unix(data.Sys.Sunset, 0)
	fmt.Printf("Sunrise: %s\n", sunriseTime.Format("03:04 PM")) // Format to HH:MM AM/PM
	fmt.Printf("Sunset:  %s\n", sunsetTime.Format("03:04 PM"))
	fmt.Printf("------------------------------\n")
}