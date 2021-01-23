package inscripcionmodel

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	calificacionesmodel "github.com/vadgun/Experimental/Modelos/CalificacionesModel"
)

//Boucher Controla el numero de boucer y quien con que usuario se dio de alta el ticket de inscripcion a la normal
type Boucher struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Alumno    calificacionesmodel.Alumno
	FechaPago time.Time `bson:"FechaPago"`
	Mensaje   string    `bson:"Mensaje"`
}
