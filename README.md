# GREQ - Simple Go Request Library
GREQ is a simple request library for go with a simple API. It is designed to be simple to use, and to (as far as possible) not require a lot of documentation. The goal when writing this was as I, for the final time, got fed up with forgetting how to do a multipart request in go and having to look it up, and started adding the request library to our go function library ([gofn](http://github.com/scheiblingco/gofn)). As it grew in complexity, I decided to split it out into its own library.

## Installation
```bash
go get github.com/clysec/greq
```

## Usage
### Simple GET Request
```go
package main

import (
    "fmt"
    "github.com/clysec/greq"
)

func main() {
    resp, err := greq.GetRequest("https://httpbin.org/get").
        WithHeader("Accept", "application/json").
        WithQueryParams(map[string]string{"key": "value"}).
        Execute()
    
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(resp)
    }
}
```

### Multipart POST
```go
package main

import (
    "fmt"
    "github.com/clysec/greq"
)

func main() {
    file, err := os.Open("file.jpg")
    if err != nil {
        fmt.Println(err)
        return
    }

    multipart, err := greq.PostRequest("https://httpbin.org/post").
        WithHeader("Accept", "application/json").
        WithMultipartFormBody([]*greq.MultipartField{
            greq.NewMultipartField("string").WithStringValue("value"),
            grew.NewMultipartField("bytes").WithBytesValue([]byte("bytes")),
            greq.NewMultipartField("file").WithReaderValue(strings.NewReader("value")).WithFilename("file.txt").WithContentType("text/plain"),
            greq.NewMultipartField("file").WithFile(file, "file.jpg")
        }).Execute()
    if err != nil {
        fmt.Println(err)
    }
}
```

### JSON POST
```go
package main

import (
    "fmt"
    "github.com/clysec/greq"
)

type MyType struct {
    Key string `json:"key"`
    Value string `json:"value"`
}

func main() {
    resp, err := greq.PostRequest("https://httpbin.org/post").
        WithHeader("Accept", "application/json").
        WithJSONBody(MyType{Key: "key", Value: "value"}).
        Execute()
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(resp)
    }
}
```

### With Basic Auth
```go
package main

import (
    "fmt"
    "github.com/clysec/greq"
    greqauth "github.com/clysec/greq/auth"
)

type MyType struct {
    Key string `json:"key"`
    Value string `json:"value"`
}

func main() {
    resp, err := greq.PostRequest("https://httpbin.org/post").
        WithHeader("Accept", "application/json").
        WithJSONBody(MyType{Key: "key", Value: "value"}).
        WithAuth(greqauth.BasicAuth{Username: "username", Password: "password"}).
        Execute()
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(resp)
    }
}
```

### With Custom HTTP Client
```go
package main

import (
    "fmt"
    "github.com/clysec/greq"
    "net/http"
)


func main() {
    client := &http.Client{
        Timeout: time.Second * 10,
    }

    resp, err := greq.PostRequest("https://httpbin.org/post").
        WithHeader("Accept", "application/json").
        WithJSONBody(MyType{Key: "key", Value: "value"}).
        WithClient(client).
        Execute()
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(resp)
    }
}
```

### With Client Certificate Authentication
```go
package main

import (
    "fmt"
    "github.com/clysec/greq"
    greqauth "github.com/clysec/greq/auth"
)

func main() {
    clientCert := greqauth.ClientCert{
        InsecureSkipVerify: true
    }.FromX509("cert.pem", "key.pem")


    resp, err := greq.GetRequest("https://httpbin.org/get").
        WithHeader("Accept", "application/json").
        WithAuth().
        Execute()
}
```