package calificacionescontroller

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/jung-kurt/gofpdf"

	"github.com/kataras/iris/v12"
	sessioncontroller "github.com/vadgun/Experimental/Controladores/SessionController"
	calificacionesmodel "github.com/vadgun/Experimental/Modelos/CalificacionesModel"
	indexmodel "github.com/vadgun/Experimental/Modelos/IndexModel"
	usuariosmodel "github.com/vadgun/Experimental/Modelos/UsuariosModel"
)

//Profile Edita el perfil del usuario logeado con la opcion de subir una imagen que puede ser usada con posterioridad
func Profile(ctx iris.Context) {

	userOn := indexmodel.GetUserOn(sessioncontroller.Sess.Start(ctx).GetString("UserID"))
	ctx.ViewData("Usuario", userOn)

	if userOn.Alumno {
		var alumno calificacionesmodel.Alumno
		alumno = calificacionesmodel.ExtraeAlumno(userOn.UserID.Hex())
		ctx.ViewData("Alumno", alumno)

	}

	if err := ctx.View("Profile.html"); err != nil {
		ctx.Application().Logger().Infof(err.Error())
	}

}

//ActualizaUsuario -> Actualiza el usuario y agrega una imagen a su perfil
func ActualizaUsuario(ctx iris.Context) {

	userid := ctx.PostValue("userid") //Usuario Alumno, Docente o Administrativo
	correo := ctx.PostValue("correousuario")
	telefono := ctx.PostValue("telefonousuario")
	var htmlcode string
	var dirPath string
	userOn := indexmodel.GetUserOn(sessioncontroller.Sess.Start(ctx).GetString("UserID")) // Usuario Logeado en el sistema MongoUser
	imagen1, _, err := ctx.FormFile("imagenusuario")
	check(err, "Error al seleccionar la imagen 1")
	if userOn.Admin {
		// dirPath = "./Recursos/Imagenes/Usuarios/Admin"
	}

	if userOn.Alumno {
		alumno := calificacionesmodel.ExtraeAlumno(userid)
		dirPath = "./Recursos/Imagenes/Usuarios/Alumnos"
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			fmt.Println("el directorio no existe")
			os.MkdirAll(dirPath, 0777)
		} else {
			fmt.Println("el directorio ya existe")
		}
		out, err := os.Create(alumno.Matricula)
		check(err, "No se puede crear el archivo revisa los privilegios de escritura.")
		defer out.Close()
		_, err = io.Copy(out, imagen1)
		check(err, "Error al escribir la imagen al directorio 1")
		alumno.Imagen = dirPath + "/" + alumno.Matricula
		alumno.CorreoE = correo
		alumno.Telefono = telefono
		calificacionesmodel.ActualizaAlumno(alumno)

	}

	if userOn.Administrativo {

	}

	if userOn.Docente {

	}

	if userOn.Director {

	}

	htmlcode += fmt.Sprintf(`<script>
	alert("Perfil Guardado");
		location.replace("/perfil");
	</script>`)

	ctx.HTML(htmlcode)

}

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
		var materias []calificacionesmodel.Materia

		tienepermiso := indexmodel.TienePermiso(0, userOn, usuario)

		if !tienepermiso {
			ctx.Redirect("/login", iris.StatusSeeOther)
		}

		if userOn.Docente || usuario.Docente {
			materias = indexmodel.IfIsDocenteBringMaterias(userOn)
			ctx.ViewData("Materias", materias)
		}

		if userOn.Admin || usuario.Admin {
			// enviar docentes para ejecutar algo similar a lo de arriba, enviar traer materias para ver calificaciones y evaluar

		}

		if userOn.Alumno || usuario.Alumno {

			var alumno calificacionesmodel.Alumno
			var materias []calificacionesmodel.Materia
			var semestre calificacionesmodel.Semestre
			var docentes []calificacionesmodel.Docente
			var nombresdocentes []string

			alumno = calificacionesmodel.ExtraeAlumno(userOn.UserID.Hex())
			materias = calificacionesmodel.ExtraeMateriasPorSemestre(alumno.CursandoSem)
			semestre = calificacionesmodel.ExtraeSemestreString(alumno.CursandoSem.Hex())
			docentes = calificacionesmodel.ExtraeDocentes(materias)

			for _, vd := range docentes {

				nombre := vd.Nombre + " " + vd.ApellidoP + " " + vd.ApellidoM
				nombresdocentes = append(nombresdocentes, nombre)
			}

			ctx.ViewData("Alumno", alumno)
			ctx.ViewData("Materias", materias)
			ctx.ViewData("Semestre", semestre)
			ctx.ViewData("Docentes", nombresdocentes)

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

		if userOn.Admin || usuario.Admin || userOn.Administrativo || usuario.Administrativo {
			Semestres := usuariosmodel.ExtraeSemestres()
			ctx.ViewData("Semestres", Semestres)
		}

		if err := ctx.View("Alumnos.html"); err != nil {
			ctx.Application().Logger().Infof(err.Error())
		}
	} else {
		ctx.Redirect("/login", iris.StatusSeeOther)
	}
}

