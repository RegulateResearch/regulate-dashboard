package generator

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"protogen/values"
	"strings"
)

func Generate() {
	classDirsTuple := values.ClassDirsTuple()
	dirServiceTuple := values.DirServiceTuple()
	log.Println("Generating Protofile")
	log.Println(values.PrintDelim('*', 15))

	for class, dirs := range classDirsTuple {
		for i := 0; i < len(dirs); i++ {
			dir := dirs[i]
			service := dirServiceTuple[dir]
			log.Println("generating protofile: ", fmt.Sprintf("%s-%s", class, service))
			generateProto(class, service)
			log.Println(values.PrintDelim('-', 15))
		}
	}

	log.Println(values.PrintDelim('=', 15))
}

func generateProto(className string, serviceName string) {
	protoFileName := fmt.Sprintf("template/%s.proto", className)
	protoFile, err := os.Open(protoFileName)
	if err != nil {
		log.Fatalln("cannot open file coba.proto")
	}
	defer protoFile.Close()

	outFileName := fmt.Sprintf("generated/coba-%s.proto", serviceName)
	outFile, err := os.Create(outFileName)
	if err != nil {
		log.Fatalf("cannot create file: %s", outFileName)
	}
	defer outFile.Close()

	generateProtoFile(protoFile, outFile, serviceName)
}

func generateProtoFile(readIO io.Reader, writeIO io.Writer, serviceName string) {
	placeholder := "<PACKAGE_NAME>"
	scanner := bufio.NewScanner(readIO)
	writer := bufio.NewWriter(writeIO)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, placeholder, serviceName, 1)
		line = line + "\n"
		_, err := writer.WriteString(line)
		if err != nil {
			log.Println("error when writing line: ", err.Error())
		}

		err = writer.Flush()
		if err != nil {
			log.Println("error when flush: ", err.Error())
		}
	}
}
