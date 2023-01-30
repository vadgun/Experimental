package indexcontroller

import (
	"fmt"
	"time"

	"github.com/kataras/iris/v12"
	sessioncontroller "github.com/vadgun/Experimental/Controladores/SessionController"
	calificacionesmodel "github.com/vadgun/Experimental/Modelos/CalificacionesModel"
	indexmodel "github.com/vadgun/Experimental/Modelos/IndexModel"
)

// Index -> Regresa la pagina de inicio
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

// Directorio -> Regresa la pagina de inicio
func Directorio(ctx iris.Context) {
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

		tienepermiso := indexmodel.TienePermiso(5, userOn, usuario)

		if !tienepermiso {
			ctx.Redirect("/login", iris.StatusSeeOther)
		}

		if userOn.Docente || usuario.Docente {
			ctx.Redirect("/login", iris.StatusSeeOther)

		}

		if userOn.Admin || usuario.Admin {
			// enviar docentes para ejecutar algo similar a lo de arriba, enviar traer materias para ver calificaciones y evaluar

		}

		if userOn.Alumno || usuario.Alumno {
			ctx.Redirect("/login", iris.StatusSeeOther)
		}

		if err := ctx.View("Directorio.html"); err != nil {
			ctx.Application().Logger().Infof(err.Error())
		}
	} else {
		ctx.Redirect("/login", iris.StatusSeeOther)
	}
}

// Reloj -> Regresa la pagina de inicio
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

// EditarDatosDeAlumno -> Editar datos en un modal?
func EditarDatosDeAlumno(ctx iris.Context) {
	idstalumno := ctx.PostValue("data")
	var alumno calificacionesmodel.Alumno
	alumno = calificacionesmodel.ExtraeAlumno(idstalumno)
	ctx.JSON(alumno)
}

// EditarAlumno -> Guarda los datos modificados del alumno previamente solicitado
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

