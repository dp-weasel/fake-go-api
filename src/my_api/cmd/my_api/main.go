package main

import (
	"my_api/internal/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/datos-persona", handler.DatosPersonaHandler)
	http.ListenAndServe(":8080", nil)
}
