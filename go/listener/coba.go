package listener

import (
	"context"
	"errors"
	"fmt"
	"frascati/exception"
	"frascati/lambda"
	"frascati/pbuf"
	"frascati/prep/logger"
	"net/http"
	"strings"
)

type CobaListener struct {
	pbuf.UnimplementedGreeterServer
	logger logger.EnhancedLogger
}

func NewCobaListener(logger logger.EnhancedLogger) *CobaListener {
	lis := CobaListener{
		logger: logger,
	}
	return &lis
}

func (l *CobaListener) SayHello(ctx context.Context, reqObj *pbuf.HelloRequest) (*pbuf.HelloResponse, error) {
	res, exc := execAndLog(ctx, reqObj, "SayHello", l.logger, func(ctx context.Context, req *pbuf.HelloRequest) (*pbuf.HelloResponse, exception.Exception) {
		resMessage := fmt.Sprintf("Hello, %s", req.GetName())

		res := &pbuf.HelloResponse{
			Header: &pbuf.ResponseHeader{
				Status:  http.StatusOK,
				Message: "success",
			},
			Message: resMessage,
		}

		return res, nil
	})

	return res, exc
}

func (l *CobaListener) SayHelloMultiple(ctx context.Context, reqObj *pbuf.HelloRequestMultiple) (*pbuf.HelloResponse, error) {
	res, exc := execAndLog(ctx, reqObj, "SayHelloMultiple", l.logger, func(ctx context.Context, req *pbuf.HelloRequestMultiple) (*pbuf.HelloResponse, exception.Exception) {
		names := lambda.MapList(req.Name, func(name string) string {
			return fmt.Sprintf("\t* %s", name)
		})

		namesStr := strings.Join(names, "\n")
		resMessage := fmt.Sprintf("Hello to you all:\n%s", namesStr)

		res := &pbuf.HelloResponse{
			Header: &pbuf.ResponseHeader{
				Status:  http.StatusOK,
				Message: "success",
			},
			Message: resMessage,
		}

		return res, nil
	})

	return res, exc
}

func (l *CobaListener) SayHelloError(ctx context.Context, reqObj *pbuf.HelloRequest) (*pbuf.HelloResponse, error) {
	res, err := execAndLog(ctx, reqObj, "SayHelloError", l.logger, func(ctx context.Context, req *pbuf.HelloRequest) (*pbuf.HelloResponse, exception.Exception) {
		resMessage := "This method will always return error"
		err := errors.New("wrong method call, this method always error")
		exc := exception.NewBaseException(exception.CAUSE_USER, "grpc", err.Error(), err)

		res := &pbuf.HelloResponse{
			Header: &pbuf.ResponseHeader{
				Status:  http.StatusBadRequest,
				Message: "bad request huehuehue",
			},
			Message: resMessage,
		}

		return res, exc
	})

	return res, err
}
