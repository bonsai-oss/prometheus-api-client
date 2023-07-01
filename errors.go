package prometheus

import (
	"fmt"
)

type ErrNoEndpoint struct{}

func (e ErrNoEndpoint) Error() string {
	return "No endpoint specified"
}

type ErrInvalidResponse struct {
	StatusCode int
	Message    string
}

func (e ErrInvalidResponse) Error() string {
	return fmt.Sprintf("Invalid response from server: %d %q", e.StatusCode, e.Message)
}
