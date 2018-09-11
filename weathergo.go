package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	// 3rd Party Request Package.
	// Link -> https://github.com/parnurzeal/gorequest
	"github.com/parnurzeal/gorequest"
)

const (
	mapboxKey  = "Insert Key Here."
	darkskyKey = "Insert Key Here."
)

func main() {
	darkSky(mapboxRequest(locationInput()))
}

func locationInput() string {
	// Initiating Reader object to read user input.
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Location: ")
	location, _ := reader.ReadString('\n')
	// Return var location for use in mapboxRequest function.
	return location
}

// mapboxRequest func takes input location and returns two float64 vars.
func mapboxRequest(location string) (lat, long float64) {
	mapboxURL := fmt.Sprintf("https://api.mapbox.com/geocoding/v5/mapbox.places/%v.json?access_token=%v", location, mapboxKey)
	// Get Request to Mapbox API.
	_, body, err := gorequest.New().Get(mapboxURL).End()
	if err != nil {
		fmt.Println("Failed to contact Mapbox API: ", err)
	}
	var Location locationObject
	// Create struct and place into variable Location.
	err2 := json.Unmarshal([]byte(body), &Location)
	if err2 != nil {
		fmt.Println("Error: ", err2)
	} else {
		// Obtain lat, long from Location struct stored in Location var.
		lat := Location.Features[0].Center[1]
		long := Location.Features[0].Center[0]
		return lat, long
	}
	return
}

// darkSky func takes two parameters. Lat and long passed in from mapboxRequest.
func darkSky(lat, long float64) {
	// Exclude all other response data execept for current weather.
	// See documentation: https://darksky.net/dev/docs#forecast-request
	excludeStr := "minutely,daily,hourly,alerts,flags"
	darkSkyURL := fmt.Sprintf("https://api.darksky.net/forecast/%v/%v,%v?exclude=%v&units=uk2", darkskyKey, lat, long, excludeStr)
	// Get Request to Dark Sky API.
	_, body, err3 := gorequest.New().Get(darkSkyURL).End()
	if err3 != nil {
		fmt.Println("Request Failed to Dark Sky API: ", err3)
	}
	var Darksky darkskyObject
	// Create struct and place into variable Darksky
	err4 := json.Unmarshal([]byte(body), &Darksky)
	if err4 != nil {
		fmt.Println("Error", err4)
	}
	// Obtain specific data from darksky struct and store in seperate vars.
	summary := Darksky.Currently.Summary
	temperature := Darksky.Currently.Temperature
	humidity := Darksky.Currently.Humidity
	windspeed := Darksky.Currently.WindSpeed
	fmt.Printf(`
Current Weather - %v
Temperature - %v°C
Humidity - %v
Wind Speed - %v Mph

Powered By Dark Sky: https://darksky.net/poweredby/

`, summary, temperature, humidity, windspeed)
}

// Struct for json resp str from Mapbox API.
type locationObject struct {
	Type     string   `json:"type"`
	Query    []string `json:"query"`
	Features []struct {
		ID         string   `json:"id"`
		Type       string   `json:"type"`
		PlaceType  []string `json:"place_type"`
		Relevance  float64  `json:"relevance"`
		Properties struct {
			Wikidata string `json:"wikidata"`
		} `json:"properties"`
		Text      string    `json:"text"`
		PlaceName string    `json:"place_name"`
		Bbox      []float64 `json:"bbox,omitempty"`
		Center    []float64 `json:"center"`
		Geometry  struct {
			Type        string    `json:"type"`
			Coordinates []float64 `json:"coordinates"`
		} `json:"geometry"`
		Context []struct {
			ID        string `json:"id"`
			ShortCode string `json:"short_code"`
			Wikidata  string `json:"wikidata"`
			Text      string `json:"text"`
		} `json:"context"`
	} `json:"features"`
	Attribution string `json:"attribution"`
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
		PrecipIntensity      float64 `json:"precipIntensity"`
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
