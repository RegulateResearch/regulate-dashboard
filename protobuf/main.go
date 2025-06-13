package main

import (
	"protogen/executor"
	"protogen/generator"
)

func main() {
	generator.Generate()
	executor.Execute()
}
