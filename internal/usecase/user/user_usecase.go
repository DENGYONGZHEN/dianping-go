package user

import "hm-dianping-go/internal/repo"

type UserUseCase struct {
	ur repo.UserRepository
}

func NewUserUseCase(ur repo.UserRepository) *UserUseCase {
	return &UserUseCase{
		ur: ur,
	}
}

func (uc *UserUseCase) Login() {

}
