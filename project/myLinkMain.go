package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/auth"
	"net/http"
	"./controllers"
	"./mymodels"
	"github.com/martini-contrib/render"
	"./storage"
)



func main() {
	m := martini.Classic()

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
	m.Post("/api/links", auth.BasicFunc(Auth), controllers.AddNewLink)
	m.Get("/:id", controllers.DoRedirect)
	m.Delete("/api/links/:id", auth.BasicFunc(Auth), controllers.DeleteLink)
	m.Get("/api/links", auth.BasicFunc(Auth), controllers.GetShortLinksByUser)



	m.Run()
}