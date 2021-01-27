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

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_MT)
	err1 := c.Find(bson.M{"Plan": plan, "Licenciatura": lic, "Semestre": sem}).All(&materias)
	if err1 != nil {
		fmt.Println(err1)
	}

	return materias

}

//ObtenerAlumnosFiltrados -> Apartir de Sem
func ObtenerAlumnosFiltrados(sem string) []Alumno {
	var alumnos []Alumno

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_AL)
	err1 := c.Find(bson.M{"CursandoSem": sem}).All(&alumnos)
	if err1 != nil {
		fmt.Println(err1)
	}

	return alumnos
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
	err1 := c.Find(bson.M{}).Select(bson.M{"Nombre": 1, "ApellidoP": 1, "ApellidoM": 1}).All(&docentes)
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