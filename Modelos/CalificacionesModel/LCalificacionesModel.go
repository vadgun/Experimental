package calificacionesmodel

import (
	"fmt"
	"log"

	conexiones "github.com/vadgun/Experimental/Conexiones"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//ObtenerMateriasFiltradas -> Apartir de Plan, Lic, Sem
func ObtenerMateriasFiltradas(plan, lic, sem string) []Materia {

	var materias []Materia
	var idsem bson.ObjectId

	idsem = bson.ObjectIdHex(sem)

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_MT)
	err1 := c.Find(bson.M{"Plan": plan, "Licenciatura": lic, "Semestre": idsem}).All(&materias)
	if err1 != nil {
		fmt.Println(err1)
	}

	return materias

}

//ObtenerAlumnosFiltrados -> Apartir de Sem
func ObtenerAlumnosFiltrados(sem string) []Alumno {
	var alumnos []Alumno

	semid := bson.ObjectIdHex(sem)
	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_AL)
	err1 := c.Find(bson.M{"CursandoSem": semid}).Sort("ApellidoP", "ApellidoM").All(&alumnos)
	if err1 != nil {
		fmt.Println(err1)
	}

	return alumnos
}

//ExtraeMateria -> Extrae la materia por idstring
func ExtraeMateria(mat string) Materia {

	var materia Materia

	idmat := bson.ObjectIdHex(mat)

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_MT)
	err1 := c.FindId(idmat).One(&materia)
	if err1 != nil {
		fmt.Println(err1)
	}

	return materia

}

func ActualizarMateria(materia Materia) {

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_MT)
	err1 := c.UpdateId(materia.ID, materia)
	if err1 != nil {
		fmt.Println(err1)
	}

}

//ActualizaConfig -> Actualiza la configuracion solicitada
func ActualizaConfig(id string, config Configuracion) {

	var idobj bson.ObjectId

	idobj = bson.ObjectIdHex(id)

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C("CONFIG")
	err1 := c.UpdateId(idobj, config)
	if err1 != nil {
		fmt.Println(err1)
	}

}

//ExtraeConfigBoleta -> Extrae parametros generales de configuracio del centro escolar
func ExtraeConfigBoleta() Configuracion {

	var configuracion Configuracion

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	c := session.DB(conexiones.MONGO_DB).C("CONFIG")
	err1 := c.Find(bson.M{"Configuracion": "General"}).One(&configuracion)
	if err1 != nil {
		fmt.Println(err1)
	}

	return configuracion

}

//ExtraeMateriasPorSemestre -> Extrae las materias por semestre y las devuelve
func ExtraeMateriasPorSemestre(sem bson.ObjectId) []Materia {

	var materias []Materia

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_MT)
	err1 := c.Find(bson.M{"Semestre": sem}).All(&materias)
	if err1 != nil {
		fmt.Println(err1)
	}

	return materias

}

//ExtraeMateriasPorSemestreID -> Extrae las materias por semestre y las devuelve
func ExtraeMateriasPorSemestreID(sem bson.ObjectId) []bson.ObjectId {

	var materias []Materia

	var bsonss []bson.ObjectId

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_MT)
	err1 := c.Find(bson.M{"Semestre": sem}).All(&materias)
	if err1 != nil {
		fmt.Println(err1)
	}

	for _, v := range materias {
		bsonss = append(bsonss, v.ID)
	}

	return bsonss

}

//PersonalDocenteActivo Devuelve el personal docente activo para ser elegible en la materia
func PersonalDocenteActivo() []Docente {

	var docentes []Docente

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_DC)
	err1 := c.Find(bson.M{}).Select(bson.M{"Nombre": 1, "ApellidoP": 1, "ApellidoM": 1, "Telefono1": 1, "CorreoE": 1}).All(&docentes)
	if err1 != nil {
		fmt.Println(err1)
	}

	return docentes

}

