package ctype

import (
	"github.com/burakkoken/api-master/expect"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const HeaderContentType = "Content-Type"

type Expect func(t *testing.T, chain *expect.Chain, response *http.Response)

func Empty() Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {
		assert.Empty(t, response.Header.Get(HeaderContentType))
	}
}

func NotEmpty() Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {
		assert.NotEmpty(t, response.Header.Get(HeaderContentType))
	}
}

func Equal(value interface{}) Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {
		assert.Equal(t, value, response.Header.Get(HeaderContentType))
	}
}

func NotEqual(value interface{}) Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {
		assert.NotEqual(t, value, response.Header.Get(HeaderContentType))
	}
}
