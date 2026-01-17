package postgres

import (
	"context"
	"hm-dianping-go/internal/domain"
	"hm-dianping-go/internal/repo"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) FindUserByPhone(ctx context.Context, phone string) (domain.User, error) {
	panic("unimplemented")
}

var _ repo.UserRepository = (*UserRepo)(nil)
