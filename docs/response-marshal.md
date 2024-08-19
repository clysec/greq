# XML/JSON Responses
Functionality to decode responses to XML and JSON is provided by the `BodyUnmarshalXML` and `BodyUnmarshalJSON` functions. These functions are used to decode the response body to a struct, map or slice.

**Request**

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
	Json    map[string]interface{} `json:"json"`
}

func main() {
    response, err := greq.GetRequest("https://httpbin.org/get").
        Execute()

    if err != nil {
        panic(err)
    }

    var data HttpbinResponse

    // Autodetect the body type (json, xml, yaml) and unmarshal it to the struct
    // Note that the body can only be accessed once. If you need to access the body multiple types, 
    // use the BodyBytes function to get the raw bytes and copy them to a new buffer.
    if err = response.BodyUnmarshal(&data); err != nil {
        panic(err)
    }

    // Or you can use the specific functions
    if err = response.BodyUnmarshalJSON(&data); err != nil {
        panic(err)
    }

    // Or you can use the specific functions
    if err = response.BodyUnmarshalXML(&data); err != nil {
        panic(err)
    }

    fmt.Println(data)
}
```