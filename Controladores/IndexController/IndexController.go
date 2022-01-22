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
	alumno.Licenciatura = ctx.PostValue("licenciaturahidden")
	alumno.Semestre = ctx.PostValue("semestrealum")
	alumno.Plan = ctx.PostValue("planhidden")
	alumno.Nss = ctx.PostValue("nss")
	alumno.Tutor = ctx.PostValue("tutor")
	alumno.OcupacionTutor = ctx.PostValue("ocupaciontutor")
	alumno.ParentescoTutor = ctx.PostValue("parentescoTutor")
	alumno.ContactoCasoEmergencia = ctx.PostValue("contactoemergencia")
	alumno.DiferenteDomicilioTutor = ctx.PostValue("difdomtutor")
	alumno.ReferenciasDomicilio = ctx.PostValue("refdomicilio")
	alumno.CorreoE = ctx.PostValue("correoe")

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

		semestres := calificacionesmodel.ExtraeSemestres()

		var primeroprimaria2012 calificacionesmodel.Semestre
		var segundoprimaria2012 calificacionesmodel.Semestre
		var terceroprimaria2012 calificacionesmodel.Semestre
		var cuartoprimaria2012 calificacionesmodel.Semestre
		var quintoprimaria2012 calificacionesmodel.Semestre
		var sextoprimaria2012 calificacionesmodel.Semestre
		var septimoprimaria2012 calificacionesmodel.Semestre
		var octavoprimaria2012 calificacionesmodel.Semestre

		var materiasprimeroprimaria2012 []calificacionesmodel.Materia
		var materiassegundoprimaria2012 []calificacionesmodel.Materia
		var materiasterceroprimaria2012 []calificacionesmodel.Materia
		var materiascuartoprimaria2012 []calificacionesmodel.Materia
		var materiasquintoprimaria2012 []calificacionesmodel.Materia
		var materiassextoprimaria2012 []calificacionesmodel.Materia
		var materiasseptimoprimaria2012 []calificacionesmodel.Materia
		var materiasoctavoprimaria2012 []calificacionesmodel.Materia

		var profesoresprimeroprimaria2012 []string
		var profesoressegundoprimaria2012 []string
		var profesoresterceroprimaria2012 []string
		var profesorescuartoprimaria2012 []string
		var profesoresquintoprimaria2012 []string
		var profesoressextoprimaria2012 []string
		var profesoresseptimoprimaria2012 []string
		var profesoresoctavoprimaria2012 []string

		var hombresprimeroprimaria2012 int //1
		var hombressegundoprimaria2012 int //2
		var hombresterceroprimaria2012 int //3
		var hombrescuartoprimaria2012 int  //4
		var hombresquintoprimaria2012 int  //5
		var hombressextoprimaria2012 int   //6
		var hombresseptimoprimaria2012 int //7
		var hombresoctavoprimaria2012 int  //8

		var mujeresprimeroprimaria2012 int
		var mujeressegundoprimaria2012 int
		var mujeresterceroprimaria2012 int
		var mujerescuartoprimaria2012 int
		var mujeresquintoprimaria2012 int
		var mujeressextoprimaria2012 int
		var mujeresseptimoprimaria2012 int
		var mujeresoctavoprimaria2012 int

		//primaria 2012
		var primeropreescolar2012 calificacionesmodel.Semestre
		var segundopreescolar2012 calificacionesmodel.Semestre
		var terceropreescolar2012 calificacionesmodel.Semestre
		var cuartopreescolar2012 calificacionesmodel.Semestre
		var quintopreescolar2012 calificacionesmodel.Semestre
		var sextopreescolar2012 calificacionesmodel.Semestre
		var septimopreescolar2012 calificacionesmodel.Semestre
		var octavopreescolar2012 calificacionesmodel.Semestre

		var materiasprimeropreescolar2012 []calificacionesmodel.Materia
		var materiassegundopreescolar2012 []calificacionesmodel.Materia
		var materiasterceropreescolar2012 []calificacionesmodel.Materia
		var materiascuartopreescolar2012 []calificacionesmodel.Materia
		var materiasquintopreescolar2012 []calificacionesmodel.Materia
		var materiassextopreescolar2012 []calificacionesmodel.Materia
		var materiasseptimopreescolar2012 []calificacionesmodel.Materia
		var materiasoctavopreescolar2012 []calificacionesmodel.Materia

		var profesoresprimeropreescolar2012 []string
		var profesoressegundopreescolar2012 []string
		var profesoresterceropreescolar2012 []string
		var profesorescuartopreescolar2012 []string
		var profesoresquintopreescolar2012 []string
		var profesoressextopreescolar2012 []string
		var profesoresseptimopreescolar2012 []string
		var profesoresoctavopreescolar2012 []string

		var hombresprimeropreescolar2012 int //1
		var hombressegundopreescolar2012 int //2
		var hombresterceropreescolar2012 int //3
		var hombrescuartopreescolar2012 int  //4
		var hombresquintopreescolar2012 int  //5
		var hombressextopreescolar2012 int   //6
		var hombresseptimopreescolar2012 int //7
		var hombresoctavopreescolar2012 int  //8

		var mujeresprimeropreescolar2012 int
		var mujeressegundopreescolar2012 int
		var mujeresterceropreescolar2012 int
		var mujerescuartopreescolar2012 int
		var mujeresquintopreescolar2012 int
		var mujeressextopreescolar2012 int
		var mujeresseptimopreescolar2012 int
		var mujeresoctavopreescolar2012 int
		//preescolar 2012

		var primeroprimaria2018 calificacionesmodel.Semestre
		var segundoprimaria2018 calificacionesmodel.Semestre
		var terceroprimaria2018 calificacionesmodel.Semestre
		var cuartoprimaria2018 calificacionesmodel.Semestre
		var quintoprimaria2018 calificacionesmodel.Semestre
		var sextoprimaria2018 calificacionesmodel.Semestre
		var septimoprimaria2018 calificacionesmodel.Semestre
		var octavoprimaria2018 calificacionesmodel.Semestre

		var materiasprimeroprimaria2018 []calificacionesmodel.Materia
		var materiassegundoprimaria2018 []calificacionesmodel.Materia
		var materiasterceroprimaria2018 []calificacionesmodel.Materia
		var materiascuartoprimaria2018 []calificacionesmodel.Materia
		var materiasquintoprimaria2018 []calificacionesmodel.Materia
		var materiassextoprimaria2018 []calificacionesmodel.Materia
		var materiasseptimoprimaria2018 []calificacionesmodel.Materia
		var materiasoctavoprimaria2018 []calificacionesmodel.Materia

		var profesoresprimeroprimaria2018 []string
		var profesoressegundoprimaria2018 []string
		var profesoresterceroprimaria2018 []string
		var profesorescuartoprimaria2018 []string
		var profesoresquintoprimaria2018 []string
		var profesoressextoprimaria2018 []string
		var profesoresseptimoprimaria2018 []string
		var profesoresoctavoprimaria2018 []string

		var hombresprimeroprimaria2018 int //1
		var hombressegundoprimaria2018 int //2
		var hombresterceroprimaria2018 int //3
		var hombrescuartoprimaria2018 int  //4
		var hombresquintoprimaria2018 int  //5
		var hombressextoprimaria2018 int   //6
		var hombresseptimoprimaria2018 int //7
		var hombresoctavoprimaria2018 int  //8

		var mujeresprimeroprimaria2018 int
		var mujeressegundoprimaria2018 int
		var mujeresterceroprimaria2018 int
		var mujerescuartoprimaria2018 int
		var mujeresquintoprimaria2018 int
		var mujeressextoprimaria2018 int
		var mujeresseptimoprimaria2018 int
		var mujeresoctavoprimaria2018 int

		//primaria 2018

		var primeropreescolar2018 calificacionesmodel.Semestre
		var segundopreescolar2018 calificacionesmodel.Semestre
		var terceropreescolar2018 calificacionesmodel.Semestre
		var cuartopreescolar2018 calificacionesmodel.Semestre
		var quintopreescolar2018 calificacionesmodel.Semestre
		var sextopreescolar2018 calificacionesmodel.Semestre
		var septimopreescolar2018 calificacionesmodel.Semestre
		var octavopreescolar2018 calificacionesmodel.Semestre

		var materiasprimeropreescolar2018 []calificacionesmodel.Materia
		var materiassegundopreescolar2018 []calificacionesmodel.Materia
		var materiasterceropreescolar2018 []calificacionesmodel.Materia
		var materiascuartopreescolar2018 []calificacionesmodel.Materia
		var materiasquintopreescolar2018 []calificacionesmodel.Materia
		var materiassextopreescolar2018 []calificacionesmodel.Materia
		var materiasseptimopreescolar2018 []calificacionesmodel.Materia
		var materiasoctavopreescolar2018 []calificacionesmodel.Materia

		var profesoresprimeropreescolar2018 []string
		var profesoressegundopreescolar2018 []string
		var profesoresterceropreescolar2018 []string
		var profesorescuartopreescolar2018 []string
		var profesoresquintopreescolar2018 []string
		var profesoressextopreescolar2018 []string
		var profesoresseptimopreescolar2018 []string
		var profesoresoctavopreescolar2018 []string

		var hombresprimeropreescolar2018 int //1
		var hombressegundopreescolar2018 int //2
		var hombresterceropreescolar2018 int //3
		var hombrescuartopreescolar2018 int  //4
		var hombresquintopreescolar2018 int  //5
		var hombressextopreescolar2018 int   //6
		var hombresseptimopreescolar2018 int //7
		var hombresoctavopreescolar2018 int  //8

		var mujeresprimeropreescolar2018 int
		var mujeressegundopreescolar2018 int
		var mujeresterceropreescolar2018 int
		var mujerescuartopreescolar2018 int
		var mujeresquintopreescolar2018 int
		var mujeressextopreescolar2018 int
		var mujeresseptimopreescolar2018 int
		var mujeresoctavopreescolar2018 int
		//preescolar 2018

		var primeroprimaria2021 calificacionesmodel.Semestre
		var segundoprimaria2021 calificacionesmodel.Semestre
		var terceroprimaria2021 calificacionesmodel.Semestre
		var cuartoprimaria2021 calificacionesmodel.Semestre
		var quintoprimaria2021 calificacionesmodel.Semestre
		var sextoprimaria2021 calificacionesmodel.Semestre
		var septimoprimaria2021 calificacionesmodel.Semestre
		var octavoprimaria2021 calificacionesmodel.Semestre

		var materiasprimeroprimaria2021 []calificacionesmodel.Materia
		var materiassegundoprimaria2021 []calificacionesmodel.Materia
		var materiasterceroprimaria2021 []calificacionesmodel.Materia
		var materiascuartoprimaria2021 []calificacionesmodel.Materia
		var materiasquintoprimaria2021 []calificacionesmodel.Materia
		var materiassextoprimaria2021 []calificacionesmodel.Materia
		var materiasseptimoprimaria2021 []calificacionesmodel.Materia
		var materiasoctavoprimaria2021 []calificacionesmodel.Materia

		var profesoresprimeroprimaria2021 []string
		var profesoressegundoprimaria2021 []string
		var profesoresterceroprimaria2021 []string
		var profesorescuartoprimaria2021 []string
		var profesoresquintoprimaria2021 []string
		var profesoressextoprimaria2021 []string
		var profesoresseptimoprimaria2021 []string
		var profesoresoctavoprimaria2021 []string

		var hombresprimeroprimaria2021 int //1
		var hombressegundoprimaria2021 int //2
		var hombresterceroprimaria2021 int //3
		var hombrescuartoprimaria2021 int  //4
		var hombresquintoprimaria2021 int  //5
		var hombressextoprimaria2021 int   //6
		var hombresseptimoprimaria2021 int //7
		var hombresoctavoprimaria2021 int  //8

		var mujeresprimeroprimaria2021 int
		var mujeressegundoprimaria2021 int
		var mujeresterceroprimaria2021 int
		var mujerescuartoprimaria2021 int
		var mujeresquintoprimaria2021 int
		var mujeressextoprimaria2021 int
		var mujeresseptimoprimaria2021 int
		var mujeresoctavoprimaria2021 int

		//primaria 2021

		var primeropreescolar2021 calificacionesmodel.Semestre
		var segundopreescolar2021 calificacionesmodel.Semestre
		var terceropreescolar2021 calificacionesmodel.Semestre
		var cuartopreescolar2021 calificacionesmodel.Semestre
		var quintopreescolar2021 calificacionesmodel.Semestre
		var sextopreescolar2021 calificacionesmodel.Semestre
		var septimopreescolar2021 calificacionesmodel.Semestre
		var octavopreescolar2021 calificacionesmodel.Semestre

		var materiasprimeropreescolar2021 []calificacionesmodel.Materia
		var materiassegundopreescolar2021 []calificacionesmodel.Materia
		var materiasterceropreescolar2021 []calificacionesmodel.Materia
		var materiascuartopreescolar2021 []calificacionesmodel.Materia
		var materiasquintopreescolar2021 []calificacionesmodel.Materia
		var materiassextopreescolar2021 []calificacionesmodel.Materia
		var materiasseptimopreescolar2021 []calificacionesmodel.Materia
		var materiasoctavopreescolar2021 []calificacionesmodel.Materia

		var profesoresprimeropreescolar2021 []string
		var profesoressegundopreescolar2021 []string
		var profesoresterceropreescolar2021 []string
		var profesorescuartopreescolar2021 []string
		var profesoresquintopreescolar2021 []string
		var profesoressextopreescolar2021 []string
		var profesoresseptimopreescolar2021 []string
		var profesoresoctavopreescolar2021 []string

		var hombresprimeropreescolar2021 int //1
		var hombressegundopreescolar2021 int //2
		var hombresterceropreescolar2021 int //3
		var hombrescuartopreescolar2021 int  //4
		var hombresquintopreescolar2021 int  //5
		var hombressextopreescolar2021 int   //6
		var hombresseptimopreescolar2021 int //7
		var hombresoctavopreescolar2021 int  //8

		var mujeresprimeropreescolar2021 int
		var mujeressegundopreescolar2021 int
		var mujeresterceropreescolar2021 int
		var mujerescuartopreescolar2021 int
		var mujeresquintopreescolar2021 int
		var mujeressextopreescolar2021 int
		var mujeresseptimopreescolar2021 int
		var mujeresoctavopreescolar2021 int
		//preescolar 2021

		for _, vv := range semestres {

			switch vv.Semestre {
			case "1":
				switch vv.Plan {
				case "2012":
					switch vv.Licenciatura {
					case "Primaria":
						primeroprimaria2012 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasprimeroprimaria2012 = calificacionesmodel.ExtraeMateriasPorSemestre(primeroprimaria2012.ID)
						profesoresprimeroprimaria2012 = calificacionesmodel.ExtraeDocentesArr(materiasprimeroprimaria2012)
						hombresprimeroprimaria2012, mujeresprimeroprimaria2012 = calificacionesmodel.HombresyMujeres(primeroprimaria2012.ID)
						break
					case "Preescolar":
						primeropreescolar2012 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasprimeropreescolar2012 = calificacionesmodel.ExtraeMateriasPorSemestre(primeropreescolar2012.ID)
						profesoresprimeropreescolar2012 = calificacionesmodel.ExtraeDocentesArr(materiasprimeropreescolar2012)
						hombresprimeropreescolar2012, mujeresprimeropreescolar2012 = calificacionesmodel.HombresyMujeres(primeropreescolar2012.ID)
						break
					}
					break
				case "2018":
					switch vv.Licenciatura {
					case "Primaria":
						primeroprimaria2018 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasprimeroprimaria2018 = calificacionesmodel.ExtraeMateriasPorSemestre(primeroprimaria2018.ID)
						profesoresprimeroprimaria2018 = calificacionesmodel.ExtraeDocentesArr(materiasprimeroprimaria2018)
						hombresprimeroprimaria2018, mujeresprimeroprimaria2018 = calificacionesmodel.HombresyMujeres(primeroprimaria2018.ID)

						break
					case "Preescolar":
						primeropreescolar2018 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasprimeropreescolar2018 = calificacionesmodel.ExtraeMateriasPorSemestre(primeropreescolar2018.ID)
						profesoresprimeropreescolar2018 = calificacionesmodel.ExtraeDocentesArr(materiasprimeropreescolar2018)
						hombresprimeropreescolar2018, mujeresprimeropreescolar2018 = calificacionesmodel.HombresyMujeres(primeropreescolar2018.ID)
						break
					}
					break
				case "2021":
					switch vv.Licenciatura {
					case "Primaria":
						primeroprimaria2021 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasprimeroprimaria2021 = calificacionesmodel.ExtraeMateriasPorSemestre(primeroprimaria2021.ID)
						profesoresprimeroprimaria2021 = calificacionesmodel.ExtraeDocentesArr(materiasprimeroprimaria2021)
						hombresprimeroprimaria2021, mujeresprimeroprimaria2021 = calificacionesmodel.HombresyMujeres(primeroprimaria2021.ID)
						break
					case "Preescolar":
						primeropreescolar2021 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasprimeropreescolar2021 = calificacionesmodel.ExtraeMateriasPorSemestre(primeropreescolar2021.ID)
						profesoresprimeropreescolar2021 = calificacionesmodel.ExtraeDocentesArr(materiasprimeropreescolar2021)
						hombresprimeropreescolar2021, mujeresprimeropreescolar2021 = calificacionesmodel.HombresyMujeres(primeropreescolar2021.ID)
						break
					}
					break
				}

				break
			case "2":
				switch vv.Plan {
				case "2012":
					switch vv.Licenciatura {
					case "Primaria":
						segundoprimaria2012 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiassegundoprimaria2012 = calificacionesmodel.ExtraeMateriasPorSemestre(segundoprimaria2012.ID)
						profesoressegundoprimaria2012 = calificacionesmodel.ExtraeDocentesArr(materiassegundoprimaria2012)
						hombressegundoprimaria2012, mujeressegundoprimaria2012 = calificacionesmodel.HombresyMujeres(segundoprimaria2012.ID)
						break
					case "Preescolar":
						segundopreescolar2012 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiassegundopreescolar2012 = calificacionesmodel.ExtraeMateriasPorSemestre(segundopreescolar2012.ID)
						profesoressegundopreescolar2012 = calificacionesmodel.ExtraeDocentesArr(materiassegundopreescolar2012)
						hombressegundopreescolar2012, mujeressegundopreescolar2012 = calificacionesmodel.HombresyMujeres(segundopreescolar2012.ID)
						break
					}

					break
				case "2018":
					switch vv.Licenciatura {
					case "Primaria":
						segundoprimaria2018 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiassegundoprimaria2018 = calificacionesmodel.ExtraeMateriasPorSemestre(segundoprimaria2018.ID)
						profesoressegundoprimaria2018 = calificacionesmodel.ExtraeDocentesArr(materiassegundoprimaria2018)
						hombressegundoprimaria2018, mujeressegundoprimaria2018 = calificacionesmodel.HombresyMujeres(segundoprimaria2018.ID)
						break
					case "Preescolar":
						segundopreescolar2018 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiassegundopreescolar2018 = calificacionesmodel.ExtraeMateriasPorSemestre(segundopreescolar2018.ID)
						profesoressegundopreescolar2018 = calificacionesmodel.ExtraeDocentesArr(materiassegundopreescolar2018)
						hombressegundopreescolar2018, mujeressegundopreescolar2018 = calificacionesmodel.HombresyMujeres(segundopreescolar2018.ID)
						break
					}

					break
				case "2021":
					switch vv.Licenciatura {
					case "Primaria":
						segundoprimaria2021 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiassegundoprimaria2021 = calificacionesmodel.ExtraeMateriasPorSemestre(segundoprimaria2021.ID)
						profesoressegundoprimaria2021 = calificacionesmodel.ExtraeDocentesArr(materiassegundoprimaria2021)
						hombressegundoprimaria2021, mujeressegundoprimaria2021 = calificacionesmodel.HombresyMujeres(segundoprimaria2021.ID)
						break
					case "Preescolar":
						segundopreescolar2021 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiassegundopreescolar2021 = calificacionesmodel.ExtraeMateriasPorSemestre(segundopreescolar2021.ID)
						profesoressegundopreescolar2021 = calificacionesmodel.ExtraeDocentesArr(materiassegundopreescolar2021)
						hombressegundopreescolar2021, mujeressegundopreescolar2021 = calificacionesmodel.HombresyMujeres(segundopreescolar2021.ID)
						break
					}
					break
				}
				break
			case "3":
				switch vv.Plan {
				case "2012":
					switch vv.Licenciatura {
					case "Primaria":
						terceroprimaria2012 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasterceroprimaria2012 = calificacionesmodel.ExtraeMateriasPorSemestre(terceroprimaria2012.ID)
						profesoresterceroprimaria2012 = calificacionesmodel.ExtraeDocentesArr(materiasterceroprimaria2012)
						hombresterceroprimaria2012, mujeresterceroprimaria2012 = calificacionesmodel.HombresyMujeres(terceroprimaria2012.ID)
						break
					case "Preescolar":
						terceropreescolar2012 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasterceropreescolar2012 = calificacionesmodel.ExtraeMateriasPorSemestre(terceropreescolar2012.ID)
						profesoresterceropreescolar2012 = calificacionesmodel.ExtraeDocentesArr(materiasterceropreescolar2012)
						hombresterceropreescolar2012, mujeresterceropreescolar2012 = calificacionesmodel.HombresyMujeres(terceropreescolar2012.ID)
						break
					}

					break
				case "2018":
					switch vv.Licenciatura {
					case "Primaria":
						terceroprimaria2018 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasterceroprimaria2018 = calificacionesmodel.ExtraeMateriasPorSemestre(terceroprimaria2018.ID)
						profesoresterceroprimaria2018 = calificacionesmodel.ExtraeDocentesArr(materiasterceroprimaria2018)
						hombresterceroprimaria2018, mujeresterceroprimaria2018 = calificacionesmodel.HombresyMujeres(terceroprimaria2018.ID)
						break
					case "Preescolar":
						terceropreescolar2018 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasterceropreescolar2018 = calificacionesmodel.ExtraeMateriasPorSemestre(terceropreescolar2018.ID)
						profesoresterceropreescolar2018 = calificacionesmodel.ExtraeDocentesArr(materiasterceropreescolar2018)
						hombresterceropreescolar2018, mujeresterceropreescolar2018 = calificacionesmodel.HombresyMujeres(terceropreescolar2018.ID)
						break
					}

					break
				case "2021":
					switch vv.Licenciatura {
					case "Primaria":
						terceroprimaria2021 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasterceroprimaria2021 = calificacionesmodel.ExtraeMateriasPorSemestre(terceroprimaria2021.ID)
						profesoresterceroprimaria2021 = calificacionesmodel.ExtraeDocentesArr(materiasterceroprimaria2021)
						hombresterceroprimaria2021, mujeresterceroprimaria2021 = calificacionesmodel.HombresyMujeres(terceroprimaria2021.ID)
						break
					case "Preescolar":
						terceropreescolar2021 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasterceropreescolar2021 = calificacionesmodel.ExtraeMateriasPorSemestre(terceropreescolar2021.ID)
						profesoresterceropreescolar2021 = calificacionesmodel.ExtraeDocentesArr(materiasterceropreescolar2021)
						hombresterceropreescolar2021, mujeresterceropreescolar2021 = calificacionesmodel.HombresyMujeres(terceropreescolar2021.ID)
						break
					}
					break
				}
				break
			case "4":
				switch vv.Plan {
				case "2012":
					switch vv.Licenciatura {
					case "Primaria":
						cuartoprimaria2012 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiascuartoprimaria2012 = calificacionesmodel.ExtraeMateriasPorSemestre(cuartoprimaria2012.ID)
						profesorescuartoprimaria2012 = calificacionesmodel.ExtraeDocentesArr(materiascuartoprimaria2012)
						hombrescuartoprimaria2012, mujerescuartoprimaria2012 = calificacionesmodel.HombresyMujeres(cuartoprimaria2012.ID)
						break
					case "Preescolar":
						cuartopreescolar2012 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiascuartopreescolar2012 = calificacionesmodel.ExtraeMateriasPorSemestre(cuartopreescolar2012.ID)
						profesorescuartopreescolar2012 = calificacionesmodel.ExtraeDocentesArr(materiascuartopreescolar2012)
						hombrescuartopreescolar2012, mujerescuartopreescolar2012 = calificacionesmodel.HombresyMujeres(cuartopreescolar2012.ID)
						break
					}

					break
				case "2018":
					switch vv.Licenciatura {
					case "Primaria":
						cuartoprimaria2018 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiascuartoprimaria2018 = calificacionesmodel.ExtraeMateriasPorSemestre(cuartoprimaria2018.ID)
						profesorescuartoprimaria2018 = calificacionesmodel.ExtraeDocentesArr(materiascuartoprimaria2018)
						hombrescuartoprimaria2018, mujerescuartoprimaria2018 = calificacionesmodel.HombresyMujeres(cuartoprimaria2018.ID)

						break
					case "Preescolar":
						cuartopreescolar2018 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiascuartopreescolar2021 = calificacionesmodel.ExtraeMateriasPorSemestre(cuartopreescolar2021.ID)
						profesorescuartopreescolar2021 = calificacionesmodel.ExtraeDocentesArr(materiascuartopreescolar2021)
						hombrescuartopreescolar2021, mujerescuartopreescolar2021 = calificacionesmodel.HombresyMujeres(cuartopreescolar2021.ID)
						break
					}

					break
				case "2021":
					switch vv.Licenciatura {
					case "Primaria":
						cuartoprimaria2021 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiascuartoprimaria2021 = calificacionesmodel.ExtraeMateriasPorSemestre(cuartoprimaria2021.ID)
						profesorescuartoprimaria2021 = calificacionesmodel.ExtraeDocentesArr(materiascuartoprimaria2021)
						hombrescuartoprimaria2021, mujerescuartoprimaria2021 = calificacionesmodel.HombresyMujeres(cuartoprimaria2021.ID)

						break
					case "Preescolar":
						cuartopreescolar2012 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiascuartopreescolar2018 = calificacionesmodel.ExtraeMateriasPorSemestre(cuartopreescolar2018.ID)
						profesorescuartopreescolar2018 = calificacionesmodel.ExtraeDocentesArr(materiascuartopreescolar2018)
						hombrescuartopreescolar2018, mujerescuartopreescolar2018 = calificacionesmodel.HombresyMujeres(cuartopreescolar2018.ID)
						break
					}
					break
				}
				break
			case "5":
				switch vv.Plan {
				case "2012":
					switch vv.Licenciatura {
					case "Primaria":
						quintoprimaria2012 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasquintoprimaria2012 = calificacionesmodel.ExtraeMateriasPorSemestre(quintoprimaria2012.ID)
						profesoresquintoprimaria2012 = calificacionesmodel.ExtraeDocentesArr(materiasquintoprimaria2012)
						hombresquintoprimaria2012, mujeresquintoprimaria2012 = calificacionesmodel.HombresyMujeres(quintoprimaria2012.ID)
						break
					case "Preescolar":
						quintopreescolar2012 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasquintopreescolar2012 = calificacionesmodel.ExtraeMateriasPorSemestre(quintopreescolar2012.ID)
						profesoresquintopreescolar2012 = calificacionesmodel.ExtraeDocentesArr(materiasquintopreescolar2012)
						hombresquintopreescolar2012, mujeresquintopreescolar2012 = calificacionesmodel.HombresyMujeres(quintopreescolar2012.ID)
						break
					}

					break
				case "2018":
					switch vv.Licenciatura {
					case "Primaria":
						quintoprimaria2018 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasquintoprimaria2018 = calificacionesmodel.ExtraeMateriasPorSemestre(quintoprimaria2018.ID)
						profesoresquintoprimaria2018 = calificacionesmodel.ExtraeDocentesArr(materiasquintoprimaria2018)
						hombresquintoprimaria2018, mujeresquintoprimaria2018 = calificacionesmodel.HombresyMujeres(quintoprimaria2018.ID)

						break
					case "Preescolar":
						quintopreescolar2018 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasquintopreescolar2018 = calificacionesmodel.ExtraeMateriasPorSemestre(quintopreescolar2018.ID)
						profesoresquintopreescolar2018 = calificacionesmodel.ExtraeDocentesArr(materiasquintopreescolar2018)
						hombresquintopreescolar2018, mujeresquintopreescolar2018 = calificacionesmodel.HombresyMujeres(quintopreescolar2018.ID)
						break
					}

					break
				case "2021":
					switch vv.Licenciatura {
					case "Primaria":
						quintoprimaria2021 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasquintoprimaria2021 = calificacionesmodel.ExtraeMateriasPorSemestre(quintoprimaria2021.ID)
						profesoresquintoprimaria2021 = calificacionesmodel.ExtraeDocentesArr(materiasquintoprimaria2021)
						hombresquintoprimaria2021, mujeresquintoprimaria2021 = calificacionesmodel.HombresyMujeres(quintoprimaria2021.ID)

						break
					case "Preescolar":
						quintopreescolar2021 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasquintopreescolar2021 = calificacionesmodel.ExtraeMateriasPorSemestre(quintopreescolar2021.ID)
						profesoresquintopreescolar2021 = calificacionesmodel.ExtraeDocentesArr(materiasquintopreescolar2021)
						hombresquintopreescolar2021, mujeresquintopreescolar2021 = calificacionesmodel.HombresyMujeres(quintopreescolar2021.ID)
						break
					}
					break
				}
				break
			case "6":
				switch vv.Plan {
				case "2012":
					switch vv.Licenciatura {
					case "Primaria":
						sextoprimaria2012 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiassextoprimaria2012 = calificacionesmodel.ExtraeMateriasPorSemestre(sextoprimaria2012.ID)
						profesoressextoprimaria2012 = calificacionesmodel.ExtraeDocentesArr(materiassextoprimaria2012)
						hombressextoprimaria2012, mujeressextoprimaria2012 = calificacionesmodel.HombresyMujeres(sextoprimaria2012.ID)
						break
					case "Preescolar":
						sextopreescolar2012 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiassextopreescolar2012 = calificacionesmodel.ExtraeMateriasPorSemestre(sextopreescolar2012.ID)
						profesoressextopreescolar2012 = calificacionesmodel.ExtraeDocentesArr(materiassextopreescolar2012)
						hombressextopreescolar2012, mujeressextopreescolar2012 = calificacionesmodel.HombresyMujeres(sextopreescolar2012.ID)
						break
					}

					break
				case "2018":
					switch vv.Licenciatura {
					case "Primaria":
						sextoprimaria2018 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiassextoprimaria2018 = calificacionesmodel.ExtraeMateriasPorSemestre(sextoprimaria2018.ID)
						profesoressextoprimaria2018 = calificacionesmodel.ExtraeDocentesArr(materiassextoprimaria2018)
						hombressextoprimaria2018, mujeressextoprimaria2018 = calificacionesmodel.HombresyMujeres(sextoprimaria2018.ID)

						break
					case "Preescolar":
						sextopreescolar2018 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiassextopreescolar2018 = calificacionesmodel.ExtraeMateriasPorSemestre(sextopreescolar2018.ID)
						profesoressextopreescolar2018 = calificacionesmodel.ExtraeDocentesArr(materiassextopreescolar2018)
						hombressextopreescolar2018, mujeressextopreescolar2018 = calificacionesmodel.HombresyMujeres(sextopreescolar2018.ID)
						break
					}

					break
				case "2021":
					switch vv.Licenciatura {
					case "Primaria":
						sextoprimaria2021 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiassextoprimaria2021 = calificacionesmodel.ExtraeMateriasPorSemestre(sextoprimaria2021.ID)
						profesoressextoprimaria2021 = calificacionesmodel.ExtraeDocentesArr(materiassextoprimaria2021)
						hombressextoprimaria2021, mujeressextoprimaria2021 = calificacionesmodel.HombresyMujeres(sextoprimaria2021.ID)
						break
					case "Preescolar":
						sextopreescolar2021 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiassextopreescolar2021 = calificacionesmodel.ExtraeMateriasPorSemestre(sextopreescolar2021.ID)
						profesoressextopreescolar2021 = calificacionesmodel.ExtraeDocentesArr(materiassextopreescolar2021)
						hombressextopreescolar2021, mujeressextopreescolar2021 = calificacionesmodel.HombresyMujeres(sextopreescolar2021.ID)
						break
					}
					break
				}
				break
			case "7":
				switch vv.Plan {
				case "2012":
					switch vv.Licenciatura {
					case "Primaria":
						septimoprimaria2012 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasseptimoprimaria2012 = calificacionesmodel.ExtraeMateriasPorSemestre(septimoprimaria2012.ID)
						profesoresseptimoprimaria2012 = calificacionesmodel.ExtraeDocentesArr(materiasseptimoprimaria2012)
						hombresseptimoprimaria2012, mujeresseptimoprimaria2012 = calificacionesmodel.HombresyMujeres(septimoprimaria2012.ID)
						break
					case "Preescolar":
						septimopreescolar2012 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasseptimopreescolar2012 = calificacionesmodel.ExtraeMateriasPorSemestre(septimopreescolar2012.ID)
						profesoresseptimopreescolar2012 = calificacionesmodel.ExtraeDocentesArr(materiasseptimopreescolar2012)
						hombresseptimopreescolar2012, mujeresseptimopreescolar2012 = calificacionesmodel.HombresyMujeres(septimopreescolar2012.ID)
						break
					}

					break
				case "2018":
					switch vv.Licenciatura {
					case "Primaria":
						septimoprimaria2018 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasseptimoprimaria2018 = calificacionesmodel.ExtraeMateriasPorSemestre(septimoprimaria2018.ID)
						profesoresseptimoprimaria2018 = calificacionesmodel.ExtraeDocentesArr(materiasseptimoprimaria2018)
						hombresseptimoprimaria2018, mujeresseptimoprimaria2018 = calificacionesmodel.HombresyMujeres(septimoprimaria2018.ID)
						break
					case "Preescolar":
						septimopreescolar2018 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasseptimopreescolar2018 = calificacionesmodel.ExtraeMateriasPorSemestre(septimopreescolar2018.ID)
						profesoresseptimopreescolar2018 = calificacionesmodel.ExtraeDocentesArr(materiasseptimopreescolar2018)
						hombresseptimopreescolar2018, mujeresseptimopreescolar2018 = calificacionesmodel.HombresyMujeres(septimopreescolar2018.ID)
						break
					}

					break
				case "2021":
					switch vv.Licenciatura {
					case "Primaria":
						septimoprimaria2021 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasseptimoprimaria2021 = calificacionesmodel.ExtraeMateriasPorSemestre(septimoprimaria2021.ID)
						profesoresseptimoprimaria2021 = calificacionesmodel.ExtraeDocentesArr(materiasseptimoprimaria2021)
						hombresseptimoprimaria2021, mujeresseptimoprimaria2021 = calificacionesmodel.HombresyMujeres(septimoprimaria2021.ID)
						break
					case "Preescolar":
						septimopreescolar2021 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasseptimopreescolar2021 = calificacionesmodel.ExtraeMateriasPorSemestre(septimopreescolar2021.ID)
						profesoresseptimopreescolar2021 = calificacionesmodel.ExtraeDocentesArr(materiasseptimopreescolar2021)
						hombresseptimopreescolar2021, mujeresseptimopreescolar2021 = calificacionesmodel.HombresyMujeres(septimopreescolar2021.ID)
						break
					}
					break
				}
				break
			case "8":
				switch vv.Plan {
				case "2012":
					switch vv.Licenciatura {
					case "Primaria":
						octavoprimaria2012 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasoctavoprimaria2012 = calificacionesmodel.ExtraeMateriasPorSemestre(octavoprimaria2012.ID)
						profesoresoctavoprimaria2012 = calificacionesmodel.ExtraeDocentesArr(materiasoctavoprimaria2012)
						hombresoctavoprimaria2012, mujeresoctavoprimaria2012 = calificacionesmodel.HombresyMujeres(octavoprimaria2012.ID)
						break
					case "Preescolar":
						octavopreescolar2012 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasoctavopreescolar2012 = calificacionesmodel.ExtraeMateriasPorSemestre(octavopreescolar2012.ID)
						profesoresoctavopreescolar2012 = calificacionesmodel.ExtraeDocentesArr(materiasoctavopreescolar2012)
						hombresoctavopreescolar2012, mujeresoctavopreescolar2012 = calificacionesmodel.HombresyMujeres(octavopreescolar2012.ID)
						break
					}

					break
				case "2018":
					switch vv.Licenciatura {
					case "Primaria":
						octavoprimaria2018 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasoctavoprimaria2018 = calificacionesmodel.ExtraeMateriasPorSemestre(octavoprimaria2018.ID)
						profesoresoctavoprimaria2018 = calificacionesmodel.ExtraeDocentesArr(materiasoctavoprimaria2018)
						hombresoctavoprimaria2018, mujeresoctavoprimaria2018 = calificacionesmodel.HombresyMujeres(octavoprimaria2018.ID)
						break
					case "Preescolar":
						octavopreescolar2018 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasoctavopreescolar2018 = calificacionesmodel.ExtraeMateriasPorSemestre(octavopreescolar2018.ID)
						profesoresoctavopreescolar2018 = calificacionesmodel.ExtraeDocentesArr(materiasoctavopreescolar2018)
						hombresoctavopreescolar2018, mujeresoctavopreescolar2018 = calificacionesmodel.HombresyMujeres(octavopreescolar2018.ID)
						break
					}

					break
				case "2021":
					switch vv.Licenciatura {
					case "Primaria":
						octavoprimaria2021 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasoctavoprimaria2021 = calificacionesmodel.ExtraeMateriasPorSemestre(octavoprimaria2021.ID)
						profesoresoctavoprimaria2021 = calificacionesmodel.ExtraeDocentesArr(materiasoctavoprimaria2021)
						hombresoctavoprimaria2021, mujeresoctavoprimaria2021 = calificacionesmodel.HombresyMujeres(octavoprimaria2021.ID)
						break
					case "Preescolar":
						octavopreescolar2021 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasoctavopreescolar2021 = calificacionesmodel.ExtraeMateriasPorSemestre(octavopreescolar2021.ID)
						profesoresoctavopreescolar2021 = calificacionesmodel.ExtraeDocentesArr(materiasoctavopreescolar2021)
						hombresoctavopreescolar2021, mujeresoctavopreescolar2021 = calificacionesmodel.HombresyMujeres(octavopreescolar2021.ID)
						break
					}
					break
				}
				break
			}

		}

		//Aqui comienza la enviadera de datos

		//2012  ---------------------------------------------------------------------
		//1
		//Primaria
		//Materias
		ctx.ViewData("Mpp2012", materiasprimeroprimaria2012)
		//Profesores
		ctx.ViewData("Ppp2012", profesoresprimeroprimaria2012)
		//Hombres y Mujeres
		ctx.ViewData("App2012", hombresprimeroprimaria2012+mujeresprimeroprimaria2012)
		ctx.ViewData("Hpp2012", hombresprimeroprimaria2012)
		ctx.ViewData("Mupp2012", mujeresprimeroprimaria2012)

		//Preescolar
		//Materias
		ctx.ViewData("Mppre2012", materiasprimeropreescolar2012)
		//Profesores
		ctx.ViewData("Pppre2012", profesoresprimeropreescolar2012)
		//Hombres y Mujeres
		ctx.ViewData("Appre2012", hombresprimeropreescolar2012+mujeresprimeropreescolar2012)
		ctx.ViewData("Hppre2012", hombresprimeropreescolar2012)
		ctx.ViewData("Muppre2012", mujeresprimeropreescolar2012)
		//2
		//Primaria
		//Materias
		ctx.ViewData("Msp2012", materiassegundoprimaria2012)
		//Profesores
		ctx.ViewData("Psp2012", profesoressegundoprimaria2012)
		//Hombres y Mujeres
		ctx.ViewData("Asp2012", hombressegundoprimaria2012+mujeressegundoprimaria2012)
		ctx.ViewData("Hsp2012", hombressegundoprimaria2012)
		ctx.ViewData("Musp2012", mujeressegundoprimaria2012)

		//Preescolar
		//Materias
		ctx.ViewData("Mspre2012", materiassegundopreescolar2012)
		//Profesores
		ctx.ViewData("Pspre2012", profesoressegundopreescolar2012)
		//Hombres y Mujeres
		ctx.ViewData("Aspre2012", hombressegundopreescolar2012+mujeressegundopreescolar2012)
		ctx.ViewData("Hspre2012", hombressegundopreescolar2012)
		ctx.ViewData("Muspre2012", mujeressegundopreescolar2012)
		//3
		//Primaria
		//Materias
		ctx.ViewData("Mtp2012", materiasterceroprimaria2012)
		//Profesores
		ctx.ViewData("Ptp2012", profesoresterceroprimaria2012)
		//Hombres y Mujeres
		ctx.ViewData("Atp2012", hombresterceroprimaria2012+mujeresterceroprimaria2012)
		ctx.ViewData("Htp2012", hombresterceroprimaria2012)
		ctx.ViewData("Mutp2012", mujeresterceroprimaria2012)

		//Preescolar
		//Materias
		ctx.ViewData("Mtpre2012", materiasterceropreescolar2012)
		//Profesores
		ctx.ViewData("Ptpre2012", profesoresterceropreescolar2012)
		//Hombres y Mujeres
		ctx.ViewData("Atpre2012", hombresterceropreescolar2012+mujeresterceropreescolar2012)
		ctx.ViewData("Htpre2012", hombresterceropreescolar2012)
		ctx.ViewData("Mutpre2012", mujeresterceropreescolar2012)
		//4
		//Primaria
		//Materias
		ctx.ViewData("Mcp2012", materiascuartoprimaria2012)
		//Profesores
		ctx.ViewData("Pcp2012", profesorescuartoprimaria2012)
		//Hombres y Mujeres
		ctx.ViewData("Acp2012", hombrescuartoprimaria2012+mujerescuartoprimaria2012)
		ctx.ViewData("Hcp2012", hombrescuartoprimaria2012)
		ctx.ViewData("Mucp2012", mujerescuartoprimaria2012)

		//Preescolar
		//Materias
		ctx.ViewData("Mcpre2012", materiascuartopreescolar2012)
		//Profesores
		ctx.ViewData("Pcpre2012", profesorescuartopreescolar2012)
		//Hombres y Mujeres
		ctx.ViewData("Acpre2012", hombrescuartopreescolar2012+mujerescuartopreescolar2012)
		ctx.ViewData("Hcpre2012", hombrescuartopreescolar2012)
		ctx.ViewData("Mucpre2012", mujerescuartopreescolar2012)
		//5
		//Primaria
		//Materias
		ctx.ViewData("Mqp2012", materiasquintoprimaria2012)
		//Profesores
		ctx.ViewData("Pqp2012", profesoresquintoprimaria2012)
		//Hombres y Mujeres
		ctx.ViewData("Aqp2012", hombresquintoprimaria2012+mujeresquintoprimaria2012)
		ctx.ViewData("Hqp2012", hombresquintoprimaria2012)
		ctx.ViewData("Muqp2012", mujeresquintoprimaria2012)

		//Preescolar
		//Materias
		ctx.ViewData("Mqpre2012", materiasquintopreescolar2012)
		//Profesores
		ctx.ViewData("Pqpre2012", profesoresquintopreescolar2012)
		//Hombres y Mujeres
		ctx.ViewData("Aqpre2012", hombresquintopreescolar2012+mujeresquintopreescolar2012)
		ctx.ViewData("Hqpre2012", hombresquintopreescolar2012)
		ctx.ViewData("Muqpre2012", mujeresquintopreescolar2012)
		//6
		//Primaria
		//Materias
		ctx.ViewData("Msxp2012", materiassextoprimaria2012)
		//Profesores
		ctx.ViewData("Psxp2012", profesoressextoprimaria2012)
		//Hombres y Mujeres
		ctx.ViewData("Asxp2012", hombressextoprimaria2012+mujeressextoprimaria2012)
		ctx.ViewData("Hsxp2012", hombressextoprimaria2012)
		ctx.ViewData("Musxp2012", mujeressextoprimaria2012)

		//Preescolar
		//Materias
		ctx.ViewData("Msxpre2012", materiassextopreescolar2012)
		//Profesores
		ctx.ViewData("Psxpre2012", profesoressextopreescolar2012)
		//Hombres y Mujeres
		ctx.ViewData("Asxpre2012", hombressextopreescolar2012+mujeressextopreescolar2012)
		ctx.ViewData("Hsxpre2012", hombressextopreescolar2012)
		ctx.ViewData("Musxpre2012", mujeressextopreescolar2012)
		//7
		//Primaria
		//Materias
		ctx.ViewData("Mstp2012", materiasseptimoprimaria2012)
		//Profesores
		ctx.ViewData("Pstp2012", profesoresseptimoprimaria2012)
		//Hombres y Mujeres
		ctx.ViewData("Astp2012", hombresseptimoprimaria2012+mujeresseptimoprimaria2012)
		ctx.ViewData("Hstp2012", hombresseptimoprimaria2012)
		ctx.ViewData("Mustp2012", mujeresseptimoprimaria2012)

		//Preescolar
		//Materias
		ctx.ViewData("Mstpre2012", materiasseptimopreescolar2012)
		//Profesores
		ctx.ViewData("Pstpre2012", profesoresseptimopreescolar2012)
		//Hombres y Mujeres
		ctx.ViewData("Astpre2012", hombresseptimopreescolar2012+mujeresseptimopreescolar2012)
		ctx.ViewData("Hstpre2012", hombresseptimopreescolar2012)
		ctx.ViewData("Mustpre2012", mujeresseptimopreescolar2012)
		//8
		//Primaria
		//Materias
		ctx.ViewData("Mop2012", materiasoctavoprimaria2012)
		//Profesores
		ctx.ViewData("Pop2012", profesoresoctavoprimaria2012)
		//Hombres y Mujeres
		ctx.ViewData("Aop2012", hombresoctavoprimaria2012+mujeresoctavoprimaria2012)
		ctx.ViewData("Hop2012", hombresoctavoprimaria2012)
		ctx.ViewData("Muop2012", mujeresoctavoprimaria2012)

		//Preescolar
		//Materias
		ctx.ViewData("Mopre2012", materiasoctavopreescolar2012)
		//Profesores
		ctx.ViewData("Popre2012", profesoresoctavopreescolar2012)
		//Hombres y Mujeres
		ctx.ViewData("Aopre2012", hombresoctavopreescolar2012+mujeresoctavopreescolar2012)
		ctx.ViewData("Hopre2012", hombresoctavopreescolar2012)
		ctx.ViewData("Muopre2012", mujeresoctavopreescolar2012)
		//2018 ----------------------------------------------------------------------
		//1
		//Primaria
		//Materias
		ctx.ViewData("Mpp2018", materiasprimeroprimaria2018)
		//Profesores
		ctx.ViewData("Ppp2018", profesoresprimeroprimaria2018)
		//Hombres y Mujeres
		ctx.ViewData("App2018", hombresprimeroprimaria2018+mujeresprimeroprimaria2018)
		ctx.ViewData("Hpp2018", hombresprimeroprimaria2018)
		ctx.ViewData("Mupp2018", mujeresprimeroprimaria2018)

		//Preescolar
		//Materias
		ctx.ViewData("Mppre2018", materiasprimeropreescolar2018)
		//Profesores
		ctx.ViewData("Pppre2018", profesoresprimeropreescolar2018)
		//Hombres y Mujeres
		ctx.ViewData("Appre2018", hombresprimeropreescolar2018+mujeresprimeropreescolar2018)
		ctx.ViewData("Hppre2018", hombresprimeropreescolar2018)
		ctx.ViewData("Muppre2018", mujeresprimeropreescolar2018)
		//2
		//Primaria
		//Materias
		ctx.ViewData("Msp2018", materiassegundoprimaria2018)
		//Profesores
		ctx.ViewData("Psp2018", profesoressegundoprimaria2018)
		//Hombres y Mujeres
		ctx.ViewData("Asp2018", hombressegundoprimaria2018+mujeressegundoprimaria2018)
		ctx.ViewData("Hsp2018", hombressegundoprimaria2018)
		ctx.ViewData("Musp2018", mujeressegundoprimaria2018)

		//Preescolar
		//Materias
		ctx.ViewData("Mspre2018", materiassegundopreescolar2018)
		//Profesores
		ctx.ViewData("Pspre2018", profesoressegundopreescolar2018)
		//Hombres y Mujeres
		ctx.ViewData("Aspre2018", hombressegundopreescolar2018+mujeressegundopreescolar2018)
		ctx.ViewData("Hspre2018", hombressegundopreescolar2018)
		ctx.ViewData("Muspre2018", mujeressegundopreescolar2018)
		//3
		//Primaria
		//Materias
		ctx.ViewData("Mtp2018", materiasterceroprimaria2018)
		//Profesores
		ctx.ViewData("Ptp2018", profesoresterceroprimaria2018)
		//Hombres y Mujeres
		ctx.ViewData("Atp2018", hombresterceroprimaria2018+mujeresterceroprimaria2018)
		ctx.ViewData("Htp2018", hombresterceroprimaria2018)
		ctx.ViewData("Mutp2018", mujeresterceroprimaria2018)

		//Preescolar
		//Materias
		ctx.ViewData("Mtpre2018", materiasterceropreescolar2018)
		//Profesores
		ctx.ViewData("Ptpre2018", profesoresterceropreescolar2018)
		//Hombres y Mujeres
		ctx.ViewData("Atpre2018", hombresterceropreescolar2018+mujeresterceropreescolar2018)
		ctx.ViewData("Htpre2018", hombresterceropreescolar2018)
		ctx.ViewData("Mutpre2018", mujeresterceropreescolar2018)
		//4
		//Primaria
		//Materias
		ctx.ViewData("Mcp2018", materiascuartoprimaria2018)
		//Profesores
		ctx.ViewData("Pcp2018", profesorescuartoprimaria2018)
		//Hombres y Mujeres
		ctx.ViewData("Acp2018", hombrescuartoprimaria2018+mujerescuartoprimaria2018)
		ctx.ViewData("Hcp2018", hombrescuartoprimaria2018)
		ctx.ViewData("Mucp2018", mujerescuartoprimaria2018)

		//Preescolar
		//Materias
		ctx.ViewData("Mcpre2018", materiascuartopreescolar2018)
		//Profesores
		ctx.ViewData("Pcpre2018", profesorescuartopreescolar2018)
		//Hombres y Mujeres
		ctx.ViewData("Acpre2018", hombrescuartopreescolar2018+mujerescuartopreescolar2018)
		ctx.ViewData("Hcpre2018", hombrescuartopreescolar2018)
		ctx.ViewData("Mucpre2018", mujerescuartopreescolar2018)
		//5
		//Primaria
		//Materias
		ctx.ViewData("Mqp2018", materiasquintoprimaria2018)
		//Profesores
		ctx.ViewData("Pqp2018", profesoresquintoprimaria2018)
		//Hombres y Mujeres
		ctx.ViewData("Aqp2018", hombresquintoprimaria2018+mujeresquintoprimaria2018)
		ctx.ViewData("Hqp2018", hombresquintoprimaria2018)
		ctx.ViewData("Muqp2018", mujeresquintoprimaria2018)

		//Preescolar
		//Materias
		ctx.ViewData("Mqpre2018", materiasquintopreescolar2018)
		//Profesores
		ctx.ViewData("Pqpre2018", profesoresquintopreescolar2018)
		//Hombres y Mujeres
		ctx.ViewData("Aqpre2018", hombresquintopreescolar2018+mujeresquintopreescolar2018)
		ctx.ViewData("Hqpre2018", hombresquintopreescolar2018)
		ctx.ViewData("Muqpre2018", mujeresquintopreescolar2018)
		//6
		//Primaria
		//Materias
		ctx.ViewData("Msxp2018", materiassextoprimaria2018)
		//Profesores
		ctx.ViewData("Psxp2018", profesoressextoprimaria2018)
		//Hombres y Mujeres
		ctx.ViewData("Hsxp2018", hombressextoprimaria2018+mujeressextoprimaria2018)
		ctx.ViewData("Hsxp2018", hombressextoprimaria2018)
		ctx.ViewData("Musxp2018", mujeressextoprimaria2018)

		//Preescolar
		//Materias
		ctx.ViewData("Msxpre2018", materiassextopreescolar2018)
		//Profesores
		ctx.ViewData("Psxpre2018", profesoressextopreescolar2018)
		//Hombres y Mujeres
		ctx.ViewData("Asxpre2018", hombressextopreescolar2018+mujeressextopreescolar2018)
		ctx.ViewData("Hsxpre2018", hombressextopreescolar2018)
		ctx.ViewData("Musxpre2018", mujeressextopreescolar2018)
		//7
		//Primaria
		//Materias
		ctx.ViewData("Mstp2018", materiasseptimoprimaria2018)
		//Profesores
		ctx.ViewData("Pstp2018", profesoresseptimoprimaria2018)
		//Hombres y Mujeres
		ctx.ViewData("Astp2018", hombresseptimoprimaria2018+mujeresseptimoprimaria2018)
		ctx.ViewData("Hstp2018", hombresseptimoprimaria2018)
		ctx.ViewData("Mustp2018", mujeresseptimoprimaria2018)

		//Preescolar
		//Materias
		ctx.ViewData("Mstpre2018", materiasseptimopreescolar2018)
		//Profesores
		ctx.ViewData("Pstpre2018", profesoresseptimopreescolar2018)
		//Hombres y Mujeres
		ctx.ViewData("Astpre2018", hombresseptimopreescolar2018+mujeresseptimopreescolar2018)
		ctx.ViewData("Hstpre2018", hombresseptimopreescolar2018)
		ctx.ViewData("Mustpre2018", mujeresseptimopreescolar2018)
		//8
		//Primaria
		//Materias
		ctx.ViewData("Mop2018", materiasoctavoprimaria2018)
		//Profesores
		ctx.ViewData("Pop2018", profesoresoctavoprimaria2018)
		//Hombres y Mujeres
		ctx.ViewData("Aop2018", hombresoctavoprimaria2018+mujeresoctavoprimaria2018)
		ctx.ViewData("Hop2018", hombresoctavoprimaria2018)
		ctx.ViewData("Muop2018", mujeresoctavoprimaria2018)

		//Preescolar
		//Materias
		ctx.ViewData("Mopre2018", materiasoctavopreescolar2018)
		//Profesores
		ctx.ViewData("Popre2018", profesoresoctavopreescolar2018)
		//Hombres y Mujeres
		ctx.ViewData("Aopre2018", hombresoctavopreescolar2018+mujeresoctavopreescolar2018)
		ctx.ViewData("Hopre2018", hombresoctavopreescolar2018)
		ctx.ViewData("Muopre2018", mujeresoctavopreescolar2018)
		//2021 ----------------------------------------------------------------------
		//1
		//Primaria
		//Materias
		ctx.ViewData("Mpp2021", materiasprimeroprimaria2021)
		//Profesores
		ctx.ViewData("Ppp2021", profesoresprimeroprimaria2021)
		//Hombres y Mujeres
		ctx.ViewData("App2021", hombresprimeroprimaria2021+mujeresprimeroprimaria2021)
		ctx.ViewData("Hpp2021", hombresprimeroprimaria2021)
		ctx.ViewData("Mupp2021", mujeresprimeroprimaria2021)

		//Preescolar
		//Materias
		ctx.ViewData("Mppre2021", materiasprimeropreescolar2021)
		//Profesores
		ctx.ViewData("Pppre2021", profesoresprimeropreescolar2021)
		//Hombres y Mujeres
		ctx.ViewData("Appre2021", hombresprimeropreescolar2021+mujeresprimeropreescolar2021)
		ctx.ViewData("Hppre2021", hombresprimeropreescolar2021)
		ctx.ViewData("Muppre2021", mujeresprimeropreescolar2021)
		//2
		//Primaria
		//Materias
		ctx.ViewData("Msp2021", materiassegundoprimaria2021)
		//Profesores
		ctx.ViewData("Psp2021", profesoressegundoprimaria2021)
		//Hombres y Mujeres
		ctx.ViewData("Asp2021", hombressegundoprimaria2021+mujeressegundoprimaria2021)
		ctx.ViewData("Hsp2021", hombressegundoprimaria2021)
		ctx.ViewData("Musp2021", mujeressegundoprimaria2021)

		//Preescolar
		//Materias
		ctx.ViewData("Mspre2021", materiassegundopreescolar2021)
		//Profesores
		ctx.ViewData("Pspre2021", profesoressegundopreescolar2021)
		//Hombres y Mujeres
		ctx.ViewData("Aspre2021", hombressegundopreescolar2021+mujeressegundopreescolar2021)
		ctx.ViewData("Hspre2021", hombressegundopreescolar2021)
		ctx.ViewData("Muspre2021", mujeressegundopreescolar2021)
		//3
		//Primaria
		//Materias
		ctx.ViewData("Mtp2021", materiasterceroprimaria2021)
		//Profesores
		ctx.ViewData("Ptp2021", profesoresterceroprimaria2021)
		//Hombres y Mujeres
		ctx.ViewData("Atp2021", hombresterceroprimaria2021+mujeresterceroprimaria2021)
		ctx.ViewData("Htp2021", hombresterceroprimaria2021)
		ctx.ViewData("Mutp2021", mujeresterceroprimaria2021)

		//Preescolar
		//Materias
		ctx.ViewData("Mtpre2021", materiasterceropreescolar2021)
		//Profesores
		ctx.ViewData("Ptpre2021", profesoresterceropreescolar2021)
		//Hombres y Mujeres
		ctx.ViewData("Atpre2021", hombresterceropreescolar2021+mujeresterceropreescolar2021)
		ctx.ViewData("Htpre2021", hombresterceropreescolar2021)
		ctx.ViewData("Mutpre2021", mujeresterceropreescolar2021)
		//4
		//Primaria
		//Materias
		ctx.ViewData("Mcp2021", materiascuartoprimaria2021)
		//Profesores
		ctx.ViewData("Pcp2021", profesorescuartoprimaria2021)
		//Hombres y Mujeres
		ctx.ViewData("Acp2021", hombrescuartoprimaria2021+mujerescuartoprimaria2021)
		ctx.ViewData("Hcp2021", hombrescuartoprimaria2021)
		ctx.ViewData("Mucp2021", mujerescuartoprimaria2021)

		//Preescolar
		//Materias
		ctx.ViewData("Mcpre2021", materiascuartopreescolar2021)
		//Profesores
		ctx.ViewData("Pcpre2021", profesorescuartopreescolar2021)
		//Hombres y Mujeres
		ctx.ViewData("Acpre2021", hombrescuartopreescolar2021+mujerescuartopreescolar2021)
		ctx.ViewData("Hcpre2021", hombrescuartopreescolar2021)
		ctx.ViewData("Mucpre2021", mujerescuartopreescolar2021)
		//5
		//Primaria
		//Materias
		ctx.ViewData("Mqp2021", materiasquintoprimaria2021)
		//Profesores
		ctx.ViewData("Pqp2021", profesoresquintoprimaria2021)
		//Hombres y Mujeres
		ctx.ViewData("Aqp2021", hombresquintoprimaria2021+mujeresquintoprimaria2021)
		ctx.ViewData("Hqp2021", hombresquintoprimaria2021)
		ctx.ViewData("Muqp2021", mujeresquintoprimaria2021)

		//Preescolar
		//Materias
		ctx.ViewData("Mqpre2021", materiasquintopreescolar2021)
		//Profesores
		ctx.ViewData("Pqpre2021", profesoresquintopreescolar2021)
		//Hombres y Mujeres
		ctx.ViewData("Aqpre2021", hombresquintopreescolar2021+mujeresquintopreescolar2021)
		ctx.ViewData("Hqpre2021", hombresquintopreescolar2021)
		ctx.ViewData("Muqpre2021", mujeresquintopreescolar2021)
		//6
		//Primaria
		//Materias
		ctx.ViewData("Msxp2021", materiassextoprimaria2021)
		//Profesores
		ctx.ViewData("Psxp2021", profesoressextoprimaria2021)
		//Hombres y Mujeres
		ctx.ViewData("Asxp2021", hombressextoprimaria2021+mujeressextoprimaria2021)
		ctx.ViewData("Hsxp2021", hombressextoprimaria2021)
		ctx.ViewData("Musxp2021", mujeressextoprimaria2021)

		//Preescolar
		//Materias
		ctx.ViewData("Msxpre2021", materiassextopreescolar2021)
		//Profesores
		ctx.ViewData("Psxpre2021", profesoressextopreescolar2021)
		//Hombres y Mujeres
		ctx.ViewData("Asxpre2021", hombressextopreescolar2021+mujeressextopreescolar2021)
		ctx.ViewData("Hsxpre2021", hombressextopreescolar2021)
		ctx.ViewData("Musxpre2021", mujeressextopreescolar2021)
		//7
		//Primaria
		//Materias
		ctx.ViewData("Mstp2021", materiasseptimoprimaria2021)
		//Profesores
		ctx.ViewData("Pstp2021", profesoresseptimoprimaria2021)
		//Hombres y Mujeres
		ctx.ViewData("Astp2021", hombresseptimoprimaria2021+mujeresseptimoprimaria2021)
		ctx.ViewData("Hstp2021", hombresseptimoprimaria2021)
		ctx.ViewData("Mustp2021", mujeresseptimoprimaria2021)

		//Preescolar
		//Materias
		ctx.ViewData("Mstpre2021", materiasseptimopreescolar2021)
		//Profesores
		ctx.ViewData("Pstpre2021", profesoresseptimopreescolar2021)
		//Hombres y Mujeres
		ctx.ViewData("Astpre2021", hombresseptimopreescolar2021+mujeresseptimopreescolar2021)
		ctx.ViewData("Hstpre2021", hombresseptimopreescolar2021)
		ctx.ViewData("Mustpre2021", mujeresseptimopreescolar2021)
		//8
		//Primaria
		//Materias
		ctx.ViewData("Mop2021", materiasoctavoprimaria2021)
		//Profesores
		ctx.ViewData("Pop2021", profesoresoctavoprimaria2021)
		//Hombres y Mujeres
		ctx.ViewData("Aop2021", hombresoctavoprimaria2021+mujeresoctavoprimaria2021)
		ctx.ViewData("Hop2021", hombresoctavoprimaria2021)
		ctx.ViewData("Muop2021", mujeresoctavoprimaria2021)

		//Preescolar
		//Materias
		ctx.ViewData("Mopre2021", materiasoctavopreescolar2021)
		//Profesores
		ctx.ViewData("Popre2021", profesoresoctavopreescolar2021)
		//Hombres y Mujeres
		ctx.ViewData("Aopre2021", hombresoctavopreescolar2021+mujeresoctavopreescolar2021)
		ctx.ViewData("Hopre2021", hombresoctavopreescolar2021)
		ctx.ViewData("Muopre2021", mujeresoctavopreescolar2021)

		//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

		fmt.Println("Bienvenido ", userOn.Nombre)
		if err := ctx.View("Semestres.html"); err != nil {
			ctx.Application().Logger().Infof(err.Error())
		}
	} else {
		ctx.Redirect("/login", iris.StatusSeeOther)
	}
}
