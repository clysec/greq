package greq

import (
	"fmt"
	"net/http"
	"strings"
)

// Sets the Authorization header with a Bearer Auth token
// The default prefix is Bearer
type BearerAuth struct {
	Token  string
	Prefix string
}

func (ha *BearerAuth) Prepare() error {
	return nil
}

func (ha *BearerAuth) Apply(addHeaderFunc func(key, value string), setTransportFunc func(transport http.RoundTripper)) error {
	if ha.Prefix == "" {
		ha.Prefix = "Bearer"
	}

	ha.Prefix = strings.TrimRight(ha.Prefix, " ")

	addHeaderFunc("Authorization", fmt.Sprintf("%s %s", ha.Prefix, ha.Token))
	return nil
}
