package domain

import (
	"github.com/stakkato95/service-engineering-go-lib/logger"
	"github.com/stakkato95/twitter-service-tweets/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbRepo interface {
	GetDb() *gorm.DB
}

type gormDbRepo struct {
	db *gorm.DB
}

func NewDbRepo() DbRepo {
	db, err := gorm.Open(postgres.Open(config.AppConfig.DbSource), &gorm.Config{})
	if err != nil {
		logger.Fatal("failed to connect database: " + err.Error())
	}

	// Migrate the schema
	if err := db.AutoMigrate(&Tweet{}, &User{}, &Subscription{}); err != nil {
		logger.Fatal("failed to migrate database: " + err.Error())
	}

	return &gormDbRepo{db}
}

func (r *gormDbRepo) GetDb() *gorm.DB {
	return r.db
}
