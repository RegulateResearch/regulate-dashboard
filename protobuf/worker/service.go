package worker

import (
	"bytes"
	"fmt"
	"os/exec"
)

type Service struct {
	dirname     string
	serviceName string
}

func NewService(dirname string, serviceName string) Service {
	return Service{
		dirname:     dirname,
		serviceName: serviceName,
	}
}

func (s Service) DirName() string {
	return s.dirname
}

func (s Service) ServiceName() string {
	return s.serviceName
}

func (s Service) execute(c Class, w serviceWriter) error {
	outPath := fmt.Sprintf("../%s/pbuf", s.dirname)

	cmd := exec.Command("protoc", []string{
		"--proto_path=generated",
		fmt.Sprintf("--go_out=%s", outPath),
		"--go_opt=paths=source_relative",
		fmt.Sprintf("--go-grpc_out=%s", outPath),
		"--go-grpc_opt=paths=source_relative",
		fmt.Sprintf("generated/%s-%s.proto", c.fileName, s.dirname),
	}...)

	var cmdStderr bytes.Buffer
	cmd.Stderr = &cmdStderr

	err := cmd.Run()
	if err != nil {
		logErr := w.writeErr(fmt.Sprintf("exec error: (%v)", cmdStderr.String()), err)
		if logErr != nil {
			return fmt.Errorf("fail executing %s-%s-%s and error logging failed, exec err = (%v) logging err = (%v)", c.fileName, s.dirname, s.serviceName, err, logErr)
		}

		return fmt.Errorf("fail executing: %s-%s-%s, err = (%v)", c.fileName, s.dirname, s.serviceName, err)
	}

	return nil
}