//ObtenerDocenteYConvertirIDMATERIA regresa el docente a partir del id Hex regresa los id hex
func ObtenerDocenteYConvertirIDMATERIA(iddocente, idmateria string) (Docente, bson.ObjectId) {
	var docente Docente

	idobjdocente := bson.ObjectIdHex(iddocente)

	idobjmateria := bson.ObjectIdHex(idmateria)

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_DC)
	err1 := c.FindId(idobjdocente).One(&docente)
	if err1 != nil {
		fmt.Println("1", err1)
	}

	return docente, idobjmateria
}

//AsignarMateria Asigna la materia si no existe
func AsignarMateria(docente Docente) bool {

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_DC)
	err1 := c.UpdateId(docente.ID, docente)
	if err1 != nil {
		fmt.Println("1", err1)
		return false
	}
	return true

}

//RevocarMateria Remueve la materia del arreglo de materias
func RevocarMateria(idmat, iddocente string) bool {
	idobjmat := bson.ObjectIdHex(idmat)
	idobjdocente := bson.ObjectIdHex(iddocente)

	var docente Docente

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_DC)
	err1 := c.FindId(idobjdocente).One(&docente)
	if err1 != nil {
		fmt.Println("1", err1)
	}
	if len(docente.Materias) == 0 {
		return false
	}

	var encontrado bool
	var index int
	for k, v := range docente.Materias {
		if v == idobjmat {
			index = k
			encontrado = true
		}
	}

	if len(docente.Materias) == 1 && encontrado {
		var arraytem []bson.ObjectId
		docente.Materias = arraytem

	} else {
		docente.Materias = RemoveIndex(docente.Materias, index)
	}

	if !encontrado {
		return false
	} else {

		err2 := c.UpdateId(idobjdocente, docente)
		if err2 != nil {
			fmt.Println("2", err1)
			return false
		}
		return true

	}

}

//RemoveIndex Remueve el index de un slice de bson Ids
func RemoveIndex(s []bson.ObjectId, index int) []bson.ObjectId {
	return append(s[:index], s[index+1:]...)
}

//ExtraeMaterias -> Regresa materias por id de docente
func ExtraeMaterias(iddocente bson.ObjectId) []Materia {

	var docente Docente

	var materias []Materia
	var materia Materia

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_DC)
	err1 := c.FindId(iddocente).Select(bson.M{"Materias": 1}).One(&docente)
	if err1 != nil {
		fmt.Println("1", err1)
	}

	d := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_MT)

	//buscar materias

	for _, v := range docente.Materias {

		err1 := d.FindId(v).One(&materia)
		if err1 != nil {
			fmt.Println("1", err1)
		}

		materias = append(materias, materia)

	}

	fmt.Println("Materias-> ", docente.Materias)
	return materias

}

//ExtraeDocentes -> Regresa el docente de materia
func ExtraeDocentes(materias []Materia) []Docente {
	var docentes []Docente
	var docente Docente

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_DC)

	for _, v := range materias {

		var matess []bson.ObjectId
		matess = append(matess, v.ID)
		// err1 := c.Find(bson.M{"Materias": bson.M{"$in": v.ID}}).One(&docente)
		err1 := c.Find(bson.M{"Materias": bson.M{"$in": matess}}).One(&docente)

		// o2 := bson.M{"$group" :bson.M{"_id": "$channel","Total": bson.M{"$sum": 1,},

		if err1 != nil {
			fmt.Println("1", err1)
			docente.Nombre = "Sin asignación"
			docentes = append(docentes, docente)
		} else {
			docentes = append(docentes, docente)
		}
	}

	return docentes
}

