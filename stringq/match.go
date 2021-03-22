package stringq

import (
	"github.com/burakkoken/api-master/context"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Match struct {
	matches []string
	names   map[string]int
	t       *testing.T
	ctx     *context.Context
}

func newMatch(t *testing.T, ctx *context.Context, matches []string, names []string) *Match {
	nameMap := map[string]int{}
	for n, name := range names {
		if name != "" {
			nameMap[name] = n
		}
	}

	return &Match{
		matches,
		nameMap,
		t,
		ctx,
	}
}

func (match *Match) Matches() []string {
	return match.matches
}

func (match *Match) Len(length int) *Match {
	assert.Len(match.t, len(match.matches), length)
	return match
}

func (match *Match) Empty() *Match {
	assert.Empty(match.t, match.matches)
	return match
}

func (match *Match) NotEmpty() *Match {
	assert.NotEmpty(match.t, match.matches)
	return match
}

func (match *Match) Name(name string) *StringQuery {
	index, ok := match.names[name]
	if !ok {
		return NewStringQuery(match.t, match.ctx, []byte(""))
	}
	return match.Index(index)
}

func (match *Match) Index(index int) *StringQuery {
	if index < 0 || index >= len(match.matches) {
		return NewStringQuery(match.t, match.ctx, []byte(""))
	}

	return NewStringQuery(match.t, match.ctx, []byte(match.matches[index]))
}
