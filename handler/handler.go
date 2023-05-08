// Package handler represents package of actual API implementations
package handler

import (
	"flight-tracking/internal/flighttracker"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	FlightTracker flighttracker.IFlightTracker
}

type calculatePathRequest struct {
	Flights [][]string `json:"flights" binding:"required"`
}

type calculatePathResponse struct {
	Result []string `json:"result"`
}

func (h *Handler) CalculatePath(c *gin.Context) {
	var req calculatePathRequest

	// check if the request json is valid
	if err := c.BindJSON(&req); err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// find the source and destination airport code
	arr, err := h.FlightTracker.FindStartAndEnd(req.Flights)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &calculatePathResponse{
		Result: arr,
	})
}
