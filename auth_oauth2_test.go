package greq_test

import (
	"fmt"
	"testing"

	"github.com/clysec/greq"
	"github.com/oauth2-proxy/mockoidc"
)

func TestOauth2ClientCredentials(t *testing.T) {
	m, _ := mockoidc.Run()
	defer m.Shutdown()

	cfg := m.Config()

	oauth2 := greq.Oauth2Auth{
		AuthType:     greq.ClientCredentials,
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		DiscoveryUrl: m.DiscoveryEndpoint(),
	}

	if err := oauth2.Prepare(); err != nil {
		t.Fatalf("Prepare() failed: %v", err)
	}

	fmt.Printf("Token: %v\n", oauth2.TokenExpired())

}
