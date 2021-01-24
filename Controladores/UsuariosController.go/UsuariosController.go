package usuarioscontroller

import (
	"fmt"

	"github.com/kataras/iris/v12"
	sessioncontroller "github.com/vadgun/Experimental/Controladores/SessionController"
	indexmodel "github.com/vadgun/Experimental/Modelos/IndexModel"
)

//Usuarios -> Regresa la pagina de inicio
func Usuarios(ctx iris.Context) {
	var usuario indexmodel.MongoUser
	var autorizado bool
	autorizado2, _ := sessioncontroller.Sess.Start(ctx).GetBoolean("Autorizado")

	if autorizado2 == false {
		usuario.Key = ctx.PostValue("pass")
		usuario.Usuario = ctx.PostValue("usuario")
		autorizado, usuario = indexmodel.VerificarUsuario(usuario)
		if autorizado {
			sessioncontroller.Sess.Start(ctx).Set("Autorizado", true)
			sessioncontroller.Sess.Start(ctx).Set("UserID", usuario.ID.Hex())
		}
	}

	if autorizado || autorizado2 {
		userOn := indexmodel.GetUserOn(sessioncontroller.Sess.Start(ctx).GetString("UserID"))
		ctx.ViewData("Usuario", userOn)

		if err := ctx.View("Usuarios.html"); err != nil {
			ctx.Application().Logger().Infof(err.Error())
		}
	} else {
		ctx.Redirect("/login", iris.StatusSeeOther)
	}
}

//SolicitarUsuario Solicita el formulario para el usuario a crear
func SolicitarUsuario(ctx iris.Context) {

	data := ctx.PostValue("data")

	var htmlcode string

	switch data {
	case "Alumno":

		htmlcode += fmt.Sprintf(`<h1>Creando formulario para Alumno</h1>`)

		break
	case "Docente":
		htmlcode += fmt.Sprintf(`<h1>Creando formulario para Docente</h1>`)
		break
	case "Administrativo":
		htmlcode += fmt.Sprintf(`<h1>Creando formulario para Administrativo</h1>`)
		break
	case "Director":
		htmlcode += fmt.Sprintf(`<h1>Creando formulario para Director</h1>`)
		break
	case "Subdirector":
		htmlcode += fmt.Sprintf(`<h1>Creando formulario para SubDirector</h1>`)
		break
	}

	ctx.HTML(htmlcode)

}