//ExtraeDocentes -> Regresa el docente de materia
func ExtraeDocentesArr(materias []Materia) []string {
	var docentes []Docente
	var docente Docente
	var docentevacio Docente

	var docentesarr []string

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_DC)

	for _, v := range materias {

		var matess []bson.ObjectId
		matess = append(matess, v.ID)
		// err1 := c.Find(bson.M{"Materias": bson.M{"$in": v.ID}}).One(&docente)
		err1 := c.Find(bson.M{"Materias": bson.M{"$in": matess}}).One(&docente)

		// o2 := bson.M{"$group" :bson.M{"_id": "$channel","Total": bson.M{"$sum": 1,},

		if err1 != nil {
			fmt.Println("1", err1)
			docentes = append(docentes, docentevacio)
		} else {
			docentes = append(docentes, docente)
		}
	}

	var cadena string
	for _, vv := range docentes {

		cadena = vv.Nombre + " " + vv.ApellidoP + " " + vv.ApellidoM + " / " + vv.Telefono1 + " / " + vv.CorreoE

		if vv.ID == "" {
			docentesarr = append(docentesarr, "Sin docente asignado")
		} else {
			docentesarr = append(docentesarr, cadena)
		}

	}

	return docentesarr
}

//HombresyMujeres -> Devuelve la cantidad de alumnos divididos en hombres y mujeres inscritos al semestre
func HombresyMujeres(semestre bson.ObjectId) (int, int) {
	var hombres int
	var mujeres int

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	var err1 error
	var err2 error

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_AL)
	hombres, err1 = c.Find(bson.M{"CursandoSem": semestre, "Sexo": "Masculino"}).Count()
	mujeres, err2 = c.Find(bson.M{"CursandoSem": semestre, "Sexo": "Femenino"}).Count()

	if err1 != nil {
		fmt.Println("1", err1)
	}

	if err2 != nil {
		fmt.Println("1", err2)
	}

	return hombres, mujeres
}

//ExtraeSemestreString -> Regresa el semestre a la peticion
func ExtraeSemestreString(semestrestring string) Semestre {

	var semestre Semestre
	var idsemestre bson.ObjectId

	idsemestre = bson.ObjectIdHex(semestrestring)

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_SM)
	err1 := c.FindId(idsemestre).One(&semestre)
	if err1 != nil {
		fmt.Println("1", err1)
	}

	return semestre
}

//ExtraeSemestre -> Regresa el semestre a la peticion
func ExtraeSemestre(idsemestre bson.ObjectId) Semestre {

	var semestre Semestre

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_SM)
	err1 := c.FindId(idsemestre).One(&semestre)
	if err1 != nil {
		fmt.Println("1", err1)
	}

	return semestre
}

//ExtraeSemestres -> Regresa todos los semestres a la peticion
func ExtraeSemestres() []Semestre {

	var semestres []Semestre

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_SM)
	err1 := c.Find(bson.M{}).All(&semestres)
	if err1 != nil {
		fmt.Println("1", err1)
	}

	return semestres
}

//GuardarCapturaCalificaciones -> Guarda las calificaciones a cada alumno
func GuardarCapturaCalificaciones(alumnos []string, calificaciones, asistencias []float64, index int) bool {

	var guardado bool

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_AL)

	for k, v := range alumnos {
		var alumno Alumno
		idalumno := bson.ObjectIdHex(v)
		err1 := c.FindId(idalumno).One(&alumno)
		if err1 != nil {
			fmt.Println("1 ERROR 1", err1)
		}

		alumno.Calificaciones[index] = calificaciones[k]
		alumno.Asistencias[index] = asistencias[k]
		err2 := c.UpdateId(alumno.ID, alumno)
		if err2 != nil {
			fmt.Println("2", err2)
		}
		guardado = true

	}

	return guardado

}

//CrearSemestre -> Crea semestre
func CrearSemestre(semestre Semestre) bool {

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_SM)
	err1 := c.Insert(semestre)
	if err1 != nil {
		fmt.Println("Error insertando semestre", err1)
		return false
	}

	return true

}

//AsignarMateriaASemestre -> Asigna materia al semestre mediande el id de materia e id de semestre
func AsignarMateriaASemestre(materia Materia, semestre Semestre) bool {

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	materia.ID = bson.NewObjectId()
	semestre.Materias = append(semestre.Materias, materia.ID)

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_MT)
	err1 := c.Insert(materia)
	if err1 != nil {
		fmt.Println("Error insertando materia", err1)
		return false
	}

	d := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_SM)
	err2 := d.UpdateId(semestre.ID, semestre)
	if err2 != nil {
		fmt.Println("Error editanto semestre", err2)
		return false
	}

	return true

}

