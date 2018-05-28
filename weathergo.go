package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	// 3rd Party Request Package.
	// Link -> https://github.com/parnurzeal/gorequest
	"github.com/parnurzeal/gorequest"
)

const (
	g_api_key = ""
	d_api_key = ""
)

func location_input() string {
	// Initiating Reader object to read user input.
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Location: ")
	input, _ := reader.ReadString('\n')
	// Replacing user input " " with + for api.
	// -1 to replace all instances of " "
	location := strings.Replace(input, " ", "+", -1)
	// Return var location for use in google_request function.
	return location
}

// google_request func takes input location and returns two float64 vars.
func google_request(location string) (lat, long float64) {
	g_api_url := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%v&key=%v", location, g_api_key)
	// Get Request to Google API.
	_, body, err := gorequest.New().Get(g_api_url).End()
	if err != nil {
		fmt.Println("Failed to contract Google API: ", err)
	}
	var Location locationObject
	// Create struct and place into variable Weather.
	err2 := json.Unmarshal([]byte(body), &Location)
	if err2 != nil {
		fmt.Println("Error: ", err2)
	} else {
		// Obtain lat, long from Weather struct stored in Weather var.
		lat := Location.Results[0].Geometry.Location.Lat
		long := Location.Results[0].Geometry.Location.Lng
		return lat, long
	}
	return
}

// dark_sky func takes two parameters. Lat and long passed in from google_request.
func dark_sky(lat, long float64) {
	// Exclude all other response data execept for current weather.
	// See documentation: https://darksky.net/dev/docs#forecast-request
	exclude_slice := []string{"minutely", "daily", "hourly", "alerts", "flags"}
	// Create one str with "," joining each item from exclude_slice.
	exclude_no_space := strings.Join(exclude_slice, ",")
	ds_api_url := fmt.Sprintf("https://api.darksky.net/forecast/%v/%v,%v?exclude=%v&units=uk2", d_api_key, lat, long, exclude_no_space)
	// Get Request to Dark Sky API.
	_, body, err3 := gorequest.New().Get(ds_api_url).End()
	if err3 != nil {
		fmt.Println("Request Failed to Dark Sky API: ", err3)
	}
	var Darksky darkskyObject
	// Create struct and place into variable Darksky
	err4 := json.Unmarshal([]byte(body), &Darksky)
	if err4 != nil {
		fmt.Println("Error", err4)
	}
	// Obtain specific data from Weather struct and store in seperate vars.
	summary := Darksky.Currently.Summary
	temperature := Darksky.Currently.Temperature
	humidity := Darksky.Currently.Humidity
	wind_speed := Darksky.Currently.WindSpeed
	fmt.Printf(`
		Current Weather - %v
		Temperature - %v°C
		Humidity - %v
		Wind Speed - %v Mph

`, summary, temperature, humidity, wind_speed)
}

func main() {
	// Location returned from user_input() for usage in make_request()
	dark_sky(google_request(location_input()))
}

// Struct for json resp str from Google API.
type locationObject struct {
	Results []struct {
		AddressComponents []struct {
			LongName  string   `json:"long_name"`
			ShortName string   `json:"short_name"`
			Types     []string `json:"types"`
		} `json:"address_components"`
		FormattedAddress string `json:"formatted_address"`
		Geometry         struct {
			Bounds struct {
				Northeast struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"northeast"`
				Southwest struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"southwest"`
			} `json:"bounds"`
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
			LocationType string `json:"location_type"`
			Viewport     struct {
				Northeast struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"northeast"`
				Southwest struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"southwest"`
			} `json:"viewport"`
		} `json:"geometry"`
		PlaceID string   `json:"place_id"`
		Types   []string `json:"types"`
	} `json:"results"`
	Status string `json:"status"`
}

// Struct for json resp str from Dark Sky API.
type darkskyObject struct {
	Currently struct {
		ApparentTemperature  float64 `json:"apparentTemperature"`
		CloudCover           float64 `json:"cloudCover"`
		DewPoint             float64 `json:"dewPoint"`
		Humidity             float64 `json:"humidity"`
		Icon                 string  `json:"icon"`
		NearestStormBearing  int     `json:"nearestStormBearing"`
		NearestStormDistance int     `json:"nearestStormDistance"`
		Ozone                float64 `json:"ozone"`
		PrecipIntensity      int     `json:"precipIntensity"`
		PrecipProbability    int     `json:"precipProbability"`
		Pressure             float64 `json:"pressure"`
		Summary              string  `json:"summary"`
		Temperature          float64 `json:"temperature"`
		Time                 int     `json:"time"`
		UvIndex              int     `json:"uvIndex"`
		Visibility           float64 `json:"visibility"`
		WindBearing          int     `json:"windBearing"`
		WindGust             float64 `json:"windGust"`
		WindSpeed            float64 `json:"windSpeed"`
	} `json:"currently"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Offset    int     `json:"offset"`
	Timezone  string  `json:"timezone"`
}
