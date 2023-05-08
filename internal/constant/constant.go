// Package constant represents error code constants
package constant

type ErrorCode int

const (
	_ ErrorCode = iota
	InvalidFlightRoute
)

var toString = map[ErrorCode]string{
	InvalidFlightRoute: "invalid flight route: %s",
}

func (e ErrorCode) String() string {
	return toString[e]
}
