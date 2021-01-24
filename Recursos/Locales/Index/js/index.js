let calificacionesit = document.getElementById("calificacionesit");
let alumnosit= document.getElementById("alumnosit");
let docentesit= document.getElementById("docentesit");
let directorioit= document.getElementById("directorioit");
let buscadorit= document.getElementById("buscadorit");
let horariosit= document.getElementById("horariosit");
let kardexit= document.getElementById("kardexit");
let inscripcionit= document.getElementById("inscripcionit");
let asignacionit= document.getElementById("asignacionit");

function EnviarAsignacion(){

    var plan = document.getElementById("plan");
    var licenciatura = document.getElementById("licenciatura");
    var semestre = document.getElementById("semestre");
    var docenteevaluado = document.getElementById("docenteevaluado");

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

    if (docenteevaluado.value != "") {


        Swal.fire({
              title: '¿Asignar esta materia?',
              text: "Vefirica",
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

//  $('.dropify').dropify();


//  function VerificarDatosInscripcion(){

//     var formulario = document.getElementById("altaPagoInscripcion");

//     if (confirm("Confirma Opción")) {
//         formulario.submit();
//         // return true;
//       } else {
//         Swal.fire({
//             title: 'Los datos son correctos?',
//             text: "Verifica una vez mas",
//             icon: 'warning',
//             showCancelButton: true,
//             confirmButtonColor: '#3085d6',
//             cancelButtonColor: '#d33',
//             confirmButtonText: 'Yes, delete it!'
//           }).then((result) => {
//             if (result.isConfirmed) {
//               Swal.fire(
//                 'Deleted!',
//                 'Your file has been deleted.',
//                 'success'
//               )
//             }
//           })
//           return false;
//       }
      

//  }

// var input = document.getElementById("buscadorcliente");
// console.log(input);

// if (input == null) {

// }else{ 
//   input.addEventListener("keyup", function(event) {
//     if (event.key === "Enter") {
//       event.preventDefault();
//       BuscaClientes();
//     }
//   });
// }

// function EditarCliente(data){

//   $.ajax({
//     url: '/obtenercliente',
//     data: { data: data },
//     type: 'POST',
//     dataType: 'html',
//     success: function(result) {
//         console.log("Operacion Realizada con Exito");
//         $("#formeditarclientes").html(result);
//     },
//     error: function(xhr, status) {
//         console.log("Error en la consulta")
//     },
//     complete: function(xhr, status) {
//         console.log("Datos de edicion de cliente obtenidos")
        
//     }
// });

// $("#ModalEditarClientes").modal("show");

// }

// function EliminarCliente(data){

// Swal.fire({
//   title: '¿Eliminar cliente?',
//   text: "Esta accion no se puede revertir",
//   icon: 'warning',
//   showCancelButton: true,
//   confirmButtonColor: '#3085d6',
//   cancelButtonColor: '#d33',
//   confirmButtonText: 'Si, eliminar'
// }).then((result) => {
//   if (result.isConfirmed) {

//     $.ajax({
//       url: '/eliminarcliente',
//       data: { data: data },
//       type: 'POST',
//       dataType: 'html',
//       success: function(result) {
//           console.log("Operacion Realizada con Exito");
//           $("#formeditarclientes").html(result);
//           setTimeout(function(){ location.replace("/clientes"); }, 2000);
//       },
//       error: function(xhr, status) {
//           console.log("Error en la consulta");
//       },
//       complete: function(xhr, status) {
//           console.log("Eliminacion de cliente completa");
          
//       }
//   });  

//   }else if (result.isDismissed) {
//     Swal.fire("El cliente no ha sido eliminado");
//   }
// })

// }

// function EditarEmpleado(data){

//   $.ajax({
//     url: '/obtenerempleado',
//     data: { data: data },
//     type: 'POST',
//     dataType: 'html',
//     success: function(result) {
//         console.log("Operacion Realizada con Exito");
//         $("#formeditarempleado").html(result);
//     },
//     error: function(xhr, status) {
//         console.log("Error en la consulta")
//     },
//     complete: function(xhr, status) {
//         console.log("Datos de edicion de empleado obtenidos")
        
//     }
// });

// $("#ModalEditarEmpleados").modal("show");

// }

// function EliminarEmpleado(data){

// Swal.fire({
//   title: '¿Eliminar empleado?',
//   text: "Esta accion no se puede revertir",
//   icon: 'warning',
//   showCancelButton: true,
//   confirmButtonColor: '#3085d6',
//   cancelButtonColor: '#d33',
//   confirmButtonText: 'Si, eliminar'
// }).then((result) => {
//   if (result.isConfirmed) {

//     $.ajax({
//       url: '/eliminarempleado',
//       data: { data: data },
//       type: 'POST',
//       dataType: 'html',
//       success: function(result) {
//           console.log("Operacion Realizada con Exito");
//           $("#formeditarempleado").html(result);
//           setTimeout(function(){ location.replace("/empleados"); }, 2000);
//       },
//       error: function(xhr, status) {
//           console.log("Error en la consulta");
//       },
//       complete: function(xhr, status) {
//           console.log("Eliminacion de empleado completa");
          
//       }
//   });  

//     Swal.fire(
//       'Eliminado!',
//       'Empleado eliminado',
//       'success'
//     )
//   }else if (result.isDismissed) {
//     Swal.fire("El empleado no ha sido eliminado");
//   }
// })

// }

// function Ubicaciones(data){
//   if (data == ""){
//     return false
//   }else{
//     alert(data);
//   }

// }

// function EliminarEspectacular(data){
   
//   Swal.fire({
//     title: '¿Eliminar espectacular?',
//     text: "Esta accion no se puede revertir",
//     icon: 'warning',
//     showCancelButton: true,
//     confirmButtonColor: '#3085d6',
//     cancelButtonColor: '#d33',
//     confirmButtonText: 'Si, eliminar'
//   }).then((result) => {
//     if (result.isConfirmed) {

//       $.ajax({
//         url: '/eliminarespectacular',
//         data: { data: data },
//         type: 'POST',
//         dataType: 'html',
//         success: function(result) {
//             console.log("Operacion Realizada con Exito");
//             $("#ContainerImagenesEspectacular").html(result);
//             setTimeout(function(){ location.replace("/espectaculares"); }, 2000);
//         },
//         error: function(xhr, status) {
//             console.log("Error en la consulta");
//         },
//         complete: function(xhr, status) {
//             console.log("Eliminacion de espectacular completa");
            
//         }
//     });  

//       Swal.fire(
//         'Eliminado!',
//         'Espectacular eliminado',
//         'success'
//       )
//     }else if (result.isDismissed) {
//       Swal.fire("El espectacular no ha sido eliminado");
//     }
//   })
  
// }


// function VerImagenesEspectacular(data){
  
//   $.ajax({
//     url: '/imagenesespectacular',
//     data: { data: data },
//     type: 'POST',
//     dataType: 'html',
//     success: function(result) {
//         console.log("Operacion Realizada con Exito");
//         $("#ContainerImagenesEspectacular").html(result);
//     },
//     error: function(xhr, status) {
//         console.log("Error en la consulta")
//     },
//     complete: function(xhr, status) {
//         console.log("Datos de imagenes para modal")
        
//     }
// });  

// $("#ModalImagenesEspectacular").modal("show");

// }

// function VerificarEspectacularDuplicado(data){

// //Crear la funcion que verifique si el numero de control esta disponible para no agregar espectaculares duplicados
// $.ajax({
//   url: '/verificaespectacular',
//   data: { data: data },
//   type: 'POST',
//   dataType: 'html',
//   success: function(result) {
//       console.log("Operacion Realizada con Exito");
//       $("#answer").html(result);
//   },
//   error: function(xhr, status) {
//       console.log("Error en la consulta")
//   },
//   complete: function(xhr, status) {
//       console.log("Verificacion "+ data +" completada")
      
//   }
// });

// }

// function GenerarFicha(data){

//   $.ajax({
//     url: '/generarfichadecliente',
//     data: { data: data },
//     type: 'POST',
//     dataType: 'html',
//     success: function(result) {
//         console.log("Operacion Realizada con Exito");
//         $("#answer").html(result);
//     },
//     error: function(xhr, status) {
//         console.log("Error en la consulta")
//     },
//     complete: function(xhr, status) {
//         console.log("Verificacion "+ data +" completada")
        
//     }
//   });
// }

// function Descargar(name) {

//     $.ajax({
//         url: 'Recursos/Archivos/Fichas' + name + '.pdf',
//         method: 'GET',
//         xhrFields: {
//             responseType: 'blob'
//         },
//         success: function(data) {
//             var a = document.createElement('a');
//             var url = window.URL.createObjectURL(data);
//             a.href = url;
//             a.download = name + '.pdf';
//             a.click();
//             window.URL.revokeObjectURL(url);
//         }
//     });
// }

// function EditarMaterial(data){

//   $.ajax({
//     url: '/obtenermaterial',
//     data: { data: data },
//     type: 'POST',
//     dataType: 'html',
//     success: function(result) {
//         console.log("Operacion Realizada con Exito");
//         $("#formeditarmaterial").html(result);
        
//     },
//     error: function(xhr, status) {
//         console.log("Error en la consulta")
//     },
//     complete: function(xhr, status) {
//         console.log("Datos de material para modal")
        
//     }
// }); 


// $("#ModalEditarMaterial").modal("show");

// }

// function EliminarMaterial(data){

//   Swal.fire({
//     title: '¿Eliminar material?',
//     text: "Esta accion no se puede revertir",
//     icon: 'warning',
//     showCancelButton: true,
//     confirmButtonColor: '#3085d6',
//     cancelButtonColor: '#d33',
//     confirmButtonText: 'Si, eliminar'
//   }).then((result) => {
//     if (result.isConfirmed) {

//       $.ajax({
//         url: '/eliminarmaterial',
//         data: { data: data },
//         type: 'POST',
//         dataType: 'html',
//         success: function(result) {
//             console.log("Operacion Realizada con Exito");
//             $("#formeditarmaterial").html(result);
//             setTimeout(function(){ location.replace("/materiales"); }, 2000);
//         },
//         error: function(xhr, status) {
//             console.log("Error en la consulta");
//         },
//         complete: function(xhr, status) {
//             console.log("Eliminacion de material completa");
            
//         }
//     });  
//     }else if (result.isDismissed) {
//       Swal.fire("El material no ha sido eliminado");
//     }
//   })

// }

// function CalculaPrecio(){

//   material = document.getElementById("material");

//   var idmaterial = "";
//   var preciomaterial = 0;
//   var res = material.value.split(":");
//    idmaterial = res[0];
//    preciomaterial = parseFloat(res[1]);

//   if (material.value != ""){
    
//     inputprecio = document.getElementById("precio");
//     inputancho = document.getElementById("ancho");
//     inputalto = document.getElementById("alto");
//     area = inputancho.value * inputalto.value
//     inputprecio.value = parseFloat(preciomaterial * area).toFixed(2);
//     instalacion = document.getElementById("instalacion");
//     costoimpreso = document.getElementById("costoimpreso");

//     costoimpreso.value = inputprecio.value;

//     if (area < 50 ){

//       instalacion.value = 800;
      
//     }else {
//       instalacion.value = 1200;
//     }




//   }else{

//     instalacion = document.getElementById("instalacion");
//     costoimpreso = document.getElementById("costoimpreso");
//     inputprecio = document.getElementById("precio");

//     instalacion.value = 0;
//     costoimpreso.value = 0;
//     inputprecio.value = 0;
    
//     return false;
//   }

// }

// function Catalogo(data){

//   container = document.getElementById("catalogocontainer");

//   switch (data) {
//     case "espectaculares":
//       container.innerHTML = `
//       <a href="Javascript:ImprimirEspectaculares('disponibles');" class="btn btn-info" role="button"> Espectaculares Disponibles</a>
//       <a href="Javascript:ImprimirEspectaculares('rentados');" class="btn btn-info" role="button">Espectaculares Ocupados</a>`
//       break;
//     case "vallas":
//       container.innerHTML = `
//       <a href="Javascript:ImprimirVallas('disponibles');" class="btn btn-warning" role="button"> Vallas Disponibles</a>
//       <a href="Javascript:ImprimirVallas('rentados');" class="btn btn-warning" role="button">Vallas Ocupadas</a>`
//       break;
//     case "vallasM":
//       container.innerHTML = `
//       <a href="Javascript:ImprimirVallasM('disponibles');" class="btn btn-dark" role="button"> Vallas Móviles Disponibles</a>
//       <a href="Javascript:ImprimirVallasM('rentados');" class="btn btn-dark" role="button">Vallas Móviles Ocupadas</a>`
//       break;
//   }

// }

// function ImprimirEspectaculares(data){
  
//   $.ajax({
//     url: '/imprimirespectaculares',
//     data: { data: data },
//     type: 'POST',
//     dataType: 'html',
//     success: function(result) {
//         console.log("Operacion Realizada con Exito");
//         $("#answer").html(result);
//     },
//     error: function(xhr, status) {
//         console.log("Error en la consulta")
//     },
//     complete: function(xhr, status) {
//         console.log("Catalogo de espectaculaes "+ data +" completado")
        
//     }
//   });


// }

// function ImprimirVallas(data){
//   $.ajax({
//     url: '/imprimirvallas',
//     data: { data: data },
//     type: 'POST',
//     dataType: 'html',
//     success: function(result) {
//         console.log("Operacion Realizada con Exito");
//         $("#answer").html(result);
//     },
//     error: function(xhr, status) {
//         console.log("Error en la consulta")
//     },
//     complete: function(xhr, status) {
//         console.log("Catalogo de vallas "+ data +" completado")
        
//     }
//   });
// }

// function ImprimirVallasM(data){
//   $.ajax({
//     url: '/imprimirvallasm',
//     data: { data: data },
//     type: 'POST',
//     dataType: 'html',
//     success: function(result) {
//         console.log("Operacion Realizada con Exito");
//         $("#answer").html(result);
//     },
//     error: function(xhr, status) {
//         console.log("Error en la consulta")
//     },
//     complete: function(xhr, status) {
//         console.log("Catalogo de vallas M "+ data +" completado")
        
//     }
//   });
// }

