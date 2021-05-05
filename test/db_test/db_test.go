package db

import (
	"log"
	"testing"
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

func Test_DB(t *testing.T) {

	MigrateBD()
	defer ResetBD()

	db, err := DatabaseConnection()
	if err != nil {
		log.Fatal(err)
		return
	}
	test_db_user(db, t)
	test_db_tag(db, t)

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
