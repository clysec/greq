package greq

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"

	"github.com/scheiblingco/gofn/errtools"
)

// The HTTP Response for the request
type GResponse struct {
	StatusCode int
	Headers    map[string][]string
	Response   *http.Response

	bodyRead bool
}

func (r *GResponse) BodyBytes() ([]byte, error) {
	if r.bodyRead {
		return nil, errtools.BodyConsumedError("body has already been read")
	}

	r.bodyRead = true
	defer r.Response.Body.Close()

	return io.ReadAll(r.Response.Body)
}

func (r *GResponse) BodyString() (string, error) {
	b, err := r.BodyBytes()
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (r *GResponse) BodyReader() (*io.ReadCloser, error) {
	if r.bodyRead {
		return nil, errtools.BodyConsumedError("body has already been read")
	}

	r.bodyRead = true
	return &r.Response.Body, nil
}

func (r *GResponse) BodyUnmarshalJson(v interface{}) error {
	if r.bodyRead {
		return errtools.BodyConsumedError("body has already been read")
	}

	r.bodyRead = true
	defer r.Response.Body.Close()

	return json.NewDecoder(r.Response.Body).Decode(v)
}

func (r *GResponse) BodyUnmarshalXml(v interface{}) error {
	if r.bodyRead {
		return errtools.BodyConsumedError("body has already been read")
	}

	r.bodyRead = true
	defer r.Response.Body.Close()

	return xml.NewDecoder(r.Response.Body).Decode(v)
}

// TODO: AutoUnmarshal, detect content type and unmarshal accordingly
// json: application/json
// xml: application/xml
// yml: application/yaml
// php: application/php

func (r *GResponse) Close() {
	if r.bodyRead {
		return
	}

	r.bodyRead = true
	defer r.Response.Body.Close()
}
