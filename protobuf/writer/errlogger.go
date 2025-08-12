package writer

import (
	"fmt"
)

type ErrLogger interface {
	WriteError(message string, err error) error
	Close() bool
}

type errLoggerImpl struct {
	writer LineWriter
}

func NewErrLogger(filename string) (ErrLogger, error) {
	writer, err := NewWriter(fmt.Sprintf("errlog/%s.log", filename))
	if err != nil {
		return nil, err
	}

	logger := errLoggerImpl{
		writer: writer,
	}

	return logger, nil
}

func (l errLoggerImpl) WriteError(message string, err error) error {
	writeErr := l.writer.Write(fmt.Sprintf("message: (%s) | error: (%s)\n", message, err))
	return writeErr
}

func (l errLoggerImpl) Close() bool {
	return l.writer.Close()
}
