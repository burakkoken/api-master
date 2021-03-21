package header

import (
	"github.com/burakkoken/api-master/expect"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

type Expect func(t *testing.T, chain *expect.Chain, response *http.Response)

func Get(key string) Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {
		chain.Next(response.Header.Get(key))
	}
}

func Empty() Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {
		assert.Empty(t, chain.GetValue())
	}
}

func NotEmpty() Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {
		assert.NotEmpty(t, chain.GetValue())
	}
}

func Equal(value interface{}) Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {
		assert.Equal(t, value, chain.GetValue())
	}
}

func NotEqual(value interface{}) Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {
		assert.NotEqual(t, value, chain.GetValue())
	}
}
