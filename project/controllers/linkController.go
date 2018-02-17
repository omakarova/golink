package controllers

import (
	//"encoding/json"
	"fmt"
	"github.com/go-martini/martini"
	"../mymodels"
	"net/http"
    "../storage"
	"github.com/martini-contrib/render"
	"io/ioutil"
	"encoding/json"

)

func AddNewLink(r render.Render, params martini.Params, req *http.Request) {
	id := getCurrentUserId(req)
	body, err := ioutil.ReadAll(req.Body)
	checkErr(err)

	var newLink mymodels.NewLink
	err = json.Unmarshal(body, &newLink)
	checkErr(err)

	if( !storage.NewLinkIsValid(newLink)){
		r.JSON(http.StatusBadRequest, "wrong url")
	}

	if(storage.DoesLongLinkExist(newLink, id)) {
		r.JSON(http.StatusNotAcceptable, "Such link has been already created by current user")
		return
	}

	newShortLink := storage.CreateNewLink(newLink, id)

	r.JSON(http.StatusOK, newShortLink)

}

func DoRedirect(r render.Render, params martini.Params, req *http.Request) {
	fmt.Println("DoRedirect")
	sholturl := params["id"]
	fmt.Println(sholturl)
	if(len(sholturl) < 1) {
		r.JSON(http.StatusNotFound, "")
	}
	longurl, linkid, err := storage.GetLongUrlAndIdByShortLink(sholturl)
	if(err != nil){
		r.JSON(http.StatusNotFound, "There are no URLs for " + sholturl)
	}
	storage.AddStat(req, linkid)
	//r.Header().Add("Location", longurl)
	//r.Status(http.StatusPermanentRedirect)
	r.Header().Add("Cache-control", "no-cache")
	r.Redirect(longurl,http.StatusPermanentRedirect)
}

func DeleteLink(r render.Render, params martini.Params, req *http.Request) {
	sholturl := params["id"]
	currentUserId := getCurrentUserId(req)
	if( !storage.DoesShortLinkExist(sholturl, currentUserId)){
		r.JSON(http.StatusNotFound, "wrong short url")
	}

	storage.DeleteLink(sholturl, currentUserId)
	r.JSON(http.StatusOK, sholturl)
}

func GetShortLinksByUser(r render.Render, params martini.Params, req *http.Request) {
	currentUserId := getCurrentUserId(req)
	links := storage.GetAllLinksByUserId(currentUserId)
	if(links == nil || len(links) < 0){
		r.JSON(http.StatusOK, "{}")
		return
	}

	//к сожалению, json не хочет нормально маршалить список из структур :(
	//поэтому обойдемся слайсом из строк
	r.JSON(http.StatusOK, links)

}

func GetLinkInfoByUser(r render.Render, params martini.Params, req *http.Request) {
	currentUserId := getCurrentUserId(req)
	sholturl := params["id"]
	linkId, longurl, err := storage.GetLinkIdByShortLinkAndUserId(sholturl, currentUserId)
	if (err != nil){
		r.JSON(http.StatusNotFound, "wrong short url")
	}
	number, err := storage.GetNumberOfClicks(linkId)
	linkInfo := mymodels.LinkInfo{ShortURL: sholturl, LongURL: longurl, NumberOfClicks: 0}
	if (err == nil){
		linkInfo.NumberOfClicks = number
	}
	r.JSON(http.StatusOK, linkInfo)
}

func GetTopReferrersByUser(r render.Render, params martini.Params, req *http.Request) {
	currentUserId := getCurrentUserId(req)

	referers := storage.GetTopReferrersByUser(currentUserId)

	r.JSON(http.StatusOK, referers)
}