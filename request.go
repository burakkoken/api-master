package apimaster

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const HeaderContentType = "Content-Type"

const (
	ContentTypeApplicationJson = "application/json"
	ContentTypeApplicationXml  = "application/xml"
	ContentTypeTextPlain       = "text/plain"
)

type Request struct {
	client      *Client
	httpRequest *http.Request
}

func newRequest(client *Client, method string, rawUrl string) *Request {
	parsedUrl, err := url.Parse(rawUrl)

	if err != nil {
		panic(err)
	}

	var request *http.Request
	request, err = http.NewRequest(method, parsedUrl.String(), nil)

	if err != nil {
		panic(err)
	}

	return &Request{
		client:      client,
		httpRequest: request,
	}
}

func (r *Request) WithRequestParameter(name string, value string) {
	queryValues := r.httpRequest.URL.Query()
	queryValues.Add(name, value)
	r.httpRequest.URL.RawQuery = queryValues.Encode()
}

func (r *Request) WithPathVariable(name string, value string) *Request {
	r.httpRequest.URL.Path = strings.ReplaceAll(r.httpRequest.URL.Path, name, value)
	return r
}

func (r *Request) WithHeader(key string, value string) *Request {
	r.httpRequest.Header.Add(key, value)
	return r
}

func (r *Request) WithHeaders(headers map[string]string) *Request {
	if headers != nil {

		for key, value := range headers {
			r.httpRequest.Header.Add(key, value)
		}

	}
	return r
}

func (r *Request) WithJson(object interface{}) *Request {
	jsonValue, err := json.Marshal(object)

	if err != nil {
		panic(err)
	}

	r.httpRequest.Header.Set(HeaderContentType, ContentTypeApplicationJson)
	r.httpRequest.Body = ioutil.NopCloser(bytes.NewReader(jsonValue))
	r.httpRequest.ContentLength = int64(len(jsonValue))

	return r
}

func (r *Request) WithText(text string) *Request {
	value := []byte(text)

	r.httpRequest.Header.Set(HeaderContentType, ContentTypeTextPlain)
	r.httpRequest.Body = ioutil.NopCloser(bytes.NewReader(value))
	r.httpRequest.ContentLength = int64(len(value))

	return r
}

func (r *Request) WithXml(object interface{}) *Request {
	xmlValue, err := xml.Marshal(object)

	if err != nil {
		panic(err)
	}

	r.httpRequest.Header.Set(HeaderContentType, ContentTypeApplicationXml)
	r.httpRequest.Body = ioutil.NopCloser(bytes.NewReader(xmlValue))
	r.httpRequest.ContentLength = int64(len(xmlValue))

	return r
}

func (r *Request) WithCookie(cookie *http.Cookie) *Request {
	r.httpRequest.AddCookie(cookie)
	return r
}

func (r *Request) Expect() *Response {
	return r.client.makeRequest(r)
}
