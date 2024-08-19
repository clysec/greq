# JWT Authentication
Generate and sign a JWT to authenticate

**Request**

```go
package main

import (
    "fmt"
    "github.com/clysec/greq"
)

func main() {
    auth := greq.JwtAuth{
        Algorithm: greq.JwtAlgorithm.HS256,
        Secret: "abcdef",
        Payload: jwt.Claims{
            "sub": "1234567890",
            "name": "John Doe",
            "admin": true,
        },
        AdditionalHeaders: map[string]interface{}{
            "kid": "1234567890",
        },
        HeaderPrefix: "Bearer",
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