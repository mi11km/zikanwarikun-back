package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/mi11km/zikanwarikun-back/config"
)

var (
	driver = config.Cfg.DB.Driver
	dns    = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
		config.Cfg.DB.User, config.Cfg.DB.Password, config.Cfg.DB.Address, config.Cfg.DB.Name)
	sourceURL = config.Cfg.DB.SourceURL
)

var Db *sql.DB

func Init() {
	db, err := sql.Open(driver, dns)
	if err != nil {
		log.Fatalf("action=connect db, err=%s", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("action=db ping, err=%s", err)
	}
	Db = db
	log.Printf("action=connect db, status=success")
}

func Migrate() {
	if err := Db.Ping(); err != nil {
		log.Fatalf("action=db ping, err=%s", err)
	}
	instance, err := mysql.WithInstance(Db, &mysql.Config{})
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
