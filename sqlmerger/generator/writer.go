package generator

import "fmt"

type writer interface {
	flushStringWriter
	writeLine(str string) error
}

type writerImpl struct {
	flushStringWriter
}

func newWriter(base flushStringWriter) writer {
	return writerImpl{
		flushStringWriter: base,
	}
}

func (w writerImpl) writeLine(str string) error {
	line := fmt.Sprintf("%s\n", str)
	_, err := w.WriteString(line)
	if err != nil {
		return fmt.Errorf("line write error: %w", err)
	}

	err = w.Flush()
	if err != nil {
		return fmt.Errorf("flush error: %w", err)
	}

	return nil
}
