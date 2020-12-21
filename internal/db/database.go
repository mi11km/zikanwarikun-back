package database

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/mi11km/zikanwarikun-back/config"
	d "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	driver = config.Cfg.DB.Driver
	dns    = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Cfg.DB.User, config.Cfg.DB.Password, config.Cfg.DB.Address, config.Cfg.DB.Name)
	sourceURL = config.Cfg.DB.SourceURL
)

var Db *gorm.DB

func Init() {
	db, err := gorm.Open(d.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalf("action=failed to connect db, err=%s", err)
	}
	db.Logger.LogMode(logger.Info)
	Db = db
	log.Printf("action=connect db, status=success")
}

func Migrate() {
	db, err := Db.DB()
	if err != nil {
		log.Fatalf("action=failed to get *sql.DB, err=%s", err)
	}
	instance, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("action=create driver instance for migrate, err=%s", err)
	}
	m, err := migrate.NewWithDatabaseInstance(sourceURL, driver, instance)
	if err != nil {
		log.Fatalf("action=read migrations, err=%s", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("action=migrate, err=%s", err)
	}
	log.Printf("action=migrate, status=success")
}
