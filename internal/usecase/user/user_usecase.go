package user

import (
	"context"
	"errors"
	"hm-dianping-go/internal/infra/logger"
	"hm-dianping-go/internal/infra/redis"
	"hm-dianping-go/internal/repo"
	"hm-dianping-go/internal/utils"

	"go.uber.org/zap"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidCode  = errors.New("invalid verify code")
	ErrPhoneInvalid = errors.New("invalid phone number")
	ErrCodeExpired  = errors.New("verify code expired")
)

type UserUseCase struct {
	ur    repo.UserRepository
	redis *redis.RedisClient
}

func NewUserUseCase(ur repo.UserRepository, redis *redis.RedisClient) *UserUseCase {
	return &UserUseCase{
		ur:    ur,
		redis: redis,
	}
}

func (uc *UserUseCase) Login(ctx context.Context, phone string) error {

	l := logger.LoggerFromContext(ctx)
	l.Info("user_usecase处理开始")

	//验证手机号是否合法
	if !utils.ValidatePhoneNumber(phone) {
		l.Info("user_usecase", zap.String("手机号不正确", phone))
		return ErrPhoneInvalid
	}

	//生成随机的验证码
	code, err := utils.GenerateCode()
	if err != nil {
		l.Info("user_usecase", zap.String("生成验证码错误", err.Error()))
		return err
	}

	if err := uc.redis.SetLoginCode(ctx, phone, code); err != nil {
		l.Info("user_usecase", zap.String("redis设置验证码错误", err.Error()))
		return err
	}

	l.Info("user_usecase处理结束", zap.String("code", code))

	return nil
}
