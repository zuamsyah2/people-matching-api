package db

import (
	"people-matching-api/helper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() (*gorm.DB, error) {
	user := helper.GetEnvWithKey("USER")
	pass := helper.GetEnvWithKey("PASS")
	host := helper.GetEnvWithKey("HOST")
	port := helper.GetEnvWithKey("PORT")
	dbname := helper.GetEnvWithKey("DBNAME")

	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
