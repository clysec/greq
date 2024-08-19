package greq

import (
	"net/http"
)

// Add an arbitrary header to the request
type HeaderAuth struct {
	Key   string
	Value string
}

func (ha *HeaderAuth) Prepare() error {
	return nil
}

func (ha *HeaderAuth) Apply(addHeaderFunc func(key, value string), setTransportFunc func(transport *http.Transport)) error {
	addHeaderFunc(ha.Key, ha.Value)
	return nil
}
