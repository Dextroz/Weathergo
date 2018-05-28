# Weathergo
A clone of my project WeatherTerm; written in Go (https://github.com/Dextroz/WeatherTerm)

## Dependencies
Weathergo is written in Golang so it is **REQUIRED**.

A working 1.10.2 [Go](https://golang.org/dl/) environment.

Weathergo requires the following dependencies:
  1. [GoRequest](https://github.com/parnurzeal/gorequest) - GoRequest -- Simplified HTTP client (inspired by nodejs SuperAgent)

## Prerequisites
To operate Weathergo you must:

  1. Obtain a [Google Geolocation API key](https://developers.google.com/maps/documentation/geolocation/intro).

  2. Obtain a [Dark Sky API Key](https://darksky.net/dev).

  3. Place **Google Geolocation API key** and **Dark Sky API** key in the respective consts found in [weathergo.go](weathergo.go):
      ```
      const (
      	g_api_key = "Insert Key Here."
      	d_api_key = "Insert Key Here."
      )
      ```
## Download Options -- Installing
Currently you can only clone or download the project ZIP file. (Recommended clone if you're going to be contributing.)

Extract and navigate to the zipfile directory and run Weathergo by executing the main entry point file (weathergo.go):
  ```
  go run weathergo.go
  ```
## Authors -- Contributors

* **Daniel Brennand** - *Author* - [Dextroz](https://github.com/Dextroz)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) for details.

## Acknowledgments
* GoRequest created by [Theeraphol Wattanavekin (parnurzeal)](https://github.com/parnurzeal) and respective contributors.
