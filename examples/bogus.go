package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/dimiro1/bogus"
	"github.com/gorilla/handlers"
)

type route struct {
	Name    string            `json:"name"`
	Path    string            `json:"path"`
	Methods []string          `json:"methods"`
	Headers map[string]string `json:"headers,omitempty"`

	// The valid options are file, url and raw
	Type   string `json:"type"`
	Body   string `json:"body"`
	Status int    `json:"status"`
}

// Renderizar o body com html/template ou text/template
// Ter acesso as todas as variáveis, Method, Headers, Status, Parameters, QueryString etc

func main() {
	var routes []route

	file, _ := os.Open("config.json")

	if err := json.NewDecoder(file).Decode(&routes); err != nil {
		log.Fatal(err)
	}

	mux := bogus.NewMux()

	for _, r := range routes {
		route := bogus.Route{
			Name:    r.Name,
			Path:    r.Path,
			Methods: r.Methods,
			Headers: r.Headers,
			Status:  r.Status,
		}

		var content []byte

		switch r.Type {
		case "file":
			file, _ = os.Open(r.Body)
			content, _ = ioutil.ReadAll(file)
		case "url":
			resp, _ := http.Get(r.Body)
			content, _ = ioutil.ReadAll(resp.Body)
		default:
			content = []byte(r.Body)
		}

		route.Body = string(content)

		mux.AddRoute(route)
	}

	addr := ":8080"

	fmt.Printf("Starting Bogus on addr %s\n", addr)
	fmt.Println(`
######                              
#     #  ####   ####  #    #  ####  
#     # #    # #    # #    # #      
######  #    # #      #    #  ####  
#     # #    # #  ### #    #      # 
#     # #    # #    # #    # #    # 
######   ####   ####   ####   ####  
`)
	fmt.Println("By: Claudemiro Alves Feitosa Neto <dimiro1@gmail.com>")
	fmt.Println("Waiting for requests...")

	log.Fatal(http.ListenAndServe(":8080",
		handlers.LoggingHandler(os.Stdout, mux)))
}
