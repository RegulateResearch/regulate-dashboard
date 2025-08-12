package main

import (
	"fmt"
	"protogen/lambda"
	"protogen/values"
	"protogen/worker"
)

func main() {
	value, err := values.ParseValue()
	if err != nil {
		fmt.Println(err.Error())
	}

	classes := value.WorkerClasses()
	lambda.ExecList(classes, func(c worker.Class) {
		err := c.Process()
		if err != nil {
			fmt.Println(err.Error())
		}
	})
}
