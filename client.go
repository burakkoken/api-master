package apimaster

import (
	"net/http"
	"net/http/cookiejar"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	cookieJar, err := cookiejar.New(nil)

	if err != nil {
		panic(err)
	}

	return &Client{
		&http.Client{
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

func (client *Client) makeRequest(request *Request) *Response {
	return nil
}
