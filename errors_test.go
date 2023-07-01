package prometheus_test

import (
	"testing"

	"github.com/bonsai-oss/prometheus-api-client"
)

// TestErrNoEndpoint tests the Error() method of the ErrNoEndpoint type.
func TestErrNoEndpoint(t *testing.T) {
	expectedErrorMessage := "No endpoint specified"
	err := prometheus.ErrNoEndpoint{}

	if err.Error() != expectedErrorMessage {
		t.Errorf("Expected error message '%s', got '%s'", expectedErrorMessage, err.Error())
	}
}

// TestErrInvalidResponse tests the Error() method of the ErrInvalidResponse type.
func TestErrInvalidResponse(t *testing.T) {
	expectedErrorMessage := "Invalid response from server: 404 \"Not Found\""
	err := prometheus.ErrInvalidResponse{StatusCode: 404, Message: "Not Found"}

	if err.Error() != expectedErrorMessage {
		t.Errorf("Expected error message '%s', got '%s'", expectedErrorMessage, err.Error())
	}
}
