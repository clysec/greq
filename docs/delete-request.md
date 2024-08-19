# DELETE Requests
DELETE requests are similar to GET requests, but they are generally used to delete data from the server. A DELETE request cannot contain a body.

**Code**

```go
package main

import (
  "fmt"
  "github.com/clysec/greq"
)

func main() {
    // You can either use the helper function greq.GetRequest
    // or create a new request using greq.NewRequest with the URL and method as parameters

    // response, err := greq.NewRequest("DELETE", "https://httpbin.org/delete").Execute()
    response, err := greq.DeleteRequest("https://httpbin.org/delete").
        WithHeader("Accept", "application/json").
        WithQueryParam("key", "value").
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

**Result**

```json
{
  "args": {
    "key": "value"
  }, 
  "headers": {
    "Accept-Encoding": "gzip", 
    "Host": "httpbin.org", 
    "User-Agent": "Clysec GREQ/1.0",
    "Accept": "application/json"
  }, 
  "origin": "1.2.3.4", 
  "url": "https://httpbin.org/delete?key=value"
}
```


