package middleware

import (
	"errors"
	"frascati/comp/logging"
	"frascati/exception"
	"time"

	"github.com/gin-gonic/gin"
)

type LoggerMiddleware struct {
	logger logging.ExceptionSupportLogger
}

func NewLoggerMiddleware(logger logging.ExceptionSupportLogger) LoggerMiddleware {
	return LoggerMiddleware{logger: logger}
}

func (m LoggerMiddleware) LogActivities(ctx *gin.Context) {
	startTime := time.Now()
	ctx.Next()
	endTime := time.Now()

	latency := endTime.Sub(startTime).String()
	reqMethod := ctx.Request.Method
	reqHost := ctx.Request.Host
	reqURI := ctx.Request.RequestURI
	statusCode := ctx.Writer.Status()
	clientIP := ctx.ClientIP()

	fields := map[string]any{
		"method":    reqMethod,
		"uri":       reqURI,
		"status":    statusCode,
		"latency":   latency,
		"client_ip": clientIP,
		"host":      reqHost,
	}

	logger := m.logger.WithFields(fields)

	if lastErr := ctx.Errors.Last(); lastErr != nil {
		var exc exception.Exception
		if errors.As(lastErr, &exc) {
			logger.LogException(exc)
		} else {
			logger.Errorf("uncaught error: %v", lastErr.Err)
		}
	} else {
		logger.Info("SUCCESS")
	}

}
