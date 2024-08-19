package greq

import (
	"encoding/base64"
	"net/http"
)

// Sets the Authorization header with a Basic Auth token
type BasicAuth struct {
	Username string
	Password string
}

func (ba *BasicAuth) Prepare() error {
	return nil
}

func (ba *BasicAuth) Apply(addHeaderFunc func(key, value string), setTransportFunc func(transport *http.Transport)) error {
	addHeaderFunc("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(ba.Username+":"+ba.Password)))
	return nil
}
