package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func showWeatherApp(w fyne.Window) {
	// myApp := app.New()
	//myWin := myApp.NewWindow("Hello")
	//myWin.Resize(fyne.NewSize(500, 300))

	res, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=noida&APPID=5d8992717e219d6ede2c77dd60a8c3b0")

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	weather, err := UnmarshalWeather(body)
	if err != nil {
		fmt.Println(err)
	}

	image := canvas.NewImageFromFile("images\\weather.png")
	image.FillMode = canvas.ImageFillOriginal

	label1 := canvas.NewText("Weather Details", color.Black)
	label1.TextStyle = fyne.TextStyle{Bold: true}
	label1.Alignment = fyne.TextAlignCenter

	label2 := canvas.NewText(fmt.Sprintf("Country: %s", weather.Sys.Country), color.Black)
	label3 := canvas.NewText(fmt.Sprintf("Wind Speed: %.2f km/h", weather.Wind.Speed), color.Black)
	label4 := canvas.NewText(fmt.Sprintf("Min Temp: %.2f", weather.Main.TempMin), color.Black)
	label5 := canvas.NewText(fmt.Sprintf("Max Temp: %.2f", weather.Main.TempMax), color.Black)
	label6 := canvas.NewText(fmt.Sprintf("Humidity: %d", weather.Main.Humidity), color.Black)
	label7 := canvas.NewText(fmt.Sprintf("Pressure: %d", weather.Main.Pressure), color.Black)
	label8 := canvas.NewText(fmt.Sprintf("Sunrise at %d", weather.Sys.Sunrise), color.Black)
	label9 := canvas.NewText(fmt.Sprintf("Sunset at %d", weather.Sys.Sunset), color.Black)

	// combo := widget.NewSelect([]string{"Mumbai", "Delhi", "Chennai"}, func (value string) {
	//      log.Println("Select set to", value)
	// })

	weatherContainer := container.NewVBox(
		label1,
		image,
		// combo,
		label2,
		label3,
		label4,
		label5,
		label6,
		label7,
		label8,
		label9,
	)
	w.SetContent(container.NewBorder(panelContent, nil, nil, nil, weatherContainer))

	w.Show()
}

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    weather, err := UnmarshalWeather(bytes)
//    bytes, err = weather.Marshal()

func UnmarshalWeather(data []byte) (Weather, error) {
	var r Weather
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Weather) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Weather struct {
	Coord      Coord            `json:"coord"`
	Weather    []WeatherElement `json:"weather"`
	Base       string           `json:"base"`
	Main       Main             `json:"main"`
	Visibility int64            `json:"visibility"`
	Wind       Wind             `json:"wind"`
	Clouds     Clouds           `json:"clouds"`
	Dt         int64            `json:"dt"`
	Sys        Sys              `json:"sys"`
	Timezone   int64            `json:"timezone"`
	ID         int64            `json:"id"`
	Name       string           `json:"name"`
	Cod        int64            `json:"cod"`
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int64   `json:"pressure"`
	Humidity  int64   `json:"humidity"`
	SeaLevel  int64   `json:"sea_level"`
	GrndLevel int64   `json:"grnd_level"`
}

type Sys struct {
	Type    int64  `json:"type"`
	ID      int64  `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type WeatherElement struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`
	Gust  float64 `json:"gust"`
}
