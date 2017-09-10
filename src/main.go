package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const configPath = "config.yml"

var db *gorm.DB
var config *ConfigStruct

func main() {
	f, err := os.OpenFile("output.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)

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

	go runTelegramBot(config.Telegram.Token)
	runHTTPSServer(
		config.WebServer.Host,
		config.WebServer.Port,
		config.WebServer.BindingPath)
}

func initMigration() {
	db.AutoMigrate(&User{}, &Lock{})
	db.Model(&Lock{}).AddUniqueIndex("idx_lock_serial", "serial")
}
