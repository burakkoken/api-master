package status

import (
	"github.com/burakkoken/api-master/expect"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

type Expect func(t *testing.T, chain *expect.Chain, response *http.Response)

func Equal(value int) Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {
		assert.Equal(t, value, response.StatusCode)
	}
}

func NotEqual(value int) Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {
		assert.NotEqual(t, value, response.StatusCode)
	}
}

func Greater(value int) Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {
		assert.Greater(t, response.StatusCode, value)
	}
}

func GreaterOrEqual(value int) Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {
		assert.GreaterOrEqual(t, response.StatusCode, value)
	}
}

func Less(value int) Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {
		assert.Less(t, response.StatusCode, value)
	}
}

func LessOrEqual(value int) Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {
		assert.LessOrEqual(t, response.StatusCode, value)
	}
}
