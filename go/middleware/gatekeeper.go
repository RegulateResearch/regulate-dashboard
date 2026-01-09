package middleware

import (
	"frascati/comp/graceful"
	"frascati/comp/logger"
	"frascati/exception"
	"frascati/response"

	"github.com/gin-gonic/gin"
)

type GatekeeperMiddleware struct {
	gatekeeper graceful.Gatekeeper
	logger     logger.EnhancedLogger
}

func NewGatekeeperMiddleware(gatekeeper graceful.Gatekeeper, logger logger.EnhancedLogger) GatekeeperMiddleware {
	return GatekeeperMiddleware{
		gatekeeper: gatekeeper,
		logger:     logger,
	}
}

func (m GatekeeperMiddleware) Process(ctx *gin.Context) {
	err := m.gatekeeper.Process(func() {
		ctx.Next()
	})

	if err != nil {
		m.logger.Infof("request attempted while gate is closed")
		ctx.AbortWithStatusJSON(exception.GetExceptionHttpStatus(err), response.NewErrorResponse("request not accepted", err))
	}
}
