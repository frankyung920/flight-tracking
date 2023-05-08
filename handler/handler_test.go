package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockFlightTracker struct {
	mock.Mock
}

func (m *mockFlightTracker) FindStartAndEnd(flights [][]string) ([]string, error) {
	args := m.Called(flights)
	return args.Get(0).([]string), args.Error(1)
}

func TestHandler_CalculatePath(t *testing.T) {
	// Define test data
	tests := []struct {
		name             string
		requestBody      []byte
		expectedStatus   int
		expectedResponse calculatePathResponse
		expectedError    error
		mockFlightData   [][]string
	}{
		{
			name: "success",
			requestBody: []byte(`{
				"flights": [
					["JFK", "SFO"],
					["SFO", "ATL"]
				]
			}`),
			expectedStatus: http.StatusOK,
			expectedResponse: calculatePathResponse{
				Result: []string{"JFK", "SFO", "ATL"},
			},
			expectedError: nil,
			mockFlightData: [][]string{
				{"JFK", "SFO"},
				{"SFO", "ATL"},
			},
		},
		{
			name: "error request",
			requestBody: []byte(`{
				"flights": [
					["JFK", "SFO"],
					["SFO", "ATL"],
					["ATL", "GSO", "IND"]
				]
			}`),
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: calculatePathResponse{},
			expectedError:    errors.New("some error"),
			mockFlightData: [][]string{
				{"JFK", "SFO"},
				{"SFO", "ATL"},
				{"ATL", "GSO", "IND"},
			},
		},
		{
			name: "bad request",
			requestBody: []byte(`{
				"badReq": [
					["JFK", "SFO"],
					["SFO", "ATL"],
					["ATL", "GSO", "IND"]
				]
			}`),
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: calculatePathResponse{},
			expectedError:    errors.New("some error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockFlightTracker := &mockFlightTracker{}
			mockFlightTracker.On("FindStartAndEnd", tt.mockFlightData).Return(tt.expectedResponse.Result, tt.expectedError)

			req, _ := http.NewRequest(http.MethodPost, "/calculate", bytes.NewBuffer(tt.requestBody))
			rr := httptest.NewRecorder()

			ctx, _ := gin.CreateTestContext(rr)
			ctx.Request = req

			handler := &Handler{FlightTracker: mockFlightTracker}

			handler.CalculatePath(ctx)

			assert.Equal(t, tt.expectedStatus, rr.Code)
			var actualResponse calculatePathResponse
			json.Unmarshal(rr.Body.Bytes(), &actualResponse)
			assert.Equal(t, tt.expectedResponse, actualResponse)
		})
	}
}
