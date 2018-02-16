package storage

import (
	"fmt"
	"net/http"
)


func AddStat(req *http.Request, linkid int) {

	referer := req.Referer()
	// insert
	stmt, err := db.Prepare("INSERT statistic SET linkid=?, referer=?, f_date_time=NOW()")
	checkErr(err)
	res, err := stmt.Exec(linkid, referer)
	checkErr(err)
	fmt.Println(res)

}

