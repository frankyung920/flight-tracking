package main

import (
	"flight-tracking/handler"
	"flight-tracking/internal/flighttracker"
	"flight-tracking/routes"
	"net/http"
	"time"
)

func main() {

	flightTracker := flighttracker.FlightTracker{}
	handler := &handler.Handler{
		FlightTracker: &flightTracker,
	}
	router := routes.CreateRouter(handler)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
