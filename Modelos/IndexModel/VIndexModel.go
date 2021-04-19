package indexmodel

import (
	"gopkg.in/mgo.v2/bson"
)

//MongoUser Controla datos de Usuario del sistema
type MongoUser struct {
	ID             bson.ObjectId `bson:"_id,omitempty"`
	Nombre         string        `bson:"Nombre"`
	Apellidos      string        `bson:"Apellidos"`
	Edad           int           `bson:"Edad"`
	Usuario        string        `bson:"Usuario"`
	Telefono       string        `bson:"Telefono"`
	Puesto         string        `bson:"Puesto"`
	Key            string        `bson:"Key"`
	Nombre2        string        `bson:"Nombre2"`
	UserID         bson.ObjectId `bson:"UserID"`
	Alumno         bool          `bson:"Alumno"`
	Docente        bool          `bson:"Docente"`
	Administrativo bool          `bson:"Administrativo"`
	Director       bool          `bson:"Director"`
	Admin          bool          `bson:"Admin"`
	Contador       bool          `bson:"Contador"`
}

//Permiso Otorga cierto control a las vistas que puede ver el usuario logeado en el sistema para evitar incongruencias de seguridad
type Permiso struct {
	ID                    bson.ObjectId `bson:"_id,omitempty"`
	Permisos              string        `bson:"Permisos"`
	Admin                 []int         `bson:"Admin"`
	Docente               []int         `bson:"Docente"`
	Alumno                []int         `bson:"Alumno"`
	Administrativo        []int         `bson:"Administrativo"`
	Director              []int         `bson:"Director"`
	Contador              []int         `bson:"Contador"`
	PermisoCalificaciones int           `bson:"PermisoCalificaciones"`
	PermisoUsuarios       int           `bson:"PermisoUsuarios"`
	PermisoAsignar        int           `bson:"PermisoAsignar"`
	PermisoInscripcion    int           `bson:"PermisoInscripcion"`
	PermisoHorarios       int           `bson:"PermisoHorarios"`
	PermisoDirectorio     int           `bson:"PermisoDirectorio"`
	PermisoKardex         int           `bson:"PermisoKardex"`
	PermisoIndex          int           `bson:"PermisoIndex`
	PermisoAlumnos        int           `bson:"PermisoAlumnos`
	PermisoDocentes       int           `bson:"PermisoDocentes`
}
