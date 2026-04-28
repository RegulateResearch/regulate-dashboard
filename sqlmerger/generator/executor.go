package generator

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type executor struct {
	path   string
	reader io.ReadCloser
}

func newExecutor(filename string) executor {
	path := fmt.Sprintf("%s/%s", fileRoot, filename)
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("cannot open file %s: %v", filename, err)
	}

	return executor{
		path:   path,
		reader: file,
	}
}

func (e executor) execute(writer writer) {
	defer e.reader.Close()
	scanner := bufio.NewScanner(e.reader)
	lineNum := 1
	for scanner.Scan() {
		line := scanner.Text()
		line = fmt.Sprintf("%s\n", line)
		_, err := writer.WriteString(line)
		if err != nil {
			log.Fatalf("cannot write line of file %s at line %d: %v", e.path, lineNum, err)
		}

		lineNum++
	}

	err := writer.writeLine("")
	if err != nil {
		log.Fatalf("cannot write new line of file %s: %v", e.path, err)
	}
}
