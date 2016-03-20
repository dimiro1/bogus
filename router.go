package bogus

type Route struct {
	name    string
	path    string
	methods []string
	headers map[string]string
	status  int
	body    string
}

func NewRoute() *Route {
	return &Route{}
}

func (r *Route) Name(name string) *Route {
	r.name = name
	return r
}

func (r *Route) Path(path string) *Route {
	r.path = path
	return r
}

func (r *Route) Methods(methods []string) *Route {
	r.methods = methods

	return r
}

func (r *Route) Headers(headers map[string]string) *Route {
	r.headers = headers
	return r
}

func (r *Route) Status(status int) *Route {
	r.status = status
	return r
}

func (r *Route) Body(body string) *Route {
	r.body = body
	return r
}
