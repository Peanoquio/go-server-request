package core

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/Peanoquio/goserverrequest/config"
)

// RequestType contains the HTTP request types for CRUD operations
type RequestType struct {
	POST   string
	GET    string
	PUT    string
	PATCH  string
	DELETE string
}

var requestType = &RequestType{
	POST:   "POST",
	GET:    "GET",
	PUT:    "PUT",
	PATCH:  "PATCH",
	DELETE: "DELETE",
}

// ServerCall is the struct class for making server-side calls
type ServerCall struct {
	config        *config.ServerCallConfig
	httpClient    *http.Client
	httpTransport *http.Transport
}

// SetConfig sets the configurations needed for this class
func (sc *ServerCall) SetConfig(config *config.ServerCallConfig) {
	// ensure that there is a default value because 0 means there is no limit
	if config.MaxIdleConns == 0 {
		config.MaxIdleConns = 50
	}
	if config.IdleConnTimeout == 0 {
		config.IdleConnTimeout = 30
	}
	if config.TLSHandshakeTimeout == 0 {
		config.TLSHandshakeTimeout = 5
	}
	if config.ExpectContinueTimeout == 0 {
		config.ExpectContinueTimeout = 2
	}
	if config.DialContextTimeout == 0 {
		config.DialContextTimeout = 10
	}
	if config.DialContextKeepAlive == 0 {
		config.DialContextKeepAlive = 10
	}
	if config.HTTPClientTimeout == 0 {
		config.HTTPClientTimeout = 10
	}

	sc.config = config
}

// Init will initialize this class so it would be able to make requests
func (sc *ServerCall) Init() {
	// NOTE: Clients and Transports are safe for concurrent use by multiple goroutines and for efficiency should only be created once and re-used.
	// transport : for control over proxies, TLS configuration, keep-alives, compression, and other settings
	sc.httpTransport = &http.Transport{
		MaxIdleConns:          sc.config.MaxIdleConns,
		IdleConnTimeout:       sc.config.IdleConnTimeout * time.Second,
		TLSHandshakeTimeout:   sc.config.TLSHandshakeTimeout * time.Second,
		ExpectContinueTimeout: sc.config.ExpectContinueTimeout * time.Second,
		DisableCompression:    sc.config.DisableCompression,
		DialContext: (&net.Dialer{
			Timeout:   sc.config.DialContextTimeout * time.Second,
			KeepAlive: sc.config.DialContextKeepAlive * time.Second,
			DualStack: sc.config.DialContextDualStack,
		}).DialContext,
	}

	// client : the HTTP client that will make the actual POST request
	sc.httpClient = &http.Client{
		Timeout:   sc.config.HTTPClientTimeout * time.Second,
		Transport: sc.httpTransport,
	}
}

// Get makes a GET request with the JSON payload to the URL
func (sc *ServerCall) Get(url string, jsonPayload map[string]interface{}) (map[string]interface{}, error) {
	jsonObj, err := sc.requestJSONObj(requestType.GET, url, jsonPayload)
	return jsonObj, err
}

// Post makes a POST request with the JSON payload to the URL
func (sc *ServerCall) Post(url string, jsonPayload map[string]interface{}) (map[string]interface{}, error) {
	jsonObj, err := sc.requestJSONObj(requestType.POST, url, jsonPayload)
	return jsonObj, err
}

// Put makes a PUT request with the JSON payload to the URL
func (sc *ServerCall) Put(url string, jsonPayload map[string]interface{}) (map[string]interface{}, error) {
	jsonObj, err := sc.requestJSONObj(requestType.PUT, url, jsonPayload)
	return jsonObj, err
}

// Patch makes a PATCH request with the JSON payload to the URL
func (sc *ServerCall) Patch(url string, jsonPayload map[string]interface{}) (map[string]interface{}, error) {
	jsonObj, err := sc.requestJSONObj(requestType.PATCH, url, jsonPayload)
	return jsonObj, err
}

// Delete makes a DELETE request with the JSON payload to the URL
func (sc *ServerCall) Delete(url string, jsonPayload map[string]interface{}) (map[string]interface{}, error) {
	jsonObj, err := sc.requestJSONObj(requestType.DELETE, url, jsonPayload)
	return jsonObj, err
}

// requestJSONObj is the helper function to make HTTP requests in JSON format (returns a JSON object as a response)
func (sc *ServerCall) requestJSONObj(method string, url string, jsonPayload map[string]interface{}) (jsonObj map[string]interface{}, err error) {
	resp, err := sc.requestJSONStr(method, url, jsonPayload)
	if err != nil {
		return nil, err
	}
	// this will ensure returning an error back to the caller once a panic happens
	// so that the request will not be left hanging without a response
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	// parse the response into a JSON object
	err = json.Unmarshal([]byte(resp), &jsonObj)
	if err != nil {
		panic(err)
	}

	return
}

// requestJSONStr is the helper function to make HTTP requests in JSON format (returns a JSON string as a response)
func (sc *ServerCall) requestJSONStr(method string, url string, jsonPayload map[string]interface{}) (jsonRes string, err error) {
	// this will ensure returning an error back to the caller once a panic happens
	// so that the request will not be left hanging without a response
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	// parse the JSON payload
	byteArr, err := json.Marshal(jsonPayload)
	if err != nil {
		panic(err)
	}

	// create a POST request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(byteArr))
	if err != nil {
		panic(err)
	} else {
		// set the header request to JSON format
		req.Header.Set("Content-Type", "application/json")
		// the HTTP client that will make the actual POST request
		resp, err := sc.httpClient.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		jsonRes = string(contents)
	}

	return
}
