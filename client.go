package apimaster

import (
	"net/http"
	"net/http/cookiejar"
	"testing"
	"time"
)

type Client struct {
	suite      *Suite
	testing    *testing.T
	httpClient *http.Client
}

func newClient() *Client {
	cookieJar, err := cookiejar.New(nil)

	if err != nil {
		panic(err)
	}

	return &Client{
		httpClient: &http.Client{
			Jar: cookieJar,
		},
	}
}

func NewClient(t *testing.T) *Client {
	if t == nil {
		panic("t must not be nil")
	}

	client := newClient()
	client.testing = t
	return client
}

func NewClientWithSuite(suite *Suite) *Client {
	if suite == nil {
		panic("suite must not be nil")
	}

	client := newClient()
	client.suite = suite
	return client
}

func (client *Client) GET(url string) *Request {
	return newRequest(client, http.MethodGet, url)
}

func (client *Client) POST(url string) *Request {
	return newRequest(client, http.MethodPost, url)
}

func (client *Client) DELETE(url string) *Request {
	return newRequest(client, http.MethodDelete, url)
}

func (client *Client) PUT(url string) *Request {
	return newRequest(client, http.MethodPut, url)
}

func (client *Client) PATCH(url string) *Request {
	return newRequest(client, http.MethodPatch, url)
}

func (client *Client) TRACE(url string) *Request {
	return newRequest(client, http.MethodTrace, url)
}

func (client *Client) OPTIONS(url string) *Request {
	return newRequest(client, http.MethodOptions, url)
}

func (client *Client) HEAD(url string) *Request {
	return newRequest(client, http.MethodHead, url)
}

func (client *Client) CONNECT(url string) *Request {
	return newRequest(client, http.MethodConnect, url)
}

func (client *Client) makeRequest(request *Request) *Response {
	startTime := time.Now()
	response, err := client.httpClient.Do(request.httpRequest)

	if err != nil {
		// TODO
	}

	elapsedTime := time.Since(startTime)
	return newResponse(client, response, elapsedTime)
}
