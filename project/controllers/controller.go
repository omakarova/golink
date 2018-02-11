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
		r.JSON(http.StatusBadRequest, "Such user has been already created earlier")
		return
	}

	fmt.Println("Hello," + nUser.Username)
	if( storage.ValidateNewUser(nUser)){
		storage.SaveUserData(nUser)
		r.JSON(http.StatusOK, nUser.Username)
	} else {
		r.JSON(http.StatusBadRequest, "too long username or password")
	}

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
