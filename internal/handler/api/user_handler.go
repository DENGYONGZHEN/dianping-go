package api

import (
	"errors"
	"hm-dianping-go/internal/handler/response"
	"hm-dianping-go/internal/infra/logger"
	"hm-dianping-go/internal/usecase/user"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase user.UserUseCase
}

func (uc *UserHandler) Login(c *gin.Context) {

	ctx := c.Request.Context()
	l := logger.LoggerFromContext(ctx)

	l.Info("user_handler处理开始")

	phone := c.Query("phone")

	err := uc.userUseCase.Login(ctx, phone)

	if err != nil {
		switch {
		case errors.Is(err, user.ErrPhoneInvalid):
			response.Fail(c, 402, err.Error())
		default:
			response.Fail(c, 500, "服务器错误")
		}
		return
	}

	response.Success(c, nil)
	l.Info("user_handler处理结束")
}
