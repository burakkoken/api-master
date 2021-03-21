package apimaster

import (
	"github.com/stretchr/testify/assert"
)

type Expect func(chain *ExpectChain, response *Response)

type ExpectFunction string

const (
	ExpectFunctionGet            ExpectFunction = "Get"
	ExpectFunctionEqual          ExpectFunction = "Equal"
	ExpectFunctionNotEqual       ExpectFunction = "NotEqual"
	ExpectFunctionGreater        ExpectFunction = "Greater"
	ExpectFunctionGreaterOrEqual ExpectFunction = "GreaterOrEqual"
	ExpectFunctionLess           ExpectFunction = "Less"
	ExpectFunctionLessOrEqual    ExpectFunction = "LessOrEqual"
	ExpectFunctionEmpty          ExpectFunction = "Empty"
	ExpectFunctionNotEmpty       ExpectFunction = "NotEmpty"
	ExpectFunctionNil            ExpectFunction = "Nil"
	ExpectFunctionNotNil         ExpectFunction = "NotNil"
	ExpectFunctionLen            ExpectFunction = "Len"
	ExpectFunctionContains       ExpectFunction = "Contains"
	ExpectFunctionBind           ExpectFunction = "Bind"
	ExpectFunctionIsValid        ExpectFunction = "IsValid"
)

type ExpectChain struct {
	name     string
	value    interface{}
	callback func(fun ExpectFunction, chain *ExpectChain, value interface{})
}

func newExpectChain(name string, value interface{}, callback func(fun ExpectFunction, chain *ExpectChain, value interface{})) *ExpectChain {
	return &ExpectChain{
		name,
		value,
		callback,
	}
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

		assert.Equal(response.client.testing, value, chain.value)
	}
}

func NotEqual(value interface{}) Expect {
	return func(chain *ExpectChain, response *Response) {
		if chain.callback != nil {
			chain.callback(ExpectFunctionNotEqual, chain, value)
		}

		assert.NotEqual(response.client.testing, value, chain.value)
	}
}

func Greater(value interface{}) Expect {
	return func(chain *ExpectChain, response *Response) {
		if chain.callback != nil {
			chain.callback(ExpectFunctionGreater, chain, value)
		}

		assert.Greater(response.client.testing, chain.value, value)
	}
}

func GreaterOrEqual(value interface{}) Expect {
	return func(chain *ExpectChain, response *Response) {
		if chain.callback != nil {
			chain.callback(ExpectFunctionGreaterOrEqual, chain, value)
		}

		assert.GreaterOrEqual(response.client.testing, chain.value, value)
	}
}

func Less(value interface{}) Expect {
	return func(chain *ExpectChain, response *Response) {
		if chain.callback != nil {
			chain.callback(ExpectFunctionLess, chain, value)
		}

		assert.Less(response.client.testing, chain.value, value)
	}
}

func LessOrEqual(value interface{}) Expect {
	return func(chain *ExpectChain, response *Response) {
		if chain.callback != nil {
			chain.callback(ExpectFunctionLessOrEqual, chain, value)
		}

		assert.LessOrEqual(response.client.testing, chain.value, value)
	}
}

func Empty() Expect {
	return func(chain *ExpectChain, response *Response) {
		if chain.callback != nil {
			chain.callback(ExpectFunctionEmpty, chain, nil)
		}

		assert.Empty(response.client.testing, chain.value)
	}
}

func NotEmpty() Expect {
	return func(chain *ExpectChain, response *Response) {
		if chain.callback != nil {
			chain.callback(ExpectFunctionNotEmpty, chain, nil)
		}

		assert.NotEmpty(response.client.testing, chain.value)
	}
}

func Nil() Expect {
	return func(chain *ExpectChain, response *Response) {
		if chain.callback != nil {
			chain.callback(ExpectFunctionNil, chain, nil)
		}

		assert.Nil(response.client.testing, chain.value)
	}
}

func NotNil() Expect {
	return func(chain *ExpectChain, response *Response) {
		if chain.callback != nil {
			chain.callback(ExpectFunctionNotNil, chain, nil)
		}

		assert.NotNil(response.client.testing, chain.value)
	}
}

func Len(length int) Expect {
	return func(chain *ExpectChain, response *Response) {
		if chain.callback != nil {
			chain.callback(ExpectFunctionLen, chain, nil)
		}

		assert.Len(response.client.testing, chain.value, length)
	}
}

func Contains(value interface{}) Expect {
	return func(chain *ExpectChain, response *Response) {
		if chain.callback != nil {
			chain.callback(ExpectFunctionContains, chain, nil)
		}

		assert.Contains(response.client.testing, chain.value, value)
	}
}

func Bind(instance interface{}) Expect {
	return func(chain *ExpectChain, response *Response) {
		if chain.callback != nil {
			chain.callback(ExpectFunctionBind, chain, instance)
		}
	}
}

func IsValid() Expect {
	return func(chain *ExpectChain, response *Response) {
		if chain.callback != nil {
			chain.callback(ExpectFunctionIsValid, chain, nil)
		}
	}
}
