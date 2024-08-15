package greq

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"strings"

	"github.com/scheiblingco/gofn/errtools"
	"github.com/scheiblingco/gofn/typetools"
)

// Add a body to the request in the form of a byte slice
func (g *GRequest) WithByteBody(body []byte) *GRequest {
	g.bodyAccepted()

	g.body = bytes.NewReader(body)
	return g
}

// Add a body to the request in the form of a string
func (g *GRequest) WithStringBody(body string) *GRequest {
	return g.WithByteBody([]byte(body))
}

// Add a body to the request in the form of a reader
func (g *GRequest) WithReaderBody(body io.Reader) *GRequest {
	g.bodyAccepted()

	g.body = body
	return g
}

// Add a JSON body to the request
// Accepts an interface{} that will be marshalled into JSON, or
// a string or string-like object with pre-marshalled JSON
func (g *GRequest) WithJSONBody(body interface{}, contentType *string) *GRequest {
	g.bodyAccepted()

	cType := "application/json"
	if contentType != nil {
		cType = *contentType
	}

	if val, ok := body.(string); ok {
		return g.WithStringBody(val).WithHeader("Content-Type", cType)
	}

	if val, ok := body.([]byte); ok {
		return g.WithByteBody(val).WithHeader("Content-Type", cType)
	}

	buf := new(bytes.Buffer)

	err := json.NewEncoder(buf).Encode(body)
	if err != nil {
		g.addError(err)
	}

	return g
}

// Adds an XML body to the request
// Accepts an interface{} that will be marshalled into XML, or
// a string or string-like object with pre-marshalled XML
func (g *GRequest) WithXMLBody(body interface{}, contentType *string) *GRequest {
	g.bodyAccepted()

	cType := "application/xml"
	if contentType != nil {
		cType = *contentType
	}

	if val, ok := body.(string); ok {
		return g.WithStringBody(val).WithHeader("Content-Type", cType)
	}

	if val, ok := body.([]byte); ok {
		return g.WithByteBody(val).WithHeader("Content-Type", cType)
	}

	buf := new(bytes.Buffer)

	err := xml.NewEncoder(buf).Encode(body)
	if err != nil {
		g.addError(err)
	}

	return g
}

// Add a application/x-www-form-urlencoded body
// Accepts the following types:
// - map[string]string
// - map[string][]string
// - map[string][]byte
// - map[string]interface{} where interface can be a string-like, numeric, string slice or boolean type
// - url.Values
func (g *GRequest) WithUrlencodedFormBody(body interface{}, contentType *string) *GRequest {
	g.bodyAccepted()

	cType := "application/x-www-form-urlencoded"
	if contentType != nil {
		cType = *contentType
	}

	data := url.Values{}

	switch val := body.(type) {
	case url.Values:
		return g.WithStringBody(val.Encode()).WithHeader("Content-Type", cType)
	case map[string]string:
		for k, v := range val {
			data.Add(k, v)
		}
		return g.WithStringBody(data.Encode()).WithHeader("Content-Type", cType)
	case map[string][]string:
		for k, v := range val {
			for _, v2 := range v {
				data.Add(k, v2)
			}
		}
		return g.WithStringBody(data.Encode()).WithHeader("Content-Type", cType)
	case map[string][]byte:
		for k, v := range val {
			data.Add(k, string(v))
		}
		return g.WithStringBody(data.Encode()).WithHeader("Content-Type", cType)
	case map[string]interface{}:
		for k, v := range val {
			if typetools.IsStringlikeType(v) || typetools.IsNumericType(v) {
				data.Add(k, typetools.EnsureString(v))
			} else {
				if vt, ok := v.([]string); ok {
					for _, v2 := range vt {
						data.Add(k, v2)
					}
				} else if vt, ok := v.(bool); ok {
					data.Add(k, fmt.Sprintf("%t", vt))
				} else {
					g.addError(errtools.InvalidTypeError(fmt.Sprintf("field %s - form value must be a string, string slice, or string-like type", k)))
				}
			}
		}
		return g.WithStringBody(data.Encode()).WithHeader("Content-Type", cType)
	}

	g.addError(errtools.InvalidTypeError("form body must be a map[string]string, map[string][]string, map[string][]byte, map[string]interface{}, or url.Values"))
	return g
}

// Add a multipart form body to the request
// Accepts a list of multipart fields
func (g *GRequest) WithMultipartFormBody(body []*MultipartField) *GRequest {
	g.bodyAccepted()

	if g.headers == nil {
		g.headers = make(map[string]string)
	}

	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)

	for _, field := range body {
		if err := field.AddToWriter(writer); err != nil {
			g.addError(err)
		}
	}

	if len(g.errs) > 0 {
		return g
	}

	if err := writer.Close(); err != nil {
		g.addError(err)
		return g
	}

	g.body = buf

	return g.WithHeader("Content-Type", writer.FormDataContentType())
}

// Add query parameters to the request
func (g *GRequest) WithQueryParams(params interface{}) *GRequest {
	data := url.Values{}

	switch val := params.(type) {
	case url.Values:
		data = val

	case map[string]string:
		for k, v := range val {
			data.Add(k, v)
		}

	case map[string][]string:
		for k, v := range val {
			for _, v2 := range v {
				data.Add(k, v2)
			}
		}

	case map[string][]byte:
		for k, v := range val {
			data.Add(k, string(v))
		}

	case map[string]interface{}:
		for k, v := range val {
			if typetools.IsStringlikeType(v) || typetools.IsNumericType(v) {
				data.Add(k, typetools.EnsureString(v))
			} else {
				if vt, ok := v.([]string); ok {
					for _, v2 := range vt {
						data.Add(k, v2)
					}
				} else if vt, ok := v.(bool); ok {
					data.Add(k, fmt.Sprintf("%t", vt))
				} else {
					g.addError(errtools.InvalidTypeError(fmt.Sprintf("field %s - form value must be a string, string slice, or string-like type", k)))
				}
			}
		}
	default:
		g.addError(errtools.InvalidTypeError("query params must be a map[string]string, map[string][]string, map[string][]byte, map[string]interface{}, or url.Values"))
	}

	if len(g.errs) > 0 {
		return g
	}

	if strings.Contains(g.Url, "?") {
		g.Url += "&"
	} else {
		g.Url += "?"
	}

	g.Url += data.Encode()

	return g
}

// TODO: Add support for GraphQL Request
// TODO: Add support for SOAP Request
// TODO: Add support for Websocket requests
