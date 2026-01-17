package handler

import (
	"hm-dianping-go/internal/infra/logger"
	"hm-dianping-go/internal/usecase/user"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase user.UserUseCase
}

func (uc *UserHandler) Login(c *gin.Context) {

	l := logger.LoggerFromContext(c.Request.Context())
	l.Info("user_handler处理开始")
	c.JSON(200, gin.H{
		"message": "pong",
	})

}
