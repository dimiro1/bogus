package bogus

// Route is the struct that represents a route
type Route struct {
	name    string
	path    string
	methods []string
	headers map[string]string
	status  int
	body    string
}

// NewRoute returns a pointer to a new Route
func NewRoute() *Route {
	return &Route{}
}

// Name set the member name in the Route struct
func (r *Route) Name(name string) *Route {
	r.name = name
	return r
}

// Path set the member path in the Route struct
func (r *Route) Path(path string) *Route {
	r.path = path
	return r
}

// Methods set the member methods in the Route struct
func (r *Route) Methods(methods []string) *Route {
	r.methods = methods
	return r
}

// Headers set the member headers in the Route struct
func (r *Route) Headers(headers map[string]string) *Route {
	r.headers = headers
	return r
}

// Status set the member status in the Route struct
func (r *Route) Status(status int) *Route {
	r.status = status
	return r
}

// Body set the member body in the Route struct
func (r *Route) Body(body string) *Route {
	r.body = body
	return r
}
