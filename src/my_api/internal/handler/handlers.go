package handler

import (
	"encoding/json"
	"fmt"
	"math"
	"my_api/internal/model"
	"net/http"
	"strconv"
)

func DatosPersonaHandler(w http.ResponseWriter, r *http.Request) {
	allPersons := []model.Person{
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
	}
	// Recuperar el número de página desde la URL
	page := r.URL.Query().Get("page")
	pageNumber, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Parámetro de página inválido", http.StatusBadRequest)
		return
	}

	const personasPerPage = 10
	startIndex := (pageNumber - 1) * personasPerPage
	endIndex := startIndex + personasPerPage

	// Validar el número de página
	if startIndex < 0 || startIndex >= len(allPersons) {
		http.Error(w, "Número de página fuera de rango", http.StatusBadRequest)
		return
	}

	endIndex = int(math.Min(float64(endIndex), float64(len(allPersons))))
	personasToReturn := allPersons[startIndex:endIndex]

	// Datos dummy para el ejemplo
	totalData := 10
	perPage := 10
	totalPages := (totalData + perPage - 1) / perPage

	data := model.PersonDataResponse{
		Page:        pageNumber,
		TotalPages:  totalPages,
		TotalSource: len(allPersons),
		Columns: []model.Column{
			{Header: "Nombre", Field: "nombre"},
			{Header: "Apellido", Field: "apellido"},
			{Header: "Correo", Field: "correo"},
		},
		Source: personasToReturn,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
}

func ListarProvinciasHandler(w http.ResponseWriter, r *http.Request) {
	// Todas las provincias (24 en este ejemplo)
	allProvinces := []model.Province{
		{NombreProvincia: "Buenos Aires", CodigoPostal: "B", CapitalProvincia: "La Plata", Superficie: 307521},
		{NombreProvincia: "Catamarca", CodigoPostal: "K", CapitalProvincia: "San Fernando del Valle de Catamarca", Superficie: 102606},
		{NombreProvincia: "Chaco", CodigoPostal: "H", CapitalProvincia: "Resistencia", Superficie: 99633},
		{NombreProvincia: "Chubut", CodigoPostal: "U", CapitalProvincia: "Rawson", Superficie: 224686},
		{NombreProvincia: "Córdoba", CodigoPostal: "X", CapitalProvincia: "Córdoba", Superficie: 165321},
		{NombreProvincia: "Corrientes", CodigoPostal: "W", CapitalProvincia: "Corrientes", Superficie: 88199},
		{NombreProvincia: "Entre Ríos", CodigoPostal: "E", CapitalProvincia: "Paraná", Superficie: 78781},
		{NombreProvincia: "Formosa", CodigoPostal: "P", CapitalProvincia: "Formosa", Superficie: 72066},
		{NombreProvincia: "Jujuy", CodigoPostal: "Y", CapitalProvincia: "San Salvador de Jujuy", Superficie: 53219},
		{NombreProvincia: "La Pampa", CodigoPostal: "L", CapitalProvincia: "Santa Rosa", Superficie: 143440},
		{NombreProvincia: "La Rioja", CodigoPostal: "F", CapitalProvincia: "La Rioja", Superficie: 89680},
		{NombreProvincia: "Mendoza", CodigoPostal: "M", CapitalProvincia: "Mendoza", Superficie: 148827},
		{NombreProvincia: "Misiones", CodigoPostal: "N", CapitalProvincia: "Posadas", Superficie: 29801},
		{NombreProvincia: "Neuquén", CodigoPostal: "Q", CapitalProvincia: "Neuquén", Superficie: 94078},
		{NombreProvincia: "Río Negro", CodigoPostal: "R", CapitalProvincia: "Viedma", Superficie: 203013},
		{NombreProvincia: "Salta", CodigoPostal: "A", CapitalProvincia: "Salta", Superficie: 155488},
		{NombreProvincia: "San Juan", CodigoPostal: "J", CapitalProvincia: "San Juan", Superficie: 89651},
		{NombreProvincia: "San Luis", CodigoPostal: "D", CapitalProvincia: "San Luis", Superficie: 76748},
		{NombreProvincia: "Santa Cruz", CodigoPostal: "Z", CapitalProvincia: "Río Gallegos", Superficie: 243943},
		{NombreProvincia: "Santa Fe", CodigoPostal: "S", CapitalProvincia: "Santa Fe", Superficie: 133007},
		{NombreProvincia: "Santiago del Estero", CodigoPostal: "G", CapitalProvincia: "Santiago del Estero", Superficie: 136351},
		{NombreProvincia: "Tierra del Fuego", CodigoPostal: "V", CapitalProvincia: "Ushuaia", Superficie: 21263},
		{NombreProvincia: "Tucumán", CodigoPostal: "T", CapitalProvincia: "San Miguel de Tucumán", Superficie: 22524},
	}

	// Recuperar el número de página desde la URL
	page := r.URL.Query().Get("page")
	pageNumber, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Parámetro de página inválido", http.StatusBadRequest)
		return
	}

	const provincesPerPage = 5
	startIndex := (pageNumber - 1) * provincesPerPage
	endIndex := startIndex + provincesPerPage

	// Validar el número de página
	if startIndex < 0 || startIndex >= len(allProvinces) {
		http.Error(w, "Número de página fuera de rango", http.StatusBadRequest)
		return
	}

	endIndex = int(math.Min(float64(endIndex), float64(len(allProvinces))))
	provincesToReturn := allProvinces[startIndex:endIndex]

	columns := []model.Column{
		{Header: "Nombre de Provincia", Field: "nombreProvincia"},
		{Header: "Codigo Postal", Field: "codigoPostal"},
		{Header: "Capital Provincia", Field: "capitalProvincia"},
		{Header: "Superficie", Field: "superficie"},
	}

	data := model.ProvinceResponse{
		Page:        pageNumber,
		TotalPages:  int(math.Ceil(float64(len(allProvinces)) / float64(provincesPerPage))),
		TotalSource: len(allProvinces),
		Columns:     columns,
		Source:      provincesToReturn,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
}
