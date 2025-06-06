package db

import (
	"log"

	"github.com/Belxz777/backgo/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(host string, port string, user string, password string, dbname string) *gorm.DB {
	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(
		&models.User{},
		&models.Training{},
		&models.GpsPoint{},
		&models.WorkoutData{},
	)

	return db
}
