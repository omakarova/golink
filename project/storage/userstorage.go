package storage

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"../mymodels"
	"../config"
	"encoding/base64"

	"errors"
)

var db *sql.DB
var err error


func init() {
	fmt.Println("init db")
	db, err = sql.Open("mysql", config.DB_CONNECTION_STRING)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func DoesUserExist(user mymodels.NewUser) bool {
	rows, err := db.Query("SELECT id FROM users where username=? AND password=?", user.Username, user.Password)
	checkErr(err)

	if(!rows.Next()) {
		return false
	}
	return true
}

func SaveUserData(user mymodels.NewUser) {
	// insert
	stmt, err := db.Prepare("INSERT users SET username=?, password=?, auth=?")
	checkErr(err)
	var siteAuth = base64.StdEncoding.EncodeToString([]byte(user.Username + ":" + user.Password))
	res, err := stmt.Exec(user.Username, user.Password, "Basic "+siteAuth)
	checkErr(err)
	fmt.Println(res)
}

func GetUserDataByAuthString(auth string) (int, string, error) {
	// query
	rows, err := db.Query("SELECT id, username FROM users where auth=?", auth)
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		err = rows.Scan(&uid, &username)
		checkErr(err)
		return uid, username, nil
	}
	return -1, "", errors.New("no such user")
}


