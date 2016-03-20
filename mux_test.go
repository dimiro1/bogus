package bogus

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_NewMux(t *testing.T) {
	m := NewMux()

	if m.mux == nil {
		t.Error("m.mux == nil, expected !nil")
	}
}

func Test_Mux_AddRoute(t *testing.T) {
	m := NewMux()

	route := Route{
		Name:    "Hello",
		Status:  200,
		Methods: []string{"GET"},
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
		Path: "/",
		Body: "Hello World",
	}

	m.AddRoute(route)

	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	m.ServeHTTP(w, r)

	if w.Code != 200 {
		t.Errorf("w.Code == %d, expected %d", w.Code, 200)
	}

	if w.Header().Get("Content-Type") != "text/html" {
		t.Errorf(`w.Header().Get("Content-Type") == %s, expected %s`, w.Header().Get("Content-Type"), "text/html")
	}

	if w.Body.String() != "Hello World" {
		t.Errorf("w.Body.String() == %s, expected %s", w.Body.String(), "Hello World")
	}
}

func Test_Mux_AddRoute_With_Template(t *testing.T) {
	m := NewMux()

	body := `
{
	"name": "{{ .Name }}",
	"path": "{{ .Path }}",
	"methods": ["{{ index .Methods 0 }}"],
	"headers": {
		"Content-Type": "{{ index .Headers "Content-Type" }}"
	},
	"params": {
		"name": "{{ index .Params "name" }}"
	},
	"status": {{ .Status }}
}`

	expectedBody := `
{
	"name": "Hello",
	"path": "/{name}",
	"methods": ["GET"],
	"headers": {
		"Content-Type": "application/json"
	},
	"params": {
		"name": "hello"
	},
	"status": 200
}`

	route := Route{
		Name:    "Hello",
		Status:  200,
		Methods: []string{"GET"},
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Path: "/{name}",
		Body: body,
	}

	m.AddRoute(route)

	r, _ := http.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()

	m.ServeHTTP(w, r)

	if w.Body.String() != expectedBody {
		t.Errorf("w.Body.String() == %s, expected %s", w.Body.String(), expectedBody)
	}
}

func Test_Mux_AddRoute_With_Invalid_Template(t *testing.T) {
	m := NewMux()

	route := Route{
		Name:    "Hello",
		Status:  200,
		Methods: []string{"GET"},
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
		Path: "/",
		Body: "{{ .Name }",
	}

	m.AddRoute(route)

	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	m.ServeHTTP(w, r)

	if w.Body.String() != "{{ .Name }" {
		t.Errorf("w.Body.String() == %s, expected %s", w.Body.String(), "{{ .Name }")
	}
}

func Test_Mux_AddRoute_invalid_method(t *testing.T) {
	m := NewMux()

	route := Route{
		Name:    "Hello",
		Status:  200,
		Methods: []string{"GET"},
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
		Path: "/",
		Body: "Hello World",
	}

	m.AddRoute(route)

	r, _ := http.NewRequest("POST", "/", nil)
	w := httptest.NewRecorder()

	m.ServeHTTP(w, r)

	if w.Code != 404 {
		t.Errorf("w.Code == %d, expected %d", w.Code, 404)
	}
}
