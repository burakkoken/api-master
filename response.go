package apimaster

import (
	"fmt"
	"net/http"
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

func (r *Response) Status(expects ...Expect) *Response {
	r.check(&ExpectChain{
		r.httpResponse.StatusCode,
		func(fun ExpectFunction, chain *ExpectChain, value interface{}) {
			switch fun {
			case ExpectFunctionEqual:
			case ExpectFunctionGreater:
			case ExpectFunctionGreaterOrEqual:
			case ExpectFunctionLess:
			case ExpectFunctionLessOrEqual:
				return
			default:
				panic(fmt.Sprintf("%s is not supported by %s", fun, "ContentType"))
			}
		},
	}, expects...)
	return r
}

func (r *Response) ElapsedTime(expects ...Expect) *Response {
	r.check(&ExpectChain{
		r.elapsedTime,
		func(fun ExpectFunction, chain *ExpectChain, value interface{}) {
			switch fun {
			case ExpectFunctionEqual:
			case ExpectFunctionGreater:
			case ExpectFunctionGreaterOrEqual:
			case ExpectFunctionLess:
			case ExpectFunctionLessOrEqual:
				return
			default:
				panic(fmt.Sprintf("%s is not supported by %s", fun, "ContentType"))
			}
		},
	}, expects...)
	return r
}

func (r *Response) Header(expects ...Expect) *Response {
	r.check(&ExpectChain{
		nil,
		func(fun ExpectFunction, chain *ExpectChain, value interface{}) {
			switch fun {
			case ExpectFunctionGet:
				switch value.(type) {
				case string:
					chain.value = r.httpResponse.Header.Get(value.(string))
				default:
				}
				return
			case ExpectFunctionEqual:
			case ExpectFunctionEmpty:
			case ExpectFunctionNotEmpty:
				return
			default:
				panic(fmt.Sprintf("%s is not supported by %s", fun, "ContentType"))
			}
		},
	}, expects...)
	return r
}

func (r *Response) Headers(expects ...Expect) *Response {
	r.check(&ExpectChain{
		r.httpResponse.Header,
		func(fun ExpectFunction, chain *ExpectChain, value interface{}) {
			switch fun {
			case ExpectFunctionGet:

				switch value.(type) {
				case string:
					key := value.(string)
					if key == "" {
						chain.value = r.httpResponse.Header
					} else {
						chain.value = r.httpResponse.Header.Get(value.(string))
					}
				default:
				}

				return
			case ExpectFunctionEqual:
			case ExpectFunctionEmpty:
			case ExpectFunctionNotEmpty:
				return
			default:
				panic(fmt.Sprintf("%s is not supported by %s", fun, "ContentType"))
			}
		},
	}, expects...)
	return r
}

func (r *Response) ContentType(expects ...Expect) *Response {
	r.check(&ExpectChain{
		r.httpResponse.Header.Get(HeaderContentType),
		func(fun ExpectFunction, chain *ExpectChain, value interface{}) {
			switch fun {
			case ExpectFunctionEqual:
			case ExpectFunctionEmpty:
			case ExpectFunctionNotEmpty:
				return
			default:
				panic(fmt.Sprintf("%s is not supported by %s", fun, "ContentType"))
			}
		},
	}, expects...)
	return r
}

func (r *Response) ContentLength(expects ...Expect) *Response {
	r.check(&ExpectChain{
		r.httpResponse.ContentLength,
		func(fun ExpectFunction, chain *ExpectChain, value interface{}) {
			switch fun {
			case ExpectFunctionEqual:
			case ExpectFunctionGreater:
			case ExpectFunctionGreaterOrEqual:
			case ExpectFunctionLess:
			case ExpectFunctionLessOrEqual:
				return
			default:
				panic(fmt.Sprintf("%s is not supported by %s", fun, "ContentLength"))
			}
		},
	}, expects...)
	return r
}

func (r *Response) Body(expects ...Expect) *Response {
	// TODO
	return r
}

func (r *Response) Raw() *http.Response {
	return r.httpResponse
}

func (r *Response) check(chainData *ExpectChain, expects ...Expect) {
	for _, expect := range expects {
		expect(chainData, r)
	}
}
