package usuariosmodel

import (
	"fmt"
	"log"

	conexiones "github.com/vadgun/Experimental/Conexiones"
	calificacionesmodel "github.com/vadgun/Experimental/Modelos/CalificacionesModel"
	indexmodel "github.com/vadgun/Experimental/Modelos/IndexModel"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//GuardaEntidadesDeAlumnos Asigna los nuevos Objects IDs para tener bien ubicado el uno al otro usuario
func GuardaEntidadesDeAlumnos(alumno calificacionesmodel.Alumno, mongouser indexmodel.MongoUser) bool {

	alumno.ID = bson.NewObjectId()
	mongouser.ID = bson.NewObjectId()

	alumno.MongoUser = mongouser.ID
	mongouser.UserID = alumno.ID

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_AL)
	err1 := c.Insert(alumno)
	if err1 != nil {
		fmt.Println(err1)
		return false
	}

	d := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_U)
	err2 := d.Insert(mongouser)
	if err2 != nil {
		fmt.Println(err2)
		return false
	}
	return true

}

//GuardaEntidadesDeDocentes Asigna los nuevos Objects IDs para tener bien ubicado el uno al otro usuario
func GuardaEntidadesDeDocentes(docente calificacionesmodel.Docente, mongouser indexmodel.MongoUser) bool {

	docente.ID = bson.NewObjectId()
	mongouser.ID = bson.NewObjectId()

	docente.MongoUser = mongouser.ID
	mongouser.UserID = docente.ID

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_DC)
	err1 := c.Insert(docente)
	if err1 != nil {
		fmt.Println(err1)
		return false
	}

	d := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_U)
	err2 := d.Insert(mongouser)
	if err2 != nil {
		fmt.Println(err2)
		return false
	}
	return true

}

//ExtraeSemestres Devuelve todos los semestres dados de alta con sus materias o no
func ExtraeSemestres() []calificacionesmodel.Semestre {

	var semestres []calificacionesmodel.Semestre

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_SM)
	err1 := c.Find(bson.M{}).Sort("Licenciatura", "Semestre", "Plan").All(&semestres)
	if err1 != nil {
		fmt.Println(err1)
	}

	return semestres

}

//GuardarAlumnosMasivamente -> Guardalos alumnos en la base de datos
func GuardarAlumnosMasivamente(alumnos []calificacionesmodel.Alumno, usuarios []indexmodel.MongoUser) bool {

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_AL)
	d := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_U)

	// for i := 0; i <= 5; i++ {
	// for k, v := range alumnos {

	for k, v := range alumnos {

		v.ID = bson.NewObjectId()
		usuarios[k].ID = bson.NewObjectId()

		v.MongoUser = usuarios[k].ID
		usuarios[k].UserID = v.ID
		err1 := c.Insert(v)
		err2 := d.Insert(usuarios[k])
		if err1 != nil {
			fmt.Println("No se pudo insertar masivamente los usuarios en la base de datos", err1)
			return false
		}
		if err2 != nil {
			fmt.Println("No se pudo insertar masivamente los usuarios en la base de datos", err2)
			return false
		}
	}

	return true
}
