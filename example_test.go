package prometheus_test

import (
	"testing"
	"time"

	"github.com/bonsai-oss/prometheus-api-client"
)

func TestInstantQuery(t *testing.T) {
	client := prometheus.Client{Endpoint: "http://demo.robustperception.io:9090"}
	result, _ := client.Query(prometheus.InstantQuery{QueryExpression: "up", Time: time.Now()})

	res := result.Result()

	if len(res) < 1 {
		t.Errorf("Not enough query results: %d", len(res))
	}

	if res[0].Values[0].Timestamp.IsZero() {
		t.Error("Result time is zero")
	}
}

func TestRangeQuery(t *testing.T) {
	client := prometheus.Client{Endpoint: "http://demo.robustperception.io:9090"}
	result, _ := client.Query(prometheus.RangeQuery{QueryExpression: "go_goroutines", Start: time.Now().Add(-1 * time.Hour), End: time.Now(), Step: 10 * time.Minute})

	res := result.Result()

	if len(res) < 1 {
		t.Errorf("Not enough query results: %d", len(res))
	}

	if res[0].Values[0].Timestamp.IsZero() {
		t.Error("Result time is zero")
	}
}
