# Basic Authentication
Adds a basic authentication header to the request.

**Request**

```go
package main

import (
    "fmt"
    "github.com/clysec/greq"
)

func main() {
    auth := greq.BasicAuth{
        Username: "user",
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