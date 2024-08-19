# Query Parameters and Headers
Below functions set query parameters and headers respectively.

## Query Parameters

::: info

The WithQueryParams function should take most map-like objects, including structs, as long as the keys are strings and values are string-like/numeric/boolean

::: 

**Request**

```go
package main

import (
    "fmt"
    "github.com/clysec/greq"
)

func main() {
    response, err := greq.GetRequest("https://httpbin.org/get").
        
        // Set a single query parameter
        WithQueryParam("key", "value").

        // Set multiple query parameters
        // Duplicate keys are supported, and can be added
        // either by passing a []string as the value
        // or adding the same field multiple times
        WithQueryParams(map[string]string{
            "key1": "value1",
            "key2": "value2",
        }).
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

**Response**
```json
{
  "args": {
    "key": [
      "value", 
      "value2"
    ], 
    "key1": "value1", 
    "key2": "value2"
  }, 
  "headers": {
    "Accept-Encoding": "gzip", 
    "Host": "httpbin.org", 
    "User-Agent": "Clysec GREQ/1.0"
  }, 
  "origin": "1.2.3.4", 
  "url": "https://httpbin.org/get?key=value&key=value2&key1=value1&key2=value2"
}
```

## Headers

**Request**

```go
package main

import (
    "fmt"
    "github.com/clysec/greq"
)

func main() {
    response, err := greq.GetRequest("https://httpbin.org/get").
        
        // Set a single header
        WithHeader("Accept", "application/json").

        // Set multiple headers
        // Duplicate fields will be overwritten
        WithHeaders(map[string]interface{}{
            "Accept-Encoding": "gzip",
            "Host": "httpbin.org",
        }).
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

**Response**
```json
{
  "args": {}, 
  "headers": {
    "Accept": "application/json", 
    "Accept-Encoding": "gzip", 
    "Host": "httpbin.org", 
    "User-Agent": "Clysec GREQ/1.0", 
    "X-Amzn-Trace-Id": "Root=1-66c3360f-5360c1287a451b5757d7962b"
  }, 
  "origin": "1.2.3.4", 
  "url": "https://httpbin.org/get"
}
```