//ObtenerAlumnosCalif -> Los envia a la pagina de regreso con calificacion
func ObtenerAlumnosCalif(ctx iris.Context) {

	semestreidstring := ctx.PostValue("semestre")

	//Necesito el ID del DOCENTE, puede provenir del mismo ajax, lo evaluo para asignarle la materia correctamente

	//Obtener Materias que cumplan con las condiciones  [ +2012  +Primaria  +1s ]
	var reprobado float64

	reprobado = 5.0

	var alumnos []calificacionesmodel.Alumno

	alumnos = calificacionesmodel.ObtenerAlumnosFiltrados(semestreidstring)

	var semestre calificacionesmodel.Semestre
	semestre = calificacionesmodel.ExtraeSemestreString(semestreidstring)

	var materias []calificacionesmodel.Materia
	materias = calificacionesmodel.ExtraeMateriasPorSemestre(semestre.ID)

	var htmlcode string

	if len(alumnos) == 0 {
		ctx.HTML("<script>Swal.fire('Sin Resultados');</script>")

	} else {

		htmlcode += fmt.Sprintf(`
	<br>
	<hr>
	<table class="table table-hover table-bordered table-lg" style="margin: auto; width: 100%s !important; font-size:14px;">
	  <thead>
	  <th>
		#
	  </th>
		<th class="textocentrado" width="30%s">
		  Nombre
		</th>`, "%%", "%%")

		for _, vm := range materias {

			htmlcode += fmt.Sprintf(`
		   <th class="textocentrado">
		   		%v
		  	</th>`, vm.Materia)
		}

		htmlcode += fmt.Sprintf(`
		</thead>
	  <tbody>
	`)

		// for _, vm := range semestre.Materias {

		// 	htmlcode += fmt.Sprintf(`
		//    <th class="textocentrado">
		// 		   %v
		// 	  </th>`, alumnos)
		// }

		for ka, v := range alumnos {
			htmlcode += fmt.Sprintf(`
		<tr>
		<td>%v</td>
		<td>%v %v %v</td>
		`, ka+1, v.ApellidoP, v.ApellidoM, v.Nombre)

			for i := 0; i < len(materias); i++ {

				if v.Calificaciones[i] <= reprobado {
					htmlcode += fmt.Sprintf(`
					<td class="reprobado">%v</td>
					`, v.Calificaciones[i])

				} else {
					htmlcode += fmt.Sprintf(`
				<td class="noreprobado">%v</td>
				`, v.Calificaciones[i])
				}

			}

			htmlcode += fmt.Sprintf(`
		</tr>`)

		}

		htmlcode += fmt.Sprintf(`
	  </tbody>
	  </table>
	`)

		ctx.HTML(htmlcode)
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
	  <th>
	  #
	  </th>
		<th class="textocentrado" width="30%s">
		  Nombre
		</th>
		<th class="textocentrado">
		 Usuario
		</th>
		<th class="textocentrado">
		Password
	   </th>
		<th class="textocentrado">
		 Licenciatura
		</th>
		<th class="textocentrado" width="15%s">
		  Acciones
		</th>
		</thead>
	  <tbody>`, "%%", "%%", "%%")

		for k, v := range alumnos {
			htmlcode += fmt.Sprintf(`
		<tr>
		<td>%v</td>
		<td>%v %v %v</td>
		<td>%v</td>
		<td>%v</td>
		<td>%v</td>
		
		<td class="textocentrado">
			<a id="myLink1" href="#" onclick="alert('%v');return false;">
				<img src="Recursos/Generales/Plugins/icons/build/svg/link-external-24.svg" height="25" alt="Ver Calificaciones" data-toggle="tooltip" title="Imprimir Calificaciones"/>
			</a>

			<a id="myLink2" href="#" onclick="GenerarBoleta('%v');return false;">
			<img src="Recursos/Generales/Plugins/icons/build/svg/file-badge-16.svg" height="25" alt="Ver Boleta" data-toggle="tooltip" title="Ver boleta"/>
			</a>		

			<a id="myLink3" href="#" onclick="alert('%v');return false;">
			<img src="Recursos/Generales/Plugins/icons/build/svg/diff-renamed-16.svg" height="25" alt="Promover" data-toggle="tooltip" title="Editar Información"/>
			</a>
		</td>

		</tr>

		`, k+1, v.ApellidoP, v.ApellidoM, v.Nombre, v.SiguienteSem, v.AnteriorSem, v.Licenciatura, v.ID.Hex(), v.ID.Hex(), v.ID.Hex())

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

		// materias := Extrae

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

	var data string
	// var iddocente string
	data = ctx.PostValue("data")
	// iddocente = ctx.PostValue("iddocente")

	var cadena []string

	cadena = strings.Split(data, ":")

	idmateria := cadena[0]
	idsemestre := cadena[1]

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

	var alumnos []string
	var calificaciones []float64
	var asistencias []float64
	var htmlcode string

	for i := 0; i < numalumnos; i++ {
		istring := strconv.Itoa(i)
		alumnos = append(alumnos, ctx.PostValue("idalumno"+istring))
		flotante, _ := ctx.PostValueFloat64("calificacion" + istring)
		asistenciafloat, _ := ctx.PostValueFloat64("asistencia" + istring)
		calificaciones = append(calificaciones, flotante)
		asistencias = append(asistencias, asistenciafloat)
	}

	//guarda esa calificacion en el alumno correspondiente en el index de calificaiones

	actualizado := calificacionesmodel.GuardarCapturaCalificaciones(alumnos, calificaciones, asistencias, index)

	if actualizado {

		htmlcode += fmt.Sprintf(`<script>
		alert("Calificaciones guardadas");
			location.replace("/calificaciones");
		</script>`)

	} else {
		htmlcode += fmt.Sprintf(`<script>
		alert("Algo salio mal");
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
	case "Docentes":

		htmlcode += fmt.Sprintf(`
			
		`)

		break
	case "Alumnos":
		htmlcode += fmt.Sprintf(`
		<form method="POST" enctype="multipart/form-data" action="/cargarmasivoalumnos" name="cargarmasivoalumnos" id="cargarmasivoalumnos">
		<div class="col-12">
			<h6 class="border-bottoms">Archivo de carga de Alumnos:</h6>
			
			<div class="form-group row">
			
				<label for="archivoalumnos" class="col-sm-3 col-form-label negrita"> Selecciona archivo : </label>
				<div class="col-sm-6">
					<input type="file" class="form-control" id="archivoalumnos" name="archivoalumnos" required>
				</div>			

			
				<div class="col-sm-3">
					<button type="submit" class="btn btn-primary">Cargar Alumnos</button>
				</div>
	
			</div>
		</div>
		</form>
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

func check(err error, mensaje string) {
	if err != nil {
		fmt.Println(err)
	}
}

//CrearUsuario -> Crea el usuario de sistema
func CrearUsuario(Plan, Nombre, semestrenum string) string {
	var user string

	nombres := strings.Split(Nombre, " ")

	user = Plan + nombres[0] + semestrenum

	return user
}

//CrearPassword -> Crea el password del sistema
func CrearPassword(cadena string) string {
	var pass string

	pass = cadena

	return pass

}

//CargarMasivoAlumnos -> Sube el archivo y lo interpreta para su conversion a la base de datos asi como la creacion de usuarios
func CargarMasivoAlumnos(ctx iris.Context) {

	var alumnos []calificacionesmodel.Alumno
	var usuarios []indexmodel.MongoUser

	layout := "2006-01-02"
	location, _ := time.LoadLocation("America/Mexico_City")

	archivo, header, err := ctx.FormFile("archivoalumnos")
	if err != nil {
		fmt.Println(err)
	}
	nombrearchivo := header.Filename

	dirpath := "./Recursos/Archivos"
	if _, err := os.Stat(dirpath); os.IsNotExist(err) {
		fmt.Println("el directorio no existe")
		os.MkdirAll(dirpath, 0777)
	} else {
		fmt.Println("el directorio ya existe")
	}
	out, err := os.Create("./Recursos/Archivos/" + nombrearchivo)
	check(err, "Unable to create the file for writing. Check your write access privilege")
	defer out.Close()
	_, err = io.Copy(out, archivo)
	check(err, "Error al escribir el archivo al directorio")

	excelFileName := dirpath + "/" + nombrearchivo

	xlFile, err := excelize.OpenFile(excelFileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	semestres := usuariosmodel.ExtraeSemestres()

	// Get all the rows in the Sheet1.
	rows, err := xlFile.GetRows("1RO PREESOLAR")
	for ks, row := range rows {
		if ks <= 46 {
			var alumno calificacionesmodel.Alumno
			var mongouser indexmodel.MongoUser
			var newwdate string
			var anio string
			var mes string
			var dia string
			var usersystem string
			var passsystem string
			var semestrenum string

			for kk, colCell := range row {

				switch kk {
				case 0:
					alumno.CorreoE = colCell
					break
				case 1:
					alumno.Matricula = colCell
					break
				case 2:
					alumno.ApellidoP = strings.ToUpper(colCell)
					break
				case 3:
					alumno.ApellidoM = strings.ToUpper(colCell)
					break
				case 4:
					alumno.Nombre = strings.ToUpper(colCell)
					break
				case 5:
					alumno.Sexo = colCell
					break
				case 6:
					alumno.Curp = strings.ToUpper(colCell)
					break
				case 7:

					if colCell != "" {
						anio = colCell[6:]
						mes = colCell[3:5]
						dia = colCell[0:2]
						newwdate = anio + "-" + mes + "-" + dia
						fechanacparsed, _ := time.ParseInLocation(layout, newwdate, location)
						alumno.FechaNac = fechanacparsed
					}
					break
				case 8:
					alumno.Telefono = colCell
					break
				case 9:
					alumno.Plan = colCell
					break
				case 10:
					if colCell == "Educación Preescolar" {
						alumno.Licenciatura = "Preescolar"
					} else if colCell == "Educación Primaria" {
						alumno.Licenciatura = "Primaria"
					}
					break
				case 11:

					//Ids de semestres a partir de  Plan(2012) Semestre(1) Licenciatura(Primaria)
					//Es donde se inscribira al alumno y obtendra sus materias para ser evaluado

					switch colCell {
					case "Primero":

						for _, semestre := range semestres {

							if semestre.Semestre == "1" && semestre.Licenciatura == alumno.Licenciatura && semestre.Plan == alumno.Plan {
								alumno.CursandoSem = semestre.ID
								alumno.Materias = semestre.Materias
								semestrenum = semestre.Semestre
								alumno.Semestre = semestre.Semestre

							}
						}

						break
					case "Segundo":
						for _, semestre := range semestres {

							if semestre.Semestre == "2" && semestre.Licenciatura == alumno.Licenciatura && semestre.Plan == alumno.Plan {
								alumno.CursandoSem = semestre.ID
								alumno.Materias = semestre.Materias
								semestrenum = semestre.Semestre
								alumno.Semestre = semestre.Semestre
							}
						}
						break
					case "Tercero":
						for _, semestre := range semestres {

							if semestre.Semestre == "3" && semestre.Licenciatura == alumno.Licenciatura && semestre.Plan == alumno.Plan {
								alumno.CursandoSem = semestre.ID
								alumno.Materias = semestre.Materias
								semestrenum = semestre.Semestre
								alumno.Semestre = semestre.Semestre
							}
						}
						break
					case "Cuarto":
						for _, semestre := range semestres {

							if semestre.Semestre == "4" && semestre.Licenciatura == alumno.Licenciatura && semestre.Plan == alumno.Plan {
								alumno.CursandoSem = semestre.ID
								alumno.Materias = semestre.Materias
								semestrenum = semestre.Semestre
								alumno.Semestre = semestre.Semestre
							}
						}
						break
					case "Quinto":
						for _, semestre := range semestres {

							if semestre.Semestre == "5" && semestre.Licenciatura == alumno.Licenciatura && semestre.Plan == alumno.Plan {
								alumno.CursandoSem = semestre.ID
								alumno.Materias = semestre.Materias
								semestrenum = semestre.Semestre
								alumno.Semestre = semestre.Semestre
							}
						}
						break
					case "Sexto":
						for _, semestre := range semestres {

							if semestre.Semestre == "6" && semestre.Licenciatura == alumno.Licenciatura && semestre.Plan == alumno.Plan {
								alumno.CursandoSem = semestre.ID
								alumno.Materias = semestre.Materias
								semestrenum = semestre.Semestre
								alumno.Semestre = semestre.Semestre
							}
						}
						break
					case "Séptimo":
						for _, semestre := range semestres {

							if semestre.Semestre == "7" && semestre.Licenciatura == alumno.Licenciatura && semestre.Plan == alumno.Plan {
								alumno.CursandoSem = semestre.ID
								alumno.Materias = semestre.Materias
								semestrenum = semestre.Semestre
								alumno.Semestre = semestre.Semestre
							}
						}
						break
					case "Octavo":
						for _, semestre := range semestres {

							if semestre.Semestre == "8" && semestre.Licenciatura == alumno.Licenciatura && semestre.Plan == alumno.Plan {
								alumno.CursandoSem = semestre.ID
								alumno.Materias = semestre.Materias
								semestrenum = semestre.Semestre
								alumno.Semestre = semestre.Semestre
							}
						}
						break

					}

					break
				case 12:
					alumno.Nss = colCell
					break
				case 13:
					alumno.TipoSangre = strings.ToUpper(colCell)
					break
				case 14:
					alumno.Tutor = strings.ToUpper(colCell)
					break
				case 15:
					alumno.Telefono = colCell
					break
				case 16:
					alumno.OcupacionTutor = strings.ToUpper(colCell)
					break
				case 17:
					alumno.ParentescoTutor = strings.ToUpper(colCell)
					break
				case 18:
					alumno.ContactoCasoEmergencia = strings.ToUpper(colCell)
					break
				case 19:
					alumno.DiferenteDomicilioTutor = strings.ToUpper(colCell)
					break
				case 20:
					alumno.Calle = strings.ToUpper(colCell)
					break
				case 21:
					alumno.Numero = colCell
					break
				case 22:
					alumno.ColAsentamiento = colCell
					break
				case 23:
					alumno.Estado = "Chiapas"
					break
				case 24:
					alumno.ReferenciasDomicilio = colCell
					break
				}

			}

			alumno.IsSystemUser = true
			usersystem = CrearUsuario(alumno.Plan, alumno.Nombre, semestrenum) //2018Magaly7
			if alumno.Curp != "" {
				passsystem = CrearPassword("probando") //CACR8612
			} else {
				passsystem = "1xk7f"
			}

			alumno.Horario = ""
			mongouser.Nombre = alumno.Nombre
			mongouser.Apellidos = alumno.ApellidoP + " " + alumno.ApellidoM
			//mongouser.Edad = CalcularEdad(alumno.FechaNac)int
			mongouser.Usuario = usersystem
			mongouser.Key = passsystem
			alumno.SiguienteSem = usersystem
			alumno.AnteriorSem = passsystem
			mongouser.Puesto = "Alumno de la Licenciatura en " + alumno.Licenciatura
			mongouser.Nombre2 = "Alumno"
			//mongouser.UserID= variable para bson package
			mongouser.Alumno = true
			mongouser.Docente = false
			mongouser.Administrativo = false
			mongouser.Director = false
			mongouser.Admin = false

			for i := 0; i < len(alumno.Materias); i++ {
				alumno.Calificaciones = append(alumno.Calificaciones, 5.0)
				alumno.Asistencias = append(alumno.Asistencias, 50)
			}
			alumnos = append(alumnos, alumno)
			usuarios = append(usuarios, mongouser)
			//Guardar Alumnos
		}
	}

	if usuariosmodel.GuardarAlumnosMasivamente(alumnos, usuarios) {
		fmt.Println("EXITO!")
	}

	htmlcode := fmt.Sprintf(`
			<script>
				alert("Alumnos Guardados - =)");
				location.replace("/calificaciones");
			</script>
		`)
	ctx.HTML(htmlcode)
}

//GenerarBoleta Obtiene la id de alumno, y genera un documento que se genera y descarga o guarda o abre.
func GenerarBoleta(ctx iris.Context) {
	data := ctx.PostValue("data")
	var htmlcode string

	var alumno calificacionesmodel.Alumno
	var materias []calificacionesmodel.Materia

	alumno = calificacionesmodel.ExtraeAlumno(data)

	materias = calificacionesmodel.ExtraeMateriasPorSemestre(alumno.CursandoSem)

	configuracionboleta := calificacionesmodel.ExtraeConfigBoleta()

	pdf := gofpdf.New("P", "mm", "Letter", `./Recursos/font`)

	tr := pdf.UnicodeTranslatorFromDescriptor("")

	pdf.AddPage()

	var opt gofpdf.ImageOptions
	pdf.ImageOptions(`./Recursos/Imagenes/seplogo.png`, 12, 9, 50, 30, false, opt, 0, "")
	pdf.ImageOptions(`./Recursos/Imagenes/Logoexp.png`, 155, 13, 50, 30, false, opt, 0, "")

	pdf.SetXY(58, 12)
	pdf.SetTextColor(155, 155, 155)
	pdf.SetFont("Arial", "B", 9)
	//ENCABEZADO DE LA BOLETA
	pdf.CellFormat(100, 10, tr("SECRETARIA DE EDUCACIÓN"), "", 0, "C", false, 0, "")
	pdf.SetXY(58, 16)
	pdf.CellFormat(100, 10, tr("SUBSECRETARIA DE EDUCACIÓN FEDERALIZADA"), "", 0, "C", false, 0, "")
	pdf.SetXY(58, 20)
	pdf.CellFormat(100, 10, tr("DIRECCION DE EDUCACION SECUNDARIA Y SUPERIOR"), "", 0, "C", false, 0, "")
	pdf.SetXY(58, 24)
	pdf.CellFormat(100, 10, tr("DEPARAMENTO DE EDUCACION NORMAL"), "", 0, "C", false, 0, "")
	pdf.SetXY(58, 28)
	pdf.CellFormat(100, 10, tr("ESCUELA NORMAL EXPERIMENTAL"), "", 0, "C", false, 0, "")
	pdf.SetXY(58, 32)
	pdf.CellFormat(100, 10, tr(`"LA ENSEÑANZA E IGNACIO MANUEL ALTAMIRANO"`), "", 0, "C", false, 0, "")
	pdf.SetXY(58, 36)
	pdf.CellFormat(100, 10, tr("CLAVE 07DNL0001X"), "", 0, "C", false, 0, "")
	pdf.SetTextColor(110, 110, 110)
	pdf.SetXY(58, 43)
	pdf.CellFormat(100, 10, tr("BOLETA DE CALIFICACIONES"), "", 0, "C", false, 0, "")

	pdf.SetDrawColor(0, 0, 0)
	pdf.SetLineWidth(0.4)
	pdf.Line(10, 10, 206, 10) //Arriba
	pdf.Line(9, 9, 207, 9)

	pdf.Line(10, 270, 206, 270) //
	pdf.Line(9, 271, 207, 271)  //Abajo bien

	pdf.Line(10, 10, 10, 270) // izq
	pdf.Line(9, 9, 9, 271)

	pdf.Line(206, 10, 206, 270) //der bien
	pdf.Line(207, 9, 207, 271)

	//CUERPO DE LA BOLETA -> ENCABEZADO
	pdf.SetLineWidth(0.3)
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetXY(25, 58)
	pdf.CellFormat(50, 5, tr("NOMBRE DEL ALUMNO:"), "", 0, "R", false, 0, "")
	pdf.SetFont("Times", "B", 10)
	pdf.CellFormat(100, 5, tr(alumno.Nombre+" "+alumno.ApellidoP+" "+alumno.ApellidoM), "1B", 0, "C", false, 0, "")
	pdf.SetXY(25, 65)
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(50, 5, tr("No. DE CONTROL:"), "", 0, "R", false, 0, "")
	pdf.SetFont("Times", "B", 10)
	pdf.CellFormat(100, 5, tr(alumno.Matricula), "1B", 0, "C", false, 0, "")

	pdf.SetXY(58, 78)
	pdf.CellFormat(100, 10, tr("LICENCIATURA EN EDUCACIÓN "+alumno.Licenciatura), "", 0, "C", false, 0, "")

	pdf.SetXY(43, 93)
	pdf.SetFont("Times", "", 10)
	pdf.CellFormat(50, 5, tr("SEMESTRE:"), "", 0, "R", false, 0, "")
	pdf.SetFont("Arial", "B", 11)
	pdf.CellFormat(50, 3, tr(alumno.Semestre), "1B", 0, "C", false, 0, "")

	pdf.SetXY(43, 98)
	pdf.SetFont("Times", "", 10)
	pdf.CellFormat(50, 5, tr("GRUPO:"), "", 0, "R", false, 0, "")
	pdf.SetFont("Arial", "B", 11)
	pdf.CellFormat(50, 3, tr("Ú N I C O"), "1B", 0, "C", false, 0, "")

	pdf.SetXY(43, 103)
	pdf.SetFont("Times", "", 10)
	pdf.CellFormat(50, 5, tr("AÑO ESCOLAR:"), "", 0, "R", false, 0, "")
	pdf.SetFont("Arial", "B", 11)
	pdf.CellFormat(50, 3, tr(configuracionboleta.AnioEscolar), "1B", 0, "C", false, 0, "")

	//CUERPO DE LA BOLETA -> MATERIAS Y CALIFICACIONES

	pdf.SetXY(25, 120)
	pdf.SetFont("Times", "", 12)
	pdf.CellFormat(140, 8, tr("M A T E R I A S"), "1", 0, "C", false, 0, "")
	pdf.CellFormat(25, 8, tr("C A L I F"), "1", 0, "C", false, 0, "")

	pixelmateria := 7.0
	iniciomaterias := 133

	for k, v := range materias {
		iniciomaterias = iniciomaterias + int(pixelmateria)

		pdf.SetFont("Arial", "", 9)
		pdf.SetXY(25, float64(iniciomaterias))
		pdf.CellFormat(140, 7, tr(v.Materia), "1", 0, "L", false, 0, "")
		pdf.CellFormat(25, 7, tr(fmt.Sprintf("%v", alumno.Calificaciones[k])), "1", 0, "C", false, 0, "")

	}

	// pdf.SetXY(25, 147)
	// pdf.CellFormat(140, 7, tr("EL SUJETO Y SU FORMACION PERSONAL"), "1", 0, "L", false, 0, "")
	// pdf.CellFormat(25, 7, tr("7"), "1", 0, "C", false, 0, "")

	// var materia string

	// materia = "HERRAMIENTAS PARA LA OBSERVACION Y ANALISIS DE LA PRACTICA EDUCATIVA"

	// pdf.SetXY(25, 154)
	// pdf.CellFormat(140, 7, tr(materia), "1", 0, "L", false, 0, "")
	// pdf.CellFormat(25, 7, tr("7"), "1", 0, "C", false, 0, "")

	//CUERPO DE LA BOLETA -> LUGAR FECHA Y FIRMAS
	pdf.SetXY(58, 210)
	pdf.CellFormat(100, 10, tr("TUXTLA CHICO, CHIAPAS A "+configuracionboleta.FechaBoleta), "", 0, "C", false, 0, "")

	pdf.SetXY(20, 245)
	pdf.SetFont("Times", "", 9)
	pdf.CellFormat(100, 5, tr(configuracionboleta.SubDirector), "", 0, "C", false, 0, "")
	pdf.SetXY(20, 250)
	pdf.CellFormat(100, 5, tr("DIRECTOR"), "", 0, "C", false, 0, "")

	pdf.SetXY(100, 245)
	pdf.CellFormat(100, 5, tr(configuracionboleta.Director), "", 0, "C", false, 0, "")
	pdf.SetXY(100, 250)
	pdf.CellFormat(100, 5, tr("DIRECTOR"), "", 0, "C", false, 0, "")

	// fileee := `.\Recursos\Archivos\` + data + `.pdf`
	fileee := `./Recursos/Archivos/` + alumno.Matricula + `.pdf`
	// fileee := `../PDFEXPE/` + data + `.pdf`

	err4 := pdf.OutputFileAndClose(fileee)
	if err4 != nil {
		fmt.Println(err4)
		fmt.Println("Ocurrio un error creando el archivo pdf")

	} else {
		htmlcode += fmt.Sprintf(`<script>
		Descargar('%v');
		</script>`, alumno.Matricula)
	}

	ctx.HTML(htmlcode)

}

//Ligar Herramienta temporal para asignar correctamente los usuarios a los alumnos que les hace falta MongoUser y a los usuarios que esta mal su UserID
func Ligar(ctx iris.Context) {

	// userOn := indexmodel.GetUserOn(sessioncontroller.Sess.Start(ctx).GetString("UserID"))

	// fmt.Println("ID ->", userOn.ID)
	// fmt.Println("Usuario Logeado : ", userOn.Usuario)

	// alumnoconusuario :=

	var alumnos []calificacionesmodel.Alumno

	var usuarios []indexmodel.MongoUser

	alumnos = calificacionesmodel.ExtraeSoloAlumnos()

	usuarios = indexmodel.ExtraeSoloUsuarios()

	fmt.Println("Usuarios encontrdos: ", len(usuarios))
	fmt.Println("Alumnos encontrdos", len(alumnos))

	encontrados := 0

	for _, v := range alumnos {

		for _, vv := range usuarios {

			// if v.SiguienteSem == vv.Usuario && v.AnteriorSem == vv.Key {
			if v.Nombre+" "+v.ApellidoP+" "+v.ApellidoM == vv.Nombre+" "+vv.Apellidos {

				if v.ID == vv.UserID && v.MongoUser == vv.ID {
					fmt.Println("Son Iguales")
					fmt.Println("alumno", v)
					fmt.Println()
					fmt.Println("usuario", vv)

				} else {
					fmt.Println("Son Diferentes... Igualando")

					v.MongoUser = vv.ID
					vv.UserID = v.ID
					calificacionesmodel.HerramientaAsignacionAlumnos(v)
					indexmodel.HerramientaAsignacionUsuarios(vv)
					encontrados++

				}
				// fmt.Println("Usuario -> ", k, " ", kk, " ", vv.Usuario)
				// fmt.Println("			", "Alumno ID ->", v.ID, "MongoUser ID->", v.MongoUser)
				// fmt.Println("			", "Usuario ID->", vv.ID, "User ID->", vv.UserID)
			}

		}

	}

	fmt.Println("Diferentes ", encontrados)

	ctx.HTML("<script>alert('Modificados'); </script> ")

}

//ImprimirCalificacion -> Imprime el pdf para la impresion de la boleta
func ImprimirCalificacion(ctx iris.Context) {

	idalumno := ctx.PostValue("data")
	var htmlcode string

	var alumno calificacionesmodel.Alumno
	var materias []calificacionesmodel.Materia
	var semestre calificacionesmodel.Semestre
	var docentes []calificacionesmodel.Docente
	var nombresdocentes []string

	alumno = calificacionesmodel.ExtraeAlumno(idalumno)
	materias = calificacionesmodel.ExtraeMateriasPorSemestre(alumno.CursandoSem)
	semestre = calificacionesmodel.ExtraeSemestreString(alumno.CursandoSem.Hex())
	docentes = calificacionesmodel.ExtraeDocentes(materias)

	for _, vd := range docentes {

		nombre := vd.Nombre + " " + vd.ApellidoP + " " + vd.ApellidoM
		nombresdocentes = append(nombresdocentes, nombre)
	}

	pdf := gofpdf.New("L", "mm", "Letter", `./Recursos/font`)

	tr := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.AddPage()
	pdf.AddFont("Montse", "", "Montse.json")
	pdf.SetFont("Montse", "", 14)
	pdf.SetXY(20, 20)
	pdf.Cell(40, 10, tr(alumno.Nombre+" "+alumno.ApellidoP+" "+alumno.ApellidoM))
	pdf.Ln(7)
	pdf.SetFont("Montse", "", 10)
	pdf.Cell(10, 4, "")
	pdf.Cell(20, 4, tr("Plan : "+semestre.Plan+"  Licenciatura en Educación "+semestre.Licenciatura))

	pdf.SetLineWidth(0.5)
	pdf.SetDrawColor(25, 25, 25)
	pdf.Line(21, 31, 240, 31)
	pdf.Ln(5)
	pdf.Cell(10, 4, "")
	pdf.SetFont("Helvetica", "", 10)
	pdf.Cell(5, 10, tr("#"))
	pdf.Cell(110, 10, tr("Materia"))
	pdf.Cell(30, 10, tr("Calificacion"))
	pdf.Cell(30, 10, tr("Asistencia"))
	pdf.Cell(40, 10, tr("Docente"))
	pdf.Ln(5)
	pdf.SetFont("Helvetica", "", 8)
	for k, v := range materias {
		pdf.Cell(10, 4, "")
		pdf.Cell(5, 10, fmt.Sprintf(`%v`, k+1))
		pdf.Cell(100, 10, tr(v.Materia))
		pdf.CellFormat(30, 10, fmt.Sprintf(`%v`, alumno.Calificaciones[k]), "", 0, "C", false, 0, "")
		pdf.CellFormat(30, 10, fmt.Sprintf(`%v`, alumno.Asistencias[k]), "", 0, "C", false, 0, "")
		// pdf.Cell(30, 10, fmt.Sprintf(`%v`, alumno.Calificaciones[k]))
		// pdf.Cell(30, 10, fmt.Sprintf(`%v`, alumno.Asistencias[k]))
		pdf.Cell(40, 10, tr(nombresdocentes[k]))
		pdf.Ln(7)

	}

	pdf.SetFont("Montse", "", 15)
	pdf.SetXY(40, 100)
	pdf.SetDrawColor(225, 225, 225)
	pdf.CellFormat(180, 7, tr(`Normal Experimental "La enseñanza e Ignacio Manuel Altamirano"`), "1", 0, "C", false, 0, "")
	pdf.Ln(7)
	pdf.SetXY(188, 25)
	pdf.CellFormat(50, 7, tr("Semestre "+semestre.Semestre+"°"), "", 0, "R", false, 0, "")

	// fileee := `.\Recursos\Archivos\` + data + `.pdf`
	fileee := `./Recursos/Archivos/` + idalumno + `.pdf`
	// fileee := `../PDFEXPE/` + data + `.pdf`

	err4 := pdf.OutputFileAndClose(fileee)
	if err4 != nil {
		fmt.Println(err4)
		fmt.Println("Ocurrio un error creando el archivo pdf")

	} else {
		htmlcode += fmt.Sprintf(`<script>
		Descargar('%v');
		</script>`, idalumno)
	}

	ctx.HTML(htmlcode)

}

//ObtenConfig -> Devuelve la configuracon solicitada
func ObtenConfig(ctx iris.Context) {

	tipoconfig := ctx.PostValue("data")

	var htmlcode string

	switch tipoconfig {
	case "General":

		configuracion := calificacionesmodel.ExtraeConfigBoleta()

		htmlcode += fmt.Sprintf(`

		<form action="/guardaconfiguracion" method="POST" >

        <div class="col-sm-12">
            <div class="form-group row">
                <label for="centroescolar" class="col-sm-1 col-form-label negrita"> Centro Escolar: </label>
                <div class="col-sm-4 col-md-4 col-lg-4">

					<input type="text" class="form-control" id="centroescolar" name="centroescolar" placeholder="Nombre del Centro Escolar" value="%v" required>
                </div>
                <label for="claveprim" class="col-sm-1 col-form-label negrita"> Clave Primaria: </label>
                <div class="col-sm-2 col-md-2 col-lg-2">
				<input type="text" class="form-control" id="claveprim" name="claveprim" placeholder="Clave de Primaria" value="%v" required>

                </div>
                <label for="claveprees" class="col-sm-2 col-form-label negrita"> Clave Preescolar: </label>
                <div class="col-sm-2 col-md-2 col-lg-2">
				<input type="text" class="form-control" id="claveprees" name="claveprees" placeholder="Clave de Preescolar" value="%v" required>
                </div>
            </div>
            <div class="form-group row">
				<label for="director" class="col-sm-1 col-form-label negrita"> Director: </label>
				<div class="col-sm-3 col-md-3 col-lg-3">
				<input type="text" class="form-control" id="director" name="director" placeholder="Nombre del Director" value="%v" required>

				</div>
				<label for="subdirector" class="col-sm-1 col-form-label negrita"> SubDirector: </label>
				<div class="col-sm-3 col-md-3 col-lg-3">
				<input type="text" class="form-control" id="subdirector" name="subdirector" placeholder="Nombre del Subdirector" value="%v" required>


				</div>

				<label for="horario" class="col-sm-1 col-form-label negrita"> Horario: </label>
                <div class="col-sm-2 col-md-2 col-lg-2">
				<input type="text" class="form-control" id="horario" name="horario" placeholder="Nombre del Subdirector" value="%v" required>
                </div>

            </div>

            <div class="form-group row">

				<label for="fechaboleta" class="col-sm-1 col-form-label negrita"> Fecha Boleta: </label>
				<div class="col-sm-3 col-md-3 col-lg-3">
				<input type="text" class="form-control" id="fechaboleta" name="fechaboleta" placeholder="Fecha en letras de la boleta a imprimir" value="%v" required>


				</div>

				<label for="anioescolar" class="col-sm-1 col-form-label negrita"> Año Escolar: </label>
                <div class="col-sm-2 col-md-2 col-lg-2">
				<input type="text" class="form-control" id="anioescolar" name="anioescolar" placeholder="Año escolar" value="%v" required>

                </div>

				<label for="domicilio" class="col-sm-1 col-form-label negrita"> Domicilio: </label>
                <div class="col-sm-4 col-md-4 col-lg-4">
				<input type="text" class="form-control" id="domicilioce" name="domicilioce" placeholder="Domicilio completo del centro escolar" value="%v" required>

            	</div>
			</div>
				<div class="form-group row">

				<label for="mensaje" class="col-sm-12 col-form-label negrita"> Mensaje Diario: </label>
                <div class="col-sm-10 col-md-10 col-lg-10">
				<input type="text" class="form-control" id="mensaje" name="mensaje" placeholder="Mensaje diario para todos los usuarios" value="%v" required>
				<input type="hidden" name="hiddenid" value="%v">
				<input type="hidden" name="configuracion" value="%v">
                </div>

            </div>


            <div class="text-center container "> 
                <button type="submit"> Guardar Configuración</button>
            </div>
            </div>
        </div>
    </form>`, configuracion.CentroEscolar, configuracion.ClavePrimaria, configuracion.ClavePreescolar, configuracion.Director, configuracion.SubDirector, configuracion.Horario, configuracion.FechaBoleta, configuracion.AnioEscolar, configuracion.Domicilio, configuracion.MensajeDiario, configuracion.ID.Hex(), configuracion.Configuracion)

		break
	}

	ctx.HTML(htmlcode)

}

//GuardaConfiguracion -> Guarda la configuracion previamente solicitada
func GuardaConfiguracion(ctx iris.Context) {

	var config calificacionesmodel.Configuracion

	var htmlcode string

	config.Configuracion = ctx.PostValue("configuracion")
	config.CentroEscolar = ctx.PostValue("centroescolar")
	config.ClavePrimaria = ctx.PostValue("claveprim")
	config.ClavePreescolar = ctx.PostValue("claveprees")
	config.Director = ctx.PostValue("director")
	config.SubDirector = ctx.PostValue("subdirector")
	config.Horario = ctx.PostValue("horario")
	config.FechaBoleta = ctx.PostValue("fechaboleta")
	config.AnioEscolar = ctx.PostValue("anioescolar")
	config.Domicilio = ctx.PostValue("domicilioce")
	config.MensajeDiario = ctx.PostValue("mensaje")
	id := ctx.PostValue("hiddenid")

	calificacionesmodel.ActualizaConfig(id, config)

	htmlcode = fmt.Sprintf(`
	<script>
		alert("Configuracion Guardada - =)");
		location.replace("/calificaciones");
	</script>
`)

	ctx.HTML(htmlcode)

}
