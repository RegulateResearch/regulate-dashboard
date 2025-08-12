package worker

import (
	"fmt"
	"protogen/writer"
	"strings"
)

type serviceWriter interface {
	Service() Service
	writeLine(line string) error
	writeErr(message string, err error) error
	close() bool
}

type serviceWriterImpl struct {
	class      Class
	service    Service
	lineWriter writer.LineWriter
	errLogger  writer.ErrLogger
	lineNum    int
}

func newServiceWriter(class Class, service Service) (serviceWriter, error) {
	filename := fmt.Sprintf("%s-%s", class.fileName, service.dirname)

	errLogger, err := writer.NewErrLogger(fmt.Sprintf("%s-generator", filename))
	if err != nil {
		return nil, fmt.Errorf("fail to create error logger %s: (%v)", filename, err)
	}

	lineWriter, err := writer.NewWriter(fmt.Sprintf("generated/%s.proto", filename))
	if err != nil {
		return nil, fmt.Errorf("fail to create protofile %s: (%v)", filename, err)
	}

	res := &serviceWriterImpl{
		class:      class,
		service:    service,
		lineWriter: lineWriter,
		errLogger:  errLogger,
		lineNum:    0,
	}

	return res, nil
}

func (w *serviceWriterImpl) Service() Service {
	return w.service
}

func (w *serviceWriterImpl) writeLine(line string) error {
	w.lineNum++
	placeholder := "<SERVICE_NAME>"
	line = strings.Replace(line, placeholder, w.service.serviceName, 1)

	writeErr := w.lineWriter.Write(line)
	if writeErr != nil {
		logMessage := fmt.Sprintf("fail to write at line %d", w.lineNum)
		logErr := w.errLogger.WriteError(logMessage, writeErr)
		if logErr != nil {
			return fmt.Errorf("fail to write error at line %d, write error: (%v), log error (%v)", w.lineNum, writeErr, logErr)
		}

		return writeErr
	}

	return nil
}

func (w *serviceWriterImpl) writeErr(message string, err error) error {
	writeErr := w.errLogger.WriteError(message, err)
	return writeErr
}

func (w *serviceWriterImpl) close() bool {
	return w.lineWriter.Close() && w.errLogger.Close()
}
