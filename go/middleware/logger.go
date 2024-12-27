package middleware

import (
	"errors"
	"frascati/exception"
	"frascati/prep/logger"
	"time"

	"github.com/gin-gonic/gin"
)

type LoggerMiddleware struct {
	logger logger.EnhancedLogger
}

func NewLoggerMiddleware(logger logger.EnhancedLogger) LoggerMiddleware {
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

	if lastErr := ctx.Errors.Last(); lastErr != nil {
		m.logger.WithFieldsInfo(fields).Error(lastErr)

		var exc exception.Exception
		if errors.As(lastErr, &exc) {
			if exc.Cause() == exception.CAUSE_INTERNAL {
				m.logger.WithFieldsError(fields).Errorf("error:\n%s", exc.Verbose())
			} else {
				m.logger.WithFieldsWarn(fields).Warnf("warning:\n%s", exc.Verbose())
			}
		}

		return
	}

	m.logger.WithFieldsInfo(fields).Infof("REQUEST %s %s SUCCESS", reqMethod, reqURI)

}
