package background

import "frascati/exception"

type Report interface {
	Name() string
	Err() exception.Exception

	// Ideally the type is determined through generic, to allow task chaining using previous result.
	// However since currently Go doesn't support type generic in method/struct function,
	// we will settle with type "any" for now, without any support for task chaining
	Result() any
	ToMap() map[string]any
}

type report struct {
	name   string
	err    exception.Exception
	result any
}

func (r report) Name() string {
	return r.name
}

func (r report) Err() exception.Exception {
	return r.err
}

func (r report) Result() any {
	return r.result
}

func (r report) ToMap() map[string]any {
	return map[string]any{
		"task name": r.name,
		"err":       r.err,
		"result":    r.result,
	}
}
