package application

import (
	"fmt"
	"os"
	"sync"
	"typograph_back/src/model"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	GlobalDB *gorm.DB
	onceDB   sync.Once
)

func InitializeDB(lvl logger.LogLevel) {
	onceDB.Do(func() {
		dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PASSWORD"))

		var err error
		GlobalDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(lvl),
		})

		if err != nil {
			log.Fatalf("Error connecting to database: %v", err)
		}

		sqlDB, err := GlobalDB.DB()
		if err != nil {
			log.Fatalf("Error getting underlying sql.DB: %v", err)
		}

		sqlDB.SetMaxIdleConns(3)

		err = GlobalDB.AutoMigrate(
			model.User{},
			model.Role{},
			model.Permission{},
			model.Paragraph{},
			model.Race{},
			model.UserRaceResult{},
		)
		if err != nil {
			log.Fatalf("Migration error: %v", err)
		}
	})
}
