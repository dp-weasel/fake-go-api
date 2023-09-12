package main

import (
	"my_api/internal/handler"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	http.HandleFunc("/datos-persona", handler.DatosPersonaHandler)

	http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(http.DefaultServeMux))
}
