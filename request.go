package greq

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/scheiblingco/gofn/errtools"
	"github.com/scheiblingco/gofn/typetools"
)

// The HTTP Method to use for the requeswt
type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	PATCH  Method = "PATCH"
	DELETE Method = "DELETE"
)

// The HTTP Request to be made
type GRequest struct {
	Url    string
	Method Method

	client  *http.Client
	headers map[string]string
	query   *url.Values
	body    io.Reader

	errs []error
}

// Add an error to the list of errors for the request
// This list is checked and validated before the request is made
func (g *GRequest) addError(err error) {
	g.errs = append(g.errs, err)
}

// Add a header to the request
func (g *GRequest) addHeader(key, value string) {
	if g.headers == nil {
		g.headers = make(map[string]string)
	}

	g.headers[key] = value
}

// Add a custom transport to the http client
func (g *GRequest) addTransport(transport *http.Transport) {
	if g.client == nil {
		g.client = &http.Client{Transport: transport}
	} else {
		g.client.Transport = transport
	}
}

// Check if a body is accepted for the request
func (g *GRequest) bodyAccepted() {
	if g.Method == GET || g.Method == DELETE {
		g.addError(errtools.BodyNotAcceptedError("cannot have a body with a GET or DELETE request"))
	}
}

// Add a custom HTTP client to the request
// This is optional, and a HTTP client will be automatically created
// if one is not provided. WithClient needs to be called first, it will panic
// if called after any other functions that have already set or modified the client
func (g *GRequest) WithClient(client *http.Client) *GRequest {
	if g.client != nil {
		panic("cannot add a client to a request that already has a client. call WithClient before any other functions")
	}

	g.client = client

	return g
}

// // Add authentication to the request
// // An Authorization type can be passed to multiple requests,
// // which is useful in the case of Oauth2 or other token-based requests
// // that can re-use the same token for multiple requests
func (g *GRequest) WithAuth(auth Authorization) *GRequest {
	if err := auth.Prepare(); err != nil {
		g.addError(err)
	}

	if err := auth.Apply(g.addHeader, g.addTransport); err != nil {
		g.addError(err)
	}

	return g
}

// Add a header to the request
// Headers are added before the body functions, meaning if you add a
// content-type header and then add a form body, the request header will
// be overridden.
// You can manually include a content-type header in the call to all functions
// that modify the header.
func (g *GRequest) WithHeader(key string, value interface{}) *GRequest {
	if key == "" {
		g.addError(errtools.InvalidKeyError("header key cannot be empty"))
	}

	if typetools.IsStringlikeType(value) {
		stv := typetools.EnsureString(value)
		if stv == "" {
			g.addError(errtools.InvalidFieldError("header value cannot be empty"))
		}

		g.addHeader(key, stv)
	} else {
		g.addError(errtools.InvalidTypeError("header value must be a string or string-like type"))
	}

	return g
}

// Add multiple headers to the request
// Any string-like object can be passed as a value, and it will be converted to a string automatically
func (g *GRequest) WithHeaders(headers map[string]interface{}) *GRequest {
	for k, v := range headers {
		if typetools.IsStringlikeType(v) || typetools.IsNumericType(v) {
			g.addHeader(k, typetools.EnsureString(v))
		} else if val, ok := v.(bool); ok {
			g.addHeader(k, fmt.Sprintf("%t", val))
		} else {
			g.addError(errtools.InvalidTypeError(fmt.Sprintf("header value for %s must be a string or string-like type", k)))
		}
	}

	return g
}

func (g *GRequest) WithQueryParam(key, value string) *GRequest {
	if g.query == nil {
		g.query = &url.Values{}
	}

	g.query.Add(key, value)

	return g
}

// Add query parameters to the request
func (g *GRequest) WithQueryParams(params interface{}) *GRequest {
	if g.query == nil {
		g.query = &url.Values{}
	}

	switch val := params.(type) {
	case url.Values:
		for k, v := range val {
			for _, v2 := range v {
				g.query.Add(k, v2)
			}
		}

	case map[string]string:
		for k, v := range val {
			g.query.Add(k, v)
		}

	case map[string][]string:
		for k, v := range val {
			for _, v2 := range v {
				g.query.Add(k, v2)
			}
		}

	case map[string][]byte:
		for k, v := range val {
			g.query.Add(k, string(v))
		}

	case map[string]interface{}:
		for k, v := range val {
			if typetools.IsStringlikeType(v) || typetools.IsNumericType(v) {
				g.query.Add(k, typetools.EnsureString(v))
			} else {
				if vt, ok := v.([]string); ok {
					for _, v2 := range vt {
						g.query.Add(k, v2)
					}
				} else if vt, ok := v.(bool); ok {
					g.query.Add(k, fmt.Sprintf("%t", vt))
				} else {
					g.addError(errtools.InvalidTypeError(fmt.Sprintf("field %s - form value must be a string, string slice, or string-like type", k)))
				}
			}
		}
	default:
		g.addError(errtools.InvalidTypeError("query params must be a map[string]string, map[string][]string, map[string][]byte, map[string]interface{}, or url.Values"))
	}

	return g
}

// Validate the request to ensure no errors have popped up during creation
func (g *GRequest) Validate() error {
	if len(g.errs) > 0 {
		return errtools.MultipleErrors(g.errs)
	}

	if len(g.Url) == 0 || (strings.Contains(g.Url, "?") && strings.Split(g.Url, "?")[0] == "") {
		return errtools.InvalidFieldError("url cannot be empty")
	}

	if g.Method == "" || !strings.Contains("GETPOSTPUTDELETEPATCH", string(g.Method)) {
		return errtools.InvalidFieldError("method must be one of GET, POST, PUT, DELETE, or PATCH")
	}

	return nil
}

// TODO: Proxy from environment
// TODO: Timeout(s)
// TODO: Redirects
// TODO: Force attempt HTTP/2
func (g *GRequest) Execute() (*GResponse, error) {
	if err := g.Validate(); err != nil {
		return nil, err
	}

	if g.query != nil && len(*g.query) != 0 {
		if strings.Contains(g.Url, "?") {
			g.Url += "&"
		} else {
			g.Url += "?"
		}

		g.Url += g.query.Encode()
	}

	userAgentFound := false
	for k := range g.headers {
		if strings.ToLower(k) == "user-agent" {
			userAgentFound = true
			break
		}
	}

	if !userAgentFound {
		g.addHeader("User-Agent", "Clysec GREQ/1.0")
	}

	req, err := http.NewRequest(string(g.Method), g.Url, g.body)
	if err != nil {
		return nil, err
	}

	if g.headers != nil {
		for k, v := range g.headers {
			req.Header.Add(k, v)
		}
	}

	if g.client == nil {
		g.client = &http.Client{}
	}

	resp, err := g.client.Do(req)
	if err != nil {
		return nil, err
	}

	return &GResponse{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Response:   resp,
		bodyRead:   false,
	}, nil
}

func NewRequest(method Method, url string) *GRequest {
	return &GRequest{
		Method: method,
		Url:    url,
	}
}

func GetRequest(url string) *GRequest {
	return NewRequest(GET, url)
}

func PostRequest(url string) *GRequest {
	return NewRequest(POST, url)
}

func PutRequest(url string) *GRequest {
	return NewRequest(PUT, url)
}

func DeleteRequest(url string) *GRequest {
	return NewRequest(DELETE, url)
}

func PatchRequest(url string) *GRequest {
	return NewRequest(PATCH, url)
}
