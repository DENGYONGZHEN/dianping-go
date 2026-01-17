package postgre

import (
	"hm-dianping-go/internal/infra/config"
	"hm-dianping-go/internal/infra/logger"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func Init(cfg config.Config) (*sqlx.DB, error) {

	logger.Log.Info("db", zap.String("url", cfg.DBSource))
	db, err := sqlx.Connect("postgres", cfg.DBSource)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)

	return db, nil
}
