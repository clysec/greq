# Certificate/mTLS Authentication
Authenticate with Mutual TLS Authentication (client certificates).

**Request**

```go
package main

import (
    "fmt"
    "github.com/clysec/greq"
)

func main() {
    // Get certificates from X509 cert/key-files
    auth := greq.NewClientCertificateAuth().FromX509("cert.pem", "key.pem")
    
    // Get certificates from X509 bytes
    auth := greq.NewClientCertificateAuth().FromX509Bytes(certBytes, keyBytes)

    // Get certificates from PKCS12 file
    auth := greq.NewClientCertificateAuth().FromPKCS12("cert.p12", "password")

    // Get certificates from PKCS12 bytes
    auth := greq.NewClientCertificateAuth().FromPKCS12Bytes(pkcs12Bytes, "password")

    // Add insecureSkipVerify
    auth = auth.WithInsecureSkipVerify(true)

    // Add CA Certificates
    auth = auth.WithCaCertificates(caCertificates)
        
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