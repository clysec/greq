package greq_test

import (
	"testing"

	"github.com/clysec/greq"
	"github.com/clysec/greq/auth"
)

type HttpbinResponse struct {
	Args    map[string]string      `json:"args"`
	Headers map[string]interface{} `json:"headers"`
	Origin  string                 `json:"origin"`
	Url     string                 `json:"url"`
	Form    map[string]interface{} `json:"form"`
	Body    string                 `json:"data"`
}

func TestGetRequest(t *testing.T) {
	resp, err := greq.GetRequest("https://httpbin.org/get").
		WithHeader("Accept", "application/json").
		WithAuth(&auth.BasicAuth{Username: "username", Password: "password"}).
		Execute()

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Unexpected status code: %d", resp.StatusCode)
	}

	if resp.Response.Header.Get("Content-Type") != "application/json" {
		t.Fatalf("Unexpected content type: %s", resp.Response.Header.Get("Content-Type"))
	}

	body := &HttpbinResponse{}

	if err := resp.BodyUnmarshalJson(body); err != nil {
		t.Fatal(err)
	}

	if body.Headers["Accept"] != "application/json" {
		t.Fatalf("Unexpected Accept header: %s", body.Headers["Accept"])
	}

	if body.Headers["Authorization"] != "Basic dXNlcm5hbWU6cGFzc3dvcmQ=" {
		t.Fatalf("Unexpected Authorization header: %s", body.Headers["Authorization"])
	}
}
