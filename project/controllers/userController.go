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

func GetCurrentUserInfo(r render.Render, params martini.Params, req *http.Request) {
	auth := req.Header.Get("Authorization")
	id, username, err := storage.GetUserDataByAuthString(auth)
	checkErr(err)
	linkscount := storage.GetLinksCountByUserId(id)
	userInfo := mymodels.UserInfo{Username: username, LinksCount: linkscount}
	r.JSON(http.StatusOK, userInfo)
}

func getCurrentUserId(req *http.Request) int {
	auth := req.Header.Get("Authorization")
	id, _, err := storage.GetUserDataByAuthString(auth)
	checkErr(err)
	return id
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
