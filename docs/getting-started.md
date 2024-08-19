# Getting Started
Below are some easy examples to get you started with GREQ

## Installation
```bash
go get github.com/clysec/greq
```

## Basic Usage
**Code**

```go
package main

import (
  "fmt"
  "github.com/clysec/greq"
)

func main() {
    // You can either use the helper functions greq.GetRequest, greq.PostRequest, greq.PutRequest, greq.PatchRequest, greq.DeleteRequest
    // or create a new request using greq.NewRequest with the URL and method as parameters

    // response, err := greq.NewRequest("GET", "https://httpbin.org/get").Execute()
    response, err := greq.GetRequest("https://httpbin.org/get").Execute()
    if err != nil {
        panic(err)
    }

    bodyString, err := response.BodyString()
	  if err != nil {
		    panic(err)
	  }


    fmt.Printf("Response Code: %d\n\n%s", response.StatusCode, response.BodyString())
}
```

**Output**

```bash
Response Code: 200

{
  "args": {}, 
  "headers": {
    "Accept-Encoding": "gzip", 
    "Host": "httpbin.org", 
    "User-Agent": "Clysec GREQ/1.0"
  }, 
  "origin": "1.2.3.4", 
  "url": "https://httpbin.org/get"
}
```

## Body Data
Body data can be sent with the request only for POST, PUT, PATCH requests. Most functions take an interface{} to allow for multiple data types, and it automatically detects the correct way to include the data.

::: tip
Even though you can pass any interface{}-compatible type to the body data functions, it needs to be serializable to JSON or XML for the respective functions and needs to be a string-like or key-value pair type for the form body functions. If you're missing any input type that you think should be included, please open an issue.
:::

**Code**

```go
package main

import (
  "fmt"
  "github.com/clysec/greq"
)

func main() {
    // Make a POST request with the application/x-www-form-urlencoded content type
    response, err := greq.PostRequest("https://httpbin.org/post").WithUrlencodedFormBody(
      map[string]string{
        "key1": "value1",
        "key2": "value2",
      },
    ).Execute()

    // Make a POST request with the application/json content type
    response, err := greq.PostRequest("https://httpbin.org/post").WithJSONBody(
      map[string]string{
        "key1": "value1",
        "key2": "value2",
      },
    ).Execute()

    // Make a POST request with the application/xml content type
    response, err := greq.PostRequest("https://httpbin.org/post").WithXMLBody(
      map[string]string{
        "key1": "value1",
        "key2": "value2",
      },
    ).Execute()
}
```

## Decoding responses
There are helper functions to decode a couple of different types of responses

::: tip
The response body is considered consumed when any of the helper functions are called and have extracted the data, and cannot be fetched again since the response body will be closed.
:::

**Code**

```go
package main

import (
  "fmt"
  "github.com/clysec/greq"
)

type HttpbinResponse struct {
	Args    map[string]string      `json:"args"`
	Headers map[string]interface{} `json:"headers"`
	Origin  string                 `json:"origin"`
	Url     string                 `json:"url"`
	Form    map[string]interface{} `json:"form"`
	Body    string                 `json:"data"`
}

func main() {
    resp, err := greq.GetRequest("https://httpbin.org/get").Execute()
    if err != nil {
        panic(err)
    }

    // Decode the response into a map[string]interface{} object
    // Note that the body is considered consumed when any of the helper
    // functions are called and have extracted the data, and cannot be fetched again
    // since the response body will be closed.
    var data map[string]interface{}

    err := resp.DecodeJSON(&data)
    if err != nil {
        panic(err)
    }

    // Decode into a struct
    var httpbinResponse HttpbinResponse
    err := resp.DecodeJSON(&httpbinResponse)
    if err != nil {
        panic(err)
    }

    // Get the body as a string
    bodystr, err := resp.BodyString()
    if err != nil {
        panic(err)
    }
}
```