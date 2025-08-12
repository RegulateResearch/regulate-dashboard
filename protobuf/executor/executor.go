package executor

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"protogen/values"
)

func Execute() {
	classDirTuple := values.ClassDirsTuple()
	dirServiceTuple := values.DirServiceTuple()

	log.Println("Executing Protoc")
	log.Println(values.PrintDelim('*', 15))

	for class, dirs := range classDirTuple {
		for i := 0; i < len(dirs); i++ {
			dir := dirs[i]
			service := dirServiceTuple[dir]
			log.Println("exec protoc: ", fmt.Sprintf("%s-%s-%s", class, dir, service))
			ExecProtoc(class, dir, service)
			log.Println(values.PrintDelim('-', 15))
		}
	}

	log.Println(values.PrintDelim('=', 15))
}

func ExecProtoc(className string, dirName string, serviceName string) {
	outPath := fmt.Sprintf("../%s/pbuf", dirName)
	// args := []string{
	// 	"--proto_path=generated",
	// 	fmt.Sprintf("--go_out=%s", outPath),
	// 	"--go_opt=paths=source_relative",
	// 	fmt.Sprintf("--go-grpc_out=%s", outPath),
	// 	"--go-grpc_opt=paths=source_relative",
	// 	fmt.Sprintf("generated/%s-%s.proto", className, serviceName),
	// }

	// cmd := exec.Command("protoc", args...)
	cmd := exec.Command("protoc", []string{
		"--proto_path=generated",
		fmt.Sprintf("--go_out=%s", outPath),
		"--go_opt=paths=source_relative",
		fmt.Sprintf("--go-grpc_out=%s", outPath),
		"--go-grpc_opt=paths=source_relative",
		fmt.Sprintf("generated/%s-%s.proto", className, serviceName),
	}...)

	var cmdStderr bytes.Buffer
	cmd.Stderr = &cmdStderr

	err := cmd.Run()
	if err != nil {
		log.Printf("fail executing: %s-%s-%s, err = %s\n", className, dirName, serviceName, err.Error())
		log.Println("stderr: ", cmdStderr.String())
	}

}
