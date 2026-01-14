package setup

import (
	"frascati/comp/logging"
	"os"

	"github.com/sirupsen/logrus"
)

func setupEnhanceLogger(warnLogFile *os.File, errLogFile *os.File) logging.EnhancedLogger {
	infoLogger := logging.NewLogrusLogger(logrus.InfoLevel, os.Stdout)
	warnLogger := logging.NewLogrusLogger(logrus.WarnLevel, warnLogFile)
	errLogger := logging.NewLogrusLogger(logrus.ErrorLevel, errLogFile)

	return logging.NewEnhancedLogger(&infoLogger, &warnLogger, &errLogger)
}
