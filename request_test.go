package greq_test

import (
	"fmt"
	"testing"

	"github.com/clysec/greq"
)

type HttpbinResponse struct {
	Args    map[string]string      `json:"args"`
	Headers map[string]interface{} `json:"headers"`
	Origin  string                 `json:"origin"`
	Url     string                 `json:"url"`
	Form    map[string]interface{} `json:"form"`
	Body    string                 `json:"data"`
	Json    map[string]interface{} `json:"json"`
}

type TestBody struct {
	Hello string `json:"hello" xml:"hello"`
}

func TestX(t *testing.T) {
	auth := greq.BasicAuth{
		Username: "user",
		Password: "password",
	}

	response, err := greq.GetRequest("https://httpbin.org/get").
		// Add basic authentication
		WithAuth(&auth).
		Execute()

	if err != nil {
		panic(err)
	}

	bodyString, err := response.BodyString()
	if err != nil {
		panic(err)
	}

	fmt.Println(bodyString)
}

func TestGetRequest(t *testing.T) {
	resp, err := greq.GetRequest("https://httpbin.org/get").
		WithHeader("Accept", "application/json").
		WithQueryParams(map[string]string{"key": "value"}).
		WithAuth(&greq.BasicAuth{Username: "username", Password: "password"}).
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
