package controllers

import (
	//"encoding/json"
	//"fmt"
	"github.com/go-martini/martini"
	//"labix.org/v2/mgo/bson"
	//"loctalk/conf"
	//"loctalk/models"
	//"loctalk/utils"
	"net/http"
)

func GetById(res http.ResponseWriter, req *http.Request, params martini.Params) {
	//id := params["id"]

	//err := post.LoadById(id)
	//if err != nil {
	//	return err.HttpCode, mu.Marshal(err)
	//}
	//return http.StatusOK, mu.Marshal(post)
	res.WriteHeader(200)
	res.Write([]byte("bubybb"))
}
