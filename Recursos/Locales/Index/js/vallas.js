

    const selectUbicacion = document.getElementById("ubicacion");
    const selectUbicaciones = document.querySelector("#ubicaciones");
    const calle = document.querySelector("#calle");
    const numero = document.querySelector("#numero");
    const colonia = document.querySelector("#colonia");
    const localidad = document.querySelector("#localidad");
    const estado = document.querySelector("#estado");
    const municipio = document.querySelector("#municipio");
    const nombre = document.querySelector("#nombre");
    const telefono = document.querySelector("#telefono");
    const celular = document.querySelector("#celular");

    selectUbicacion.addEventListener("change",function(e){
        e.preventDefault()
        if(selectUbicacion.value == "guardada"){
            selectUbicaciones.classList.add("d-block")
            calle.classList.add("d-none")
            numero.classList.add("d-none")
            colonia.classList.add("d-none")
            localidad.classList.add("d-none")
            estado.classList.add("d-none")
            municipio.classList.add("d-none")
            
        }else{
            selectUbicaciones.classList.remove("d-block")
            calle.classList.remove("d-none")
            numero.classList.remove("d-none")
            colonia.classList.remove("d-none")
            localidad.classList.remove("d-none")
            estado.classList.remove("d-none")
            municipio.classList.remove("d-none")
        }
    })

    window.propietario.addEventListener("change", function(e){
        e.preventDefault()
        if(this.value == "REGISTRADO"){
            window.propietariosReg.classList.add("d-block");
            nombre.classList.add("d-none");
            celular.classList.add("d-none");
            telefono.classList.add("d-none");
        }else{
            window.propietariosReg.classList.remove("d-block")
            nombre.classList.remove("d-none")
            celular.classList.remove("d-none")
            telefono.classList.remove("d-none")

        }
    })
