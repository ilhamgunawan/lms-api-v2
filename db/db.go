package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

var db *gorp.DbMap

func Init() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_SECRET"), os.Getenv("DB_NAME"))

	var err error
	db, err = ConnectDB(dbinfo)
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB(dataSourceName string) (*gorp.DbMap, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	//dbmap.TraceOn("[gorp]", log.New(os.Stdout, "golang-gin:", log.Lmicroseconds)) //Trace database requests

	// Register struct to table
	dbmap.AddTableWithName(UserAccount{}, "user_account")
	dbmap.AddTableWithName(UserLoginData{}, "user_login_data")
	dbmap.AddTableWithName(Course{}, "course")
	dbmap.AddTableWithName(Topic{}, "topic")
	dbmap.AddTableWithName(Mission{}, "mission")
	dbmap.AddTableWithName(MissionType{}, "mission_type")
	dbmap.AddTableWithName(CourseUser{}, "course_user")
	dbmap.AddTableWithName(TopicUser{}, "topic_user")
	dbmap.AddTableWithName(MissionUser{}, "mission_user")

	return dbmap, nil
}

func GetDB() *gorp.DbMap {
	return db
}
