# PATCH Requests
PATCH requests can include data in a range of different forms. Below are some basic examples of how to use the PATCH request function, for more information on the different body types and how to use them see the body data section in the menu to the left

**Code**

```go
package main

import (
  "fmt"
  "github.com/clysec/greq"
)

func main() {
    // You can either use the helper function greq.PatchRequest
    // or create a new request using greq.NewRequest with the URL and method as parameters

    // response, err := greq.NewRequest("PATCH", "https://httpbin.org/patch").Execute()
    response, err := greq.PatchRequest("https://httpbin.org/patch").
        WithHeader("Accept", "application/json").
        WithQueryParam("key", "value").
        WithStringBody(`{"hello": "world"}`).
        WithHeader("Content-Type", "application/json").
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
  "json": {
    "hello": "world"
  },
  "data": "{\"hello\": \"world\"}",
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

## Multipart Form Request

**Code**

```go
package main

import (
  "fmt"
  "github.com/clysec/greq"
)

func main() {
    multipartData := []*greq.MultipartField{
		greq.NewMultipartField("field1").WithStringValue("value1.1"),
		greq.NewMultipartField("field1").WithStringValue("value1.2"),
		greq.NewMultipartField("field2").WithBytesValue([]byte("value2")),
		greq.NewMultipartField("field3").WithContentType("text/plain").WithStringValue("value3").WithFilename("file.txt"),
	}

	response, err := greq.PatchRequest("https://httpbin.org/post").
		WithHeader("Accept", "application/json").
		WithQueryParam("key", "value").
		WithMultipartFormBody(multipartData).
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
    "key": "value"
  }, 
  "data": "", 
  "files": {
    "field3": "value3"
  }, 
  "form": {
    "field1": [
      "value1.1", 
      "value1.2"
    ],
    "field2": "value2"
  }, 
  "headers": {
    "Accept": "application/json", 
    "Accept-Encoding": "gzip", 
    "Content-Length": "476", 
    "Content-Type": "multipart/form-data; boundary=ad35b49540e6d537ffc122ae052b24efbf38d6977d9763046d81def68b99", 
    "Host": "httpbin.org", 
    "User-Agent": "Clysec GREQ/1.0"
  }, 
  "origin": "1.2.3.4", 
  "url": "https://httpbin.org/post?key=value"
}
```


## JSON Request

**Code**

```go
package main

import (
  "fmt"
  "github.com/clysec/greq"
)

func main() {
    response, err := greq.PatchRequest("https://httpbin.org/put").
        WithHeader("Accept", "application/json").
        WithQueryParam("key", "value").
        WithJSONBody(
          map[string]string{
            "hello": "world",
          },
        ).
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
  "json": {
    "hello": "world"
  },
  "data": "{\"hello\": \"world\"}",
  "headers": {
    "Accept-Encoding": "gzip", 
    "Host": "httpbin.org", 
    "User-Agent": "Clysec GREQ/1.0",
    "Accept": "application/json"
  }, 
  "origin": "1.2.3.4",
  "url": "https://httpbin.org/patch?key=value"
}

```