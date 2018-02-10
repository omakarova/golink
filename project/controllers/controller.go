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
	if err != nil {
		panic(err)
	}
	var nUser mymodels.NewUser
	err = json.Unmarshal(body, &nUser)
	if err != nil {
		panic(err)
	}
	fmt.Println("Hello world!" + nUser.Username)
	storage.SaveUserdata(nUser)
	r.JSON(http.StatusOK, nUser.Username)
}

func GetCurrentUserInfo(r render.Render, params martini.Params, req *http.Request) {
	auth := req.Header.Get("Authorization")
	r.JSON(http.StatusOK, auth)
}
