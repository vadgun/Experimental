package main

import (
	"github.com/kataras/iris/v12"
	asignacioncontroller "github.com/vadgun/Experimental/Controladores/AsignacionController"
	calificacionescontroller "github.com/vadgun/Experimental/Controladores/CalificacionesController"
	indexcontroller "github.com/vadgun/Experimental/Controladores/IndexController"
	inscripcioncontroller "github.com/vadgun/Experimental/Controladores/InscripcionController"
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

	app.Post("/calificaciones", calificacionescontroller.Calificaciones)
	app.Get("/calificaciones", calificacionescontroller.Calificaciones)

	app.Post("/inscripcion", inscripcioncontroller.Inscripcion)
	app.Get("/inscripcion", inscripcioncontroller.Inscripcion)

	app.Post("/asignacion", asignacioncontroller.Asignacion)
	app.Get("/asignacion", asignacioncontroller.Asignacion)

	app.Post("/obtenerMaterias", asignacioncontroller.ObtenerMaterias)

	app.Post("/asignarMateriaADocente", asignacioncontroller.AsignarMaterias)

	app.Post("/alumnos", indexcontroller.Index)
	app.Get("/alumnos", indexcontroller.Index)

	app.Post("/docentes", indexcontroller.Index)
	app.Get("/docentes", indexcontroller.Index)

	app.Post("/directorio", indexcontroller.Index)
	app.Get("/directorio", indexcontroller.Index)

	app.Post("/buscador", indexcontroller.Index)
	app.Get("/buscador", indexcontroller.Index)

	app.Post("/horarios", indexcontroller.Index)
	app.Get("/horarios", indexcontroller.Index)

	app.Post("/kardex", indexcontroller.Index)
	app.Get("/kardex", indexcontroller.Index)

	app.Run(iris.Addr(":8080"))
}
