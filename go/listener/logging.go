package listener

import (
	"context"
	"frascati/exception"
	"frascati/prep/logger"
	"time"
)

func execAndLog[Req any, Res any](ctx context.Context, req Req, reqName string, logger logger.EnhancedLogger, proc func(context.Context, Req) (Res, exception.Exception)) (Res, exception.Exception) {
	start := time.Now()
	res, exc := proc(ctx, req)
	end := time.Now()

	latency := end.Sub(start).String()
	entry := map[string]any{
		"method":  "GRPC",
		"request": reqName,
		"latency": latency,
	}

	if exc != nil {
		if exc.Cause() == exception.CAUSE_INTERNAL {
			logger.WithFieldsError(entry).WithField("errorObj", exc.ToMap()).Errorf("error:\n%s", exc.Error())
		} else {
			logger.WithFieldsWarn(entry).WithField("errorObj", exc.ToMap()).Warnf("warning:\n%s", exc.Error())
		}
	} else {
		logger.WithFieldsInfo(entry).Infof("REQUEST %s SUCCESS", reqName)
	}

	return res, exc
}
