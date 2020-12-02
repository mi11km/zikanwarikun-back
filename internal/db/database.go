package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// todo configファイルから読み込む
const (
	driver    = "mysql"
	dns       = "user:password@tcp(localhost)/zikanwarikun?charset=utf8&parseTime=true"
	sourceURL = "file://internal/db/migrations/mysql"
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
