package prometheus

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

// QueryResult - The result of a prometheus query and query_range request
// The Values field is used for both instant and range queries.
// For instant queries, the Values field will contain a single Value.
type QueryResult struct {
	Metric map[string]string `json:"metric"`
	Values []Value           `json:"values"`
}

// UnmarshalJSON - Unmarshal the json response from prometheus
func (qr *QueryResult) UnmarshalJSON(data []byte) error {
	type alias QueryResult
	aux := struct {
		*alias
		Value Value `json:"value"`
	}{
		alias: (*alias)(qr),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if !aux.Value.Timestamp.IsZero() {
		qr.Values = append(qr.Values, aux.Value)
	}
	return nil
}

// QueryResponse - The return data from a simple prometheus time series query
type QueryResponse struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string        `json:"resultType"`
		Result     []QueryResult `json:"result"`
	} `json:"data"`
}

// Query - fetch the data from prometheus API
func (c Client) Query(query Query) (*QueryResponse, error) {
	query.values()

	var path string
	switch query.(type) {
	case InstantQuery:
		path = "api/v1/query"
	case RangeQuery:
		path = "api/v1/query_range"
	}

	request, requestCreateError := http.NewRequest(http.MethodPost, func(result string, _ error) string { return result }(url.JoinPath(c.Endpoint, path)), strings.NewReader(query.values().Encode()))
	if requestCreateError != nil {
		return nil, requestCreateError
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, responseError := http.DefaultClient.Do(request)
	if responseError != nil {
		return nil, responseError
	}

	var obj QueryResponse
	if err := json.NewDecoder(res.Body).Decode(&obj); err != nil {
		return nil, err
	}
	return &obj, nil
}

// Result - Only return the result struct of the query
func (obj *QueryResponse) Result() []QueryResult {
	return obj.Data.Result
}
