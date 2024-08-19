# URL-Encoded Form Data

## Example

**Request**

::: info
WithUrlencodedFormBody accepts the following data types:

String-like (string, []byte, numbers)
Maps (map[string]interface{} where interface is string, []byte, numbers, []string or [][]byte)
// TODO: Ensure bson.M works
Map-like types/type aliases (bson.M, url.Values, etc)

:::


```go
package main

import (
  "fmt"
  "github.com/clysec/greq"
)

func main() {
    resp, err := greq.PostRequest("https://httpbin.org/post").
        WithUrlencodedFormBody(map[string]string{
            "hello": "world",
        }, nil).
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
  "data": "", 
  "files": {}, 
  "form": {
    "hello": "world"
  }, 
  "headers": {
    "Accept-Encoding": "gzip", 
    "Content-Length": "11", 
    "Content-Type": "application/x-www-form-urlencoded", 
    "Host": "httpbin.org", 
    "User-Agent": "Clysec GREQ/1.0"
  }, 
  "json": null, 
  "origin": "1.2.3.4", 
  "url": "https://httpbin.org/post"
}
```