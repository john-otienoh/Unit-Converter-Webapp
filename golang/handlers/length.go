package handlers

import (
	"html/template"
	"main/services"
	"net/http"
	"strconv"
)

var lenghtPageTemplate = template.Must(
	template.ParseFiles("templates/length.html", "templates/base.html"),
)

type LenghtPageData struct {
	Value     float64
	FromUnit  string
	ToUnit    string
	Result    float64
	ErrorMsg  string
	Submitted bool
}

func LengthHandler(w http.ResponseWriter, r *http.Request) {
	data := LenghtPageData{}

	if r.Method == http.MethodPost {
		r.ParseForm()
		valueStr := r.FormValue("value")
		fromUnit := r.FormValue("from")
		toUnit := r.FormValue("to")
		val, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			data.ErrorMsg = "Invalid input Value"
		} else {
			result, convErr := services.ConvertLength(val, fromUnit, toUnit)
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
	lenghtPageTemplate.Execute(w, data)
}
