# Raw String/Bytes
The body can be fetched as a string or bytes using the `BodyString` and `BodyBytes` functions. The `BodyString` function returns the body as a string, while the `BodyBytes` function returns the body as a byte slice.

**Request**

```go
package main

import (
    "fmt"
    "github.com/clysec/greq"
)

func main() {
    response, err := greq.GetRequest("https://httpbin.org/get").
        Execute()

    if err != nil {
        panic(err)
    }

    // Get the body as a string
    bodyString, err := response.BodyString()
    if err != nil {
        panic(err)
    }

    fmt.Println(bodyString)

    // Get the body as a byte slice
    bodyBytes, err := response.BodyBytes()
    if err != nil {
        panic(err)
    }

    fmt.Println(bodyBytes)
}
```