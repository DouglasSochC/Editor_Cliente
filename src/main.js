console.log("JS Loaded")

//Aqui se declara hacia donde se envia el post
const url = "http://localhost:5555"
//Este nos permite utilizar el elemento del inputForm
var inputForm = document.getElementById("inputForm")
//Estas variables nos permiten utilizar los editores de la api ACE y tambien
//ayuda en la interaccion del mismo.
var editor = ace.edit("Editor");
var consolapy = ace.edit("consolapython");
var consolajs = ace.edit("consolajavascript");

inputForm.addEventListener("submit", (e) => {

    //prevent auto submission
    e.preventDefault()
    const formdata = new FormData(inputForm)
    formdata.append("codigojava", editor.getValue());
    //body: formdata
    fetch(url, {
        method: "POST",
        body: formdata,
    }).then(
        response => response.text()
    ).then(
        (data) => { console.log("Esta es la data que retorna: " + data); consolapy.setValue(data) }
    ).catch(
        error => console.error(error)
    )
})