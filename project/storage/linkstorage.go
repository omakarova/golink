package storage

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"../mymodels"
	"../config"
    "math/rand"
	"time"
	"errors"
    s "strings"
	//"container/list"
)


func DoesLongLinkExist(link mymodels.NewLink, userId int) bool {
	rows, err := db.Query("SELECT id FROM links where longurl=? AND userid=?", link.URL, userId)
	checkErr(err)

	if(!rows.Next()) {
		return false
	}
	return true
}

func DoesShortLinkExist(shortlink string, userId int) bool {
	rows, err := db.Query("SELECT id FROM links where shorturl=? AND userid=?", shortlink, userId)
	checkErr(err)

	if(!rows.Next()) {
		return false
	}
	return true
}

func CreateNewLink(link mymodels.NewLink, userId int) *mymodels.NewLinkResponse {

	longlink := link.URL
	if(!s.HasPrefix(longlink, "http://") && !s.HasPrefix(longlink, "https://")) {
		longlink = "http://" + longlink
	}
	shortLink := "m" + generateRandomStr()
	// insert
	stmt, err := db.Prepare("INSERT links SET longurl=?, shorturl=?, userid=?")
	checkErr(err)
	res, err := stmt.Exec(longlink, shortLink, userId)
	checkErr(err)
	fmt.Println(res)

	newLinkResponse := mymodels.NewLinkResponse{ShortURL: shortLink}
	return &newLinkResponse
}

func GetLongUrlAndIdByShortLink(shortUrl string) (string, int, error) {
	rows, err := db.Query("SELECT id, longurl FROM links where shorturl=?", shortUrl)
	checkErr(err)

	if(rows.Next()) {
		var longurl string
		var linkid int
		err = rows.Scan(&linkid, &longurl)
		checkErr(err)
		return longurl, linkid, nil
	}
	return "", -1, errors.New("no such link")
}

func GetLinkIdByShortLinkAndUserId(shortUrl string, userId int) (int, string, error) {
	rows, err := db.Query("SELECT id, longurl FROM links where shorturl=? AND userid=?", shortUrl, userId)
	checkErr(err)

	if(rows.Next()) {
		var linkid int
		var longurl string
		err = rows.Scan(&linkid, &longurl)
		checkErr(err)
		return linkid, longurl, nil
	}
	return -1, "", errors.New("no such link")
}

func DeleteLink(shortlink string, userId int) {
	_, err := db.Query("DELETE FROM links where shorturl=? AND userid=?", shortlink, userId)
	checkErr(err)
}

func GetAllLinksByUserId(userId int) []string {
	rows, err := db.Query("SELECT shorturl FROM links where userid=?", userId)
	checkErr(err)
	alist := make([]string, 0, 20)
	for rows.Next() {
		var vLink string
		err = rows.Scan(&vLink)
		checkErr(err)
		alist = append(alist, vLink)
	}
	return alist
}

func GetLinksCountByUserId(userId int) int {
	rows, err := db.Query("SELECT count(id) FROM links where userid=?", userId)
	checkErr(err)
	if(rows.Next()) {
		var numberOfLinks int
		err = rows.Scan(&numberOfLinks)
		checkErr(err)
		return numberOfLinks
	}
	return 0
}

func generateRandomStr() string {
	pool := "0123456789abcdefghijklmnopqrstuvwxyz"
	length := config.SHORT_LINK_LEN - 1
	b := make([]byte, length)
	rg := rand.New(rand.NewSource(time.Now().Unix()))
	for i, _ := range b {
		b[i] = pool[rg.Intn(len(pool))]
	}
	r := string(b)
	return r
}
