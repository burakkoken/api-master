package apimaster

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"time"
)

type Response struct {
	client       *Client
	httpResponse *http.Response
	elapsedTime  time.Duration
	validator    *validator.Validate
}

func newResponse(client *Client, httpResponse *http.Response, elapsedTime time.Duration) *Response {
	return &Response{
		client:       client,
		httpResponse: httpResponse,
		elapsedTime:  elapsedTime,
		validator:    validator.New(),
	}
}

func (r *Response) Status(expects ...Expect) *Response {
	expectChain := newExpectChain("Status",
		r.httpResponse.StatusCode,
		func(fun ExpectFunction, chain *ExpectChain, value interface{}) {
			switch fun {
			case ExpectFunctionEqual:
			case ExpectFunctionNotEqual:
			case ExpectFunctionGreater:
			case ExpectFunctionGreaterOrEqual:
			case ExpectFunctionLess:
			case ExpectFunctionLessOrEqual:
				return
			default:
				panic(fmt.Sprintf("%s is not supported by %s", fun, "Status"))
			}
		},
	)

	r.check(expectChain, expects...)
	return r
}

func (r *Response) ElapsedTime(expects ...Expect) *Response {
	expectChain := newExpectChain("ElapsedTime",
		r.elapsedTime,
		func(fun ExpectFunction, chain *ExpectChain, value interface{}) {
			switch fun {
			case ExpectFunctionEqual:
			case ExpectFunctionNotEqual:
			case ExpectFunctionGreater:
			case ExpectFunctionGreaterOrEqual:
			case ExpectFunctionLess:
			case ExpectFunctionLessOrEqual:
				return
			default:
				panic(fmt.Sprintf("%s is not supported by %s", fun, "ElapsedTime"))
			}
		},
	)

	r.check(expectChain, expects...)
	return r
}

func (r *Response) Header(expects ...Expect) *Response {
	expectChain := newExpectChain("Header",
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
			case ExpectFunctionNotEqual:
			case ExpectFunctionEmpty:
			case ExpectFunctionNotEmpty:
				return
			default:
				panic(fmt.Sprintf("%s is not supported by %s", fun, "Header"))
			}
		},
	)

	r.check(expectChain, expects...)
	return r
}

func (r *Response) Headers(expects ...Expect) *Response {
	expectChain := newExpectChain("Headers",
		r.httpResponse.Header,
		func(fun ExpectFunction, chain *ExpectChain, value interface{}) {
			switch fun {
			case ExpectFunctionEqual:
			case ExpectFunctionNotEqual:
			case ExpectFunctionNil:
			case ExpectFunctionNotNil:
			case ExpectFunctionLen:
			case ExpectFunctionContains:
				return
			default:
				panic(fmt.Sprintf("%s is not supported by %s", fun, "Headers"))
			}
		},
	)

	r.check(expectChain, expects...)
	return r
}

func (r *Response) ContentType(expects ...Expect) *Response {
	expectChain := newExpectChain("ContentType",
		r.httpResponse.Header.Get(HeaderContentType),
		func(fun ExpectFunction, chain *ExpectChain, value interface{}) {
			switch fun {
			case ExpectFunctionEqual:
			case ExpectFunctionNotEqual:
			case ExpectFunctionEmpty:
			case ExpectFunctionNotEmpty:
				return
			default:
				panic(fmt.Sprintf("%s is not supported by %s", fun, "ContentType"))
			}
		},
	)

	r.check(expectChain, expects...)
	return r
}

func (r *Response) ContentLength(expects ...Expect) *Response {
	expectChain := newExpectChain("ContentLength",
		r.httpResponse.ContentLength,
		func(fun ExpectFunction, chain *ExpectChain, value interface{}) {
			switch fun {
			case ExpectFunctionEqual:
			case ExpectFunctionNotEqual:
			case ExpectFunctionGreater:
			case ExpectFunctionGreaterOrEqual:
			case ExpectFunctionLess:
			case ExpectFunctionLessOrEqual:
				return
			default:
				panic(fmt.Sprintf("%s is not supported by %s", fun, "ContentLength"))
			}
		},
	)

	r.check(expectChain, expects...)
	return r
}

func (r *Response) Body(expects ...Expect) *Response {

	defer r.httpResponse.Body.Close()
	bytes, err := ioutil.ReadAll(r.httpResponse.Body)
	assert.NoError(r.client.testing, err)

	expectChain := newExpectChain("Body",
		bytes,
		func(fun ExpectFunction, chain *ExpectChain, value interface{}) {
			switch fun {
			case ExpectFunctionNil:
			case ExpectFunctionNotNil:
			case ExpectFunctionLen:
				return
			case ExpectFunctionBind:
				r.bindBody(chain, value)
				chain.value = value
				return
			case ExpectFunctionIsValid:
				r.validate(chain)
				return
			default:
				panic(fmt.Sprintf("%s is not supported by %s", fun, "Body"))
			}
		},
	)

	r.check(expectChain, expects...)
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

func (r *Response) bindBody(chain *ExpectChain, value interface{}) {
	switch chain.value.(type) {
	case []byte:
		contentType := r.httpResponse.Header.Get(HeaderContentType)

		if contentType == ContentTypeApplicationJson {
			err := json.Unmarshal(chain.value.([]byte), value)
			if err != nil {
				panic(err)
			}
		} else if contentType == ContentTypeApplicationXml {
			err := xml.Unmarshal(chain.value.([]byte), value)
			if err != nil {
				panic(err)
			}
		}
	default:
	}
}

func (r *Response) validate(chain *ExpectChain) {
	err := r.validator.Struct(chain.value)

	if err != nil {

		var errMessage string
		if _, ok := err.(*validator.InvalidValidationError); ok {
			errMessage = err.Error()

		} else {
			for index, err := range err.(validator.ValidationErrors) {

				errMessage += errMessage +
					fmt.Sprintf("%d. '%s' validation error on %s (%s), \nTag Value: %v, \nValue : %v\n",
						index+1,
						err.Tag(),
						err.StructNamespace(),
						err.Type(),
						err.Param(),
						err.Value(),
					)

			}
		}

		assert.NoError(r.client.testing, err, errMessage)
	}
}
