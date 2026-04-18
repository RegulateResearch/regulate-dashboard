package generator

import "io"

type flushStringWriter interface {
	io.StringWriter
	Flush() error
}
