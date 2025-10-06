package services

import "fmt"

func ConvertTemperature(value float64, fromUnit, toUnit string) (float64, error) {
	var celsius float64

	switch fromUnit {
	case "Celsius":
		celsius = value
	case "Fahrenheit":
		celsius = (value - 32) * 5 / 9
	case "Kelvin":
		celsius = value - 273.15
	default:
		return 0, fmt.Errorf("unsupported fromUnit: %s", fromUnit)
	}

	switch toUnit {
	case "Celsius":
		return celsius, nil
	case "Fahrenheit":
		return (celsius*9/5 + 32), nil
	case "Kelvin":
		return celsius + 273.15, nil
	default:
		return 0, fmt.Errorf("unsupported toUnit: %s", toUnit)
	}
}
