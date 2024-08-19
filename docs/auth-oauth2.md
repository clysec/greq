# Oauth2 Authentication
Authenticates with Oauth2. There are a number of different flows supported, see the examples below

## Client Credentials
**Request**

```go
package main

import (
    "fmt"
    "github.com/clysec/greq"
)

func main() {
    // This authentication object can be reused for multiple requests.
    // It will manage token renewal and caching automatically
    auth := greq.Oauth2Auth{
        AuthType:           greq.ClientCredentials,
        Scopes:             []string{"openid", "profile"},
        ClientID:           "my_client_id",
        ClientSecret:       "my_client_secret",

        // You can specify the discovery endpoint, or manually specify the endpoint URLs
        DiscoveryEndpoint:  "https://idpea.org/.well-known/openid-configuration",
        
        // You can specify the discovery endpoint, or manually specify the endpoint URLs
        TokenEndpoint:     "https://idpea.org/token",
    }
        
    response, err := greq.GetRequest("https://httpbin.org/get").
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
```

## Authorization Code
**Request**

```go
package main

import (
    "fmt"
    "github.com/clysec/greq"
)

func main() {
    // This authentication object can be reused for multiple requests.
    // It will manage token renewal and caching automatically

    // For the authorization code flow, you will need to include http://localhost:[port]/callback in the redirect URIs of your IDP
    // you can specify this port in the ListenCallback function
    auth := greq.Oauth2Auth{
        AuthType:               greq.AuthorizationCode,
        Scopes:                 []string{"openid", "profile"},
        ClientID:               "my_client_id",
        ClientSecret:           "my_client_secret",

        // You can specify the discovery endpoint, or manually specify the endpoint URLs
        DiscoveryEndpoint:      "https://idpea.org/.well-known/openid-configuration",
        
        // You can specify the discovery endpoint, or manually specify the endpoint URLs
        AuthorizationEndpoint:  "https://idpea.org/authorize",
        TokenEndpoint:          "https://idpea.org/token",
    }

    // Since this is an authorization code flow, you will need to open a browser to the authorization endpoint
    // and get the code from the redirect URL
    fmt.Println("Open this URL in your browser: ", auth.GetAuthorizationURL())

    err := auth.ListenCallback(":8080")
    if err != nil {
        panic(err)
    }

    // For renewals, the script will stop and wait for the user to visit the callback URL
    // If you want to control this manually, you can call the GetAuthorizationURL() and ListenCallback()
    // functions manually
    response, err := greq.GetRequest("https://httpbin.org/get").
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
```


## Device Code Authentication
**Request**

```go
package main

import (
    "fmt"
    "github.com/clysec/greq"
)

func main() {
    // This authentication object can be reused for multiple requests.
    // It will manage token renewal and caching automatically
    auth := greq.Oauth2Auth{
        AuthType:               greq.DeviceCode,
        Scopes:                 []string{"openid", "profile"},
        ClientID:               "my_client_id",
        ClientSecret:           "my_client_secret",

        // You can specify the discovery endpoint, or manually specify the endpoint URLs
        DiscoveryEndpoint:      "https://idpea.org/.well-known/openid-configuration",
        
        // You can specify the discovery endpoint, or manually specify the endpoint URLs
        AuthorizationEndpoint:  "https://idpea.org/authorize",
        TokenEndpoint:          "https://idpea.org/token"
    }

    // Since this is an device auth flow, you will need to open a browser to the authorization endpoint
    // and get the code from the redirect URL
    fmt.Println("Open this URL in your browser: ", auth.GetDeviceAuthorizationURL())

    // You can also get the URL and code separately
    fmt.Printf("URL: %s\nCode: %s\n", auth.GetDeviceBaseUrl(), auth.GetDeviceCode())

    // Or display a QR code in the terminal
    auth.DisplayDeviceQRCode()

    err := auth.ListenCallback(":8080")
    if err != nil {
        panic(err)
    }

    // For renewals, the script will stop and wait for the user to visit the callback URL
    // If you want to control this manually, you can call the GetAuthorizationURL() and ListenCallback()
    // functions manually
    response, err := greq.GetRequest("https://httpbin.org/get").
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
```