package stringq

import (
	"github.com/burakkoken/api-master/context"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

type StringQuery struct {
	data string
	t    *testing.T
	ctx  *context.Context
}

func NewStringQuery(t *testing.T, ctx *context.Context, data []byte) *StringQuery {
	return &StringQuery{
		string(data),
		t,
		ctx,
	}
}

func (query *StringQuery) Empty() *StringQuery {
	assert.Empty(query.t, query.data)
	return query
}

func (query *StringQuery) NotEmpty() *StringQuery {
	assert.NotEmpty(query.t, query.data)
	return query
}

func (query *StringQuery) Equal(value string) *StringQuery {
	assert.Equal(query.t, value, query.data)
	return query
}

func (query *StringQuery) NotEqual(value interface{}) *StringQuery {
	assert.NotEqual(query.t, value, query.data)
	return query
}

func (query *StringQuery) Contains(value string) *StringQuery {
	assert.Contains(query.t, query.data, value)
	return query
}

func (query *StringQuery) Len(length int) *StringQuery {
	assert.Len(query.t, query.data, length)
	return query
}

func (query *StringQuery) PutContext(contextKey string) *StringQuery {
	if query.ctx != nil {
		query.ctx.Put(contextKey, query.data)
	}

	return query
}

func (query *StringQuery) String() string {
	return query.data
}

func (query *StringQuery) Match(expression string) *Match {
	regex := regexp.MustCompile(expression)
	matches := regex.FindStringSubmatch(query.data)

	if matches == nil {
		return newMatch(query.t, query.ctx, make([]string, 0), nil)
	}

	return newMatch(query.t, query.ctx, matches, regex.SubexpNames())
}

func (query *StringQuery) MatchAll(expression string) []Match {
	regex := regexp.MustCompile(expression)

	matches := regex.FindAllStringSubmatch(query.data, -1)
	if matches == nil {
		return []Match{}
	}

	result := make([]Match, 0)
	for _, match := range matches {
		result = append(result, *newMatch(query.t, query.ctx, match, regex.SubexpNames()))
	}

	return result
}
