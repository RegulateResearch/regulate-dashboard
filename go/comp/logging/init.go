package logging

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func InitLogger(warnOut io.Writer, errOut io.Writer) ExceptionSupportLogger {
	consoleLog := newLogrusLogger(logrus.InfoLevel, os.Stdout)
	warnLog := newLogrusLogger(logrus.WarnLevel, warnOut)
	errLog := newLogrusLogger(logrus.ErrorLevel, errOut)

	return NewExceptionSupportLogger(&consoleLog, &warnLog, &errLog)
}
