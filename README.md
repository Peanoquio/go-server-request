# goserverrequest
![Build Status](https://img.shields.io/travis/rust-lang/rust.svg)
A module built on top of Go net/http module to communicate with other servers

## Introduction

This library will make HTTP requests (CRUD operations) to a server (supporting RESTful APIs) based on the URL and the request payload.

---

## Tutorial

To make requests, simply follow this code example.
```go
import (
	"github.com/Peanoquio/goserverrequest/config"
	"github.com/Peanoquio/goserverrequest/factory"
)

// the configuration
config := &config.ServerCallConfig{
  MaxIdleConns:          50,
  IdleConnTimeout:       30,
  TLSHandshakeTimeout:   5,
  ExpectContinueTimeout: 2,
  //DisableCompression:    false,
  DialContextTimeout:   10,
  DialContextKeepAlive: 10,
  DialContextDualStack: true,
  HTTPClientTimeout:    10,
}

// create a new object instance through the factory
serverCallUtil := factory.NewServerCallUtil(config)

// the JSON payload for the request
jsonPayload := map[string]interface{}{
  "test":     "123",
  "greeting": "hello world!",
}

// make an POST API call
resp, err = serverCallUtil.Post("http://httpbin.org/post", jsonPayload)
if err != nil {
  t.Errorf(err.Error())
} else {
  t.Log("response from POST:", resp)
}
```
---

### Testing

To run the test script, execute this on your command line:
```shell
go test -run TestRequest
```

