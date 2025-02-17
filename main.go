package main

import (
	"fmt"
	"os"
	database "people-matching-api/db"
	_ "people-matching-api/docs"
	"people-matching-api/helper"
	"people-matching-api/router"
	"sync"

	"gorm.io/gorm"
)

var dbConnOnce sync.Once

func main() {
	helper.LoadEnv()
	db, err := database.Connection()
	if err != nil {
		fmt.Printf("error connecting database: %s", err.Error())
		os.Exit(1)
	}
	var dbConn *gorm.DB
	dbConnOnce.Do(func() {
		dbConn = db
	})

	database.Migrate(dbConn)

	r := router.UserRoute(dbConn)
	r.Run(":8080")
}
