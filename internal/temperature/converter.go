package temperature

type Unit string

const (
	Celsius    Unit = "CELSIUS"
	Fahrenheit Unit = "FAHRENHEIT"
)

type Temperature struct {
	value float64
	unit  Unit
}

func NewTemperature(value float64, unit Unit) Temperature {
	return Temperature{
		value: value,
		unit:  unit,
	}
}

func (t Temperature) Value() float64 {
	return t.value
}

func (t Temperature) Unit() Unit {
	return t.unit
}

type Converter interface {
	ToCelsius(temperature Temperature) Temperature
	ToFahrenheit(temperature Temperature) Temperature
}