package prometheus

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"
)

func TestQueryReachable(t *testing.T) {
	client := Client{Endpoint: "http://demo.robustperception.io:9090"}
	result, _ := client.Query(InstantQuery{QueryExpression: "up"})
	if result.Status != "success" {
		t.Error("Query not successful")
	}
}

func TestQueryNoEndpoint(t *testing.T) {
	client := Client{}
	result, queryError := client.Query(InstantQuery{QueryExpression: "up"})
	if queryError == nil {
		t.Error("Query should not be successful")
	}
	if result != nil {
		t.Error("result should be nil")
	}
}

func TestInstantResultTime(t *testing.T) {
	client := Client{Endpoint: "http://demo.robustperception.io:9090"}
	result, _ := client.Query(InstantQuery{QueryExpression: "up"})
	if result.Result()[0].Values[0].Timestamp.IsZero() {
		t.Error("Result time is zero")
	}
}

func TestRangeResultTime(t *testing.T) {
	client := Client{Endpoint: "http://demo.robustperception.io:9090"}
	result, _ := client.Query(RangeQuery{QueryExpression: "up", Start: time.Now().Add(-1 * time.Hour), End: time.Now(), Step: 1 * time.Minute})
	if result.Result()[0].Values[0].Timestamp.IsZero() {
		t.Error("Result time is zero")
	}
}

func TestRangeQuery(t *testing.T) {
	client := Client{Endpoint: "http://demo.robustperception.io:9090"}
	result, _ := client.Query(RangeQuery{QueryExpression: "node_load5", Start: time.Now().Add(-1 * time.Hour), End: time.Now(), Step: 1 * time.Minute})

	if len(result.Result()) < 1 {
		t.Error("Not enough query results")
	}
}

func TestRangeResultTimeZero(t *testing.T) {
	client := Client{Endpoint: "http://demo.robustperception.io:9090"}
	result, _ := client.Query(RangeQuery{QueryExpression: "up", Start: time.Now().Add(-1 * time.Hour), End: time.Now(), Step: 1 * time.Minute})
	for _, res := range result.Result() {
		for _, val := range res.Values {
			if val.Timestamp.IsZero() {
				t.Error("Result time is zero")
			}
		}
	}
}

func TestQueryResponseTypes(t *testing.T) {
	client := Client{Endpoint: "http://demo.robustperception.io:9090"}

	for _, testCase := range []struct {
		query         Query
		expectedType  string
		expectSuccess bool
	}{
		{expectSuccess: true, query: InstantQuery{QueryExpression: "up"}, expectedType: "vector"},
		{expectSuccess: true, query: InstantQuery{QueryExpression: "sum(up)"}, expectedType: "vector"},
		{expectSuccess: true, query: InstantQuery{QueryExpression: "up{job=\"prometheus\"}"}, expectedType: "vector"},
		{expectSuccess: true, query: InstantQuery{QueryExpression: "up{job=\"prometheus\"}[5m]"}, expectedType: "matrix"},
		{expectSuccess: true, query: InstantQuery{QueryExpression: "rate(up[5m])"}, expectedType: "vector"},
		{expectSuccess: true, query: RangeQuery{QueryExpression: "up", Start: time.Now().Add(-1 * time.Hour), End: time.Now(), Step: 20 * time.Minute}, expectedType: "matrix"},
		{expectSuccess: true, query: RangeQuery{QueryExpression: "sum(up)", Start: time.Now().Add(-1 * time.Hour), End: time.Now(), Step: 20 * time.Minute}, expectedType: "matrix"},
		{expectSuccess: true, query: RangeQuery{QueryExpression: "up{job=\"prometheus\"}", Start: time.Now().Add(-1 * time.Hour), End: time.Now(), Step: 20 * time.Minute}, expectedType: "matrix"},
		{expectSuccess: false, query: RangeQuery{QueryExpression: "up{job=\"prometheus\"}[5m]", Start: time.Now().Add(-1 * time.Hour), End: time.Now(), Step: 20 * time.Minute}, expectedType: "matrix"},
		{expectSuccess: true, query: RangeQuery{QueryExpression: "rate(up[5m])", Start: time.Now().Add(-1 * time.Hour), End: time.Now(), Step: 20 * time.Minute}, expectedType: "matrix"},
	} {
		t.Run("", func(t *testing.T) {
			response, _ := client.Query(testCase.query)

			if testCase.expectSuccess && response.Status != "success" {
				t.Errorf("Expected success, got %s", response.Status)
			}
			if !testCase.expectSuccess && response.Status == "success" {
				t.Errorf("Expected failure, got %s", response.Status)
			}

			buf := new(bytes.Buffer)
			encoder := json.NewEncoder(buf)
			encoder.SetIndent("", "  ")
			_ = encoder.Encode(response)

			t.Log(buf.String())

			if !testCase.expectSuccess {
				return
			}

			if response.Data.ResultType != testCase.expectedType {
				t.Errorf("Expected result type %s, got %s", testCase.expectedType, response.Data.ResultType)
			}
			// check if the result is not empty
			if len(response.Data.Result) == 0 {
				t.Error("Result is empty")
			}

			// check if result data is not empty
			for _, res := range response.Data.Result {
				if len(res.Values) == 0 {
					t.Error("Result data is empty")
				}
			}

			// check if the timestamp is not zero
			for _, res := range response.Data.Result {
				for _, val := range res.Values {
					if val.Timestamp.IsZero() {
						t.Error("Result time is zero")
					}
				}
			}
		})
	}
}
