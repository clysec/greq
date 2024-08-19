# multipart/form-data Body

## Notice
The Multipart body function needs a slice of MultipartField-objects. You can generate these from a map using the `greq.MultipartFieldsFromMap` function, or manually field by field. 

## Field Types
The MultipartField supports multiple value types. You can use the same key more than once, the existing data will be appended.

```go
// String field. Can add multiple fields 
stringField := greq.NewMultipartField("key").WithStringValue("value")

byteField := greq.NewMultipartField("key").WithBytesValue([]byte("value"))

readerField := greq.NewMultipartField("key").WithReaderValue(*io.Reader)

pipeField := greq.NewMultipartField("key").WithPipeValue(*io.PipeReader)

// This field gets the filename automatically from the os.File handle
osfile, _ := os.Open("file.txt")
fileField := greq.NewMultipartField("file").WithFile(osfile, "text/plain")

// You can also specify the filename manually
// This has to be done after calling WithFile
fileField := greq.NewMultipartField("file").WithFile(osfile, "text/plain").WithFilename("file.txt")

// You can also compose a file manually
fileField := greq.NewMultipartField("file").WithStringContent("Hello World").WithFilename("file.txt").WithContentType("text/plain")
```

## Make a multipart request

**Request**
```go

  multipartFields := []*greq.MultipartField{
    greq.NewMultipartField("hello").WithStringValue("world"),
    greq.NewMultipartField("test").WithBytesValue([]byte("test")),
    greq.NewMultipartField("abc").WithIntValue(123),
    greq.NewMultipartField("def").WithFloatValue(123.456),
  }

	resp, err := greq.PostRequest("https://httpbin.org/post").
		WithMultipartFormBody(multipartFields).
		Execute()

	if err != nil {
		panic(err)
	}

	bodyString, err := resp.BodyString()
	if err != nil {
		panic(err)
	}

	fmt.Println(bodyString)

```

**Response**

```json
{
  "args": {}, 
  "data": "", 
  "files": {}, 
  "form": {
    "abc": "123", 
    "def": "123.456", 
    "hello": "world", 
    "test": "test"
  }, 
  "headers": {
    "Accept-Encoding": "gzip", 
    "Content-Length": "536", 
    "Content-Type": "multipart/form-data; boundary=0e3fe7a96261c7b815428ae56d9e792882369e66ce46ae948f9a6f48b196", 
    "Host": "httpbin.org", 
    "User-Agent": "Clysec GREQ/1.0"
  }, 
  "json": null, 
  "origin": "1.2.3.4", 
  "url": "https://httpbin.org/post"
}
```

## Create from map

**Request**

```go
package main

import (
  "fmt"
  "github.com/clysec/greq"
)

func main() {
   multipartFields, err := greq.MultipartFieldsFromMap(map[string]interface{}{
		"hello": "world",
		"test":  []byte("test"),
		"abc":   123,
		"def":   123.456,
	})

	if err != nil {
		panic(err)
	}

	resp, err := greq.PostRequest("https://httpbin.org/post").
		WithMultipartFormBody(multipartFields).
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
    "abc": "123", 
    "def": "123.456", 
    "hello": "world", 
    "test": "test"
  }, 
  "headers": {
    "Accept-Encoding": "gzip", 
    "Content-Length": "536", 
    "Content-Type": "multipart/form-data; boundary=0e3fe7a96261c7b815428ae56d9e792882369e66ce46ae948f9a6f48b196", 
    "Host": "httpbin.org", 
    "User-Agent": "Clysec GREQ/1.0"
  }, 
  "json": null, 
  "origin": "1.2.3.4", 
  "url": "https://httpbin.org/post"
}
```