package main

import (
	"log"
	"main/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/length", handlers.LengthHandler)
	http.HandleFunc("/weight", handlers.WeigthHandler)
	http.HandleFunc("/temperature", handlers.TemperatureHandler)
	log.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
