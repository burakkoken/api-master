package expect

import "github.com/burakkoken/api-master/context"

type Chain struct {
	value        interface{}
	chainContext *context.Context
}

func NewChain() *Chain {
	return &Chain{
		chainContext: context.NewContext(),
	}
}

func (chain *Chain) Next(value interface{}) *Chain {
	chain.value = value
	return chain
}

func (chain *Chain) GetValue() interface{} {
	return chain.value
}

func (chain *Chain) GetChainContext() *context.Context {
	return chain.chainContext
}
