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
}
