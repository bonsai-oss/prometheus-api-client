package prometheus

import (
	"reflect"
	"testing"
)

func TestWithEndpoint(t *testing.T) {
	expectedEndpoint := "http://example.com"
	option := WithEndpoint(expectedEndpoint)
	client := Client{}
	option(&client)

	if client.Endpoint != expectedEndpoint {
		t.Errorf("Expected endpoint '%s', got '%s'", expectedEndpoint, client.Endpoint)
	}
}

func TestWithHeaders(t *testing.T) {
	expectedHeaders := map[string]string{
		"Header1": "Value1",
		"Header2": "Value2",
	}
	option := WithHeaders(expectedHeaders)
	client := Client{}
	option(&client)

	if !reflect.DeepEqual(client.headers, expectedHeaders) {
		t.Errorf("Expected headers '%v', got '%v'", expectedHeaders, client.headers)
	}
}

func TestNewClient(t *testing.T) {
	expectedEndpoint := "http://example.com"
	expectedHeaders := map[string]string{
		"Header1": "Value1",
		"Header2": "Value2",
	}
	client, err := NewClient(WithEndpoint(expectedEndpoint), WithHeaders(expectedHeaders))

	if err != nil {
		t.Errorf("Expected no error, got '%s'", err.Error())
	}

	if client.Endpoint != expectedEndpoint {
		t.Errorf("Expected endpoint '%s', got '%s'", expectedEndpoint, client.Endpoint)
	}

	if !reflect.DeepEqual(client.headers, expectedHeaders) {
		t.Errorf("Expected headers '%v', got '%v'", expectedHeaders, client.headers)
	}
}

func TestNewClientNoEndpoint(t *testing.T) {
	_, err := NewClient()

	if err == nil {
		t.Error("Expected error, got nil")
	}

	if _, ok := err.(ErrNoEndpoint); !ok {
		t.Errorf("Expected error of type 'ErrNoEndpoint', got '%s'", reflect.TypeOf(err))
	}

	if err.Error() != "No endpoint specified" {
		t.Errorf("Expected error message 'No endpoint specified', got '%s'", err.Error())
	}
}
