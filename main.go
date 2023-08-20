package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Current Current `json:"current"`
}

type Current struct {
	Temp float64 `json:"temp_c"`
}

func main() {
	t := time.NewTicker(time.Minute * 10)
	for {
		select {
		case <-t.C:
			writeToFile(getWeather())
		}
	}
}

func writeToFile(resp Response) {
	file, err := os.Create("/data/weather.txt")
	if err != nil {
		log.Printf("Create => %s", err)
	}
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			log.Printf("file.Close => %s", err)
		}
	}(file)

	output := fmt.Sprintf("%.1f°", math.Round(resp.Current.Temp))
	if _, err := file.WriteString(output); err != nil {
		log.Printf("WriteString => %s", err)
	}
}

func getWeather() Response {
	query := "http://api.weatherapi.com/v1/current.json?key=" + os.Getenv("API_Key") +
		"&q=" + os.Getenv("CITY") + "&aqi=no"
	resp, err := http.Get(query)
	if err != nil {
		log.Printf("Get=> %s", err)
	}
	if resp.StatusCode != 200 {
		log.Printf("StatusCode => %s", resp.Status)
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			log.Printf("Body.Close => %s", err)
		}
	}(resp.Body)

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Printf("Decode => %s", err)
	}
	return response
}
