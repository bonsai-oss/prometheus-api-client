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
