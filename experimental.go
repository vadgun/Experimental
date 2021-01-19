package main

import (
	"github.com/kataras/iris/v12"
	indexcontroller "github.com/vadgun/Experimental/Controladores/IndexController"
	logincontroller "github.com/vadgun/Experimental/Controladores/LoginController"
)

func main() {
	app := iris.New()
	app.HandleDir("/Recursos", "./Recursos")
	app.RegisterView(iris.HTML("./Vistas", ".html").Reload(true))
	app.Get("/", logincontroller.Getlogin)
	app.Get("/login", logincontroller.Getlogin)
	app.Post("/login", logincontroller.Getlogin)
	app.Get("/logout", logincontroller.Getlogout)

	app.Get("/index", indexcontroller.Index)
	app.Post("/index", indexcontroller.Index)
	app.Get("/perfil", indexcontroller.Index)

	app.Run(iris.Addr(":8080"))
}
