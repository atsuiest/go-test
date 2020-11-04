package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Model
type Producto struct {
	ID     int64  `json:"id"`
	Nombre string `json:"nombre"`
	Desc   string `json:"desc"`
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
	log.Fatal(http.ListenAndServe(":8080", setHeaders(r)))
}

func setHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-Type", "application/json, text/plain")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Content-Type, Apikey, "+
				"Access-Control-Allow-Origin, Access-Control-Allow-Headers, Access-Control-Allow-Methods,  Origin, x-ibm-client-id, X-Authorization-Type, X-System-Type")
			w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Strict-Transport-Security", "max-age=86400; includeSubDomains")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		next.ServeHTTP(w, r)
	})
}
