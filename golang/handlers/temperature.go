package handlers

import (
	"html/template"
	"main/services"
	"net/http"
	"strconv"
)

var tempPageTemplate = template.Must(template.ParseFiles("templates/temperature.html", "templates/base.html"))

type TempPageData struct {
	Value     float64
	FromUnit  string
	ToUnit    string
	Result    float64
	ErrorMsg  string
	Submitted bool
}

func TemperatureHandler(w http.ResponseWriter, r *http.Request) {
	data := TempPageData{}

	if r.Method == http.MethodPost {
		r.ParseForm()
		valueStr := r.FormValue("value")
		fromUnit := r.FormValue("from")
		toUnit := r.FormValue("to")

		val, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			data.ErrorMsg = "Invalid input value"
		} else {
			result, convErr := services.ConvertTemperature(val, fromUnit, toUnit)
			if convErr != nil {
				data.ErrorMsg = convErr.Error()
			} else {
				data.Result = result
				data.Value = val
				data.FromUnit = fromUnit
				data.ToUnit = toUnit
				data.Submitted = true
			}
		}
	}

	tempPageTemplate.Execute(w, data)
}
