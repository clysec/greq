package greq

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type JwtAlgorithm string

const (
	HS256 JwtAlgorithm = "HS256"
	HS384 JwtAlgorithm = "HS384"
	HS512 JwtAlgorithm = "HS512"
	RS256 JwtAlgorithm = "RS256"
	RS384 JwtAlgorithm = "RS384"
	RS512 JwtAlgorithm = "RS512"
	PS256 JwtAlgorithm = "PS256"
	PS384 JwtAlgorithm = "PS384"
	PS512 JwtAlgorithm = "PS512"
	ES256 JwtAlgorithm = "ES256"
	ES384 JwtAlgorithm = "ES384"
	ES512 JwtAlgorithm = "ES512"
)

type JwtAuth struct {
	Algorithm         JwtAlgorithm           `json:"algorithm"`
	Secret            interface{}            `json:"jwtSecret"`
	Payload           jwt.Claims             `json:"payload"`
	AdditionalHeaders map[string]interface{} `json:"additionalHeaders"`

	HeaderPrefix string `json:"headerPrefix"`

	method jwt.SigningMethod
	jwt    *jwt.Token
}

func (ja *JwtAuth) Prepare() error {
	method := jwt.GetSigningMethod(string(ja.Algorithm))
	if method == nil {
		return fmt.Errorf("invalid jwt algorithm: %s", ja.Algorithm)
	}

	ja.method = method

	ja.jwt = jwt.NewWithClaims(ja.method, ja.Payload)

	for k, v := range ja.AdditionalHeaders {
		ja.jwt.Header[k] = v
	}

	return nil
}

func (ja *JwtAuth) Apply(addHeaderFunc func(key, value string), setTransportFunc func(transport http.RoundTripper)) error {
	tokenString, err := ja.jwt.SignedString(ja.Secret)
	if err != nil {
		return err
	}

	if ja.HeaderPrefix == "" {
		ja.HeaderPrefix = "Bearer"
	}

	ja.HeaderPrefix = strings.TrimRight(ja.HeaderPrefix, " ")

	addHeaderFunc("Authorization", fmt.Sprintf("%s %s", ja.HeaderPrefix, tokenString))

	return nil
}
