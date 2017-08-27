package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const configPath = "config.yml"

var db *gorm.DB
var config *ConfigStruct

func main() {
	runHttpsServer(
		config.WebServer.Host,
		config.WebServer.Port,
		config.WebServer.BindingPath)
}

func init() {
	var err error
	config, err = LoadConfig(configPath)
	if err != nil {
		log.Fatal("config: can't read config")
	}

	db, err = gorm.Open(
		"postgres",
		fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
			config.DataBase.Host,
			config.DataBase.User,
			config.DataBase.Name,
			config.DataBase.Pass))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	initMigration()
}

func initMigration() {
	db.AutoMigrate(&User{}, &Lock{})
	db.Model(&Lock{}).AddUniqueIndex("idx_lock_serial", "serial")
}
