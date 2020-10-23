package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var Port = ":5555"
func main(){

	http.HandleFunc("/",ServeFiles)
	fmt.Println("Serving @ : ","http://localhost"+Port)
	log.Fatal(http.ListenAndServe(Port,nil))
}

func ServeFiles(w http.ResponseWriter, r *http.Request){

	switch r.Method{

		//Aqui se obtiene la pagina principal al momento de iniciar el servidor
		case "GET":

			path := r.URL.Path
			//path: Esta variable nos sirve para iniciar el servidor
			//Si nos fijamos inicia con '/', ya despues le asignamos la direccion del get
			if path == "/"{
				path = "./src/index.html"
			}else{
				path = "."+path
			}
			http.ServeFile(w,r,path)

		//Este metodo 'POST', es donde se ingresan datos y
		//este nos deben de retornar algun mensaje
		case "POST":

			r.ParseMultipartForm(0)
			
			message := r.FormValue("codigojava")

			//fmt.Println("----------------------------------")
			//fmt.Println("Message from Client: ",message)
			// respond to client's request
			fmt.Fprintf(w, "Server: %s \n", message+ " | " + time.Now().Format(time.RFC3339))
			//fmt.Fprintf(w, "Esto es lo que se retorna")
		default:
			fmt.Fprintf(w,"Request type other than GET or POSt not supported")
		}
}