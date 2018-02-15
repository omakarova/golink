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
	fmt.Println(res)

	newLinkResponse := mymodels.NewLinkResponse{ShortURL: shortLink}
	return &newLinkResponse
}

func GetLongUrlByShortLink(shortUrl string) (string, error) {
	rows, err := db.Query("SELECT longurl FROM links where shorturl=?", shortUrl)
	checkErr(err)

	if(rows.Next()) {
		var longurl string
		err = rows.Scan(&longurl)
		checkErr(err)
		return longurl, nil
	}
	return "", errors.New("no such link")
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