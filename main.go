package main

import "fmt"

// BasicAuth is used to configure a http basic auth on the endpoint
type BasicAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Route is used to configure the route
type Route struct {
	Name    string            `json:"name"`
	Route   string            `json:"route"`
	Methods []string          `json:"methods"`
	Headers map[string]string `json:"headers,omitempty"`

	// Body is converted to a url.URL to extract the scheme
	// The valid protocols are file: and http:
	// if none of these protocols are found, then the response body will be the raw string
	Body   string `json:"body"`
	Status int    `json:"status"`

	Auth BasicAuth `json:"auth,omitempty"`
}

// Routes is a list of configured routes
type Routes []Route

func main() {
	fmt.Println("Hello Bogus")
}
