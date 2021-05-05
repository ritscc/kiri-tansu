package main

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"testing"

	model "./db/model"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	gomysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	DBNAME     = "kiritan"
	USERNAME   = "kiritan"
	PASSWORD   = "kiritan"
	IPADDRESS  = "127.0.0.1"
	PORTNUMBER = "3307"
)

func Test_User_DB(t *testing.T) {

	MigrateBD()

	db, err := DatabaseConnection()
	if err != nil {
		log.Fatal(err)
		return
	}

	result := db.Create(&model.User{
		ID: 1,
		Nickname: "Hoge",
		Role:     1,
	}).Error

	if result != nil {
		t.Error(result)
		return
	}

	testCase := []struct {
		expected model.User
	}{
		{
			expected: model.User{
				ID:       1,
				Nickname: "Hoge",
				Role:     1,
			},
		},
	}

	for i, testCase := range testCase {
		var actual model.User
		db.First(&actual)
		expected := testCase.expected
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("errorNumber:%d actual: %v, expect: %v", i, actual, expected)
			db.Delete(&actual)
			return
		}
		db.Delete(&actual)
	}

}

func DatabaseConnection() (*gorm.DB, error) {
	dsn := USERNAME + ":" + PASSWORD + "@tcp(" + IPADDRESS + ":" + PORTNUMBER + ")/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(gomysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	return db, err
}

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

	m, err := migrate.NewWithDatabaseInstance("file://db/migration", "mysql", driver)
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
