package calificacionesmodel

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Un maestro solo podra llamar a 40 alumnos por grupo su consulta no puede ser tan grande
//tambien podemos ocultar los campos en la propia consulta para evitar campos incecesarios a la hora de hacer las consultas
//Tal seria el caso para la integracion de la boleta que campos elegir para traer al alumno y sus calificaciones

//Boleta impresa por medio de un ID.
type Boleta struct {
	ID     bson.ObjectId `bson:"_id,omitempty"`
	Alumno Alumno
}

//Alumno ligado al usuario de sistema sus datos personales y perfil dentro de la institución
type Alumno struct {
	ID              bson.ObjectId `bson:"_id,omitempty"`
	MongoUser       bson.ObjectId `bson:"MongoUser"`
	IsSystemUser    bool          `bson:"IsSystemUser"`
	Matricula       string        `bson:"Matricula"`
	Nombre          string        `bson:"Nombre"`
	ApellidoP       string        `bson:"ApellidoP"`
	ApellidoM       string        `bson:"ApellidoM"`
	FechaNac        time.Time     `bson:"FechaNac"`
	Curp            string        `bson:"Curp"`
	Calle           string        `bson:"Calle"`
	ColAsentamiento string        `bson:"ColAsentamiento"`
	Municipio       string        `bson:"Municipio"`
	Estado          string        `bson:"Estado"`
	Telefono        string        `bson:"Telefono"`
	TipoSangre      string        `bson:"TipoSangre"`
	Licenciatura    string        `bson:"Licenciatura"` //Que sea un Documento Licenciatura
	Plan            string        `bson:"Plan"`
	CursandoSem     string        `bson:"CursandoSem"`
	SiguienteSem    string        `bson:"SiguienteSem"`
	AnteriorSem     string        `bson:"AnteriorSem"`
	InicioSem       string        `bson:"InicioSem"`
	Imagen          bson.ObjectId `bson:"Imagen"`
	Horario         string        `bson:"Horario"`
}

//Docente y su perfil dentro del sistema
type Docente struct {
	ID              bson.ObjectId   `bson:"_id,omitempty"`
	MongoUser       bson.ObjectId   `bson:"MongoUser"`
	IsSystemUser    bool            `bson:"IsSystemUser"`
	Nombre          string          `bson:"Nombre"`
	ApellidoP       string          `bson:"ApellidoP"`
	ApellidoM       string          `bson:"ApellidoM"`
	FechaNac        time.Time       `bson:"FechaNac"`
	Curp            string          `bson:"Curp"`
	Rfc             string          `bson:"Rfc"`
	Calle           string          `bson:"Calle"`
	ColAsentamiento string          `bson:"ColAsentamiento"`
	Municipio       string          `bson:"Municipio"`
	Estado          string          `bson:"Estado"`
	Telefono1       string          `bson:"Telefono1"`
	Telefono2       string          `bson:"Telefono2"`
	TipoSangre      string          `bson:"TipoSangre"`
	Imagen          bson.ObjectId   `bson:"Imagen"`
	Grupos          string          `bson:"Grupos"`
	Materias        []bson.ObjectId `bson:"Materias"`
	Horario         string          `bson:"Horario"`
	CapturaInicio   time.Time       `bson:"CapturaInicio"`
	CapturaFin      time.Time       `bson:"CapturaFin"`
}

// //Director y su perfil dentro del sistema
// type Director struct {
// 	ID              bson.ObjectId `bson:"_id,omitempty"`
// 	MongoUser      bson.ObjectId `bson:"MongoUser"`
// 	IsSystemUser    bool          `bson:"IsSystemUser"`
// 	Nombre          string        `bson:"Nombre"`
// 	ApellidoP       string        `bson:"ApellidoP"`
// 	ApellidoM       string        `bson:"ApellidoM"`
// 	FechaNac        time.Time     `bson:"FechaNac"`
// 	Curp            string        `bson:"Curp"`
// 	Rfc             string        `bson:"Rfc"`
// 	Calle           string        `bson:"Calle"`
// 	ColAsentamiento string        `bson:"ColAsentamiento"`
// 	Municipio       string        `bson:"Municipio"`
// 	Estado          string        `bson:"Estado"`
// 	Telefono1       string        `bson:"Telefono1"`
// 	Telefono2       string        `bson:"Telefono2"`
// 	TipoSangre      string        `bson:"TipoSangre"`
// 	Grupos          string        `bson:"Grupos"`
// 	Imagen          string        `bson:"Imagen"`
// 	Horario         string        `bson:"Horario"`
// }

// //Administrativo Administrativo y su perfil dentro del sistema
// type Administrativo struct {
// 	ID              bson.ObjectId `bson:"_id,omitempty"`
// 	MongoUser      bson.ObjectId `bson:"MongoUser"`
// 	IsSystemUser    bool          `bson:"IsSystemUser"`
// 	Nombre          string        `bson:"Nombre"`
// 	ApellidoP       string        `bson:"ApellidoP"`
// 	ApellidoM       string        `bson:"ApellidoM"`
// 	FechaNac        time.Time     `bson:"FechaNac"`
// 	Curp            string        `bson:"Curp"`
// 	Rfc             string        `bson:"Rfc"`
// 	Calle           string        `bson:"Calle"`
// 	ColAsentamiento string        `bson:"ColAsentamiento"`
// 	Municipio       string        `bson:"Municipio"`
// 	Estado          string        `bson:"Estado"`
// 	Telefono1       string        `bson:"Telefono1"`
// 	Telefono2       string        `bson:"Telefono2"`
// 	TipoSangre      string        `bson:"TipoSangre"`
// 	Grupos          string        `bson:"Grupos"`
// 	Imagen          string        `bson:"Imagen"`
// 	Horario         string        `bson:"Horario"`
// }

//CalificacionAlumno seran las unicas para el en cada materia.
type CalificacionAlumno struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
}

//CalificacionMaestro seran las de unicamente su grupo solicitado bajo un filtro de consultas o botones previamente consultados por ser un docente
type CalificacionMaestro struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
}

//CalificacionAdmon seran las entregadas al personal administrativo para su verificacion e impresion de boletas.
type CalificacionAdmon struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
}

//Materia y sus caracteristicas a utilizar
type Materia struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	Materia      string        `bson:"Materia"`
	Plan         string        `bson:"Plan"`
	Licenciatura string        `bson:"Licenciatura"`
	Semestre     bson.ObjectId `bson:"Semestre"`
	Horas        string        `bson:"Horas"`
	Creditos     string        `bson:"Creditos"`
}

//Semestre Controlara el id donde se interconectaran los alumnos las materias y los docentes, para crear algo llamado ordenes de captura de calificaciones
type Semestre struct {
	ID           bson.ObjectId   `bson:"_id,omitempty"`
	Semestre     string          `bson:"Semestre"`
	Licenciatura string          `bson:"Licenciatura"`
	Plan         string          `bson:"Plan"`
	Materias     []bson.ObjectId `bson:"Materias"`
}
