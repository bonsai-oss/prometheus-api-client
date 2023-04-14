package prometheus

// Client - A simple prometheus api client
type Client struct {
	Endpoint string

	// headers - A map of headers to add to the requests
	headers map[string]string
}

type ClientOption func(*Client)

// NewClient - Create a new prometheus client with the given options
func NewClient(options ...ClientOption) (Client, error) {
	client := Client{headers: map[string]string{}}
	for _, option := range options {
		option(&client)
	}

	// Validate the client options
	if client.Endpoint == "" {
		return Client{}, ErrNoEndpoint{}
	}

	return client, nil
}

// WithEndpoint - Set the endpoint for the client
func WithEndpoint(endpoint string) ClientOption {
	return func(c *Client) {
		c.Endpoint = endpoint
	}
}

// WithHeaders - Set the additional headers for the client
func WithHeaders(headers map[string]string) ClientOption {
	return func(c *Client) {
		c.headers = headers
	}
}
