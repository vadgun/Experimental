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

//GuardarCapturaCalificaciones -> Guarda las calificaciones a cada alumno
func GuardarCapturaCalificaciones(alumnos []string, calificaciones []float64, index int) bool {

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
		fmt.Println(index)
		fmt.Println("Alumno", alumno.Calificaciones)
		fmt.Println("Alumno", alumno.Calificaciones[index])

		alumno.Calificaciones[index] = calificaciones[k]
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
