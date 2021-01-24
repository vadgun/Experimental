package asignacioncontroller

import (
	"fmt"

	"github.com/kataras/iris/v12"
	sessioncontroller "github.com/vadgun/Experimental/Controladores/SessionController"
	calificacionesmodel "github.com/vadgun/Experimental/Modelos/CalificacionesModel"
	indexmodel "github.com/vadgun/Experimental/Modelos/IndexModel"
)

//Asignacion -> Regresa la pagina de inicio
func Asignacion(ctx iris.Context) {
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

		docentes := calificacionesmodel.PersonalDocenteActivo()
		ctx.ViewData("Docentes", docentes)

		if err := ctx.View("Asignacion.html"); err != nil {
			ctx.Application().Logger().Infof(err.Error())
		}
	} else {
		ctx.Redirect("/login", iris.StatusSeeOther)
	}
}

//ObtenerMaterias Devuelve las materias dependiendo de la consulta
func ObtenerMaterias(ctx iris.Context) {

	licenciatura := ctx.PostValue("licenciatura")
	semestre := ctx.PostValue("semestre")
	plan := ctx.PostValue("plan")

	fmt.Println("Lic ->", licenciatura)
	fmt.Println("Sem ->", semestre)
	fmt.Println("Plan ->", plan)

	//Necesito el ID del DOCENTE, puede provenir del mismo ajax, lo evaluo para asignarle la materia correctamente

	var htmlcode string

	htmlcode += fmt.Sprintf(`
	<br>
	<hr>
	<table class="table table-hover table-bordered table-lg" style="margin: auto; width: 94%s !important; font-size:14px;">
	  <thead>
		<th class="textocentrado">
		  Materia
		</th>
		<th class="textocentrado">
		 Horas
		</th>
		<th class="textocentrado">
		 Créditos
		</th>
		<th class="textocentrado">
		Asignada a:
		  </th>	
		<th class="textocentrado">
		  Acciones
		</th>
		</thead>
	  <tbody>`, "%%")

	//Obtener Materias que cumplan con las condiciones  [ +2012  +Primaria  +1s ]

	var materias []calificacionesmodel.Materia

	materias = calificacionesmodel.ObtenerMateriasFiltradas(plan, licenciatura, semestre)

	for k, v := range materias {

		fmt.Println("k ->", k)
		fmt.Println("v ->", v)
		htmlcode += fmt.Sprintf(`
		<tr>
		<td>%v</td>
		<td>%v</td>
		<td>%v</td>
		<td>%v</td>
		<td>
			<a id="myLink" href="#" onclick="AsignarMateria('%v');return false;">
				<img src="Recursos/Generales/Plugins/icons/build/svg/plus-circle-16.svg" height="25" alt="Asignar Materia"/>
			</a>

			<a id="myLink" href="#" onclick="AsignarMateria('%v');return false;">
			<img src="Recursos/Generales/Plugins/icons/build/svg/no-entry-16.svg" height="25" alt="Revocar Materia"/>
			</a>		
		
		</td>

		</tr>

		`, v.Materia, v.Horas, v.Creditos, v.Asignada, v.ID.Hex(), v.ID.Hex())

	}

	htmlcode += fmt.Sprintf(`
	  </tbody>
	  </table>
	`)

	ctx.HTML(htmlcode)

}

//AsignarMaterias Asigna la materia seleccionada y si ya la tiene responde que ha sido seleccionada
func AsignarMaterias(ctx iris.Context) {

	data := ctx.PostValue("data")
	iddocente := ctx.PostValue("iddocente")

	fmt.Println("Data ->", data)
	fmt.Println("Data ->", iddocente)

	//Necesito el ID del DOCENTE y el ID de la MATERIA
	//ya tengo el ID de MATERIA

	ctx.HTML("<script>alert('Hola');</script>")

}
