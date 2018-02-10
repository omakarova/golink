package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/auth"
	"fmt"
	"net/http"
	"./controllers"
	"./mymodels"
	"github.com/martini-contrib/render"
	"encoding/json"
	"io/ioutil"
)



func main() {
	m := martini.Classic()
	fmt.Println("Hello world!")
	m.Use(func(w http.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	})

	m.Use(render.Renderer())
	//m.Use(auth.BasicFunc(func(username, password string) bool {return true}))

	Auth := func(username, password string) bool {return true}
	// ROUTES
	//users
	m.Post("/users", func(r render.Render, params martini.Params, req *http.Request) {

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
		r.JSON(http.StatusOK, nUser.Username)
	})

	m.Get("/user", auth.BasicFunc(Auth), controllers.GetCurrentUserInfo)

	// users
	//m.Get("/api/v1/users", controllers.GetUsers)
	//m.Get("/api/v1/users/:id", controllers.GetById)
	//m.Post("/api/v1/users", controllers.CreateUser)




	m.Run()
}