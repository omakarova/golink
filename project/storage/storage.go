package storage

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"../mymodels"
	"encoding/base64"

)

var db *sql.DB
var err error


func init() {
	fmt.Println("init")
	db, err = sql.Open("mysql", "mylink:123@/mylink?charset=utf8")
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func SaveUserdata(user mymodels.NewUser) {
	// insert
	stmt, err := db.Prepare("INSERT users SET name=?,auth=?")
	checkErr(err)
	var siteAuth = base64.StdEncoding.EncodeToString([]byte(user.Username + ":" + user.Password))
	res, err := stmt.Exec(user.Username, "Basic "+siteAuth)
	checkErr(err)
	fmt.Println(res)
}
