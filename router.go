package bogus

// Route is the struct that represents a route
type Route struct {
	Name    string
	Path    string
	Methods []string
	Headers map[string]string
	Status  int
	Body    string
}
