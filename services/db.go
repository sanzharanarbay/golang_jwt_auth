package services

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/joho/godotenv/autoload"
)

var Db *gorm.DB
var err error
var e error

func init() {

	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbType := os.Getenv("DB_TYPE")

	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, DbUser, DbName, DbPassword)
	conn, err := gorm.Open(dbType, dbUri)

	Db = conn
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	fmt.Println("Connected to database on " + dbHost)

}

func GetDB() *gorm.DB {
	return Db
}
