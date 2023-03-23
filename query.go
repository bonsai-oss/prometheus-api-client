package prometheus

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// QueryResult - Only results
type QueryResult struct {
	Metric map[string]string `json:"metric"`
	Value  []interface{}     `json:"value"`
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
func (c Client) Query(metric string) *QueryResponse {
	if c.Endpoint == "" {
		c.Endpoint = "http://demo.robustperception.io:9090"
	}
	res, _ := http.Get(fmt.Sprintf("%v/api/v1/query?query=%v", c.Endpoint, metric))
	data, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	var obj QueryResponse
	err := json.Unmarshal(data, &obj)
	if err != nil {
		log.Fatalln(err)
	}
	return &obj
}

// Result - Only return the result struct of the query
func (obj *QueryResponse) Result() []QueryResult {
	return obj.Data.Result
}
