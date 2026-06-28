package main

import (
	"fmt"
	"testing"
)

type WeatherServiceStub struct {
	Temperature int
}

// Forecast сообщает ожидаемую дневную температуру на завтра.
func (ws *WeatherServiceStub) Forecast() int {
	return ws.Temperature
}

type testCase struct {
	deg  int
	want string
}

var tests []testCase = []testCase{
	{-10, "холодно"},
	{0, "холодно"},
	{5, "холодно"},
	{10, "прохладно"},
	{15, "идеально"},
	{20, "жарко"},
}

func TestForecast(t *testing.T) {
	service := &WeatherServiceStub{}
	weather := Weather{service}
	for _, test := range tests {
		name := fmt.Sprintf("%v", test.deg)
		t.Run(name, func(t *testing.T) {
			service.Temperature = test.deg
			got := weather.Forecast()
			if got != test.want {
				t.Errorf("%s: got %s, want %s", name, got, test.want)
			}
		})
	}
}
