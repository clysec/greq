# Raw Body/Bytes
To just send plain text or bytes in the body, you can use the helper functions below. You can set the content type via the header modification functions, or via the `WithContentType` function.

## Raw String

**Request**

```go
package main

import (
    "fmt"
    "github.com/clysec/greq"
)

func main() {
    resp, err := greq.PostRequest("https://httpbin.org/post").
		WithStringBody("Hello World").
		Execute()

	if err != nil {
		panic(err)
	}

	bodyString, err := resp.BodyString()
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
  "data": "Hello World", 
  "files": {}, 
  "form": {}, 
  "headers": {
    "Accept-Encoding": "gzip", 
    "Content-Length": "11", 
    "Host": "httpbin.org", 
    "User-Agent": "Clysec GREQ/1.0"
  }, 
  "json": null, 
  "origin": "1.2.3.4", 
  "url": "https://httpbin.org/post"
}
```

## Raw Bytes

**Request**

```go
package main

import (
    "fmt"
    "github.com/clysec/greq"
)

func main() {
    resp, err := greq.PostRequest("https://httpbin.org/post").
		WithByteBody([]byte("Hello World")).
        WithContentType("text/plain").
		Execute()

	if err != nil {
		panic(err)
	}

	bodyString, err := resp.BodyString()
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
  "data": "Hello World", 
  "files": {}, 
  "form": {}, 
  "headers": {
    "Accept-Encoding": "gzip", 
    "Content-Length": "11", 
    "Host": "httpbin.org", 
    "User-Agent": "Clysec GREQ/1.0"
  }, 
  "json": null, 
  "origin": "1.2.3.4", 
  "url": "https://httpbin.org/post"
}
```