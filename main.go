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
	fmt.Fprintf(w, "<html> <body> Hola <br> Mundo</body></html>")
}

func main() {
	Productos = []Producto{
		Producto{Id: "1", Nombre: "Monitor 17 pulgadas", Descripcion: "Conexión HDMI - Full HD", Cantidad: 3},
		Producto{Id: "2", Nombre: "Teclado USB", Descripcion: "Color negro y teclas multimedia", Cantidad: 7},
	}
	iniciarServidor()
}

// Estructura del la tabla producto

type Producto struct {
	Id          string `json:id`
	Nombre      string `json:nombre`
	Descripcion string `json:descripcion`
	Cantidad    int    `json:cantidad`
}

// Array global de productos:
var Productos []Producto

// Metodo para encontrar todos los productos

func findAllProductos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Solicitud atendida: findAllProductos")
	json.NewEncoder(w).Encode(Productos)
}

func createNewProducto(w http.ResponseWriter, r *http.Request) {
	// Se obtiene el body desde el request y
	// se deserializa en una variable producto:
	reqBody, _ := ioutil.ReadAll(r.Body)
	var producto Producto
	json.Unmarshal(reqBody, &producto)
	// adicionamos en el array el nuevo producto:
	Productos = append(Productos, producto)
	json.NewEncoder(w).Encode(producto)
}

// Función get para consultar producto por id

func findProductoById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Solicitud atendida: findProductoById")
	vars := mux.Vars(r)
	key := vars["id"]
	for _, producto := range Productos {
		if producto.Id == key {
			json.NewEncoder(w).Encode(producto)
		}
	}
}

// Función borrar un producto mediante su id

func deleteProducto(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Solicitud atendida: deleteProducto")
	vars := mux.Vars(r)
	key := vars["id"]
	// buscar el producto a eliminar:
	for index, producto := range Productos {
		if producto.Id == key {
			// borrar del array:
			Productos = append(Productos[:index], Productos[index+1:]...)
		}
	}
}

// Actualizar los datos de un producto

func updateProducto(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Solicitud atendida: updateProducto")
	// Se obtiene el body desde el request y
	// se deserializa en una variable producto:
	reqBody, _ := ioutil.ReadAll(r.Body)
	var producto Producto
	json.Unmarshal(reqBody, &producto)
	key := producto.Id
	// buscar el producto a actualizar:
	for index, p := range Productos {
		if p.Id == key {
			// actualizar el array:
			Productos[index] = producto
			break
		}
	}
	json.NewEncoder(w).Encode(producto)
}

// Se agregan las rutas y se inicia el servidor

func iniciarServidor() {
	/* fmt.Println("API REST simple con lenguaje go.")
	http.HandleFunc("/", homePage)
	http.HandleFunc("/productos", findAllProductos)
	log.Fatal(http.ListenAndServe(":8080", nil)) */

	fmt.Println("API REST simple con lenguaje go.")
	ruteador := mux.NewRouter().StrictSlash(true)
	ruteador.HandleFunc("/", homePage)
	ruteador.HandleFunc("/productos", findAllProductos)
	ruteador.HandleFunc("/producto", createNewProducto).Methods("POST")
	ruteador.HandleFunc("/producto", updateProducto).Methods("PUT")
	ruteador.HandleFunc("/producto/{id}", deleteProducto).Methods("DELETE")
	ruteador.HandleFunc("/producto/{id}", findProductoById)
	log.Fatal(http.ListenAndServe(":8080", ruteador))
}
