package main

import (
	"context"
	"hm-dianping-go/internal/infra/config"
	"hm-dianping-go/internal/infra/logger"
	"hm-dianping-go/internal/infra/postgre"
	"hm-dianping-go/internal/infra/redis"
	"hm-dianping-go/internal/middleware"
	"hm-dianping-go/internal/routes"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// logger init
	if err := logger.New(); err != nil {
		panic(err)
	}
	defer logger.Log.Sync()

	//config init
	cfg, err := config.LoadConfig(".")
	if err != nil {
		logger.Log.Fatal("cannot load config", zap.Error(err))
	}

	// db init
	db, err := postgre.Init(cfg.DBSource)
	if err != nil {
		logger.Log.Fatal("cannot connect to db", zap.Error(err))
	}
	defer db.Close()

	//redis init
	rdb, err := redis.New(cfg.RedisAddress)
	if err != nil {
		logger.Log.Fatal("cannot connect to redis", zap.Error(err))
	}
	defer rdb.Close()

	router := gin.New()
	router.Use(
		middleware.LoggerMiddle(),
		gin.Recovery(),
	)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	routes.RegisteRouter(router, db, rdb)

	srv := &http.Server{
		Addr:    cfg.HTTPServerAddress,
		Handler: router,
	}

	go func() {
		logger.Log.Info("server started", zap.String("port", cfg.HTTPServerAddress))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatal("listen failed", zap.Error(err))
		}
	}()

	// graceful shutdown
	// syscall.SIGINT   // 人为中断（Ctrl + C）
	// syscall.SIGTERM,	// 	kill <pid>（不加 -9）	// Docker / Kubernetes 停止容器	 // systemd 停止服务	// 云平台缩容 / 重启 Pod
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	defer stop()

	<-ctx.Done()
	logger.Log.Info("shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Log.Error("server shutdown failed", zap.Error(err))
	}

	logger.Log.Info("server exited gracefully")
}
