package typing

import "context"

type Context interface {
	context.Context
	Set(key string, value any)
	Get(key string) (any, bool)
}
