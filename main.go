package main

import (
	"fmt"
	"io"
	"net/http"
)

type Weather struct{
	Location struct{
		Name string `json:"name"`
		Country string `json:"country"`
	}`json:"location"`
	Current struct{
		TempC float64 `json:"temp_c"`
		Condition struct{
			Text string `json:"text"`
		}`json:"condition"`
	}`json:"current"`
	Forecast struct{
	Forecastday []struct{
		Date string `json:"date"`
		Day struct{
			MaxtempC float64 `json:"maxtemp_c"`
			MintempC float64 `json:"mintemp_c"`
			Condition struct{
				Text string `json:"text"`
			}`json:"condition"`
		}`json:"day"`
	}`json:"forecastday"`
	}`json:"forecast"`
}

func main(){
	res,err := http.Get("https://api.weatherapi.com/v1/forecast.json?key=fd43b9fafb65445e9b261906250403&q=Abuja&days=1")
	if err != nil{
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200{
		panic("Weather API not available")
	}

	body,err :=io.ReadAll(res.Body)
	if err != nil{
		panic(err)
	}
	
	fmt.Println(string(body))
}



