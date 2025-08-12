package values

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"protogen/lambda"
	"protogen/worker"
)

type service struct {
	DirName     string `json:"dir_name"`
	ServiceName string `json:"service_name"`
}

type class struct {
	Filename        string   `json:"file_name"`
	ServiceDirNames []string `json:"service_dir_names"`
}

type Value struct {
	Services []service `json:"services"`
	Classes  []class   `json:"classes"`
}

func ParseValue() (Value, error) {
	var valueObj Value

	jsonFile, err := os.Open("values.json")
	if err != nil {
		return Value{}, fmt.Errorf("cannot open json file, error: %v", err)
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return Value{}, fmt.Errorf("cannot read json file, error: %v", err)
	}

	err = json.Unmarshal(byteValue, &valueObj)
	if err != nil {
		return Value{}, fmt.Errorf("cannot parse json file, error: %v", err)
	}

	return valueObj, nil
}

func (v Value) WorkerClasses() []worker.Class {
	serviceMap := make(map[string]worker.Service)
	lambda.ExecList(v.Services, func(s service) {
		serviceMap[s.DirName] = worker.NewService(s.DirName, s.ServiceName)
	})

	workerClasses := lambda.MapList(v.Classes, func(c class) worker.Class {
		services := lambda.MapList(c.ServiceDirNames, func(dirname string) worker.Service {
			return serviceMap[dirname]
		})

		return worker.NewClass(c.Filename, services)
	})

	return workerClasses
}
