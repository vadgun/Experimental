package inscripcioncontroller

import (
	"github.com/kataras/iris/v12"
	sessioncontroller "github.com/vadgun/Experimental/Controladores/SessionController"
	indexmodel "github.com/vadgun/Experimental/Modelos/IndexModel"
)

//Inscripcion -> Regresa la pagina de inicio
func Inscripcion(ctx iris.Context) {
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

	// "PermisoCalificaciones" : 0,
	// "PermisoUsuarios" : 1,
	// "PermisoAsignar" : 2,
	// "PermisoInscripcion" : 3,
	// "PermisoHorarios" : 4,
	// "PermisoDirectorio" : 5,
	// "PermisoKardex" : 6,
	// "PermisoIndex" : 7

	if autorizado || autorizado2 {
		userOn := indexmodel.GetUserOn(sessioncontroller.Sess.Start(ctx).GetString("UserID"))
		ctx.ViewData("Usuario", userOn)

		tienepermiso := indexmodel.TienePermiso(3, userOn, usuario)

		if !tienepermiso {
			ctx.Redirect("/login", iris.StatusSeeOther)
		}

		if err := ctx.View("Inscripcion.html"); err != nil {
			ctx.Application().Logger().Infof(err.Error())
		}
	} else {
		ctx.Redirect("/login", iris.StatusSeeOther)
	}
}

//GuardarInscripcion -> Guarda la inscripcion
func GuardarInscripcion(ctx iris.Context) {

	ctx.Redirect("/login", iris.StatusSeeOther)

}
