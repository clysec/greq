package greq

import (
	"crypto/tls"
	"encoding/base64"
	"net/http"

	"github.com/Azure/go-ntlmssp"
)

type NTLMAuth struct {
	Username string
	Password string

	ForceHttp11        bool
	InsecureSkipVerify bool
}

func (na *NTLMAuth) Prepare() error {
	return nil
}

func (na *NTLMAuth) Apply(addHeaderFunc func(key, value string), setTransportFunc func(transport http.RoundTripper)) error {
	transport := ntlmssp.Negotiator{
		RoundTripper: &http.Transport{},
	}

	if na.ForceHttp11 {
		transport.RoundTripper.(*http.Transport).TLSNextProto = map[string]func(string, *tls.Conn) http.RoundTripper{}
	}

	if na.InsecureSkipVerify {
		transport.RoundTripper.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	addHeaderFunc("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(na.Username+":"+na.Password)))
	setTransportFunc(transport)

	return nil
}

// TODO: Implement
// https://github.com/vadimi/go-http-ntlm
// https://authenticationtest.com/HTTPAuth/
