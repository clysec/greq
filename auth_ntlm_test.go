package greq_test

import (
	"fmt"
	"testing"

	"github.com/clysec/greq"
)

func TestNtlmAuth(t *testing.T) {
	auth := greq.NTLMAuth{
		Username: "user",
		Password: "pass",
	}

	resp, err := greq.GetRequest("https://authenticationtest.com/HTTPAuth/").WithAuth(&auth).Execute()
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("expected status code 200, got %d", resp.StatusCode)
	}

	fmt.Println(resp.BodyString())
}
