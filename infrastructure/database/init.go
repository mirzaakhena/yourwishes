package database

import (
	"fmt"
	"time"
	"yourwishes/infrastructure/config"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewSQLiteDefault() (db *gorm.DB) {

	db, err := gorm.Open(sqlite.Open("local.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return db
}

func NewPostgresDefault() (db *gorm.DB) {

	cfg, err := config.ReadConfig()
	if err != nil {
		panic(err.Error())
	}

	if cfg.User == "" || cfg.Password == "" || cfg.Database == "" {
		panic(fmt.Errorf("user or password ord databaseName is empty"))
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%v", cfg.Host, cfg.Port, cfg.User, cfg.Database, cfg.Password, cfg.SSLMode)

	loggerMode := logger.Silent

	if cfg.LogMode {
		loggerMode = logger.Info
	}

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(loggerMode),
	})
	if err != nil {
		panic(err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err.Error())
	}

	sqlDB.SetMaxIdleConns(10)

	sqlDB.SetMaxOpenConns(10)

	sqlDB.SetConnMaxLifetime(10 * time.Second)

	return db
}
