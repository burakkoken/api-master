package apimaster

import (
	"github.com/burakkoken/api-master/body"
	"github.com/burakkoken/api-master/clength"
	"github.com/burakkoken/api-master/ctype"
	"github.com/burakkoken/api-master/duration"
	"github.com/burakkoken/api-master/expect"
	"github.com/burakkoken/api-master/header"
	"github.com/burakkoken/api-master/headers"
	"github.com/burakkoken/api-master/jsonq"
	"github.com/burakkoken/api-master/status"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

type Response struct {
	client       *Client
	httpResponse *http.Response
	elapsedTime  time.Duration
}

func newResponse(client *Client, httpResponse *http.Response, elapsedTime time.Duration) *Response {
	return &Response{
		client:       client,
		httpResponse: httpResponse,
		elapsedTime:  elapsedTime,
	}
}

func (r *Response) Status(expects ...status.Expect) *Response {
	for _, expectFunc := range expects {
		expectFunc(r.client.testing, nil, r.httpResponse)
	}

	return r
}

func (r *Response) ElapsedTime(expects ...duration.Expect) *Response {
	for _, expectFunc := range expects {
		expectFunc(r.client.testing, nil, r.elapsedTime)
	}

	return r
}

func (r *Response) Header(expects ...header.Expect) *Response {
	chain := expect.NewChain()

	for _, expectFunc := range expects {
		expectFunc(r.client.testing, chain, r.httpResponse)
	}

	return r
}

func (r *Response) Headers(expects ...headers.Expect) *Response {
	for _, expectFunc := range expects {
		expectFunc(r.client.testing, nil, r.httpResponse)
	}

	return r
}

func (r *Response) ContentType(expects ...ctype.Expect) *Response {
	for _, expectFunc := range expects {
		expectFunc(r.client.testing, nil, r.httpResponse)
	}

	return r
}

func (r *Response) ContentLength(expects ...clength.Expect) *Response {
	for _, expectFunc := range expects {
		expectFunc(r.client.testing, nil, r.httpResponse)
	}

	return r
}

func (r *Response) Body(expects ...body.Expect) *ResponseQuery {

	defer r.httpResponse.Body.Close()
	bytes, err := ioutil.ReadAll(r.httpResponse.Body)
	assert.NoError(r.client.testing, err)

	chain := expect.NewChain().Next(bytes)
	chain.GetContext().Put(body.ContextKeyBody, bytes)

	for _, expectFunc := range expects {
		expectFunc(r.client.testing, chain, r.httpResponse)
	}

	return newResponseQuery(r.client.testing, bytes)
}

func (r *Response) Raw() *http.Response {
	return r.httpResponse
}

type ResponseQuery struct {
	body []byte
	t    *testing.T
}

func newResponseQuery(t *testing.T, body []byte) *ResponseQuery {
	return &ResponseQuery{
		body,
		t,
	}
}

func (responseQuery *ResponseQuery) JsonQuery() *jsonq.JsonQuery {
	return jsonq.NewJsonQuery(responseQuery.t, responseQuery.body)
}
