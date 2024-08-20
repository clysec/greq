package greq_test

import (
	"testing"

	"github.com/clysec/greq"
)

func TestBearerAuth(t *testing.T) {
	auth1 := greq.BearerAuth{
		Token: "my-access-token",
	}

	auth2 := greq.BearerAuth{
		Token:  "my-access-token",
		Prefix: "Bearer",
	}

	resp1, err := greq.GetRequest("https://httpbin.org/bearer").WithAuth(&auth1).Execute()
	if err != nil {
		t.Fatal(err)
	}

	resp2, err := greq.GetRequest("https://httpbin.org/bearer").WithAuth(&auth2).Execute()
	if err != nil {
		t.Fatal(err)
	}

	if resp1.StatusCode != resp2.StatusCode || resp1.StatusCode != 200 {
		t.Errorf("expected status code 200, got %d and %d", resp1.StatusCode, resp2.StatusCode)
	}

	var out1 map[string]interface{}
	var out2 map[string]interface{}

	if err := resp1.BodyUnmarshalJson(&out1); err != nil {
		t.Fatal(err)
	}

	if err := resp2.BodyUnmarshalJson(&out2); err != nil {
		t.Fatal(err)
	}

	if v, ok := out1["authenticated"].(bool); !ok || !v {
		t.Fatalf("Not authenticated or error occured: %v", out1)
	}

	if v, ok := out2["authenticated"].(bool); !ok || !v {
		t.Fatalf("Not authenticated or error occured: %v", out2)
	}

	if v, ok := out1["token"].(string); !ok || v != "my-access-token" {
		t.Fatalf("Token not found or incorrect: %v", out1)
	}

	if v, ok := out2["token"].(string); !ok || v != "my-access-token" {
		t.Fatalf("Token not found or incorrect: %v", out2)
	}
}
