package auth

import (
	"fmt"
	"net/http"
)

type AwsSignatureAuth struct {
	AccessKey    string
	SecretKey    string
	Region       string
	ServiceName  string
	SessionToken string
}

func (a *AwsSignatureAuth) Prepare() error {
	// TODO: Implement
	return fmt.Errorf("not implemented")
}

func (a *AwsSignatureAuth) Apply(addHeaderFunc func(key, value string), setTransportFunc func(transport *http.Transport)) error {
	// TODO: Implement
	return fmt.Errorf("not implemented")
}
