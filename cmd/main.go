package main

import (
	"hm-dianping-go/internal/infra/config"
	"hm-dianping-go/internal/infra/logger"
	"hm-dianping-go/internal/infra/postgre"
	"hm-dianping-go/internal/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	if err := logger.New(); err != nil {
		panic(err)
	}

	defer logger.Log.Sync()

	cfg, err := config.LoadConfig(".")
	if err != nil {
		logger.Log.Fatal("cannot load config", zap.Error(err))
	}

	_, err = postgre.Init(cfg)
	if err != nil {
		logger.Log.Fatal("cannot connnect to db", zap.Error(err))
	}

	router := gin.New()
	router.Use(middleware.LoggerMiddle())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	if err := router.Run(); err != nil {
		// logger.Log.Fatalf("failed to run server: %v", err)
	}
}
