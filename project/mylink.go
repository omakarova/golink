package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/auth"
	"fmt"
	"net/http"
	"./controllers"
)

func main() {
	m := martini.Classic()
	fmt.Println("Hello world!")
	m.Use(func(w http.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	})

	m.Use(auth.BasicFunc(func(username, password string) bool {return true}))
	// ROUTES
	//m.Get("/", controllers.Home)

	// users
	//m.Get("/api/v1/users", controllers.GetUsers)
	m.Get("/api/v1/users/:id", controllers.GetById)
	//m.Post("/api/v1/users", controllers.CreateUser)




	m.Run()
}