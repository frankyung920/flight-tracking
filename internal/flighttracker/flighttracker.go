// Package flighttracker represents package flight tracker functionality
package flighttracker

import (
	"flight-tracking/internal/constant"
	"fmt"
	"log"
)

const (
	START = "START"
	END   = "END"
)

type IFlightTracker interface {
	FindStartAndEnd(routes [][]string) ([]string, error)
}

type FlightTracker struct {
}

func (fr *FlightTracker) FindStartAndEnd(routes [][]string) ([]string, error) {
	starts := make([]string, 0)
	ends := make([]string, 0)

	for _, v := range routes {

		// check if route of a flight has source and destination airport code only
		if len(v) != 2 {
			return nil, fmt.Errorf(constant.InvalidFlightRoute.String(), "abnormal route in flight routes")
		}

		starts = append(starts, v[0])
		ends = append(ends, v[1])
	}
	// find airport code exist in end but not in start
	diffStart := fr.findDiff(ends, starts)
	log.Printf("Source airport code(s) does not exist in destination airport codes: %v", diffStart)
	if len(diffStart) != 1 {
		return nil, fmt.Errorf(constant.InvalidFlightRoute.String(), "more than 1 start / no start found")
	}

	// find airport code exist in ends but not in starts
	diffEnd := fr.findDiff(starts, ends)
	log.Printf("Destination airport code(s) does not exist in source airport codes: %v", diffEnd)
	if len(diffEnd) != 1 {
		return nil, fmt.Errorf(constant.InvalidFlightRoute.String(), "more than 1 end / no end found")
	}

	return []string{diffStart[0], diffEnd[0]}, nil
}

func (fr *FlightTracker) findDiff(arr1, arr2 []string) []string {
	arr := make([]string, 0)
	m := make(map[string]bool, len(arr1))
	for _, v := range arr1 {
		m[v] = true
	}
	for _, v := range arr2 {
		if !m[v] {
			arr = append(arr, v)
		}
	}
	return arr
}
