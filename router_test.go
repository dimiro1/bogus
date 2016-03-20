package bogus

import "testing"

func Test_Route_Name(t *testing.T) {
	expected := "Hello"

	route := NewRoute()
	route.Name(expected)

	if route.name != expected {
		t.Errorf("route.name == %s, expected %s", route.name, expected)
	}
}

func Test_Route_Path(t *testing.T) {
	expected := "/hello"

	route := NewRoute()
	route.Path(expected)

	if route.path != expected {
		t.Errorf("route.path == %s, expected %s", route.path, expected)
	}
}

func Test_Route_Methods(t *testing.T) {
	expected := []string{"GET", "POST"}

	route := NewRoute()
	route.Methods(expected)

	stringInSlice(t, "GET", route.methods)
	stringInSlice(t, "POST", route.methods)
}

func Test_Route_Headers(t *testing.T) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	route := NewRoute()
	route.Headers(headers)

	_, ok := route.headers["Content-Type"]

	if !ok {
		t.Errorf(`route.headers["Content-Type"] is not present`)
	}
}

func Test_Route_Status(t *testing.T) {
	expected := 200

	route := NewRoute()
	route.Status(expected)

	if route.status != expected {
		t.Errorf("route.status == %d, expected %d", route.status, expected)
	}
}

func Test_Route_Body(t *testing.T) {
	expected := "Hello World"

	route := NewRoute()
	route.Body(expected)

	if route.body != expected {
		t.Errorf("route.body == %s, expected %s", route.body, expected)
	}
}

func stringInSlice(t *testing.T, a string, list []string) {
	for _, b := range list {
		if b == a {
			return
		}
	}

	t.Errorf("%v does not contains %s", list, a)
}
