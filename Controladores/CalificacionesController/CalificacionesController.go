package calificacionescontroller

import (
	"fmt"

	"github.com/kataras/iris/v12"
	sessioncontroller "github.com/vadgun/Experimental/Controladores/SessionController"
	calificacionesmodel "github.com/vadgun/Experimental/Modelos/CalificacionesModel"
	indexmodel "github.com/vadgun/Experimental/Modelos/IndexModel"
	usuariosmodel "github.com/vadgun/Experimental/Modelos/UsuariosModel"
)

//Calificaciones -> Regresa la pagina de inicio
func Calificaciones(ctx iris.Context) {
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

		tienepermiso := indexmodel.TienePermiso(0, userOn, usuario)

		if !tienepermiso {
			ctx.Redirect("/login", iris.StatusSeeOther)
		}

		if err := ctx.View("Calificaciones.html"); err != nil {
			ctx.Application().Logger().Infof(err.Error())
		}
	} else {
		ctx.Redirect("/login", iris.StatusSeeOther)
	}
}

//Alumnos -> Regresa la pagina de inicio
func Alumnos(ctx iris.Context) {
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
	// "PermisoAlumnos" : 8
	// "PermisoDocentes" : 9

	if autorizado || autorizado2 {
		userOn := indexmodel.GetUserOn(sessioncontroller.Sess.Start(ctx).GetString("UserID"))
		ctx.ViewData("Usuario", userOn)

		tienepermiso := indexmodel.TienePermiso(8, userOn, usuario)

		if !tienepermiso {
			ctx.Redirect("/login", iris.StatusSeeOther)
		}

		Semestres := usuariosmodel.ExtraeSemestres()
		ctx.ViewData("Semestres", Semestres)

		if err := ctx.View("Alumnos.html"); err != nil {
			ctx.Application().Logger().Infof(err.Error())
		}
	} else {
		ctx.Redirect("/login", iris.StatusSeeOther)
	}
}

//ObtenerAlumnos -> Los envia a la pagina de regreso
func ObtenerAlumnos(ctx iris.Context) {

	semestre := ctx.PostValue("semestre")

	//Necesito el ID del DOCENTE, puede provenir del mismo ajax, lo evaluo para asignarle la materia correctamente

	//Obtener Materias que cumplan con las condiciones  [ +2012  +Primaria  +1s ]

	var alumnos []calificacionesmodel.Alumno
	// var semestre calificacionesmodel.Semestre

	alumnos = calificacionesmodel.ObtenerAlumnosFiltrados(semestre)

	// alumnos, semestre = calificacionesmodel.ObtenerAlumnosFiltradosYTraerSemestre(semestre)

	var htmlcode string

	if len(alumnos) == 0 {
		ctx.HTML("<script>Swal.fire('Sin Resultados');</script>")

	} else {

		htmlcode += fmt.Sprintf(`
	<br>
	<hr>
	<table class="table table-hover table-bordered table-lg" style="margin: auto; width: 100%s !important; font-size:14px;">
	  <thead>
		<th class="textocentrado" width="30%s">
		  Nombre
		</th>
		<th class="textocentrado">
		 Semestre
		</th>
		<th class="textocentrado">
		 Licenciatura
		</th>
		<th class="textocentrado" width="15%s">
		  Acciones
		</th>
		</thead>
	  <tbody>`, "%%", "%%", "%%")

		for _, v := range alumnos {
			htmlcode += fmt.Sprintf(`
		<tr>
		<td>%v %v %v</td>
		<td>%v</td>
		<td>%v</td>
		
		<td class="textocentrado">
			<a id="myLink1" href="#" onclick="alert('%v');return false;">
				<img src="Recursos/Generales/Plugins/icons/build/svg/link-external-24.svg" height="25" alt="Ver Calificaciones" data-toggle="tooltip" title="Imprimir Calificaciones"/>
			</a>

			<a id="myLink2" href="#" onclick="alert('%v');return false;">
			<img src="Recursos/Generales/Plugins/icons/build/svg/file-badge-16.svg" height="25" alt="Ver Boleta" data-toggle="tooltip" title="Ver boleta"/>
			</a>		

			<a id="myLink3" href="#" onclick="alert('%v');return false;">
			<img src="Recursos/Generales/Plugins/icons/build/svg/diff-renamed-16.svg" height="25" alt="Promover" data-toggle="tooltip" title="Promover de curso"/>
			</a>
		</td>

		</tr>

		`, v.Nombre, v.ApellidoP, v.ApellidoM, v.CursandoSem, v.Licenciatura, v.ID.Hex(), v.ID.Hex(), v.ID.Hex())

		}

		htmlcode += fmt.Sprintf(`
	  </tbody>
	  </table>
	`)

		ctx.HTML(htmlcode)
	}

}

//Docentes -> Regresa la pagina de inicio
func Docentes(ctx iris.Context) {
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
	// "PermisoAlumnos" : 8
	// "PermisoDocentes" : 9

	if autorizado || autorizado2 {
		userOn := indexmodel.GetUserOn(sessioncontroller.Sess.Start(ctx).GetString("UserID"))
		ctx.ViewData("Usuario", userOn)

		tienepermiso := indexmodel.TienePermiso(9, userOn, usuario)

		if !tienepermiso {
			ctx.Redirect("/login", iris.StatusSeeOther)
		}

		docentes := calificacionesmodel.PersonalDocenteActivo()
		ctx.ViewData("Docentes", docentes)

		if err := ctx.View("Docentes.html"); err != nil {
			ctx.Application().Logger().Infof(err.Error())
		}
	} else {
		ctx.Redirect("/login", iris.StatusSeeOther)
	}
}
