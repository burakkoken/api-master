package expect

type Chain struct {
	value interface{}
	ctx   *Context
}

func NewChain() *Chain {
	return &Chain{
		ctx: newContext(),
	}
}

func (chain *Chain) Next(value interface{}) *Chain {
	chain.value = value
	return chain
}

func (chain *Chain) GetValue() interface{} {
	return chain.value
}

func (chain *Chain) GetContext() *Context {
	return chain.ctx
}
