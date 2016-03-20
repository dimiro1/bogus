package bogus

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

// Mux is struct that knows how to build handlers
type Mux struct {
	mux *mux.Router
}

// NewMux returns a new Mux
func NewMux() *Mux {
	return &Mux{
		mux: mux.NewRouter(),
	}
}

// AddRoute adds a new Route
func (m *Mux) AddRoute(route Route) *Mux {
	m.mux.
		NewRoute().
		Methods(route.Methods...).
		Name(route.Name).
		Path(route.Path).
		HandlerFunc(routeHandler(route))

	return m
}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.mux.ServeHTTP(w, r)
}

func routeHandler(route Route) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		for key, value := range route.Headers {
			w.Header().Add(key, value)
		}

		w.WriteHeader(route.Status)

		t, err := template.New(route.Name).Parse(route.Body)

		// Error, rendering the raw response
		if err != nil {
			fmt.Fprint(w, route.Body)
			return
		}

		t.Execute(w, struct {
			Name    string
			Path    string
			Methods []string
			Headers map[string]string
			Params  map[string]string
			Status  int
		}{
			route.Name,
			route.Path,
			route.Methods,
			route.Headers,
			params,
			route.Status,
		})
	}
}
