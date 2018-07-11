package model

import (
	"math/rand"
	"strconv"
	"time"

	"log"

	"reflect"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
	"github.com/web-demo/config"
)

// DB : DB pooling
var DB *sqlx.DB

// DBInit : InitDB
func DBInit() {
	config := config.GetConfig()
	dialect := (*config.DB).Dialect
	connection := (*config.DB).Username + ":" + (*config.DB).Password + "@tcp(" + (*config.DB).IP + ":" + (*config.DB).Port + ")/" + (*config.DB).Table + "?parseTime=true"
	DB, _ = sqlx.Open(dialect, connection)
	DB.SetMaxOpenConns(2000)
	DB.SetMaxIdleConns(1000)
	DB.SetConnMaxLifetime(14400)

	if err := DB.Ping(); err != nil {
		log.Println("DB ping err to: ", (*config.DB).IP)
		log.Println("Err: ", err)
	} else {
		log.Println("DB connected OK to: ", (*config.DB).IP)
	}

}

// GenID : generate ID
func GenID() string {
	uuid, _ := uuid.NewV4()
	return uuid.String()
}

// RandString return a rand string with length n
func RandString(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// GetCurrentTimeStampUnixTime : get current timestamp in milliseconds
func GetCurrentTimeStampUnixTime() string {
	t := GetCurrentTimeStamp()
	return GetUnixTime(t)
}

// GetUnixTime : convert time.Time to timestamp string in milliseconds
func GetUnixTime(aTime time.Time) string {
	return strconv.FormatInt(aTime.UnixNano()/1000000, 10)
}

// GetCurrentTimeStamp : get current timestamp in time.Time
func GetCurrentTimeStamp() time.Time {
	// remove millisecond for sql time sync
	return time.Unix(time.Now().UnixNano()/1000000000, 0)
}

func GetCreateSQLPreString(table string) string {
	r := "INSERT into " + table + " ( "
	var p interface{}
	if table == "user" {
		p = User{}
	} else if table == "post" {
		p = Post{}
	} else if table == "comment" {
		p = Comment{}
	}

	pv := reflect.ValueOf(p)
	for i := 0; i < pv.Type().NumField(); i++ {
		field := pv.Type().Field(i).Tag.Get("db")
		if field != "-" {

			if i == pv.Type().NumField()-1 {
				r += field + ")values("
			} else {
				r += field + ", "
			}
		}
	}
	for i := 0; i < pv.Type().NumField(); i++ {
		field := pv.Type().Field(i).Tag.Get("db")
		if field != "-" {
			if i == pv.Type().NumField()-1 {
				r += "?)"
			} else {
				r += "?,"
			}
		}
	}

	return r
}
