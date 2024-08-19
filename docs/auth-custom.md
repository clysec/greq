# Custom Authentication
To add custom authentication modules to a request, you need to create a struct that implements the `Auth` interface. This interface has two methods:

```go

type Auth interface {
    Prepare() error
    Apply(addHeaderFunc func(key, value string), setTransportFunc(transport *http.Transport))
}
```

The `Prepare` method is called before the request is executed. For Oauth2, this part performs the authentication steps. The `Apply` method is called when the request is executed, and adds the required metadata to the request headers or transport.

Below are some examples of custom auth modules compatible with GREQ.


## Basic Auth
```go
type BasicAuth struct {
    Username string
    Password string
}

func (b BasicAuth) Prepare() error {
    return nil
}

func (b BasicAuth) Apply(addHeaderFunc func(key, value string), setTransportFunc func(transport *http.Transport)) {
    
    // Adds a request Authorization header with the basic auth credentials
    addHeaderFunc(
        "Authorization", 
        "Basic "+base64.StdEncoding.EncodeToString(
            []byte(b.Username+":"+b.Password),
        ),
    )
}
```

## mTLS Authentication
```go
type CertificateAuth struct {
    CertFile string
    KeyFile string
}

func (c CertificateAuth) Prepare() error {
    return nil
}

func (c CertificateAuth) Apply(addHeaderFunc func(key, value string), setTransportFunc func(transport *http.Transport)) {
    // Load the certificate and key files
    cert, err := tls.LoadX509KeyPair(c.CertFile, c.KeyFile)
    if err != nil {
        panic(err)
    }

    // Create a new transport with the certificate
    transport := &http.Transport{
        TLSClientConfig: &tls.Config{
            Certificates: []tls.Certificate{cert},
        },
    }

    // Set the transport
    setTransportFunc(transport)
}
```