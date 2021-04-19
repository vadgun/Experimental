let calificacionesit = document.getElementById("calificacionesit");
let alumnosit= document.getElementById("alumnosit");
let docentesit= document.getElementById("docentesit");
let directorioit= document.getElementById("directorioit");
let buscadorit= document.getElementById("buscadorit");
let horariosit= document.getElementById("horariosit");
let kardexit= document.getElementById("kardexit");
let inscripcionit= document.getElementById("inscripcionit");
let asignacionit= document.getElementById("asignacionit");
let usuariosit= document.getElementById("usuariosit");
let profileit= document.getElementById("profileit");
let relojit= document.getElementById("relojit");





function EnviarAsignacion(){

    var plan = document.getElementById("plan");
    var licenciatura = document.getElementById("licenciatura");
    var semestre = document.getElementById("semestre");

    console.log(plan.value);
    console.log(licenciatura.value);
    console.log(semestre.value);


if (licenciatura.value!= "" && plan.value!= "" && semestre.value != ""){

    
    $.ajax({
        url: '/obtenerMaterias',
        data: { licenciatura: licenciatura.value, plan:plan.value, semestre:semestre.value },
        type: 'POST',
        dataType: 'html',
        success: function(result) {
            console.log("Operacion Realizada con Exito");
            $("#AsignacionTable").html(result);
        },
        error: function(xhr, status) {
            console.log("Error en la consulta");
        },
        complete: function(xhr, status) {
            console.log("Filtro de Matarias realizado");
            
        }
    });
    
}   
}


function AsignarMateria(data){

    var docenteevaluado = document.getElementById("docenteevaluado");
    var optionSelected = docenteevaluado.options[docenteevaluado.selectedIndex].text;
    

    if (docenteevaluado.value != "") {


        Swal.fire({
              title: '¿Asignar esta materia?',
              text: optionSelected,
              icon: 'warning',
              showCancelButton: true,
              confirmButtonColor: '#3085d6',
              cancelButtonColor: '#d33',
              confirmButtonText: 'Si, continuar'
            }).then((result) => {
              if (result.isConfirmed) {
            
                $.ajax({
                  url: '/asignarMateriaADocente',
                  data: { data: data, iddocente:docenteevaluado.value },
                  type: 'POST',
                  dataType: 'html',
                  success: function(result) {
                      console.log("Operacion Realizada con Exito");
                        $("#answer").html(result);

                  },
                  error: function(xhr, status) {
                      console.log("Error en la consulta");
                  },
                  complete: function(xhr, status) {
                      console.log("La asignacion de la materia ha sido completada");
                      
                  }
              });  
            
              }else if (result.isDismissed) {
                Swal.fire("La materia no ha sido asignada");
              }
            })

        
        return false
    }else{

        Swal.fire("Selecciona un docente antes de asignar una materia");

    }
}

function RevocarMateria(data){

    var docenteevaluado = document.getElementById("docenteevaluado");
    var optionSelected = docenteevaluado.options[docenteevaluado.selectedIndex].text;
    

    if (docenteevaluado.value != "") {


        Swal.fire({
              title: '¿Revocar esta materia?',
              text: optionSelected,
              icon: 'warning',
              showCancelButton: true,
              confirmButtonColor: '#3085d6',
              cancelButtonColor: '#d33',
              confirmButtonText: 'Si, continuar'
            }).then((result) => {
              if (result.isConfirmed) {
            
                $.ajax({
                  url: '/revocarMateriaADocente',
                  data: { data: data, iddocente:docenteevaluado.value },
                  type: 'POST',
                  dataType: 'html',
                  success: function(result) {
                      console.log("Operacion Realizada con Exito");
                        $("#answer").html(result);

                  },
                  error: function(xhr, status) {
                      console.log("Error en la consulta");
                  },
                  complete: function(xhr, status) {
                      console.log("La revocacion de la materia ha sido completada");
                      
                  }
              });  
            
              }else if (result.isDismissed) {
                Swal.fire("La materia no ha sido revocada");
              }
            })

        
        return false
    }else{

        Swal.fire("Selecciona un docente antes de revocar una materia");

    }
}

function Usuarios(data){

    // Swal.fire(data);
  $.ajax({
    url: '/solicitudUsuario',
    data: { data: data },
    type: 'POST',
    dataType: 'html',
    success: function(result) {
        console.log("Operacion Realizada con Exito");
        $("#UsuariosContainer").html(result);
    },
    error: function(xhr, status) {
        console.log("Error en la consulta")
    },
    complete: function(xhr, status) {
        console.log("Formulario para "+data+ " solicitado")
        
    }
});


}


