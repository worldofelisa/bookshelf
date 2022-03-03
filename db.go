package main

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
)

// generateDSN sets the Data Source Name using the package Viper to retrieve information from the config.yaml file
func generateDSN() string {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	return strings.Join([]string{
		viper.GetString("db.username"),
		":",
		viper.GetString("db.password"),
		"@tcp(",
		viper.GetString("db.host"),
		")/",
		viper.GetString("db.db_name"),
	}, "")
}

// connectToDB initializes gorm.DB with my existing DB connection
//returns the reference to the original variable with connection still open to DB
func connectToDB() *gorm.DB {
	//open up mySQL using the generateDSN (pre-generated)
	sqlDB, err := sql.Open("mysql", generateDSN())
	fatalErrorHandler(err)

	//initializes gorm.db, which is an ORM
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	fatalErrorHandler(err)

	return gormDB
}

//migrateDB will create/update the tables but will not delete unused columns
func migrateDB(conn *gorm.DB) {
	err := conn.AutoMigrate(
		&Book{},
		&Author{},
		&BookAuthor{},
		&Genre{},
		&Tag{},
		&BookTag{},
		&Shelf{},
		&ShelfBook{},
		&User{},
		&Review{},
		&ReadStatus{},
		&PageTracker{})
	fatalErrorHandler(err)
}
