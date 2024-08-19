# JSON/XML Body Marshalling
GREQ includes helper functions for marshalling JSON and/or XML body data. These functions use the Go standard library JSON marshaller

## JSON Body Example


::: info
WithJSONBody accepts any value that can be passed to json.Marshal, and additionally pre-rendered json in the form of string or []byte
::: 

```go
package main

import (
  "fmt"
  "github.com/clysec/greq"
)

func main() {
    data := map[string]string{
		"hello": "world",
	}

	response, err := greq.PostRequest("https://httpbin.org/post").
		WithHeader("Accept", "application/json").
		WithQueryParam("key", "value").
		WithJSONBody(data, nil).
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
  "url": "https://httpbin.org/post?key=value"
}
```

## XML Body Example

::: info
WithXMLBody accepts any value that can be passed to xml.Marshal, and additionally pre-rendered xml in the form of string or []byte
::: 

```go
package main

import (
  "fmt"
  "github.com/clysec/greq"
)

type TestBody struct {
	Hello string `xml:"hello"`
}

func main() {
    response, err := greq.PostRequest("https://httpbin.org/put").
        WithHeader("Accept", "application/json").
        WithQueryParam("key", "value").
        WithXMLBody(TestBody{Hello: "world"}, nil).
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
  "data": "<TestBody><hello>world</hello></TestBody>", 
  "files": {}, 
  "form": {}, 
  "headers": {
    "Accept": "application/json", 
    "Accept-Encoding": "gzip", 
    "Content-Length": "41", 
    "Content-Type": "application/xml", 
    "Host": "httpbin.org", 
    "User-Agent": "Clysec GREQ/1.0", 
    "X-Amzn-Trace-Id": "Root=1-66c32716-3a3f0d0e67055e6c52b5bc84"
  }, 
  "json": null, 
  "origin": "185.242.229.246", 
  "url": "https://httpbin.org/post?key=value"
}

```
