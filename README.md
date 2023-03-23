# prometheus api client

## Usage
```go
package main

import (
	"fmt"

	Prometheus "github.com/bonsai-oss/prometheus-api-client"
)

func main() {
	client := Prometheus.Client{Endpoint: "http://localhost:9090"}
	data := client.Query("up").Result()
	fmt.Printf("%+v", data)
}
```
