package db

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrateBD() {
	db, err := sql.Open("mysql", "kiritan:kiritan@tcp(localhost:3307)/kiritan")

	if err != nil {
		fmt.Println("DBの接続に失敗しました")
		fmt.Printf("%s\n", err)
		return
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})

	if err != nil {
		fmt.Println("Hoge")
		fmt.Printf("%s\n", err)
		return
	}

	m, err := migrate.NewWithDatabaseInstance("file://../../db/migration", "mysql", driver)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		fmt.Printf("%s\n", err)
		return
	}

	db.Close()
}

func ResetBD() {
	db, err := sql.Open("mysql", "kiritan:kiritan@tcp(localhost:3307)/kiritan")

	if err != nil {
		fmt.Println("DBの接続に失敗しました")
		fmt.Printf("%s\n", err)
		return
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})

	if err != nil {
		fmt.Println("Hoge")
		fmt.Printf("%s\n", err)
		return
	}

	m, err := migrate.NewWithDatabaseInstance("file://../../db/migration", "mysql", driver)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	err = m.Down()
	if err != nil && err != migrate.ErrNoChange {
		fmt.Printf("%s\n", err)
		return
	}

	db.Close()
}
