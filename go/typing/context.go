package typing

import "context"

type Context interface {
	context.Context
	Set(key string, value any)
	Get(key string) (value any, ok bool)
}

type dictionaryContext struct {
	context.Context
	dict map[string]any
}

func NewDictionaryContext(ctx context.Context) Context {
	dict := make(map[string]any)
	return dictionaryContext{
		Context: ctx,
		dict:    dict,
	}
}

func (ctx dictionaryContext) Set(key string, value any) {
	ctx.dict[key] = value
}

func (ctx dictionaryContext) Get(key string) (value any, ok bool) {
	value, ok = ctx.dict[key]
	return value, ok
}
