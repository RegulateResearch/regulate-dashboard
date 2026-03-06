package setup

import (
	"frascati/comp/background"
	"frascati/comp/logging"
)

func setupBackgroundProcessor(logger logging.ExceptionSupportLogger) background.Processor {
	return background.NewProcessor(logger)
}
