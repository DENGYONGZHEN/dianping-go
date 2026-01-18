package api

import "hm-dianping-go/internal/usecase/user"

type APIHandler struct {
	UserHandler UserHandler
}

func NewAPIHandler(userUseCase *user.UserUseCase) *APIHandler {
	return &APIHandler{
		UserHandler: UserHandler{
			userUseCase: *userUseCase,
		},
	}
}
