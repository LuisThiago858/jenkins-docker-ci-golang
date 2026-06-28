package temperature

type ConverterService struct{}

func NewConverterService() Converter {
	return ConverterService{}
}

func (c ConverterService) ToCelsius(temperature Temperature) Temperature {
	if temperature.Unit() == Celsius {
		return temperature
	}

	value := (temperature.Value() - 32) * 5 / 9

	return NewTemperature(value, Celsius)


func (c ConverterService) ToFahrenheit(temperature Temperature) Temperature {
	if temperature.Unit() == Fahrenheit {
		return temperature
	}

	value := (temperature.Value() * 9 / 5) + 32

	return NewTemperature(value, Fahrenheit)
}