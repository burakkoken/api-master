package apimaster

import (
	"github.com/burakkoken/api-master/context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"testing"
	"time"
)

type Client struct {
	testing    *testing.T
	ctx        *context.Context
	httpClient *http.Client
}

func NewClient(t *testing.T) *Client {
	if t == nil {
		panic("t must not be nil")
	}

	cookieJar, err := cookiejar.New(nil)

	if err != nil {
		panic(err)
	}

	return &Client{
		testing: t,
		ctx:     context.NewContext(),
		httpClient: &http.Client{
			Jar: cookieJar,
		},
	}
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

func (client *Client) GetContext() *context.Context {
	return client.ctx
}

func (client *Client) GetCookies(u *url.URL) []*http.Cookie {
	return client.httpClient.Jar.Cookies(u)
}

func (client *Client) makeRequest(request *Request) *Response {
	startTime := time.Now()
	response, err := client.httpClient.Do(request.httpRequest)

	if err != nil {
		assert.NoError(client.testing, err, "request failed!")
	}

	elapsedTime := time.Since(startTime)
	return newResponse(client, response, elapsedTime)
}
