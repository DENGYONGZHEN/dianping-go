package repo

import (
	"context"
	"hm-dianping-go/internal/domain"
)

type UserRepository interface {
	FindUserByPhone(ctx context.Context, phone string) (domain.User, error)
}
