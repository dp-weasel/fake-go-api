package handler

import (
	"encoding/json"
	"fmt"
	"my_api/internal/model"
	"net/http"
	"strconv"
)

func DatosPersonaHandler(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	pageNumber, err := strconv.Atoi(page)

	if err != nil || pageNumber < 1 {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}

	// Datos dummy para el ejemplo
	totalData := 45
	perPage := 10
	totalPages := (totalData + perPage - 1) / perPage

	data := model.PersonDataResponse{
		Page:        pageNumber,
		TotalPages:  totalPages,
		TotalSource: totalData,
		Columns: []model.Column{
			{Header: "Nombre", Field: "nombre"},
			{Header: "Apellido", Field: "apellido"},
			{Header: "Correo", Field: "correo"},
		},
		Source: []model.Person{
			{ID: "1", Nombre: "Juan", Apellido: "Diaz", Correo: "algo@algo.com"},
			{ID: "2", Nombre: "María", Apellido: "González", Correo: "maria@email.com"},
			{ID: "3", Nombre: "Carlos", Apellido: "López", Correo: "carlos@correo.com"},
			{ID: "4", Nombre: "Ana", Apellido: "Martínez", Correo: "ana@gmail.com"},
			{ID: "5", Nombre: "Luis", Apellido: "Rodríguez", Correo: "luis@hotmail.com"},
			{ID: "6", Nombre: "Laura", Apellido: "Pérez", Correo: "laura@correo.com"},
			{ID: "7", Nombre: "Pedro", Apellido: "Sanchez", Correo: "pedro@email.com"},
			{ID: "8", Nombre: "Sofía", Apellido: "Ramírez", Correo: "sofia@gmail.com"},
			{ID: "9", Nombre: "Diego", Apellido: "Hernández", Correo: "diego@hotmail.com"},
			{ID: "10", Nombre: "Isabella", Apellido: "Luna", Correo: "isabella@correo.com"},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
}
