package services

import "fmt"

func ConvertLength(value float64, fromUnit, toUnit string) (float64, error) {
	valueInMeters, err := toMeters(value, fromUnit)
	if err != nil {
		return 0, err
	}
	result, err := fromMeters(valueInMeters, toUnit)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func toMeters(value float64, unit string) (float64, error) {
	switch unit {
	case "millimeter":
		return value / 1000, nil
	case "centimeter":
		return value / 100, nil
	case "meter":
		return value, nil
	case "kilometer":
		return value * 1000, nil
	case "inch":
		return value * 0.0254, nil
	case "feet":
		return value * 0.3048, nil
	case "yard":
		return value * 0.9144, nil
	case "mile":
		return value * 1609.34, nil
	default:
		return 0, fmt.Errorf("unsupported fromUnit: %s", unit)
	}
}

func fromMeters(value float64, unit string) (float64, error) {
	switch unit {
	case "millimeter":
		return value * 1000, nil
	case "centimeter":
		return value * 100, nil
	case "meter":
		return value, nil
	case "kilometre":
		return value / 1000, nil
	case "inch":
		return value / 0.0254, nil
	case "feett":
		return value / 0.3048, nil
	case "yard":
		return value / 0.9144, nil
	case "mile":
		return value / 1609.34, nil
	default:
		return 0, fmt.Errorf("unsupported toUnit: %s", unit)
	}

}
