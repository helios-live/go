package api

import (
	"github.com/ideatocode/go/api/transport"
)

// Instance represents a single api instance
type Instance struct {
	list []Endpoint
}

// New returns a new Instance
func New() *Instance {
	return &Instance{}
}

// All returns a slice of All the Endpoints
func (i *Instance) All() []Endpoint {

	// var list endpoint.Endpoint
	// list = makeListServersEndpoint(svc)
	return i.list
}

// Add registers a single Endpoint to the Instance
func (i *Instance) Add(s Endpoint) {
	i.list = append(i.list, s)
}

// Use adds all the Instances endpoints to the sm ServeMux
func (i *Instance) Use(ts transport.Transport) {
	for _, ep := range i.All() {
		ts.Register(ep.Entry(), ts.Decode(ep.Decoder()), ep.Path())
	}
}
