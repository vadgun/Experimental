package indexmodel

import (
	"fmt"
	"log"

	conexiones "github.com/vadgun/Experimental/Conexiones"
	calificacionesmodel "github.com/vadgun/Experimental/Modelos/CalificacionesModel"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//VerificarUsuario Autentifica al usuario en la base de datos
func VerificarUsuario(usuario MongoUser) (bool, MongoUser) {
	var encontrado bool
	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_U)
	err1 := c.Find(bson.M{"Usuario": usuario.Usuario, "Key": usuario.Key}).One(&usuario)
	if err1 != nil {
		fmt.Println(err1)
	}

	if usuario.Nombre == "" {
		encontrado = false
	} else {
		encontrado = true
	}

	return encontrado, usuario
}

//GetUserOn Se extrae el usuario logeado
func GetUserOn(user string) MongoUser {

	var usuarioOn MongoUser
	usrobjid := bson.ObjectIdHex(user)

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_U)
	err1 := c.FindId(usrobjid).One(&usuarioOn)
	if err1 != nil {
		fmt.Println(err1)
	}

	return usuarioOn
}

//TienePermiso Verifica si puede o no ver la vista correspondiente al usuario logeado en el sistema
func TienePermiso(indexpermiso int, user1, user2 MongoUser) bool {

	var isadmin bool
	var isalumno bool
	var isdocente bool
	var isadministrativo bool
	var isdirector bool
	var iscontador bool

	var permisos Permiso

	if user1.Nombre != "" || user2.Nombre != "" {

		if user1.Admin == true || user1.Admin {
			isadmin = true
		}

		if user1.Docente == true || user1.Docente {
			isdocente = true
		}
		if user1.Alumno == true || user1.Alumno {
			isalumno = true
		}

		if user1.Administrativo == true || user1.Administrativo {
			isadministrativo = true
		}
		if user1.Director == true || user1.Director {
			isdirector = true
		}

		if user1.Contador == true || user1.Contador {
			iscontador = true
		}
	}
	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_PR)
	err1 := c.Find(bson.M{"Permisos": "ConfiguracionesDePermisos"}).One(&permisos)
	if err1 != nil {
		fmt.Println(err1)
	}

	if isadmin {

		if permisos.Admin[indexpermiso] == 1 {
			fmt.Println("Es admin y tiene permiso")
			return true
		} else {
			fmt.Println("Es admin y no tiene permiso")
			return false
		}

	}

	if isalumno {
		if permisos.Alumno[indexpermiso] == 1 {
			fmt.Println("Es alumno y tiene permiso")
			return true
		} else {
			fmt.Println("Es alumno y no tiene permiso")
			return false
		}
	}

	if isadministrativo {
		if permisos.Administrativo[indexpermiso] == 1 {
			fmt.Println("Es administrativo y tiene permiso")

			return true
		} else {
			fmt.Println("Es administrativo y no tiene permiso")

			return false
		}
	}

	if isdirector {
		if permisos.Director[indexpermiso] == 1 {
			fmt.Println("Es director y tiene permiso")
			return true
		} else {
			fmt.Println("Es director y no tiene permiso")
			return false
		}
	}

	if isdocente {
		if permisos.Docente[indexpermiso] == 1 {
			fmt.Println("Es docente y tiene permiso")
			return true
		} else {
			fmt.Println("Es docente y no tiene permiso")
			return false
		}
	}

	if iscontador {
		if permisos.Contador[indexpermiso] == 1 {
			fmt.Println("Es contador y tiene permiso")
			return true
		} else {
			fmt.Println("Es contador y no tiene permiso")
			return false
		}
	}

	return false

}

//IfIsDocenteBringMaterias Regresa la botonera si es docente o un mensaje si no lo es
func IfIsDocenteBringMaterias(userOn MongoUser) []calificacionesmodel.Materia {

	var idmateriasdeldocente []calificacionesmodel.Materia

	idmateriasdeldocente = calificacionesmodel.ExtraeMaterias(userOn.UserID)

	//Si tiene 5 ids para las materias debes traer los datos de esas 5 materias.
	//el primer paso es traer esas 5 materias para filtrarlas y el docente pueda elegir cual materia ver y rellenar consecutivamente

	return idmateriasdeldocente

}

//ExtraeSoloUsuarios Herramienta temporal
func ExtraeSoloUsuarios() []MongoUser {

	var usuarios []MongoUser

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_U)
	err1 := c.Find(bson.M{"Alumno": true}).All(&usuarios)
	if err1 != nil {
		fmt.Println(err1)
	}

	return usuarios

}

//HerramientaAsignacionUsuarios ->
func HerramientaAsignacionUsuarios(usuario MongoUser) {

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_U)
	err2 := c.UpdateId(usuario.ID, usuario)
	if err2 != nil {
		fmt.Println("No se encontro el usuario en la base de datos", err2)
	}

}
