package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"myapp/config"
	"myapp/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbpg *gorm.DB
var err error

func Init() {
	db, err := gorm.Open(postgres.Open(config.POSTGRESQL_CONNECTION), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
				Colorful:                  false,       // Disable color
			},
		),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("db connected!")
	models.AutoMigrate(db)
	dbpg = db

}

func DbManager() *gorm.DB {
	return dbpg
}
