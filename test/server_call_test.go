package test

import (
	"testing"

	"github.com/Peanoquio/goserverrequest/config"
	"github.com/Peanoquio/goserverrequest/factory"
)

// TestRequest is the test script for making requests to another server
func TestRequest(t *testing.T) {
	// test making server calls
	testServerCall(t)
}

// testServerCall will test making server calls to other servers based on the URL
func testServerCall(t *testing.T) {
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

	serverCallUtil := factory.NewServerCallUtil(config)

	// make an POST API call to an invalid URL
	resp, err := serverCallUtil.Post("http://whereami/whatamidoing", map[string]interface{}{"test": 123})
	if err != nil {
		t.Log("error from invalid POST:", err.Error())
		t.Log("response from invalid POST:", resp)
	} else {
		t.Log("response from invalid POST:", resp)
		t.Errorf("should return an error from invalid POST")
	}

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

	// make a PUT API call
	resp, err = serverCallUtil.Put("http://httpbin.org/put", jsonPayload)
	if err != nil {
		t.Errorf(err.Error())
	} else {
		t.Log("response from PUT:", resp)
	}

	// make a PATCH API call
	resp, err = serverCallUtil.Patch("http://httpbin.org/patch", jsonPayload)
	if err != nil {
		t.Errorf(err.Error())
	} else {
		t.Log("response from PATCH:", resp)
	}

	// make a GET API call
	resp, err = serverCallUtil.Get("http://httpbin.org/get", jsonPayload)
	if err != nil {
		t.Errorf(err.Error())
	} else {
		t.Log("response from GET:", resp)
	}

	// make a DELETE API call
	resp, err = serverCallUtil.Delete("http://httpbin.org/delete", jsonPayload)
	if err != nil {
		t.Errorf(err.Error())
	} else {
		t.Log("response from DELETE:", resp)
	}
}
