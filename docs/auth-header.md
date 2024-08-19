# Header Authentication
Adds authentication data as a header to the request.

**Request**

```go
package main

import (
    "fmt"
    "github.com/clysec/greq"
)

func main() {
    auth := greq.HeaderAuth{
        Key: "Authorization",
        Value: "Token token=abcdefg",
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