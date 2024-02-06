package db

import (
	"fmt"
	"log/slog"
	"os"
	"sync"

	"github.com/gaoozi/auc/config"
  "github.com/gaoozi/auc/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func connect() {
	var err error

	config := config.GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Database.User, config.Database.Password, config.Database.Url, config.Database.Name,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("database connect is not available: ", err)
		os.Exit(1)
	}

  db.AutoMigrate(&model.User{})
	slog.Info("database connection is available")
}

func GetDb() *gorm.DB {
	once.Do(connect)
	return db
}
