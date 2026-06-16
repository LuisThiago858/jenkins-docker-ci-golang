package temperature

import "testing"

func TestShouldConvertFahrenheitToCelsius(t *testing.T) {
	converter := NewConverterService()

	input := NewTemperature(32, Fahrenheit)
	result := converter.ToCelsius(input)

	expectedValue := 0.0

	if result.Value() != expectedValue {
		t.Errorf("expected %.2f, got %.2f", expectedValue, result.Value())
	}

	if result.Unit() != Celsius {
		t.Errorf("expected unit %s, got %s", Celsius, result.Unit())
	}
}

func TestShouldConvertBoilingPointFromFahrenheitToCelsius(t *testing.T) {
	converter := NewConverterService()

	input := NewTemperature(212, Fahrenheit)
	result := converter.ToCelsius(input)

	expectedValue := 100.0

	if result.Value() != expectedValue {
		t.Errorf("expected %.2f, got %.2f", expectedValue, result.Value())
	}
}

func TestShouldConvertCelsiusToFahrenheit(t *testing.T) {
	converter := NewConverterService()

	input := NewTemperature(0, Celsius)
	result := converter.ToFahrenheit(input)

	expectedValue := 32.0

	if result.Value() != expectedValue {
		t.Errorf("expected %.2f, got %.2f", expectedValue, result.Value())
	}

	if result.Unit() != Fahrenheit {
		t.Errorf("expected unit %s, got %s", Fahrenheit, result.Unit())
	}
}

func TestShouldConvertBoilingPointFromCelsiusToFahrenheit(t *testing.T) {
	converter := NewConverterService()

	input := NewTemperature(100, Celsius)
	result := converter.ToFahrenheit(input)

	expectedValue := 212.0

	if result.Value() != expectedValue {
		t.Errorf("expected %.2f, got %.2f", expectedValue, result.Value())
	}
}