function SolicitarAlumnos(){
    var semestre = document.getElementById("semestre");

if (semestre.value != ""){

    
    $.ajax({
        url: '/obtenerAlumnos',
        data: { semestre:semestre.value },
        type: 'POST',
        dataType: 'html',
        success: function(result) {
            console.log("Operacion Realizada con Exito");
            $("#AlumnosTable").html(result);
        },
        error: function(xhr, status) {
            console.log("Error en la consulta");
        },
        complete: function(xhr, status) {
            console.log("Consulta de alumnos realizado");
            
        }
    });
    
}}

function SolicitarAlumnosCal(){ 
var semestre = document.getElementById("semestrecalif");

if (semestre.value != ""){

    
    $.ajax({
        url: '/obtenerAlumnosCalif',
        data: { semestre:semestre.value },
        type: 'POST',
        dataType: 'html',
        success: function(result) {
            console.log("Operacion Realizada con Exito");
            $("#AlumnosTable").html(result);
        },
        error: function(xhr, status) {
            console.log("Error en la consulta");
        },
        complete: function(xhr, status) {
            console.log("Consulta de alumnos realizado");
            
        }
    });
    
}}
function CrearFormulario(data){

    $.ajax({
        url: '/crearformulario',
        data: { data:data },
        type: 'POST',
        dataType: 'html',
        success: function(result) {
            console.log("Operacion Realizada con Exito");
            $("#CalificacionesTable").html(result);
        },
        error: function(xhr, status) {
            console.log("Error en la consulta");
        },
        complete: function(xhr, status) {
            console.log("Formulario " +data+ " Devuelto");
            
        }
    });

    
}



function AgregarCalificacion(data){
    
    var iddocente = document.getElementById("iddocente");
    
    $.ajax({
        url: '/agregarcalificacion',
        data: { data:data, iddocente: iddocente.value},
        type: 'POST',
        dataType: 'html',
        success: function(result) {
            console.log("Operacion Realizada con Exito");
            $("#CalificacionesTable").html(result);
        },
        error: function(xhr, status) {
            console.log("Error en la consulta");
        },
        complete: function(xhr, status) {
            console.log("Calificacion Agregada");
            
        }
    });

}

function CambiaAtributos(data){
var input = document.getElementById(data);
input.classList.remove("btn-dark");
input.classList.add("btn-warning");
}




function GenerarBoleta(data){

    $.ajax({
        url: '/generarboleta',
        data: { data:data},
        type: 'POST',
        dataType: 'html',
        success: function(result) {
            console.log("Operacion Realizada con Exito");
            $("#answer").html(result);
        },
        error: function(xhr, status) {
            console.log("Error generando boleta");
        },
        complete: function(xhr, status) {
            console.log("Generar Boleta Terminado");
            
        }
    });
}
function Descargar(name) {

    $.ajax({
        url: 'Recursos/Archivos/' + name + '.pdf',
        method: 'GET',
        xhrFields: {
            responseType: 'blob'
        },
        success: function(data) {
            var a = document.createElement('a');
            var url = window.URL.createObjectURL(data);
            a.href = url;
            a.download = name + '.pdf';
            a.click();
            window.URL.revokeObjectURL(url);

            // eliminardocumento(name);

        }
    });
}

function LigarUsuarios(){

    
    $.ajax({
        url: '/ligarusuarios',
        data: {},
        type: 'POST',
        dataType: 'html',
        success: function(result) {
            console.log("Operacion LigarUsuarios con Exito");
            $("#answer").html(result);
        },
        error: function(xhr, status) {
            console.log("Error LigarUsuarios");
        },
        complete: function(xhr, status) {
            console.log("GLigarUsuarios Terminado");
            
        }
    });

}

function ImprimirCalificacion(data){

    $.ajax({
        url: '/imprimircalificacion',
        data: {data:data},
        type: 'POST',
        dataType: 'html',
        success: function(result) {
            console.log("Operacion Imprimir calif con Exito");
            $("#answer").html(result);
        },
        error: function(xhr, status) {
            console.log("Error Imprimir calif");
        },
        complete: function(xhr, status) {
            console.log("Imprimir calif Terminado");
            
        }
    });

}

function Config(data){
    $.ajax({
        url: '/obtenconfig',
        data: {data:data},
        type: 'POST',
        dataType: 'html',
        success: function(result) {
            console.log("Operacion config con Exito");
            $("#CalificacionesTable").html(result);
        },
        error: function(xhr, status) {
            console.log("Error config");
        },
        complete: function(xhr, status) {
            console.log("config Terminado");
            
        }
    });
}

$('.dropify').dropify();  