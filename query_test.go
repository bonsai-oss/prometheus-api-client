package prometheus

import (
	"testing"
)

func TestQueryReachable(t *testing.T) {
	client := Client{Endpoint: "http://demo.robustperception.io:9090"}
	result := client.Query("up")
	if result.Status != "success" {
		t.Error("Query not successful")
	}
}

func TestQueryNoEndpoint(t *testing.T) {
	client := Client{}
	result := client.Query("up")
	if result.Status != "success" {
		t.Error("Query not successful")
	}
}

func TestResult(t *testing.T) {
	client := Client{Endpoint: "http://demo.robustperception.io:9090"}
	result := client.Query("up").Result()
	if len(result) < 1 {
		t.Error("Not enough query results")
	}
}
