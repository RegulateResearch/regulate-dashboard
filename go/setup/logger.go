package setup

import (
	"frascati/comp/logging"
	"io"
)

func setupLogger(warnLogFile io.Writer, errLogFile io.Writer) logging.ExceptionSupportLogger {
	return logging.InitLogger(warnLogFile, errLogFile)
}
