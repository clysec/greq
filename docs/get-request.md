# GET Requests
GET requests are the simplest type of request. They are used to retrieve data from the server. Below is an example of a simple GET request using GREQ.

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

    // response, err := greq.NewRequest("GET", "https://httpbin.org/get").Execute()
    response, err := greq.GetRequest("https://httpbin.org/get").
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
  "url": "https://httpbin.org/get"
}
```


