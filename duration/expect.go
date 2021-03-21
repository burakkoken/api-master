package duration

import (
	"github.com/burakkoken/api-master/expect"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type Expect func(t *testing.T, chain *expect.Chain, actual time.Duration)

func Equal(value time.Duration) Expect {
	return func(t *testing.T, chain *expect.Chain, actual time.Duration) {
		assert.Equal(t, value, actual)
	}
}

func NotEqual(value time.Duration) Expect {
	return func(t *testing.T, chain *expect.Chain, actual time.Duration) {
		assert.NotEqual(t, value, actual)
	}
}

func Greater(value time.Duration) Expect {
	return func(t *testing.T, chain *expect.Chain, actual time.Duration) {
		assert.Greater(t, actual, value)
	}
}

func GreaterOrEqual(value time.Duration) Expect {
	return func(t *testing.T, chain *expect.Chain, actual time.Duration) {
		assert.GreaterOrEqual(t, actual, value)
	}
}

func Less(value time.Duration) Expect {
	return func(t *testing.T, chain *expect.Chain, actual time.Duration) {
		assert.Less(t, actual, value)
	}
}

func LessOrEqual(value time.Duration) Expect {
	return func(t *testing.T, chain *expect.Chain, actual time.Duration) {
		assert.LessOrEqual(t, actual, value)
	}
}
