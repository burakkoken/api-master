package context

type Context struct {
	values map[string]interface{}
}

func NewContext() *Context {
	return &Context{
		make(map[string]interface{}),
	}
}

func (ctx *Context) Get(key string) interface{} {
	return ctx.values[key]
}

func (ctx *Context) Put(key string, value interface{}) *Context {
	ctx.values[key] = value
	return ctx
}

func (ctx *Context) Delete(key string) *Context {
	delete(ctx.values, key)
	return ctx
}

func (ctx *Context) Clear() *Context {
	ctx.values = make(map[string]interface{})
	return ctx
}
