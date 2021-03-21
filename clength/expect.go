package clength

import (
	"github.com/burakkoken/api-master/expect"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

type Expect func(t *testing.T, chain *expect.Chain, response *http.Response)

func Equal(value time.Duration) Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {
		assert.Equal(t, value, response.ContentLength)
	}
}

func NotEqual(value time.Duration) Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {
		assert.NotEqual(t, value, response.ContentLength)
	}
}

func Greater(value time.Duration) Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {
		assert.Greater(t, response.ContentLength, value)
	}
}

func GreaterOrEqual(value time.Duration) Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {
		assert.GreaterOrEqual(t, response.ContentLength, value)
	}
}

func Less(value time.Duration) Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {
		assert.Less(t, response.ContentLength, value)
	}
}

func LessOrEqual(value time.Duration) Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {
		assert.LessOrEqual(t, response.ContentLength, value)
	}
}
