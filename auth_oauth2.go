package greq

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Oauth2AuthType string

const (
	AuthorizationCode   Oauth2AuthType = "authorization_code"
	PasswordCredentials Oauth2AuthType = "password"
	ClientCredentials   Oauth2AuthType = "client_credentials"
	DeviceCode          Oauth2AuthType = "device_code"
)

type Oauth2Token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresAt    int    `json:"expires_at"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
}

type OidcDiscovery struct {
	Issuer                 string   `json:"issuer"`
	AuthorizationEndpoint  string   `json:"authorization_endpoint"`
	TokenEndpoint          string   `json:"token_endpoint"`
	UserinfoEndpoint       string   `json:"userinfo_endpoint"`
	JwksUri                string   `json:"jwks_uri"`
	RegistrationEndpoint   string   `json:"registration_endpoint"`
	IntrospectionEndpoint  string   `json:"introspection_endpoint"`
	EndSessionEndpoint     string   `json:"end_session_endpoint"`
	CheckSessionIframe     string   `json:"check_session_iframe"`
	GrantTypesSupported    []string `json:"grant_types_supported"`
	ResponseTypesSupported []string `json:"response_types_supported"`
	ClaimsSupported        []string `json:"claims_supported"`
	ScopesSupported        []string `json:"scopes_supported"`
}

func (od *OidcDiscovery) IsGrantTypeSupported(grantType string) bool {
	for _, gt := range od.GrantTypesSupported {
		if gt == grantType {
			return true
		}
	}
	return false
}

func (od *OidcDiscovery) IsResponseTypeSupported(responseType string) bool {
	for _, rt := range od.ResponseTypesSupported {
		if rt == responseType {
			return true
		}
	}
	return false
}

func (od *OidcDiscovery) IsClaimSupported(claim string) bool {
	for _, c := range od.ClaimsSupported {
		if c == claim {
			return true
		}
	}
	return false
}

func (od *OidcDiscovery) IsScopeSupported(scope string) bool {
	for _, s := range od.ScopesSupported {
		if s == scope {
			return true
		}
	}
	return false
}

type Oauth2Auth struct {
	AuthType Oauth2AuthType `json:"auth_type"`

	ClientID          string `json:"client_id"`
	ClientSecret      string `json:"client_secret"`
	CredentialsInBody bool   `json:"credentials_in_body"`

	RedirectURL string   `json:"redirect_url"`
	Scopes      []string `json:"scopes"`

	DiscoveryUrl string `json:"discovery_url"`

	AuthorizationUrl string `json:"authorization_url"`
	TokenUrl         string `json:"token_url"`
	UserinfoUrl      string `json:"userinfo_url"`

	AdditionalBodyFields map[string]string `json:"additional_body_fields"`

	discovery *OidcDiscovery
	token     *Oauth2Token
}

func (oa *Oauth2Auth) TokenExpired() bool {
	return oa.token.ExpiresAt < int(time.Now().Unix())
}

func (oa *Oauth2Auth) Token() *Oauth2Token {
	return oa.token
}

func (oa *Oauth2Auth) Prepare() error {
	if oa.token != nil && oa.token.AccessToken != "" && !oa.TokenExpired() {
		return nil
	}

	if oa.AuthType == "" {
		return fmt.Errorf("auth_type is required")
	}

	if oa.AuthType == ClientCredentials {
		if oa.ClientID == "" || oa.ClientSecret == "" || (oa.DiscoveryUrl == "" && oa.TokenUrl == "") {
			return fmt.Errorf("client_id, client_secret and discovery_url or authorization_url and token_url are required")
		}

		if oa.TokenUrl == "" {
			metadata, err := GetRequest(oa.DiscoveryUrl).Execute()
			if err != nil {
				return err
			}

			if err := metadata.BodyUnmarshalJson(&oa.discovery); err != nil {
				return err
			}

			oa.TokenUrl = oa.discovery.TokenEndpoint
			oa.AuthorizationUrl = oa.discovery.AuthorizationEndpoint

			if !oa.discovery.IsGrantTypeSupported("client_credentials") {
				return fmt.Errorf("client_credentials grant type is not supported by this provider")
			}
		}

		request := PostRequest(oa.TokenUrl)

		if oa.CredentialsInBody {
			body := map[string]string{
				"client_id":     oa.ClientID,
				"client_secret": oa.ClientSecret,
				"grant_type":    "client_credentials",
			}

			if len(oa.Scopes) > 0 {
				body["scope"] = strings.Join(oa.Scopes, " ")
			}

			if len(oa.AdditionalBodyFields) > 0 {
				for k, v := range oa.AdditionalBodyFields {
					body[k] = v
				}
			}

			request = request.WithUrlencodedFormBody(body, nil)
		} else {
			request = request.WithAuth(&BasicAuth{
				Username: oa.ClientID,
				Password: oa.ClientSecret,
			})

			body := map[string]string{
				"grant_type": "client_credentials",
			}

			if len(oa.AdditionalBodyFields) > 0 {
				for k, v := range oa.AdditionalBodyFields {
					body[k] = v
				}
			}

			if len(oa.Scopes) > 0 {
				body["scope"] = strings.Join(oa.Scopes, " ")
			}

			request = request.WithUrlencodedFormBody(body, nil)
		}

		resp, err := request.Execute()
		if err != nil {
			return err
		}

		if resp.StatusCode != 200 {
			bodystr, _ := resp.BodyString()
			return fmt.Errorf("unexpected status code: %d - %s", resp.StatusCode, bodystr)
		}

		if err := resp.BodyUnmarshalJson(&oa.token); err != nil {
			return err
		}

		if oa.token.ExpiresAt == 0 && oa.token.ExpiresIn != 0 {
			oa.token.ExpiresAt = int(time.Now().Unix()) + oa.token.ExpiresIn
		}

		return nil
	}

	// TODO: Implement Rest
	return fmt.Errorf("not implemented")

}

func (oa *Oauth2Auth) Apply(addHeaderFunc func(key, value string), setTransportFunc func(transport http.RoundTripper)) error {
	if oa.token == nil || oa.TokenExpired() {
		err := oa.Prepare()
		if err != nil {
			return err
		}
	}

	if oa.token.TokenType == "" {
		addHeaderFunc("Authorization", fmt.Sprintf("Bearer %s", oa.token.AccessToken))
		return nil
	}

	addHeaderFunc("Authorization", fmt.Sprintf("%s %s", oa.token.TokenType, oa.token.AccessToken))

	return nil
}
