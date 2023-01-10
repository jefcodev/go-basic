package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html> <body> Proformas</body></html>")
}

func main() {
	Proformas = []Proforma{
		Proforma{Id: "1", NumProforma: "PRO-001-001-00000191", Cliente: "Jose Perez", Cantidad: 3, Precio: 158.10, Total: 473.30},
		Proforma{Id: "2", NumProforma: "PRO-001-001-00000192", Cliente: "Maria Alveaer", Cantidad: 2, Precio: 100.10, Total: 300.00},
	}
	iniciarServidor()
}

// Estructura del la tabla prformas

type Proforma struct {
	Id          string  `json:id`
	NumProforma string  `json:numProforma`
	Cliente     string  `json:cliente`
	Cantidad    int     `json:cantidad`
	Precio      float32 `json:precio`
	Total       float32 `json:total`
}

// Array global de proforma:
var Proformas []Proforma

// Metodo para encontrar todos los proformas

func findAllProformas(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Solicitud atendida: findAllProformas")
	json.NewEncoder(w).Encode(Proformas)
}

func createNewProforma(w http.ResponseWriter, r *http.Request) {
	// Se obtiene el body desde el request y
	// se deserializa en una variable proforma:
	reqBody, _ := ioutil.ReadAll(r.Body)
	var proforma Proforma
	json.Unmarshal(reqBody, &proforma)
	// adicionamos en el array la nueva proforma:
	Proformas = append(Proformas, proforma)
	json.NewEncoder(w).Encode(proforma)
}

// Función borrar un proformas mediante su id

func deleteProforma(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Solicitud atendida: deleteProforma")
	vars := mux.Vars(r)
	key := vars["id"]
	// buscar la proforma a eliminar:
	for index, proforma := range Proformas {
		if proforma.Id == key {
			// borrar del array:
			Proformas = append(Proformas[:index], Proformas[index+1:]...)
		}
	}
}

// Actualizar los datos de una proforma

func updateProforma(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Solicitud atendida: updateProforma")
	// Se obtiene el body desde el request y
	// se deserializa en una variable proforma:
	reqBody, _ := ioutil.ReadAll(r.Body)
	var proforma Proforma
	json.Unmarshal(reqBody, &proforma)
	key := proforma.Id
	// buscar la proforma a actualizar:
	for index, p := range Proformas {
		if p.Id == key {
			// actualizar el array:
			Proformas[index] = proforma
			break
		}
	}
	json.NewEncoder(w).Encode(proforma)
}

// Se agregan las rutas y se inicia el servidor

func iniciarServidor() {
	fmt.Println("API REST simple con lenguaje go.")
	ruteador := mux.NewRouter().StrictSlash(true)
	ruteador.HandleFunc("/", homePage)
	ruteador.HandleFunc("/proformas", findAllProformas)
	ruteador.HandleFunc("/proforma", createNewProforma).Methods("POST")
	ruteador.HandleFunc("/proforma", updateProforma).Methods("PUT")
	ruteador.HandleFunc("/proforma/{id}", deleteProforma).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", ruteador))
}
