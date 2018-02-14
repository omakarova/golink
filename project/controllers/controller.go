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

func AddNewUser(r render.Render, params martini.Params, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)
	checkErr(err)

	var nUser mymodels.NewUser
	err = json.Unmarshal(body, &nUser)
	checkErr(err)

	if(storage.DoesUserExist(nUser)) {
		r.JSON(http.StatusNotAcceptable, "Such user has been already created earlier")
		return
	}

	fmt.Println("Hello," + nUser.Username)
	if( storage.NewUserIsValid(nUser)){
		storage.SaveUserData(nUser)
		r.JSON(http.StatusOK, nUser.Username)
	} else {
		r.JSON(http.StatusBadRequest, "too long username or password")
	}
}

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

	if(storage.DoesLinkExist(newLink, id)) {
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
	longurl, err := storage.GetLongUrl(sholturl)
	if(err != nil){
		r.JSON(http.StatusNotFound, "There are no URLs for " + sholturl)
	}
	//r.Header().Add("Location", longurl)
	//r.Status(http.StatusPermanentRedirect)
	r.Redirect(longurl,http.StatusPermanentRedirect)
}

func getCurrentUserId(req *http.Request) int {
	auth := req.Header.Get("Authorization")
	id, err := storage.GetUserIdByAuthString(auth)
    checkErr(err)
    return id
}

func GetCurrentUserInfo(r render.Render, params martini.Params, req *http.Request) {
	auth := req.Header.Get("Authorization")
	r.JSON(http.StatusOK, auth)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
