package usuarioscontroller

import (
	"fmt"
	"time"

	"github.com/kataras/iris/v12"
	sessioncontroller "github.com/vadgun/Experimental/Controladores/SessionController"
	calificacionesmodel "github.com/vadgun/Experimental/Modelos/CalificacionesModel"
	indexmodel "github.com/vadgun/Experimental/Modelos/IndexModel"
	usuariosmodel "github.com/vadgun/Experimental/Modelos/UsuariosModel"
	"gopkg.in/mgo.v2/bson"
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

		tienepermiso := indexmodel.TienePermiso(1, userOn, usuario)

		if !tienepermiso {
			ctx.Redirect("/login", iris.StatusSeeOther)
		}

		if err := ctx.View("Usuarios.html"); err != nil {
			ctx.Application().Logger().Infof(err.Error())
		}
	} else {
		ctx.Redirect("/login", iris.StatusSeeOther)
	}
}

//AltaDeUsuario Recibe el formulario de alta de usuario y lo guarda en la base de datos creando la logica de 2 entidades
func AltaDeUsuario(ctx iris.Context) { // la 1ra es del docente y la 2da del mongouser, se conectan entre si.

	tipoUsuario := ctx.PostValue("tipoUsuario")
	var htmlcode string

	switch tipoUsuario {
	case "Alumno":

		var alumno calificacionesmodel.Alumno
		var mongouser indexmodel.MongoUser

		//alumno.MongoUser = Esta variable sera asignada en el modelo por el paquete bson
		alumno.IsSystemUser = true
		alumno.Matricula = ctx.PostValue("nummatricula")
		alumno.Nombre = ctx.PostValue("nombrealumno")
		alumno.ApellidoP = ctx.PostValue("apaterno")
		alumno.ApellidoM = ctx.PostValue("amaterno")
		alumno.Sexo = ctx.PostValue("sexo")
		//Evalular la fecha y agregarla correctamente.
		layout := "2006-01-02"
		location, _ := time.LoadLocation("America/Mexico_City")

		fechanac := ctx.PostValue("fechanac")
		fechanacparsed, _ := time.ParseInLocation(layout, fechanac, location)
		alumno.FechaNac = fechanacparsed
		alumno.Curp = ctx.PostValue("curp")
		alumno.Calle = ctx.PostValue("calle")
		alumno.Numero = ctx.PostValue("numero")
		alumno.ColAsentamiento = ctx.PostValue("colonia")
		alumno.Municipio = ctx.PostValue("municipio")
		alumno.Estado = ctx.PostValue("estado")
		alumno.Telefono = ctx.PostValue("telefono")
		alumno.TipoSangre = ctx.PostValue("tiposangre")
		idsemestre := bson.ObjectIdHex(ctx.PostValue("semestre"))

		//Con el id de semestre crear los arreglos correspondientes a las materias y calificaciones en 5
		semestre := calificacionesmodel.ExtraeSemestre(idsemestre)
		alumno.Plan = semestre.Plan
		alumno.Licenciatura = semestre.Licenciatura

		alumno.Materias = semestre.Materias

		for i := 0; i < len(alumno.Materias); i++ {
			alumno.Calificaciones = append(alumno.Calificaciones, 5.0)
		}

		alumno.CursandoSem = idsemestre
		alumno.CorreoE = ctx.PostValue("correoe")
		//Hacer la logica para el semestre siguiente, anterior, y el cual inicio, no urge pueden ir null por ahora
		// alumno.SiguienteSem=ctx.PostValue("")
		// alumno.AnteriorSem=ctx.PostValue("")
		// alumno.InicioSem=ctx.PostValue("")
		//alumno.Imagen= variable para bson package
		alumno.Horario = ""

		mongouser.Nombre = ctx.PostValue("nombrealumno")
		mongouser.Apellidos = ctx.PostValue("apaterno") + " " + ctx.PostValue("amaterno")
		//mongouser.Edad = CalcularEdad(alumno.FechaNac)int
		mongouser.Usuario = ctx.PostValue("nameuser")
		mongouser.Key = ctx.PostValue("passuser")
		mongouser.Puesto = "Alumno de la Licenciatura" + ctx.PostValue("licenciatura")
		mongouser.Nombre2 = "Alumno"
		//mongouser.UserID= variable para bson package
		mongouser.Alumno = true
		mongouser.Docente = false
		mongouser.Administrativo = false
		mongouser.Director = false
		mongouser.Admin = false

		if usuariosmodel.GuardaEntidadesDeAlumnos(alumno, mongouser) {
			htmlcode += fmt.Sprintf(`<script>
		alert("Alumno Guardado");
		location.replace("/usuarios");
		</script>`)
		}

		break
	case "Docente":

		//Hacer lo mismo que para a alumno

		var docente calificacionesmodel.Docente
		var mongouser indexmodel.MongoUser

		docente.IsSystemUser = true
		docente.Nombre = ctx.PostValue("nombredocente")
		docente.ApellidoP = ctx.PostValue("apaterno")
		docente.ApellidoM = ctx.PostValue("amaterno")
		layout := "2006-01-02"
		location, _ := time.LoadLocation("America/Mexico_City")

		fechanac := ctx.PostValue("fechanac")
		fechanacparsed, _ := time.ParseInLocation(layout, fechanac, location)
		docente.FechaNac = fechanacparsed
		docente.Curp = ctx.PostValue("curp")
		docente.Rfc = ctx.PostValue("rfc")
		docente.Calle = ctx.PostValue("calle")
		docente.ColAsentamiento = ctx.PostValue("colonia")
		docente.Municipio = ctx.PostValue("municipio")
		docente.Estado = ctx.PostValue("estado")
		docente.Telefono1 = ctx.PostValue("telefono")
		docente.Telefono2 = ctx.PostValue("telefono2")
		docente.TipoSangre = ctx.PostValue("tiposangre")
		//		docente.Grupos=
		//		docente.Materias=
		//		docente.Horario=
		docente.CorreoE = ctx.PostValue("correoe")

		fechacapinicio := ctx.PostValue("fechacapinicio")
		fechacapfin := ctx.PostValue("fechacapfin")

		fechacapinicioparsed, _ := time.ParseInLocation(layout, fechacapinicio, location)
		fechacapfinparsed, _ := time.ParseInLocation(layout, fechacapfin, location)

		docente.CapturaInicio = fechacapinicioparsed
		docente.CapturaFin = fechacapfinparsed

		//ID variable para bson
		mongouser.Nombre = ctx.PostValue("nombredocente")
		mongouser.Apellidos = ctx.PostValue("apaterno") + " " + ctx.PostValue("amaterno")
		mongouser.Edad = 0
		mongouser.Usuario = ctx.PostValue("nameuser")
		mongouser.Telefono = ctx.PostValue("telefono")
		mongouser.Puesto = "Docente frente a grupo"
		mongouser.Key = ctx.PostValue("passuser")
		mongouser.Nombre2 = "Docente"
		//UserID variable para bson
		mongouser.Alumno = false
		mongouser.Docente = true
		mongouser.Administrativo = false
		mongouser.Director = false
		mongouser.Admin = false

		if usuariosmodel.GuardaEntidadesDeDocentes(docente, mongouser) {
			htmlcode += fmt.Sprintf(`<script>
		alert("Docente Guardado");
		location.replace("/usuarios");
		</script>`)
		}

		break
	case "Administrativo":
		break
	}

	ctx.HTML(htmlcode)

}

