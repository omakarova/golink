package controllers

import (
	//"encoding/json"
	"fmt"
	"github.com/go-martini/martini"
	"../mymodels"
	"net/http"

	"github.com/martini-contrib/render"
)

func AddNewUser(u1 mymodels.NewUser, res http.ResponseWriter, req *http.Request) {
    fmt.Println("AddNewUser")
	res.WriteHeader(200)
	res.Write([]byte("bubybb" + u1.Username))
}

func GetCurrentUserInfo(r render.Render, params martini.Params, req *http.Request) {
	auth := req.Header.Get("Authorization")
	r.JSON(http.StatusOK, auth)
}
