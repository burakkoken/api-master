package apimaster

import (
	"github.com/stretchr/testify/assert"
	"reflect"
)

type Expect func(chain *ExpectChain, response *Response)

type ExpectFunction string

const (
	ExpectFunctionGet            ExpectFunction = "Get"
	ExpectFunctionEqual          ExpectFunction = "Equal"
	ExpectFunctionGreater        ExpectFunction = "Greater"
	ExpectFunctionGreaterOrEqual ExpectFunction = "GreaterOrEqual"
	ExpectFunctionLess           ExpectFunction = "Less"
	ExpectFunctionLessOrEqual    ExpectFunction = "LessOrEqual"
	ExpectFunctionEmpty          ExpectFunction = "Empty"
	ExpectFunctionNotEmpty       ExpectFunction = "NotEmpty"
	ExpectFunctionNil            ExpectFunction = "Nil"
	ExpectFunctionNotNil         ExpectFunction = "NotNil"
	ExpectFunctionLen            ExpectFunction = "Len"
)

type ExpectChain struct {
	value    interface{}
	callback func(fun ExpectFunction, chain *ExpectChain, value interface{})
}

func Get(value ...interface{}) Expect {
	return func(chain *ExpectChain, response *Response) {
		if chain.callback != nil && len(value) != 0 {
			chain.callback(ExpectFunctionGet, chain, value[0])
		}
	}
}

func Equal(value interface{}) Expect {
	return func(chain *ExpectChain, response *Response) {
		if chain.callback != nil {
			chain.callback(ExpectFunctionEqual, chain, value)
		}

		if response.client.suite != nil {
			response.client.suite.Equal(value, chain.value)
		} else {
			assert.Equal(response.client.testing, chain.value, value)
		}
	}
}

func Greater(value interface{}) Expect {
	return func(chain *ExpectChain, response *Response) {
		if chain.callback != nil {
			chain.callback(ExpectFunctionGreater, chain, value)
		}

		if response.client.suite != nil {
			response.client.suite.Greater(value, chain.value)
		} else {
			assert.Greater(response.client.testing, chain.value, value)
		}
	}
}

func GreaterOrEqual(value interface{}) Expect {
	return func(chain *ExpectChain, response *Response) {
		if chain.callback != nil {
			chain.callback(ExpectFunctionGreaterOrEqual, chain, value)
		}

		if response.client.suite != nil {
			response.client.suite.GreaterOrEqual(value, chain.value)
		} else {
			assert.GreaterOrEqual(response.client.testing, chain.value, value)
		}
	}
}

func Less(value interface{}) Expect {
	return func(chain *ExpectChain, response *Response) {
		if chain.callback != nil {
			chain.callback(ExpectFunctionLess, chain, value)
		}

		if response.client.suite != nil {
			response.client.suite.Less(value, chain.value)
		} else {
			assert.Less(response.client.testing, chain.value, value)
		}
	}
}

func LessOrEqual(value interface{}) Expect {
	return func(chain *ExpectChain, response *Response) {
		if chain.callback != nil {
			chain.callback(ExpectFunctionLessOrEqual, chain, value)
		}

		if response.client.suite != nil {
			response.client.suite.LessOrEqual(value, chain.value)
		} else {
			assert.LessOrEqual(response.client.testing, chain.value, value)
		}
	}
}

func Empty() Expect {
	return func(chain *ExpectChain, response *Response) {
		if chain.callback != nil {
			chain.callback(ExpectFunctionEmpty, chain, nil)
		}

		if response.client.suite != nil {
			response.client.suite.Empty(chain.value)
		} else {
			assert.Empty(response.client.testing, chain.value)
		}
	}
}

func NotEmpty() Expect {
	return func(chain *ExpectChain, response *Response) {
		if chain.callback != nil {
			chain.callback(ExpectFunctionNotEmpty, chain, nil)
		}

		if response.client.suite != nil {
			response.client.suite.NotEmpty(chain.value)
		} else {
			assert.NotEmpty(response.client.testing, chain.value)
		}
	}
}

func Nil() Expect {
	return func(chain *ExpectChain, response *Response) {
		if chain.callback != nil {
			chain.callback(ExpectFunctionNil, chain, nil)
		}

		if response.client.suite != nil {
			response.client.suite.Nil(chain.value)
		} else {
			assert.Nil(response.client.testing, chain.value)
		}
	}
}

func NotNil() Expect {
	return func(chain *ExpectChain, response *Response) {
		if chain.callback != nil {
			chain.callback(ExpectFunctionNotNil, chain, nil)
		}

		if response.client.suite != nil {
			response.client.suite.NotNil(chain.value)
		} else {
			assert.NotNil(response.client.testing, chain.value)
		}
	}
}

func Len() Expect {
	return func(chain *ExpectChain, response *Response) {
		if chain.callback != nil {
			chain.callback(ExpectFunctionLen, chain, nil)
		}

		v := reflect.ValueOf(chain.value)
		switch v.Kind() {
		case reflect.String:
			chain.value = len(v.String())
		case reflect.Array:
		case reflect.Slice:
		case reflect.Map:
			chain.value = v.Len()
		default:
			panic("unsupported type")
		}
	}
}