//SolicitarUsuario Solicita el formulario para el usuario a crear
func SolicitarUsuario(ctx iris.Context) {

	data := ctx.PostValue("data")

	var htmlcode string

	switch data {
	case "Alumno":

		semestres := usuariosmodel.ExtraeSemestres()

		htmlcode += fmt.Sprintf(`
		<form action="/altadeusuario" method="POST" enctype="multipart/form-data" id="formularioAlumno" name="formularioAlumno" >
        <div class="col-lg-12">
            <h6 class="border-bottoms-c"> Datos personales del alumno: </h6>
            
            <div class="form-group row">
                <label for="nummatricula" class="col-sm-1 col-form-label negrita"> #Matricula: </label>
                <div class="col-sm-3 col-md-3 col-lg-3">
                    <input type="text" class="form-control" id="nummatricula" name="nummatricula" placeholder="Introduce matricula del alumno" minlength="12" maxlength="12" value="" required>
                </div>
                <label for="tiposangre" class="col-sm-2 col-form-label negrita"> Tipo de sangre: </label>
                <div class="col-sm-2 col-md-2 col-lg-2">
                    <input type="text" class="form-control" id="tiposangre" name="tiposangre" placeholder="Tipo de sangre" value="" >
                </div>
                <label for="sexo" class="col-sm-1 col-form-label negrita"> Sexo: </label>
                <div class="col-sm-2 col-md-2 col-lg-2">
                    <select class="form-control" id="sexo" name="sexo" value="" required>
                        <option value="">Sexo</option>
                        <option value="Masculino">Masculino</option>
                        <option value="Femenino">Femenino</option>

                    </select>  
                    
                </div>


            </div>        
            <div class="form-group row">
                <label for="nombrealumno" class="col-sm-1 col-form-label negrita"> Nombre: </label>
                <div class="col-sm-3 col-md-3 col-lg-3">
                    <input type="text" class="form-control" id="nombrealumno" name="nombrealumno" placeholder="Nombre del alumno" value="" required>
                </div>

                <label for="apaterno" class="col-sm-1 col-form-label negrita"> APaterno: </label>
                <div class="col-sm-3 col-md-3 col-lg-3">
                    <input type="text" class="form-control" id="apaterno" name="apaterno" placeholder="Apellido Paterno" value="" required>
                </div>
                <label for="amaterno" class="col-sm-1 col-form-label negrita"> AMaterno: </label>
                <div class="col-sm-3 col-md-3 col-lg-3">
                    <input type="text" class="form-control" id="amaterno" name="amaterno" placeholder="Apellido Materno" value="" required>
                </div>
            </div>

            <div class="form-group row">
                <label for="fechanac" class="col-sm-3 col-form-label negrita"> Fecha de Nacimiento: </label>
                <div class="col-sm-4 col-md-4 col-lg-4">
                    <input type="date" class="form-control" id="fechanac" name="fechanac" value="" required>
                </div>

                <label for="curp" class="col-sm-1 col-form-label negrita"> CURP: </label>
                <div class="col-sm-4 col-md-4 col-lg-4">
                    <input type="text" class="form-control" id="curp" name="curp" placeholder="18 digitos de la CURP" maxlength="18" minlength="18" value="" required>
                </div>
            </div>

            <hr>
            <h6 class="border-bottoms-c"> Datos domiciliares del alumno: </h6>

            <div class="form-group row">
                <label for="calle" class="col-sm-1 col-form-label negrita"> Calle: </label>
                <div class="col-sm-3 col-md-3 col-lg-3">
                    <input type="text" class="form-control" id="calle" name="calle" placeholder="Calle" value="" required>
                </div>

                <label for="numero" class="col-sm-1 col-form-label negrita"> Número: </label>
                <div class="col-sm-2 col-md-2 col-lg-2">
                    <input type="text" class="form-control" id="numero" name="numero" placeholder="Número" value="" required>
                </div>
                <label for="colonia" class="col-sm-1 col-form-label negrita"> Colonia: </label>
                <div class="col-sm-4 col-md-4 col-lg-4">
                    <input type="text" class="form-control" id="colonia" name="colonia" placeholder="Colonia/Asentamiento" value="" required>
                </div>
            </div>

            <div class="form-group row">
                <label for="municipio" class="col-sm-1 col-form-label negrita"> Municipio: </label>
                <div class="col-sm-3 col-md-3 col-lg-3">
                    <input type="text" class="form-control" id="municipio" name="municipio" placeholder="Municipio" value="" required>
                </div>

                <label for="estado" class="col-sm-1 col-form-label negrita"> Estado: </label>
                <div class="col-sm-2 col-md-2 col-lg-2">
                    <input type="text" class="form-control" id="estado" name="estado" placeholder="Estado" value="" required>
                </div>
                <label for="telefono" class="col-sm-1 col-form-label negrita"> Télefono: </label>
                <div class="col-sm-4 col-md-4 col-lg-4">
                    <input type="text" class="form-control" id="telefono" name="telefono" placeholder="10 digitos" maxlength="10" minlength="10" value="" required>
                </div>
            </div>
            <hr>
            <h6 class="border-bottoms-c"> Datos escolares del alumno: </h6>
            <div class="form-group row">
                <label for="semestre" class="col-sm-2 col-form-label negrita"> Semestre: </label>
                <div class="col-sm-10 col-md-10 col-lg-10">
                    <select class="form-control" id="semestre" name="semestre" value="" >
						<option value="">Selecciona</option>`)

		for _, v := range semestres {
			htmlcode += fmt.Sprintf(`
			<option value="%v">%v - %v - %v - %v materias </option>`, v.ID.Hex(), v.Semestre, v.Plan, v.Licenciatura, len(v.Materias))
		}

		htmlcode += fmt.Sprintf(`
                    </select>
                </div>
            </div>
            <hr>
            <h6 class="border-bottoms-c"> Datos del sistema del alumno: </h6>

            <div class="form-group row">
                <label for="nameuser" class="col-sm-2 col-form-label negrita"> Nombre de Usuario: </label>
                <div class="col-sm-3 col-md-3 col-lg-3">
                    <input type="text" class="form-control" id="nameuser" name="nameuser" placeholder="Nombre de usuario" value="" required>
                </div>

                <label for="passuser" class="col-sm-3 col-form-label negrita"> Contraseña preestablecida: </label>
                <div class="col-sm-3 col-md-3 col-lg-3">
                    <input type="text" class="form-control" id="passuser" name="passuser" placeholder="Contraseña" value="" required>
                </div>
            </div>

            <div class="form-group row">
                <label for="correoe" class="col-sm-2 col-form-label negrita"> Correo electronico: </label>
                <div class="col-sm-4 col-md-4 col-lg-4">
                    <input type="email" class="form-control" id="correoe" name="correoe" placeholder="Correo electronico" value="" required>
                </div>

                <label for="tipoUsuario" class="col-sm-2 col-form-label negrita"> Tipo de Usuario: </label>
                <div class="col-sm-4 col-md-4 col-lg-4">
                    <input type="email" class="form-control" id="tipoUsuario" name="tipoUsuario" value="Alumno" readonly required>
                    <input type="hidden" name="tipoUsuario1" value="Alumno">
                </div>
            </div>

			<!--    <div class="form-group row">
			<label for="imagen" class="col-sm-3 col-form-label negrita"> Imagen: </label>
			<div class="col-sm-2 col-md-4 col-lg-5">
				<input type="file" class="dropify" data-allowed-file-extensions="jpg jpeg" id="imagen" name="imagen" required />
			</div>
			</div>
	-->



            <div class="form-group row centrado">                
                   <button type="submit"> Guardar Alumno</button>
            </div>
        </div>

    
    </form>

		`)
		break
	case "Docente":
		htmlcode += fmt.Sprintf(`
		<form action="/altadeusuario" method="POST" enctype="multipart/form-data" id="formularioAlumno" name="formularioAlumno" >
        <div class="col-lg-12">
            <h6 class="border-bottoms-d"> Datos personales del docente: </h6>
            
            <div class="form-group row">
                <label for="nummatricula" class="col-sm-1 col-form-label negrita"> #Matricula: </label>
                <div class="col-sm-5 col-md-5 col-lg-5">
                    <input type="text" class="form-control" id="nummatricula" name="nummatricula" placeholder="Introduce matricula del docente" value="" required>
                </div>
                <label for="tiposangre" class="col-sm-2 col-form-label negrita"> Tipo de sangre: </label>
                <div class="col-sm-4 col-md-4 col-lg-4">
                    <input type="text" class="form-control" id="tiposangre" name="tiposangre" placeholder="Tipo de sangre del docente" value="" >
                </div>

            </div>        
            <div class="form-group row">
                <label for="nombredocente" class="col-sm-1 col-form-label negrita"> Nombre: </label>
                <div class="col-sm-3 col-md-3 col-lg-3">
                    <input type="text" class="form-control" id="nombredocente" name="nombredocente" placeholder="Nombre del docente" value="" required>
                </div>

                <label for="apaterno" class="col-sm-1 col-form-label negrita"> APaterno: </label>
                <div class="col-sm-3 col-md-3 col-lg-3">
                    <input type="text" class="form-control" id="apaterno" name="apaterno" placeholder="Apellido Paterno" value="" required>
                </div>
                <label for="amaterno" class="col-sm-1 col-form-label negrita"> AMaterno: </label>
                <div class="col-sm-3 col-md-3 col-lg-3">
                    <input type="text" class="form-control" id="amaterno" name="amaterno" placeholder="Apellido Materno" value="" required>
                </div>
            </div>

            <div class="form-group row">
                <label for="fechanac" class="col-sm-3 col-form-label negrita"> Fecha de Nacimiento: </label>
                <div class="col-sm-4 col-md-4 col-lg-4">
                    <input type="date" class="form-control" id="fechanac" name="fechanac" value="" required>
                </div>

                <label for="curp" class="col-sm-1 col-form-label negrita"> CURP: </label>
                <div class="col-sm-4 col-md-4 col-lg-4">
                    <input type="text" class="form-control" id="curp" name="curp" placeholder="18 digitos de la CURP" maxlength="18" minlength="18" value="" required>
                </div>
            </div>

            <hr>
            <h6 class="border-bottoms-d"> Datos domiciliares del docente: </h6>

            <div class="form-group row">
                <label for="calle" class="col-sm-1 col-form-label negrita"> Calle: </label>
                <div class="col-sm-3 col-md-3 col-lg-3">
                    <input type="text" class="form-control" id="calle" name="calle" placeholder="Calle" value="" required>
                </div>

                <label for="numero" class="col-sm-1 col-form-label negrita"> Número: </label>
                <div class="col-sm-2 col-md-2 col-lg-2">
                    <input type="text" class="form-control" id="numero" name="numero" placeholder="Número" value="" required>
                </div>
                <label for="colonia" class="col-sm-1 col-form-label negrita"> Colonia: </label>
                <div class="col-sm-4 col-md-4 col-lg-4">
                    <input type="text" class="form-control" id="colonia" name="colonia" placeholder="Colonia/Asentamiento" value="" required>
                </div>
            </div>

            <div class="form-group row">
                <label for="municipio" class="col-sm-1 col-form-label negrita"> Municipio: </label>
                <div class="col-sm-3 col-md-3 col-lg-3">
                    <input type="text" class="form-control" id="municipio" name="municipio" placeholder="Municipio" value="" required>
                </div>

                <label for="estado" class="col-sm-1 col-form-label negrita"> Estado: </label>
                <div class="col-sm-2 col-md-2 col-lg-2">
                    <input type="text" class="form-control" id="estado" name="estado" placeholder="Estado" value="" required>
                </div>
                <label for="telefono" class="col-sm-1 col-form-label negrita"> Télefonos: </label>
                <div class="d-flex col-sm-4 col-md-4 col-lg-4">
                    <input type="text" class="form-control" id="telefono" name="telefono" placeholder="10 digitos" maxlength="10" minlength="10" value="" required>
                       <input type="text" class="form-control" id="telefono2" name="telefono2" placeholder="10 digitos" maxlength="10" minlength="10" value="" required>
                
                </div>
            </div>
            <hr>
            <h6 class="border-bottoms-d"> Datos del sistema del docente: </h6>
            <div class="form-group row">
                <label for="nameuser" class="col-sm-2 col-form-label negrita"> Nombre de Usuario: </label>
                <div class="col-sm-3 col-md-3 col-lg-3">
                    <input type="text" class="form-control" id="nameuser" name="nameuser" placeholder="Nombre de usuario" value="" required>
                </div>
                <label for="passuser" class="col-sm-3 col-form-label negrita"> Contraseña preestablecida: </label>
                <div class="col-sm-3 col-md-3 col-lg-3">
                    <input type="text" class="form-control" id="passuser" name="passuser" placeholder="Contraseña" value="" required>
                </div>
            </div>
            <div class="form-group row">
                <label for="correoe" class="col-sm-2 col-form-label negrita"> Correo electronico: </label>
                <div class="col-sm-4 col-md-4 col-lg-4">
                    <input type="email" class="form-control" id="correoe" name="correoe" placeholder="Correo electronico" value="" required>
                </div>
                <label for="tipoUsuario" class="col-sm-2 col-form-label negrita"> Tipo de Usuario: </label>
                <div class="col-sm-4 col-md-4 col-lg-4">
                    <input type="email" class="form-control" id="tipoUsuario" name="tipoUsuario" value="Docente" readonly required>
                </div>
            </div>

            <div class="form-group row">
            <label for="correoe" class="col-sm-2 col-form-label negrita"> Permiso de Captura Inicial: </label>
            <div class="col-sm-4 col-md-4 col-lg-4">
                <input type="date" class="form-control" id="fechacapinicio" name="fechacapinicio" value="" required>
            </div>
            <label for="tipoUsuario" class="col-sm-2 col-form-label negrita"> Permiso de Captura Final:  </label>
            <div class="col-sm-4 col-md-4 col-lg-4">
                <input type="hidden" name="tipoUsuario1" value="Docente">
                <input type="date" class="form-control" id="fechacapfin" name="fechacapfin" required>
            </div>
        </div>

 


        <!--    <div class="form-group row">
                <label for="imagen" class="col-sm-3 col-form-label negrita"> Imagen: </label>
                <div class="col-sm-2 col-md-4 col-lg-5">
                    <input type="file" class="dropify" data-allowed-file-extensions="jpg jpeg" id="imagen" name="imagen" required />
                </div>
				</div>
		-->
            <div class="form-group row centrado">                
                   <button type="submit"> Guardar Docente</button>
            </div>
        </div>    
    </form>
		`)
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
