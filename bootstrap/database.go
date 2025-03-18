package bootstrap

import (
	"fmt"

	"go-fiber/core/logs"
	"go-fiber/domain/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabaseConnection(env *Env) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Vientiane",
		env.Database.DBHost,
		env.Database.Username,
		env.Database.Password,
		env.Database.DBName,
		env.Database.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("can not connect to database")
	}
	logs.Info("database connection success")

	// Migrate(db)

	return db
}

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
	//Migrate table here
	&entities.Author{},
	&entities.Book{},
	)
	if err != nil {
		logs.Error(err.Error())
	}
	logs.Info("Migrate successfully")
}
