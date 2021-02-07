package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
	"log"
	"encoding/json"
)

var summaries = [10]string{"Freezing", "Bracing", "Chilly", "Cool", "Mild", "Warm", "Balmy", "Hot", "Sweltering", "Scorching"}

type WeatherForecast struct {
	Date         time.Time
	TemperatureC int
	TemperatureF int
	Summary      string
}

func (w *WeatherForecast) weatherDiscovery() {
	s1 := rand.NewSource(time.Now().UnixNano())
	random := rand.New(s1)
	randomNumber := random.Intn(10)
	randomTemperature := random.Intn(50)

	w.Summary = summaries[randomNumber]
	w.Date = time.Now()
	w.TemperatureC = randomTemperature
	w.TemperatureF = 32 + int(float64(w.TemperatureC)/0.5556)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func handlerHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func handlerWeatherForecast(w http.ResponseWriter, r *http.Request) {
	weather := WeatherForecast{}
	weather.weatherDiscovery()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(weather)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", handlerHelloWorld)
	http.HandleFunc("/weather", handlerWeatherForecast)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
