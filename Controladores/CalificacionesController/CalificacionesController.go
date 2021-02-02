package calificacionescontroller

import (
	"fmt"
	"strconv"
	"strings"

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

		if userOn.Docente || usuario.Docente {
			var materias []calificacionesmodel.Materia
			materias = indexmodel.IfIsDocenteBringMaterias(userOn)
			ctx.ViewData("Materias", materias)
		}

		if userOn.Admin || usuario.Admin {
			// enviar docentes para ejecutar algo similar a lo de arriba, enviar traer materias para ver calificaciones y evaluar

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

		`, v.ApellidoP, v.ApellidoM, v.Nombre, v.CursandoSem, v.Licenciatura, v.ID.Hex(), v.ID.Hex(), v.ID.Hex())

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

//AgregarCalificacion -> Regresa una tabla para capturar la materia con una lista de alumnos inscritos a ese semestre
func AgregarCalificacion(ctx iris.Context) {

	//Tengo el id de materia. tengo el id de semestre y tengo el id de docente

	//Si genero una captura de calificaciones

	data := ctx.PostValue("data")

	cadena := strings.Split(data, ":")

	idmateria := cadena[0]
	idsemestre := cadena[1]

	iddocente := ctx.PostValue("iddocente")

	fmt.Println("id mat ->", idmateria)
	fmt.Println("id sem ->", idsemestre)
	fmt.Println("id doc ->", iddocente)

	alumnos := calificacionesmodel.ObtenerAlumnosFiltrados(idsemestre)
	materia := calificacionesmodel.ExtraeMateria(idmateria)

	var htmlcode string
	var index int

	htmlcode += fmt.Sprintf(`
	<form action="/guardarcalificaciones" method="POST" >
	`)

	htmlcode += fmt.Sprintf(`
	<br>
	<hr>
	<table class="table table-hover table-bordered table-lg" style="margin: auto; width: 70%s !important; font-size:14px;">
	  <thead>
		<th class="textocentrado" width="30%s">
			Alumno
		</th>
		<th class="textocentrado">
			%v
		</th>
		<th class="textocentrado">
		%s de Asistencia
	</th>

		</thead>
	  <tbody>`, "%%", "%%", materia.Materia, "%%")

	for k, v := range alumnos {

		htmlcode += fmt.Sprintf(`
		<tr>
		<td>%v %v %v 
		<input type="hidden" name="idalumno%v" value="%v">
		</td>`, v.ApellidoP, v.ApellidoM, v.Nombre, k, v.ID.Hex())

		for kk, vv := range v.Materias {

			if materia.ID == vv {
				index = kk
			}
		}

		htmlcode += fmt.Sprintf(`
				<td class="text-center"> 
					<input type="number" class="form-control letrasGrandes" name="calificacion%v" value="%v">
					
				</td>
				<td class="text-center"> 
				<input type="number" class="form-control letrasGrandes" name="asistencia%v" value="%v">
				</td>
		</tr>`, k, v.Calificaciones[index], k, v.Asistencias[index])

		fmt.Println(index)
	}

	htmlcode += fmt.Sprintf(`
	</tbody>
	</table>

 `)

	htmlcode += fmt.Sprintf(`
	<input type="hidden" name="materiafilter" value="%v">
	<input type="hidden" name="index" value="%v">
	<input type="hidden" name="numalumnos" value="%v">
	<br>
    <div class="text-center container "> 
 <button type="submit"> Guardar Calificaciones</button>
 </div>

 </form>
 `, materia.ID.Hex(), index, len(alumnos))

	ctx.HTML(htmlcode)

}

//GuardarCalificaciones Guarda la peticion del docente para guardar materias de los alumnos por materia
func GuardarCalificaciones(ctx iris.Context) {

	numalumnos, _ := ctx.PostValueInt("numalumnos")
	index, _ := ctx.PostValueInt("index")
	materiafilter := ctx.PostValue("materiafilter")

	var alumnos []string
	var calificaciones []float64
	var htmlcode string

	fmt.Println("Numero de alumnos", numalumnos, " - ", index, " - ", materiafilter)
	for i := 0; i < numalumnos; i++ {
		istring := strconv.Itoa(i)
		alumnos = append(alumnos, ctx.PostValue("idalumno"+istring))
		flotante, _ := ctx.PostValueFloat64("calificacion" + istring)
		calificaciones = append(calificaciones, flotante)
	}

	fmt.Println(alumnos)
	fmt.Println(calificaciones)

	//guarda esa calificacion en el alumno correspondiente en el index de calificaiones

	actualizado := calificacionesmodel.GuardarCapturaCalificaciones(alumnos, calificaciones, index)

	if actualizado {

		htmlcode += fmt.Sprintf(`<script>
		alert("Calificaciones Guardadas");
		location.replace("/calificaciones");
		</script>`)

	} else {
		htmlcode += fmt.Sprintf(`<script>
		alert("Ocurrio un error");
		location.replace("/calificaciones");
		</script>`)

	}

	ctx.HTML(htmlcode)

}

//CrearFormulario -> Regresa un formulario correspondiente al boton 'Materia' 'Semestre'
func CrearFormulario(ctx iris.Context) {

	data := ctx.PostValue("data")

	var htmlcode string

	switch data {
	case "Materia":

		semestres := usuariosmodel.ExtraeSemestres()

		htmlcode += fmt.Sprintf(`
		<form action="/guardarmateria" method="POST" >

        <div class="col-sm-12">
            <div class="form-group row">
                <label for="materia" class="col-sm-1 col-form-label negrita"> Materia: </label>
                <div class="col-sm-3 col-md-3 col-lg-4">
                    <input type="text" class="form-control" id="materia" name="materia" placeholder="Introduce nombre de la materia" value="" required>
                </div>
                <label for="plan" class="col-sm-1 col-form-label negrita"> Plan: </label>
                <div class="col-sm-2 col-md-2 col-lg-2">
                    <select class="form-control" id="plan" name="plan" value="" required>
                        <option value="">selecciona</option>
                        <option value="2012">2012</option>
                        <option value="2018">2018</option>
                        <option value="2021">2021</option>
                    </select>  
                </div>
                <label for="licenciatura" class="col-sm-2 col-form-label negrita"> Licenciatura: </label>
                <div class="col-sm-2 col-md-2 col-lg-2">
                    <select class="form-control" id="licenciatura" name="licenciatura" value="" required>
                        <option value="">selecciona</option>
                        <option value="Primaria">Primaria</option>
                        <option value="Preescolar">Preescolar</option>
                    </select>  
                </div>
            </div>
            <div class="form-group row">
                <label for="semestreid" class="col-sm-1 col-form-label negrita"> Semestre: </label>
                <div class="col-sm-3 col-md-3 col-lg-3">
                    <select class="form-control" id="semestreid" name="semestreid" value="" required>
						<option value="">selecciona</option>`)

		for _, v := range semestres {
			htmlcode += fmt.Sprintf(`
						<option value="%v">%v - %v - %v - Cuenta con %v materias</option>
					`, v.ID.Hex(), v.Semestre, v.Licenciatura, v.Plan, len(v.Materias))
		}

		htmlcode += fmt.Sprintf(`
                    </select>
                </div>
                <label for="Creditos" class="col-sm-1 col-form-label negrita"> Créditos: </label>
                <div class="col-sm-3 col-md-3 col-lg-3">
                    <input type="number" class="form-control"  step="any" id="creditos" name="creditos" placeholder="Número Créditos" value="" required>
                </div>
                <label for="horas" class="col-sm-1 col-form-label negrita"> Horas: </label>
                <div class="col-sm-3 col-md-3 col-lg-3">
                    <input type="number" class="form-control" step="any" id="horas" name="horas" value="" placeholder="Cantidad de horas" required>
                </div>
            </div>
            <div class="form-group row">

            <div class="text-center container "> 
                <button type="submit"> Guardar Materia</button>
            </div>
            </div>
        </div>
    </form>

		
		`)

		break
	case "Semestre":

		htmlcode += fmt.Sprintf(`

		<form action="/guardarsemestre" method="POST" >

        <div class="col-sm-12">
            <div class="form-group row">
                <label for="materia" class="col-sm-1 col-form-label negrita"> Semestre: </label>
                <div class="col-sm-2 col-md-2 col-lg-2">
                    <select class="form-control" id="semestre" name="semestre" value="" required>
                        <option value="">selecciona</option>
                        <option value="1">1</option>
                        <option value="2">2</option>
                        <option value="3">3</option>
                        <option value="4">4</option>
                        <option value="5">5</option>
                        <option value="6">6</option>
                        <option value="7">7</option>
                        <option value="8">8</option>

                    </select> 

                </div>
                <label for="plan" class="col-sm-1 col-form-label negrita"> Plan: </label>
                <div class="col-sm-2 col-md-2 col-lg-2">
                    <select class="form-control" id="plan" name="plan" value="" required>
                        <option value="">selecciona</option>
                        <option value="2012">2012</option>
                        <option value="2018">2018</option>
                        <option value="2021">2021</option>
                    </select>  
                </div>
                <label for="licenciatura" class="col-sm-2 col-form-label negrita"> Licenciatura: </label>
                <div class="col-sm-2 col-md-2 col-lg-2">
                    <select class="form-control" id="licenciatura" name="licenciatura" value="" required>
                        <option value="">selecciona</option>
                        <option value="Primaria">Primaria</option>
                        <option value="Preescolar">Preescolar</option>
                    </select>  
                </div>
            </div>

            <div class="text-center container "> 
                <button type="submit"> Guardar Semestre</button>
            </div>
            </div>
        </div>
    </form>
		
		`)

		break
	case "Alumnos":

		htmlcode += fmt.Sprintf(`
			carga de alumnos
		`)

		break
	case "Docentes":
		htmlcode += fmt.Sprintf(`
			carga de docentes
		 `)

		break
	}

	ctx.HTML(htmlcode)

}

//GuardarMateria -> Asigna la materia al semestre
func GuardarMateria(ctx iris.Context) {

	var materia calificacionesmodel.Materia
	var semestre calificacionesmodel.Semestre

	materia.Materia = ctx.PostValue("materia")
	materia.Plan = ctx.PostValue("plan")
	materia.Licenciatura = ctx.PostValue("licenciatura")
	materia.Creditos = ctx.PostValue("creditos")
	materia.Horas = ctx.PostValue("horas")

	semestre = calificacionesmodel.TraerSemestre(ctx.PostValue("semestreid"))
	materia.Semestre = semestre.ID

	var htmlcode string

	guardado := calificacionesmodel.AsignarMateriaASemestre(materia, semestre)

	if guardado {

		htmlcode += fmt.Sprintf(`<script>
		alert("Materia asignada al semestre");
		location.replace("/calificaciones");
		</script>`)

	} else {
		htmlcode += fmt.Sprintf(`<script>
		alert("Ocurrio un error");
		location.replace("/calificaciones");
		</script>`)

	}

	ctx.HTML(htmlcode)

}

//GuardarSemestre -> Guarda el semestre donde se asignaran las materias
func GuardarSemestre(ctx iris.Context) {

	var semestre calificacionesmodel.Semestre

	semestre.Licenciatura = ctx.PostValue("licenciatura")
	semestre.Semestre = ctx.PostValue("semestre")
	semestre.Plan = ctx.PostValue("plan")

	var htmlcode string

	guardado := calificacionesmodel.CrearSemestre(semestre)

	if guardado {

		htmlcode += fmt.Sprintf(`<script>
		alert("Semestre Guardado");
		location.replace("/calificaciones");
		</script>`)

	} else {
		htmlcode += fmt.Sprintf(`<script>
		alert("Ocurrio un error");
		location.replace("/calificaciones");
		</script>`)

	}

	ctx.HTML(htmlcode)

}
