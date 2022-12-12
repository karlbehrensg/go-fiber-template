package database

import (
	"os"

	"github.com/karlbehrensg/go-fiber-template/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	loggerGorm "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// GetDbClient generates the client for the database
func GetDbClient() *gorm.DB {

	dataSource := os.ExpandEnv("host=${DB_HOST} user=${DB_USER} password=${DB_PASSWORD} dbname=${DB_NAME} port=${DB_PORT} sslmode=${DB_SSL_MODE} TimeZone=${DB_TIME_ZONE}")
	client, err := gorm.Open(postgres.Open(dataSource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   os.ExpandEnv("${DB_SCHEMA}."),
			SingularTable: false,
		},
		Logger: loggerGorm.Default.LogMode(loggerGorm.Silent),
	})
	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info("Database connected")

	return client
}

// Migrate create the tables in the database
func Migrate(client *gorm.DB, models ...interface{}) {
	client.AutoMigrate(models...)
}
