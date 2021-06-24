package indexcontroller

import (
	"fmt"
	"time"

	"github.com/kataras/iris/v12"
	sessioncontroller "github.com/vadgun/Experimental/Controladores/SessionController"
	calificacionesmodel "github.com/vadgun/Experimental/Modelos/CalificacionesModel"
	indexmodel "github.com/vadgun/Experimental/Modelos/IndexModel"
)

//Index -> Regresa la pagina de inicio
func Index(ctx iris.Context) {
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
		fmt.Println("Bienvenido ", userOn.Nombre)
		if err := ctx.View("Index.html"); err != nil {
			ctx.Application().Logger().Infof(err.Error())
		}
	} else {
		ctx.Redirect("/login", iris.StatusSeeOther)
	}
}

//Reloj -> Regresa la pagina de inicio
func Reloj(ctx iris.Context) {
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
		fmt.Println("Bienvenido ", userOn.Nombre)
		if err := ctx.View("Reloj.html"); err != nil {
			ctx.Application().Logger().Infof(err.Error())
		}
	} else {
		ctx.Redirect("/login", iris.StatusSeeOther)
	}
}

//EditarDatosDeAlumno -> Editar datos en un modal?
func EditarDatosDeAlumno(ctx iris.Context) {
	idstalumno := ctx.PostValue("data")
	var alumno calificacionesmodel.Alumno
	alumno = calificacionesmodel.ExtraeAlumno(idstalumno)
	ctx.JSON(alumno)
}

//EditarAlumno -> Guarda los datos modificados del alumno previamente solicitado
func EditarAlumno(ctx iris.Context) {

	idstalumno := ctx.PostValue("datas1")

	var alumno calificacionesmodel.Alumno

	alumno = calificacionesmodel.ExtraeAlumno(idstalumno)

	alumno.Matricula = ctx.PostValue("matricula")
	alumno.Nombre = ctx.PostValue("nombre")
	alumno.ApellidoP = ctx.PostValue("apellidop")
	alumno.ApellidoM = ctx.PostValue("apellidom")

	layout := "2006-01-02"
	location, _ := time.LoadLocation("America/Mexico_City")

	fechanac := ctx.PostValue("fechanac")

	alumno.FechaNac, _ = time.ParseInLocation(layout, fechanac, location)

	alumno.Curp = ctx.PostValue("curp")
	alumno.Calle = ctx.PostValue("calle")
	alumno.Numero = ctx.PostValue("numero")
	alumno.ColAsentamiento = ctx.PostValue("colAsentamiento")
	alumno.Municipio = ctx.PostValue("municipio")
	alumno.Estado = ctx.PostValue("estado")
	alumno.Telefono = ctx.PostValue("telefono")
	alumno.TipoSangre = ctx.PostValue("tipoSangre")
	alumno.Sexo = ctx.PostValue("sexo")
	alumno.Licenciatura = ctx.PostValue("licenciatura")
	alumno.Semestre = ctx.PostValue("semestrealum")
	alumno.Plan = ctx.PostValue("plan")
	alumno.Nss = ctx.PostValue("nss")
	alumno.Tutor = ctx.PostValue("tutor")
	alumno.OcupacionTutor = ctx.PostValue("ocupaciontutor")
	alumno.ParentescoTutor = ctx.PostValue("parentescoTutor")
	alumno.ContactoCasoEmergencia = ctx.PostValue("contactoemergencia")
	alumno.DiferenteDomicilioTutor = ctx.PostValue("difdomtutor")
	alumno.ReferenciasDomicilio = ctx.PostValue("refdomicilio")
	alumno.CorreoE = ctx.PostValue("correoe")

	fmt.Println(alumno)

	calificacionesmodel.ActualizaAlumno(alumno)

	var htmlcode string

	htmlcode = fmt.Sprintf(`
	<script>
		alert("Alumno Guardado =)");
		location.replace("/alumnos");
	</script>
`)

	ctx.HTML(htmlcode)

}

//Semestres -> Regresa la pagina de semestres
func Semestres(ctx iris.Context) {
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
		fmt.Println("Bienvenido ", userOn.Nombre)
		if err := ctx.View("Semestres.html"); err != nil {
			ctx.Application().Logger().Infof(err.Error())
		}
	} else {
		ctx.Redirect("/login", iris.StatusSeeOther)
	}
}
