# NTLM Authentication
Authenticate with NTLM authentication

**Request**

```go
package main

import (
    "fmt"
    "github.com/clysec/greq"
)

func main() {
    auth := greq.NtlmAuth{
        Domain:   "INTERNAL",
        User:     "username",
        Password: "password",
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