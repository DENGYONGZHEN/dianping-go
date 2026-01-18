package routes

import (
	"hm-dianping-go/internal/handler/api"
	"hm-dianping-go/internal/infra/redis"
	"hm-dianping-go/internal/repo/postgres"
	"hm-dianping-go/internal/usecase/user"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisteRouter(router *gin.Engine, db *sqlx.DB, rdb *redis.RedisClient) {

	userRepo := postgres.NewUserRepo(db)
	userUseCase := user.NewUserUseCase(userRepo, rdb)
	h := api.NewAPIHandler(userUseCase)

	{
		user := router.Group("/user")
		user.POST("/code", h.UserHandler.Login)
	}
}
