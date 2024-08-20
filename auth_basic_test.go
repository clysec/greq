package greq_test

import (
	"testing"

	"github.com/clysec/greq"
)

func TestBasicAuth(t *testing.T) {
	auth := greq.BasicAuth{
		Username: "user",
		Password: "pass",
	}

	resp, err := greq.GetRequest("https://httpbin.org/basic-auth/user/pass").WithAuth(&auth).Execute()
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("expected status code 200, got %d", resp.StatusCode)
	}

	var out map[string]interface{}

	if err := resp.BodyUnmarshalJson(&out); err != nil {
		t.Fatal(err)
	}

	if v, ok := out["authenticated"].(bool); !ok || !v {
		t.Fatalf("Not authenticated or error occured: %v", out)
	}
}

func TestHiddenBasicAuth(t *testing.T) {
	auth := greq.BasicAuth{
		Username: "user",
		Password: "pass",
	}

	resp, err := greq.GetRequest("https://httpbin.org/hidden-basic-auth/user/pass").WithAuth(&auth).Execute()
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("expected status code 200, got %d", resp.StatusCode)
	}

	var out map[string]interface{}

	if err := resp.BodyUnmarshalJson(&out); err != nil {
		t.Fatal(err)
	}

	if v, ok := out["authenticated"].(bool); !ok || !v {
		t.Fatalf("Not authenticated or error occured: %v", out)
	}
}
