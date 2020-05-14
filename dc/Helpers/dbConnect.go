package Helpers

import (
	"dataclips/Models"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetVarFromEnv() string {
	host := os.Getenv("Db_host")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("Db_port")
	if port == "" {
		port = "5432"
	}
	user := os.Getenv("Db_user")
	if user == "" {
		user = "postgres"
	}
	password := os.Getenv("Db_pass")
	if password == "" {
		password = "1234"
	}
	dbname := os.Getenv("Db_name")
	if dbname == "" {
		dbname = "Yuso"
	}
	if os.Getenv("DATABASE_URL") != "" {
		ConnectionString := os.Getenv("DATABASE_URL")
		return ConnectionString
	}
	ConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	return ConnectionString
}
func DbConnect() *gorm.DB {
	connection := GetVarFromEnv()
	db, err := gorm.Open("postgres", connection)
	if err != nil {
		fmt.Print(err)
	}
	// defer db.Close()
	return db
}

func Migration() {
	db := DbConnect()
	db.AutoMigrate(&Models.Dataclip{})

}
