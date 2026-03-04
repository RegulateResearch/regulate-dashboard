package listener

import (
	"context"
	"frascati/comp/logging"
	"frascati/exception"
	"time"
)

func execAndLog[Req any, Res any](ctx context.Context, req Req, reqName string, logger logging.ExceptionSupportLogger, proc func(context.Context, Req) (Res, exception.Exception)) (Res, exception.Exception) {
	start := time.Now()
	res, exc := proc(ctx, req)
	end := time.Now()

	latency := end.Sub(start).String()
	entry := map[string]any{
		"method":  "GRPC",
		"request": reqName,
		"latency": latency,
	}

	logger = logger.WithFields(entry)

	if exc != nil {
		logger.LogException(exc)
	} else {
		logger.Infof("REQUEST %s SUCCESS", reqName)
	}

	return res, exc
}
