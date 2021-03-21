package headers

import (
	"github.com/burakkoken/api-master/context"
	"github.com/burakkoken/api-master/expect"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

type Expect func(t *testing.T, ctx *context.Context, chain *expect.Chain, response *http.Response)

func Contains(key string) Expect {
	return func(t *testing.T, ctx *context.Context, chain *expect.Chain, response *http.Response) {
		assert.Contains(t, response.Header, key)
	}
}

func Len(length int) Expect {
	return func(t *testing.T, ctx *context.Context, chain *expect.Chain, response *http.Response) {
		assert.Len(t, response.Header, length)
	}
}

func PutContext(contextKey string) Expect {
	return func(t *testing.T, ctx *context.Context, chain *expect.Chain, response *http.Response) {
		if ctx != nil {
			ctx.Put(contextKey, response.Header)
		}
	}
}
