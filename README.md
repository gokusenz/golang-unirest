# Unitrest for Golang

> Unirest is a set of Lightweight HTTP Request Client Libraries.

## Installation

As a library

```shell
go get github.com/gokusenz/unirest
```

## Usage

```go
package main

import (
    "github.com/gokusenz/unirest"
    "log"
    "os"
)

func main() {
  headers := make(map[string][]string)

  req := unirest.Request{
    BaseURL:     "www.google.com",
    Endpoint:    "/api/client",
    QueryString: "",
    Headers:     headers,
  }

  res, err := unirest.Get(req)
  if err != nil {
    fmt.Println(err)
  }

  // now do something with response or whatever
}
```