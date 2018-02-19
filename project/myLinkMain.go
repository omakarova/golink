package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/auth"
	"net/http"
	"./controllers"
	"./mymodels"
	"github.com/martini-contrib/render"
	"./storage"
	"github.com/BurntSushi/toml"
	"./config"
)



func main() {
	m := martini.Classic()

	if _, err := toml.DecodeFile("config.toml", &config.Config); err != nil {
		config.Config = config.ConfigT{config.DB_CONNECTION_STRING_CONST, config.SHORT_LINK_LEN_CONST}
	}

	storage.InitDB (config.Config.DB_CONNECTION_STRING)

	m.Use(func(w http.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	})

	m.Use(render.Renderer())
	//m.Use(auth.BasicFunc(func(username, password string) bool {return true}))

	Auth := func(username, password string) bool {
		user := mymodels.NewUser{Username: username, Password: password}
		return storage.DoesUserExist(user)
	}

	// ROUTES

	//users
	m.Post("/api/users", controllers.AddNewUser)
	m.Get("/api/user", auth.BasicFunc(Auth), controllers.GetCurrentUserInfo)

	//links
	m.Get("/:id", controllers.DoRedirect)
	m.Post("/api/links", auth.BasicFunc(Auth), controllers.AddNewLink)
	m.Delete("/api/links/:id", auth.BasicFunc(Auth), controllers.DeleteLink)
	m.Get("/api/links", auth.BasicFunc(Auth), controllers.GetShortLinksByUser)
	m.Get("/api/links/:id", auth.BasicFunc(Auth), controllers.GetLinkInfoByUser)

	//statistics
	m.Get("/api/stat/topref", auth.BasicFunc(Auth), controllers.GetTopReferrersByUser)
    m.Get("/api/stat/interval/:id", auth.BasicFunc(Auth), controllers.GetLinksStatByUser)

	m.Run()
}