//TraerSemestre -> Trae el Semestre para su modificacion
func TraerSemestre(idsemestre string) Semestre {

	var semestre Semestre

	objidsemestre := bson.ObjectIdHex(idsemestre)

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_SM)
	err1 := c.FindId(objidsemestre).One(&semestre)
	if err1 != nil {
		fmt.Println("No se encontro el semestre en la base de datos", err1)
	}

	return semestre

}

//ActualizaAlumno Actualiza los datos del alumno
func ActualizaAlumno(alumno Alumno) {

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_AL)
	err1 := c.UpdateId(alumno.ID, alumno)
	if err1 != nil {
		fmt.Println("alumno no actualizado", err1)
	}

}

//GuardaKardex Guarda la informacion de un documento de kardex de un semestre de un alumno
func GuardaKardex(kardex Kardex) {
	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_KR)
	err1 := c.Insert(kardex)
	if err1 != nil {
		fmt.Println("kardex no guardado", err1)
	}
}

//ExtraeAlumno -> Regresa el alumno por idstring
func ExtraeAlumno(idalum string) Alumno {

	var alumno Alumno

	objidalumno := bson.ObjectIdHex(idalum)

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_AL)
	err1 := c.FindId(objidalumno).One(&alumno)
	if err1 != nil {
		fmt.Println("No se encontro el alumno en la base de datos", err1)
	}

	return alumno

}

//ExtraeSoloAlumnos -> Herramienta Temporal
func ExtraeSoloAlumnos() []Alumno {

	var alumnos []Alumno

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_AL)
	err1 := c.Find(bson.M{}).All(&alumnos)
	if err1 != nil {
		fmt.Println("No se encontro el alumno en la base de datos", err1)
	}

	return alumnos

}

//HerramientaAsignacionAlumnos -> Herramienta Temporal
func HerramientaAsignacionAlumnos(alumno Alumno) {

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_AL)
	err1 := c.UpdateId(alumno.ID, alumno)
	if err1 != nil {
		fmt.Println("No se encontro el alumno en la base de datos", err1)
	}
}

//ExtraeDocente -> Regresa el docente a la peticion de asignar materia
func ExtraeDocente(iddocente string) Docente {

	var docente Docente

	objiddocente := bson.ObjectIdHex(iddocente)

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_DC)
	err1 := c.FindId(objiddocente).One(&docente)
	if err1 != nil {
		fmt.Println("No se encontro el docente en la base de datos", err1)
	}

	return docente

}

func ActualizarMateriasDocente(docente Docente) {

	var materiasVacias []bson.ObjectId
	docente.Materias = materiasVacias
	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_DC)
	err1 := c.UpdateId(docente.ID, docente)
	if err1 != nil {
		fmt.Println("error eliminando materias de docente ", err1)
	}

}

//SiguienteSemestre -> Regresa el siguiente semestre a partir del numero de semestre, licenciatura y plan (1,Primaria,2012)
func SiguienteSemestre(numsemestre, lic, plan string) bson.ObjectId {

	var semestre Semestre

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	var busqueda string

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_SM)
	err1 := c.Find(bson.M{"Semestre": numsemestre, "Licenciatura": lic, "Plan": plan}).One(&semestre)
	if err1 != nil {
		fmt.Println("No se encontro el semestre en la base de datos", err1, "+||||+", busqueda, lic, plan)
	}

	return semestre.ID

}

func RemoverUsuarioAlumno(idAlumno, idUser bson.ObjectId) bool {

	var eliminado bool

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_AL)
	d := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_U)

	c.RemoveId(idAlumno)
	d.RemoveId(idUser)

	return eliminado

}

func ExtraeBusquedaDeMaterias(buscar string) []Materia {

	var materias []Materia

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Buscando :", buscar)

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_MT)
	err1 := c.Find(bson.M{"Materia": bson.RegEx{buscar, "i"}}).Sort("Materia").All(&materias)
	if err1 != nil {
		fmt.Println(err1)
	}

	return materias

}
