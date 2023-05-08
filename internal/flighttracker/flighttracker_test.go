package flighttracker

import (
	"errors"
	"reflect"
	"testing"
)

func TestFlightTracker_FindStartAndEnd(t *testing.T) {
	tests := []struct {
		name   string
		fr     *FlightTracker
		routes [][]string
		want   []string
		err    error
	}{
		{
			name: "valid input with one start and one end",
			routes: [][]string{
				{"SFO", "LAX"},
				{"LAX", "JFK"},
				{"JFK", "BOS"},
			},
			want: []string{"SFO", "BOS"},
			err:  nil,
		},
		{
			name: "valid input with multiple routes",
			routes: [][]string{
				{"SFO", "LAX"},
				{"LAX", "JFK"},
				{"JFK", "BOS"},
				{"BOS", "MIA"},
				{"MIA", "ORD"},
			},
			want: []string{"SFO", "ORD"},
			err:  nil,
		},
		{
			name: "invalid input with abnormal route",
			routes: [][]string{
				{"SFO", "LAX"},
				{"LAX", "JFK"},
				{"JFK", "BOS", "MIA"},
			},
			want: nil,
			err:  errors.New("invalid flight route: abnormal route in flight routes"),
		},
		{
			name: "invalid input with no start found",
			routes: [][]string{
				{"JFK", "LAX"},
				{"LAX", "JFK"},
				{"JFK", "BOS"},
				{"BOS", "MIA"},
			},
			want: nil,
			err:  errors.New("invalid flight route: more than 1 start / no start found"),
		},
		{
			name: "invalid input with multiple starts found",
			routes: [][]string{
				{"SFO", "LAX"},
				{"ATL", "JFK"},
				{"JFK", "BOS"},
				{"LAX", "MIA"},
			},
			want: nil,
			err:  errors.New("invalid flight route: more than 1 start / no start found"),
		},
		{
			name: "invalid input with no end found",
			routes: [][]string{
				{"SFO", "LAX"},
				{"LAX", "LAX"},
				{"LAX", "LAX"},
			},
			want: nil,
			err:  errors.New("invalid flight route: more than 1 end / no end found"),
		},
		{
			name: "invalid input with multiple ends found",
			routes: [][]string{
				{"SFO", "LAX"},
				{"LAX", "JFK"},
				{"LAX", "MIA"},
			},
			want: nil,
			err:  errors.New("invalid flight route: more than 1 end / no end found"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fr := &FlightTracker{}
			got, err := fr.FindStartAndEnd(tt.routes)
			if err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("FlightTracker.FindStartAndEnd() error = %v, wantErr %v", err, tt.err)
					return
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlightTracker.FindStartAndEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
