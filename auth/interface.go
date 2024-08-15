package auth

import "net/http"

type Authorization interface {
	Prepare() error
	Apply(addHeaderFunc func(key, value string), setTransportFunc func(transport *http.Transport)) error
}