// Semestres -> Regresa la pagina de semestres
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

		var primeroprimaria2022 calificacionesmodel.Semestre
		var segundoprimaria2022 calificacionesmodel.Semestre
		var terceroprimaria2022 calificacionesmodel.Semestre
		var cuartoprimaria2022 calificacionesmodel.Semestre
		var quintoprimaria2022 calificacionesmodel.Semestre
		var sextoprimaria2022 calificacionesmodel.Semestre
		var septimoprimaria2022 calificacionesmodel.Semestre
		var octavoprimaria2022 calificacionesmodel.Semestre

		var materiasprimeroprimaria2022 []calificacionesmodel.Materia
		var materiassegundoprimaria2022 []calificacionesmodel.Materia
		var materiasterceroprimaria2022 []calificacionesmodel.Materia
		var materiascuartoprimaria2022 []calificacionesmodel.Materia
		var materiasquintoprimaria2022 []calificacionesmodel.Materia
		var materiassextoprimaria2022 []calificacionesmodel.Materia
		var materiasseptimoprimaria2022 []calificacionesmodel.Materia
		var materiasoctavoprimaria2022 []calificacionesmodel.Materia

		var profesoresprimeroprimaria2022 []string
		var profesoressegundoprimaria2022 []string
		var profesoresterceroprimaria2022 []string
		var profesorescuartoprimaria2022 []string
		var profesoresquintoprimaria2022 []string
		var profesoressextoprimaria2022 []string
		var profesoresseptimoprimaria2022 []string
		var profesoresoctavoprimaria2022 []string

		var hombresprimeroprimaria2022 int //1
		var hombressegundoprimaria2022 int //2
		var hombresterceroprimaria2022 int //3
		var hombrescuartoprimaria2022 int  //4
		var hombresquintoprimaria2022 int  //5
		var hombressextoprimaria2022 int   //6
		var hombresseptimoprimaria2022 int //7
		var hombresoctavoprimaria2022 int  //8

		var mujeresprimeroprimaria2022 int
		var mujeressegundoprimaria2022 int
		var mujeresterceroprimaria2022 int
		var mujerescuartoprimaria2022 int
		var mujeresquintoprimaria2022 int
		var mujeressextoprimaria2022 int
		var mujeresseptimoprimaria2022 int
		var mujeresoctavoprimaria2022 int

		//primaria 2022

		var primeropreescolar2022 calificacionesmodel.Semestre
		var segundopreescolar2022 calificacionesmodel.Semestre
		var terceropreescolar2022 calificacionesmodel.Semestre
		var cuartopreescolar2022 calificacionesmodel.Semestre
		var quintopreescolar2022 calificacionesmodel.Semestre
		var sextopreescolar2022 calificacionesmodel.Semestre
		var septimopreescolar2022 calificacionesmodel.Semestre
		var octavopreescolar2022 calificacionesmodel.Semestre

		var materiasprimeropreescolar2022 []calificacionesmodel.Materia
		var materiassegundopreescolar2022 []calificacionesmodel.Materia
		var materiasterceropreescolar2022 []calificacionesmodel.Materia
		var materiascuartopreescolar2022 []calificacionesmodel.Materia
		var materiasquintopreescolar2022 []calificacionesmodel.Materia
		var materiassextopreescolar2022 []calificacionesmodel.Materia
		var materiasseptimopreescolar2022 []calificacionesmodel.Materia
		var materiasoctavopreescolar2022 []calificacionesmodel.Materia

		var profesoresprimeropreescolar2022 []string
		var profesoressegundopreescolar2022 []string
		var profesoresterceropreescolar2022 []string
		var profesorescuartopreescolar2022 []string
		var profesoresquintopreescolar2022 []string
		var profesoressextopreescolar2022 []string
		var profesoresseptimopreescolar2022 []string
		var profesoresoctavopreescolar2022 []string

		var hombresprimeropreescolar2022 int //1
		var hombressegundopreescolar2022 int //2
		var hombresterceropreescolar2022 int //3
		var hombrescuartopreescolar2022 int  //4
		var hombresquintopreescolar2022 int  //5
		var hombressextopreescolar2022 int   //6
		var hombresseptimopreescolar2022 int //7
		var hombresoctavopreescolar2022 int  //8

		var mujeresprimeropreescolar2022 int
		var mujeressegundopreescolar2022 int
		var mujeresterceropreescolar2022 int
		var mujerescuartopreescolar2022 int
		var mujeresquintopreescolar2022 int
		var mujeressextopreescolar2022 int
		var mujeresseptimopreescolar2022 int
		var mujeresoctavopreescolar2022 int
		//preescolar 2022

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
				case "2022":
					switch vv.Licenciatura {
					case "Primaria":
						primeroprimaria2022 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasprimeroprimaria2022 = calificacionesmodel.ExtraeMateriasPorSemestre(primeroprimaria2022.ID)
						profesoresprimeroprimaria2022 = calificacionesmodel.ExtraeDocentesArr(materiasprimeroprimaria2022)
						hombresprimeroprimaria2022, mujeresprimeroprimaria2022 = calificacionesmodel.HombresyMujeres(primeroprimaria2022.ID)
						break
					case "Preescolar":
						primeropreescolar2022 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasprimeropreescolar2022 = calificacionesmodel.ExtraeMateriasPorSemestre(primeropreescolar2022.ID)
						profesoresprimeropreescolar2022 = calificacionesmodel.ExtraeDocentesArr(materiasprimeropreescolar2022)
						hombresprimeropreescolar2022, mujeresprimeropreescolar2022 = calificacionesmodel.HombresyMujeres(primeropreescolar2022.ID)
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
				case "2022":
					switch vv.Licenciatura {
					case "Primaria":
						segundoprimaria2022 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiassegundoprimaria2022 = calificacionesmodel.ExtraeMateriasPorSemestre(segundoprimaria2022.ID)
						profesoressegundoprimaria2022 = calificacionesmodel.ExtraeDocentesArr(materiassegundoprimaria2022)
						hombressegundoprimaria2022, mujeressegundoprimaria2022 = calificacionesmodel.HombresyMujeres(segundoprimaria2022.ID)
						break
					case "Preescolar":
						segundopreescolar2022 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiassegundopreescolar2022 = calificacionesmodel.ExtraeMateriasPorSemestre(segundopreescolar2022.ID)
						profesoressegundopreescolar2022 = calificacionesmodel.ExtraeDocentesArr(materiassegundopreescolar2022)
						hombressegundopreescolar2022, mujeressegundopreescolar2022 = calificacionesmodel.HombresyMujeres(segundopreescolar2022.ID)
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
				case "2022":
					switch vv.Licenciatura {
					case "Primaria":
						terceroprimaria2022 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasterceroprimaria2022 = calificacionesmodel.ExtraeMateriasPorSemestre(terceroprimaria2022.ID)
						profesoresterceroprimaria2022 = calificacionesmodel.ExtraeDocentesArr(materiasterceroprimaria2022)
						hombresterceroprimaria2022, mujeresterceroprimaria2022 = calificacionesmodel.HombresyMujeres(terceroprimaria2022.ID)
						break
					case "Preescolar":
						terceropreescolar2022 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasterceropreescolar2022 = calificacionesmodel.ExtraeMateriasPorSemestre(terceropreescolar2022.ID)
						profesoresterceropreescolar2022 = calificacionesmodel.ExtraeDocentesArr(materiasterceropreescolar2022)
						hombresterceropreescolar2022, mujeresterceropreescolar2022 = calificacionesmodel.HombresyMujeres(terceropreescolar2022.ID)
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
						materiascuartopreescolar2022 = calificacionesmodel.ExtraeMateriasPorSemestre(cuartopreescolar2022.ID)
						profesorescuartopreescolar2022 = calificacionesmodel.ExtraeDocentesArr(materiascuartopreescolar2022)
						hombrescuartopreescolar2022, mujerescuartopreescolar2022 = calificacionesmodel.HombresyMujeres(cuartopreescolar2022.ID)
						break
					}

					break
				case "2022":
					switch vv.Licenciatura {
					case "Primaria":
						cuartoprimaria2022 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiascuartoprimaria2022 = calificacionesmodel.ExtraeMateriasPorSemestre(cuartoprimaria2022.ID)
						profesorescuartoprimaria2022 = calificacionesmodel.ExtraeDocentesArr(materiascuartoprimaria2022)
						hombrescuartoprimaria2022, mujerescuartoprimaria2022 = calificacionesmodel.HombresyMujeres(cuartoprimaria2022.ID)

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
				case "2022":
					switch vv.Licenciatura {
					case "Primaria":
						quintoprimaria2022 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasquintoprimaria2022 = calificacionesmodel.ExtraeMateriasPorSemestre(quintoprimaria2022.ID)
						profesoresquintoprimaria2022 = calificacionesmodel.ExtraeDocentesArr(materiasquintoprimaria2022)
						hombresquintoprimaria2022, mujeresquintoprimaria2022 = calificacionesmodel.HombresyMujeres(quintoprimaria2022.ID)

						break
					case "Preescolar":
						quintopreescolar2022 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasquintopreescolar2022 = calificacionesmodel.ExtraeMateriasPorSemestre(quintopreescolar2022.ID)
						profesoresquintopreescolar2022 = calificacionesmodel.ExtraeDocentesArr(materiasquintopreescolar2022)
						hombresquintopreescolar2022, mujeresquintopreescolar2022 = calificacionesmodel.HombresyMujeres(quintopreescolar2022.ID)
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
				case "2022":
					switch vv.Licenciatura {
					case "Primaria":
						sextoprimaria2022 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiassextoprimaria2022 = calificacionesmodel.ExtraeMateriasPorSemestre(sextoprimaria2022.ID)
						profesoressextoprimaria2022 = calificacionesmodel.ExtraeDocentesArr(materiassextoprimaria2022)
						hombressextoprimaria2022, mujeressextoprimaria2022 = calificacionesmodel.HombresyMujeres(sextoprimaria2022.ID)
						break
					case "Preescolar":
						sextopreescolar2022 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiassextopreescolar2022 = calificacionesmodel.ExtraeMateriasPorSemestre(sextopreescolar2022.ID)
						profesoressextopreescolar2022 = calificacionesmodel.ExtraeDocentesArr(materiassextopreescolar2022)
						hombressextopreescolar2022, mujeressextopreescolar2022 = calificacionesmodel.HombresyMujeres(sextopreescolar2022.ID)
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
				case "2022":
					switch vv.Licenciatura {
					case "Primaria":
						septimoprimaria2022 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasseptimoprimaria2022 = calificacionesmodel.ExtraeMateriasPorSemestre(septimoprimaria2022.ID)
						profesoresseptimoprimaria2022 = calificacionesmodel.ExtraeDocentesArr(materiasseptimoprimaria2022)
						hombresseptimoprimaria2022, mujeresseptimoprimaria2022 = calificacionesmodel.HombresyMujeres(septimoprimaria2022.ID)
						break
					case "Preescolar":
						septimopreescolar2022 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasseptimopreescolar2022 = calificacionesmodel.ExtraeMateriasPorSemestre(septimopreescolar2022.ID)
						profesoresseptimopreescolar2022 = calificacionesmodel.ExtraeDocentesArr(materiasseptimopreescolar2022)
						hombresseptimopreescolar2022, mujeresseptimopreescolar2022 = calificacionesmodel.HombresyMujeres(septimopreescolar2022.ID)
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
				case "2022":
					switch vv.Licenciatura {
					case "Primaria":
						octavoprimaria2022 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasoctavoprimaria2022 = calificacionesmodel.ExtraeMateriasPorSemestre(octavoprimaria2022.ID)
						profesoresoctavoprimaria2022 = calificacionesmodel.ExtraeDocentesArr(materiasoctavoprimaria2022)
						hombresoctavoprimaria2022, mujeresoctavoprimaria2022 = calificacionesmodel.HombresyMujeres(octavoprimaria2022.ID)
						break
					case "Preescolar":
						octavopreescolar2022 = calificacionesmodel.ExtraeSemestre(vv.ID)
						materiasoctavopreescolar2022 = calificacionesmodel.ExtraeMateriasPorSemestre(octavopreescolar2022.ID)
						profesoresoctavopreescolar2022 = calificacionesmodel.ExtraeDocentesArr(materiasoctavopreescolar2022)
						hombresoctavopreescolar2022, mujeresoctavopreescolar2022 = calificacionesmodel.HombresyMujeres(octavopreescolar2022.ID)
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
		//2022 ----------------------------------------------------------------------
		//1
		//Primaria
		//Materias
		ctx.ViewData("Mpp2022", materiasprimeroprimaria2022)
		//Profesores
		ctx.ViewData("Ppp2022", profesoresprimeroprimaria2022)
		//Hombres y Mujeres
		ctx.ViewData("App2022", hombresprimeroprimaria2022+mujeresprimeroprimaria2022)
		ctx.ViewData("Hpp2022", hombresprimeroprimaria2022)
		ctx.ViewData("Mupp2022", mujeresprimeroprimaria2022)

		//Preescolar
		//Materias
		ctx.ViewData("Mppre2022", materiasprimeropreescolar2022)
		//Profesores
		ctx.ViewData("Pppre2022", profesoresprimeropreescolar2022)
		//Hombres y Mujeres
		ctx.ViewData("Appre2022", hombresprimeropreescolar2022+mujeresprimeropreescolar2022)
		ctx.ViewData("Hppre2022", hombresprimeropreescolar2022)
		ctx.ViewData("Muppre2022", mujeresprimeropreescolar2022)
		//2
		//Primaria
		//Materias
		ctx.ViewData("Msp2022", materiassegundoprimaria2022)
		//Profesores
		ctx.ViewData("Psp2022", profesoressegundoprimaria2022)
		//Hombres y Mujeres
		ctx.ViewData("Asp2022", hombressegundoprimaria2022+mujeressegundoprimaria2022)
		ctx.ViewData("Hsp2022", hombressegundoprimaria2022)
		ctx.ViewData("Musp2022", mujeressegundoprimaria2022)

		//Preescolar
		//Materias
		ctx.ViewData("Mspre2022", materiassegundopreescolar2022)
		//Profesores
		ctx.ViewData("Pspre2022", profesoressegundopreescolar2022)
		//Hombres y Mujeres
		ctx.ViewData("Aspre2022", hombressegundopreescolar2022+mujeressegundopreescolar2022)
		ctx.ViewData("Hspre2022", hombressegundopreescolar2022)
		ctx.ViewData("Muspre2022", mujeressegundopreescolar2022)
		//3
		//Primaria
		//Materias
		ctx.ViewData("Mtp2022", materiasterceroprimaria2022)
		//Profesores
		ctx.ViewData("Ptp2022", profesoresterceroprimaria2022)
		//Hombres y Mujeres
		ctx.ViewData("Atp2022", hombresterceroprimaria2022+mujeresterceroprimaria2022)
		ctx.ViewData("Htp2022", hombresterceroprimaria2022)
		ctx.ViewData("Mutp2022", mujeresterceroprimaria2022)

		//Preescolar
		//Materias
		ctx.ViewData("Mtpre2022", materiasterceropreescolar2022)
		//Profesores
		ctx.ViewData("Ptpre2022", profesoresterceropreescolar2022)
		//Hombres y Mujeres
		ctx.ViewData("Atpre2022", hombresterceropreescolar2022+mujeresterceropreescolar2022)
		ctx.ViewData("Htpre2022", hombresterceropreescolar2022)
		ctx.ViewData("Mutpre2022", mujeresterceropreescolar2022)
		//4
		//Primaria
		//Materias
		ctx.ViewData("Mcp2022", materiascuartoprimaria2022)
		//Profesores
		ctx.ViewData("Pcp2022", profesorescuartoprimaria2022)
		//Hombres y Mujeres
		ctx.ViewData("Acp2022", hombrescuartoprimaria2022+mujerescuartoprimaria2022)
		ctx.ViewData("Hcp2022", hombrescuartoprimaria2022)
		ctx.ViewData("Mucp2022", mujerescuartoprimaria2022)

		//Preescolar
		//Materias
		ctx.ViewData("Mcpre2022", materiascuartopreescolar2022)
		//Profesores
		ctx.ViewData("Pcpre2022", profesorescuartopreescolar2022)
		//Hombres y Mujeres
		ctx.ViewData("Acpre2022", hombrescuartopreescolar2022+mujerescuartopreescolar2022)
		ctx.ViewData("Hcpre2022", hombrescuartopreescolar2022)
		ctx.ViewData("Mucpre2022", mujerescuartopreescolar2022)
		//5
		//Primaria
		//Materias
		ctx.ViewData("Mqp2022", materiasquintoprimaria2022)
		//Profesores
		ctx.ViewData("Pqp2022", profesoresquintoprimaria2022)
		//Hombres y Mujeres
		ctx.ViewData("Aqp2022", hombresquintoprimaria2022+mujeresquintoprimaria2022)
		ctx.ViewData("Hqp2022", hombresquintoprimaria2022)
		ctx.ViewData("Muqp2022", mujeresquintoprimaria2022)

		//Preescolar
		//Materias
		ctx.ViewData("Mqpre2022", materiasquintopreescolar2022)
		//Profesores
		ctx.ViewData("Pqpre2022", profesoresquintopreescolar2022)
		//Hombres y Mujeres
		ctx.ViewData("Aqpre2022", hombresquintopreescolar2022+mujeresquintopreescolar2022)
		ctx.ViewData("Hqpre2022", hombresquintopreescolar2022)
		ctx.ViewData("Muqpre2022", mujeresquintopreescolar2022)
		//6
		//Primaria
		//Materias
		ctx.ViewData("Msxp2022", materiassextoprimaria2022)
		//Profesores
		ctx.ViewData("Psxp2022", profesoressextoprimaria2022)
		//Hombres y Mujeres
		ctx.ViewData("Asxp2022", hombressextoprimaria2022+mujeressextoprimaria2022)
		ctx.ViewData("Hsxp2022", hombressextoprimaria2022)
		ctx.ViewData("Musxp2022", mujeressextoprimaria2022)

		//Preescolar
		//Materias
		ctx.ViewData("Msxpre2022", materiassextopreescolar2022)
		//Profesores
		ctx.ViewData("Psxpre2022", profesoressextopreescolar2022)
		//Hombres y Mujeres
		ctx.ViewData("Asxpre2022", hombressextopreescolar2022+mujeressextopreescolar2022)
		ctx.ViewData("Hsxpre2022", hombressextopreescolar2022)
		ctx.ViewData("Musxpre2022", mujeressextopreescolar2022)
		//7
		//Primaria
		//Materias
		ctx.ViewData("Mstp2022", materiasseptimoprimaria2022)
		//Profesores
		ctx.ViewData("Pstp2022", profesoresseptimoprimaria2022)
		//Hombres y Mujeres
		ctx.ViewData("Astp2022", hombresseptimoprimaria2022+mujeresseptimoprimaria2022)
		ctx.ViewData("Hstp2022", hombresseptimoprimaria2022)
		ctx.ViewData("Mustp2022", mujeresseptimoprimaria2022)

		//Preescolar
		//Materias
		ctx.ViewData("Mstpre2022", materiasseptimopreescolar2022)
		//Profesores
		ctx.ViewData("Pstpre2022", profesoresseptimopreescolar2022)
		//Hombres y Mujeres
		ctx.ViewData("Astpre2022", hombresseptimopreescolar2022+mujeresseptimopreescolar2022)
		ctx.ViewData("Hstpre2022", hombresseptimopreescolar2022)
		ctx.ViewData("Mustpre2022", mujeresseptimopreescolar2022)
		//8
		//Primaria
		//Materias
		ctx.ViewData("Mop2022", materiasoctavoprimaria2022)
		//Profesores
		ctx.ViewData("Pop2022", profesoresoctavoprimaria2022)
		//Hombres y Mujeres
		ctx.ViewData("Aop2022", hombresoctavoprimaria2022+mujeresoctavoprimaria2022)
		ctx.ViewData("Hop2022", hombresoctavoprimaria2022)
		ctx.ViewData("Muop2022", mujeresoctavoprimaria2022)

		//Preescolar
		//Materias
		ctx.ViewData("Mopre2022", materiasoctavopreescolar2022)
		//Profesores
		ctx.ViewData("Popre2022", profesoresoctavopreescolar2022)
		//Hombres y Mujeres
		ctx.ViewData("Aopre2022", hombresoctavopreescolar2022+mujeresoctavopreescolar2022)
		ctx.ViewData("Hopre2022", hombresoctavopreescolar2022)
		ctx.ViewData("Muopre2022", mujeresoctavopreescolar2022)

		//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

		fmt.Println("Bienvenido ", userOn.Nombre)
		if err := ctx.View("Semestres.html"); err != nil {
			ctx.Application().Logger().Infof(err.Error())
		}
	} else {
		ctx.Redirect("/login", iris.StatusSeeOther)
	}
}
