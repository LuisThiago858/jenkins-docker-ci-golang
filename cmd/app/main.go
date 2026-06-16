package main

import (
	"fmt"

	"github.com/LuisThiago858/jenkins-docker-ci-golang/internal/temperature"
)

func main() {
	converter := temperature.NewConverterService()

	fahrenheit := temperature.NewTemperature(32, temperature.Fahrenheit)
	celsius := converter.ToCelsius(fahrenheit)

	fmt.Printf("%.2f °F = %.2f °C\n", fahrenheit.Value(), celsius.Value())

	originalCelsius := temperature.NewTemperature(100, temperature.Celsius)
	convertedFahrenheit := converter.ToFahrenheit(originalCelsius)

	fmt.Printf("%.2f °C = %.2f °F\n", originalCelsius.Value(), convertedFahrenheit.Value())
}