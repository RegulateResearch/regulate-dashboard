package worker

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"protogen/generictype"
	"protogen/lambda"
	"protogen/writer"
)

type Class struct {
	fileName  string
	services  []Service
	writers   []serviceWriter
	errLogger writer.ErrLogger
	reader    io.Reader
}

func NewClass(fileName string, services []Service) Class {
	return Class{
		fileName: fileName,
		services: services,
	}
}

func (c Class) FileName() string {
	return c.fileName
}

func (c Class) Process() error {
	errLogger, err := writer.NewErrLogger(c.fileName)
	if err != nil {
		return fmt.Errorf("cannot create error log file for %s, error : (%v)", c.fileName, err)
	}

	c.errLogger = errLogger
	defer errLogger.Close()

	serviceWriters, err := c.prepServiceWriters()
	if err != nil {
		errLogger.WriteError("writer generation failure", err)
		return fmt.Errorf("cannot create writers and loggers for class %s, error: (%v)", c.fileName, err)
	}

	defer lambda.ExecList(serviceWriters, func(w serviceWriter) {
		w.close()
	})

	c.writers = serviceWriters

	generateErr := c.generate()
	if generateErr != nil {
		return generateErr
	}

	executeErr := c.execute()
	if executeErr != nil {
		return executeErr
	}

	return nil
}

func (c Class) execute() error {
	serviceExecErrPairs := lambda.MapList(c.writers, func(w serviceWriter) generictype.Pair[Service, error] {
		s := w.Service()
		err := s.execute(c, w)
		return generictype.NewPair(s, err)
	})

	serviceExecErrPairs = lambda.FilterList(serviceExecErrPairs, func(p generictype.Pair[Service, error]) bool {
		err := p.B()
		return err != nil
	})

	if len(serviceExecErrPairs) > 0 {
		lambda.ExecList(serviceExecErrPairs, func(p generictype.Pair[Service, error]) {
			service := p.A()
			err := p.B()
			c.errLogger.WriteError(fmt.Sprintf("protogen failed, service = %s", service.serviceName), err)
		})

		return fmt.Errorf("executing protogen failed, see every log for detail")
	}

	return nil
}

func (c Class) generate() error {
	protoFileName := fmt.Sprintf("template/%s.proto", c.fileName)
	protoFile, err := os.Open(protoFileName)
	if err != nil {
		finalErr := fmt.Errorf("cannot open file %s: %v", protoFileName, err)
		return finalErr
	}
	defer protoFile.Close()

	c.reader = protoFile

	writeErr := c.readWriteProto()
	return writeErr
}

func (c Class) readWriteProto() error {
	scanner := bufio.NewScanner(c.reader)
	for scanner.Scan() {
		line := scanner.Text()

		writerErrPairs := lambda.MapList(c.writers, func(w serviceWriter) generictype.Pair[serviceWriter, error] {
			err := w.writeLine(line)
			return generictype.NewPair(w, err)
		})

		pairsWithErr := lambda.FilterList(writerErrPairs, func(p generictype.Pair[serviceWriter, error]) bool {
			err := p.B()
			return err != nil
		})

		if len(pairsWithErr) > 0 {
			lambda.ExecList(pairsWithErr, func(p generictype.Pair[serviceWriter, error]) {
				service := p.A().Service()
				writeErr := p.B()

				errMessage := fmt.Sprintf("writing fails for service %s", service.DirName())
				c.errLogger.WriteError(errMessage, writeErr)
			})

			return errors.New("proto files generation error, see every log for details")
		}

	}

	return nil
}

func (c Class) prepServiceWriters() ([]serviceWriter, error) {
	writerErrPairs := lambda.MapList(c.services, func(s Service) generictype.Pair[serviceWriter, error] {
		serviceWriter, err := newServiceWriter(c, s)
		return generictype.NewPair(serviceWriter, err)
	})

	pairsWithErr := lambda.FilterList(writerErrPairs, func(p generictype.Pair[serviceWriter, error]) bool {
		err := p.B()
		return err != nil
	})

	if len(pairsWithErr) > 0 {
		lambda.ExecList(pairsWithErr, func(p generictype.Pair[serviceWriter, error]) {
			err := p.B()
			c.errLogger.WriteError("file generation failure", err)
		})

		return nil, errors.New("fail to create service writers, please look at log")
	}

	writers := lambda.MapList(writerErrPairs, func(p generictype.Pair[serviceWriter, error]) serviceWriter {
		writer := p.A()
		return writer
	})

	return writers, nil
}
