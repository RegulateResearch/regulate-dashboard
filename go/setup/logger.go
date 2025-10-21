package setup

import (
	"frascati/comp/logger"
	"os"

	"github.com/sirupsen/logrus"
)

func setupEnhanceLogger(warnLogFile *os.File, errLogFile *os.File) logger.EnhancedLogger {
	infoLogger := logger.NewLogrusLogger(logrus.InfoLevel, os.Stdout)
	warnLogger := logger.NewLogrusLogger(logrus.WarnLevel, warnLogFile)
	errLogger := logger.NewLogrusLogger(logrus.ErrorLevel, errLogFile)

	return logger.NewEnhancedLogger(&infoLogger, &warnLogger, &errLogger)
}
