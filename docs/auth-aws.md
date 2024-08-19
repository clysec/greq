# AWS Authentication
Adds an AWS authentication signature to the request.

**Request**

```go
package main

import (
    "fmt"
    "github.com/clysec/greq"
)

func main() {
    auth := greq.AwsAuth{
        AccessKey: "ABC",
        SecretKey: "DEF",
        Region: "us-east-1",
        ServiceName: "execute-api",
        SessionToken: "GHI",
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