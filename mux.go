package bogus

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Mux struct {
	mux *mux.Router
}

func NewMux() *Mux {
	return &Mux{
		mux: mux.NewRouter(),
	}
}

func (m *Mux) AddRoute(route *Route) *Mux {
	m.mux.
		NewRoute().
		Methods(route.methods...).
		Name(route.name).
		Path(route.path).
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for key, value := range route.headers {
				w.Header().Add(key, value)
			}

			w.WriteHeader(route.status)

			fmt.Fprint(w, route.body)
		})

	return m
}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.mux.ServeHTTP(w, r)
}

func (m *Mux) Serve(addr string) {
	fmt.Printf("Starting Bogus on addr %s\n", addr)
	log.Fatal(http.ListenAndServe(addr,
		handlers.LoggingHandler(os.Stdout, m)))
}
