package migration

import (
	"database/sql"
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	ErrMigrate = errors.New("migration error")
)

func Migrate() error {
	db, err := sql.Open("mysql", "kiritan:kiritan@tcp(localhost:3306)/kiritan")

	if err != nil {
		log.Fatalln("DBの接続に失敗しました")
		log.Fatal(err)
		return ErrMigrate
	}

	defer db.Close()

	driver, err := mysql.WithInstance(db, &mysql.Config{})

	if err != nil {
		log.Fatal(err)
		return ErrMigrate
	}

	m, err := migrate.NewWithDatabaseInstance("file://db/migration", "mysql", driver)
	if err != nil {
		log.Fatalln("マイグレーションの準備に失敗しました")
		log.Fatal(err)
		return ErrMigrate
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
		return ErrMigrate
	}

	if err == migrate.ErrNoChange {
		log.Println("マイグレーションは実行されませんでした")
	} else {
		log.Println("新たにマイグレーションを実行しました")
	}

	return nil
}
