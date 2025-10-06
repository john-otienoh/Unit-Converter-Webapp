package services

import "fmt"

func ConvertWeight(value float64, fromUnit, toUnit string) (float64, error) {
	valueInKg, err := toKg(value, fromUnit)
	if err != nil {
		return 0, err
	}
	result, err := fromKg(valueInKg, toUnit)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func toKg(value float64, unit string) (float64, error) {
	switch unit {
	case "mg":
		return value / 1_000_000, nil
	case "g":
		return value / 1000, nil
	case "kg":
		return value, nil
	case "oz":
		return value * 0.0283495, nil
	case "lb":
		return value * 0.453592, nil
	default:
		return 0, fmt.Errorf("unsupported fromUnit: %s", unit)
	}
}

func fromKg(value float64, unit string) (float64, error) {
	switch unit {
	case "mg":
		return value * 1_000_000, nil
	case "g":
		return value * 1000, nil
	case "kg":
		return value, nil
	case "oz":
		return value / 0.0283495, nil
	case "lb":
		return value / 0.453592, nil
	default:
		return 0, fmt.Errorf("unsupported toUnit: %s", unit)
	}
}
