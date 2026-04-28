package generator

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Merger struct {
	exs    []executor
	writer io.WriteCloser
}

func NewMerger(filenames ...string) Merger {
	exs := make([]executor, len(filenames))
	for i := 0; i < len(filenames); i++ {
		exs[i] = newExecutor(filenames[i])
	}

	resPath := fmt.Sprintf("%s/%s", fileRoot, "initdb.sql")
	writer, err := os.Create(resPath)
	if err != nil {
		log.Fatalf("cannot generate file %s: %v", resPath, err)
	}

	return Merger{
		exs:    exs,
		writer: writer,
	}
}

func (m Merger) Execute() {
	defer m.writer.Close()
	bufWriter := bufio.NewWriter(m.writer)
	lineWriter := newWriter(bufWriter)
	for i := 0; i < len(m.exs); i++ {
		m.exs[i].execute(lineWriter)
	}
}
