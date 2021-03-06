package storage

import (
	"fmt"
	"net/http"
	"database/sql"
)


func AddStat(req *http.Request, linkid int) {

	referer := req.Referer()
	// insert
	stmt, err := db.Prepare("INSERT statistics SET linkid=?, referer=?, f_date_time=NOW()")
	checkErr(err)
	res, err := stmt.Exec(linkid, referer)
	checkErr(err)
	fmt.Println(res)

}

func GetNumberOfClicks(linkId int) (int, error) {
	rows, err := db.Query("SELECT count(id) FROM statistics where linkid=?", linkId)
	if (err != nil) {
		return -1, err
	}
	if(rows.Next()) {
		var numberOfClicks int
		err = rows.Scan(&numberOfClicks)
		checkErr(err)
		return numberOfClicks, nil
	}
	return 0, nil
}

func GetTopReferrersByUser(currentUserId int) []string {
	rows, err := db.Query("select referer, count(referer) from statistics as stat inner join links as li" +
								 " on stat.linkid=li.id " +
								 "where li.userid=? and referer is not null and referer !=\"\" " +
								 "group by referer order by count(referer) desc limit 0,20", currentUserId)
	checkErr(err)
	alist := make([]string, 0, 20)
	for rows.Next() {
		var referer string
		var count int
		err = rows.Scan(&referer, &count)
		checkErr(err)
		alist = append(alist, referer)
	}
	return alist
}

func GetStatDataForUser(currentUserId int, interval string) map[string]int {
	var rows *sql.Rows
	var err error
	switch interval {
		case "d": rows, err = db.Query("SELECT COUNT(`id`), DATE_FORMAT(`f_date_time`, '%Y %m %d') as dat " +
				"FROM statistics where linkid in(select id from links where userid=?) GROUP BY dat" +
				" ORDER BY dat DESC;", currentUserId)
	    case "h": rows, err = db.Query("SELECT COUNT(`id`), DATE_FORMAT(`f_date_time`, '%Y %m %d %H') as dat " +
			"FROM statistics where linkid in(select id from links where userid=?) GROUP BY dat" +
			" ORDER BY dat DESC;", currentUserId)
	case "i": rows, err = db.Query("SELECT COUNT(`id`), DATE_FORMAT(`f_date_time`, '%Y %m %d %H %i') as dat " +
		"FROM statistics where linkid in(select id from links where userid=?) GROUP BY dat" +
		" ORDER BY dat DESC;", currentUserId)
	default: rows, err = db.Query("SELECT COUNT(`id`), DATE_FORMAT(`f_date_time`, '%Y %m %d %H') as dat " +
		"FROM statistics where linkid in(select id from links where userid=?) GROUP BY dat" +
		" ORDER BY dat DESC;", currentUserId)
	}
	checkErr(err)
	amap := make(map[string]int)
	for rows.Next() {
		var interval string
		var count int
		err = rows.Scan(&count, &interval)
		checkErr(err)
		amap[interval] = count
	}
	return amap
}

