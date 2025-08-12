package writer

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type LineWriter interface {
	Write(str string) error
	Close() bool
}

type lineWriterImpl struct {
	file   io.WriteCloser
	writer *bufio.Writer
}

func NewWriter(pathstr string) (LineWriter, error) {
	file, err := os.Create(pathstr)
	if err != nil {
		return nil, fmt.Errorf("fail to create file %s, error: (%v)", pathstr, err)
	}

	res := lineWriterImpl{
		file:   file,
		writer: bufio.NewWriter(file),
	}

	return res, nil
}

func (w lineWriterImpl) Write(str string) error {
	line := fmt.Sprintf("%s\n", str)

	_, err := w.writer.WriteString(line)
	if err != nil {
		return fmt.Errorf("fail to write line, text: %s, error: (%v)", str, err)
	}

	err = w.writer.Flush()
	if err != nil {
		return fmt.Errorf("fail to flush: (%v)", err)
	}

	return nil
}

func (w lineWriterImpl) Close() bool {
	err := w.file.Close()
	return err == nil
}
