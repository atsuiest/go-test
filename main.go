package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Model
type Producto struct {
	ID     int64
	Nombre string
	Desc   string
}

// Slice = Array
var productos []Producto

// Definir handler getProductos
func getProductos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(productos)
}

func main() {
	// Inicializar router
	r := mux.NewRouter()

	// Mock de datos
	productos = append(productos, Producto{ID: 1, Nombre: "Lapiz", Desc: "Lapiz genérico"})
	productos = append(productos, Producto{ID: 2, Nombre: "Teclado", Desc: "Teclado genius"})
	productos = append(productos, Producto{ID: 3, Nombre: "Mouse", Desc: "Mouse hp"})

	// Definir handler
	r.HandleFunc("/api/productos", getProductos).Methods("GET")

	// Servir app
	log.Fatal(http.ListenAndServe(":8080", r))
}
