package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/HEEPOKE/fiber-hexagonal/pkg/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		configs.Cfg.DBHost, configs.Cfg.DBUserName, configs.Cfg.DBPassword, configs.Cfg.DBName, configs.Cfg.DBPort, configs.Cfg.DBssl, configs.Cfg.DBTimezone)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Error,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
