package main

import "fmt"

type responseType string

var (
	file     responseType = "file"
	redirect              = "redirect"
	raw                   = "raw"
)

type Header struct {
	ContentType string `json:"Content-Type"`
}

type Response struct {
	Status  int          `json:"status"`
	Type    responseType `json:"type"`
	Content string       `json:"content"`
}

type Auth struct {
	Type     string `json:"type"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Route struct {
	Name     string   `json:"name"`
	Route    string   `json:"route"`
	Method   string   `json:"method"`
	Headers  []Header `json:"headers,omitempty"`
	Response Response `json:"response"`
	Auth     Auth     `json:"auth,omitempty"`
}

type Routes []Route

func main() {
	fmt.Println("Hello Bogus")
}
