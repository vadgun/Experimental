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

//PersonalDocenteActivo Devuelve el personal docente activo para ser elegible en la materia
func PersonalDocenteActivo() []Docente {

	var docentes []Docente

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_DC)
	err1 := c.Find(bson.M{}).Select(bson.M{"Nombre": 1, "ApellidoP": 1}).All(&docentes)
	if err1 != nil {
		fmt.Println(err1)
	}

	return docentes

}
