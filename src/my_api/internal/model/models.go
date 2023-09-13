package model

type Column struct {
	Header string `json:"header"`
	Field  string `json:"field"`
}

type Person struct {
	ID       string `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Correo   string `json:"correo"`
}

type PersonDataResponse struct {
	Page        int      `json:"page"`
	TotalPages  int      `json:"total_pages"`
	TotalSource int      `json:"total_source"`
	Columns     []Column `json:"columns"`
	Source      []Person `json:"source"`
}

type Province struct {
	NombreProvincia  string `json:"nombre_provincia"`
	CodigoPostal     string `json:"codigo_postal"`
	CapitalProvincia string `json:"capital_provincia"`
	Superficie       int    `json:"superficie"`
}

type ProvinceResponse struct {
	Page        int        `json:"page"`
	TotalPages  int        `json:"total_pages"`
	TotalSource int        `json:"total_source"`
	Columns     []Column   `json:"columns"`
	Source      []Province `json:"source"`
}
