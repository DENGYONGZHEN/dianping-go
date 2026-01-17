package middleware

import (
	"hm-dianping-go/internal/infra/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func LoggerMiddle() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		requestId := uuid.NewString()

		l := logger.Log.With(
			zap.String("request_id", requestId),
			zap.String("path", ctx.Request.URL.Path),
		)

		oldCtx := ctx.Request.Context()

		newCtx := logger.WithContext(oldCtx, l)

		ctx.Request = ctx.Request.WithContext(newCtx)

		ctx.Next()
	}

}
