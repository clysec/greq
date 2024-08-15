package auth

import (
	"fmt"
	"net/http"
)

type Oauth2Auth struct {
}

func (oa *Oauth2Auth) Prepare() error {
	// TODO: Implement
	return fmt.Errorf("not implemented")
}

// TODO: Finish
func (oa *Oauth2Auth) Apply(addHeaderFunc func(key, value string), setTransportFunc func(transport *http.Transport)) error {
	// TODO: Implement
	return fmt.Errorf("not implemented")
}
