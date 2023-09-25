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
	PageSize    int      `json:"page_size"`
	Columns     []Column `json:"columns"`
	Source      []Person `json:"source"`
}

type Province struct {
	NombreProvincia  string `json:"nombreProvincia"`
	CodigoPostal     string `json:"codigoPostal"`
	CapitalProvincia string `json:"capitalProvincia"`
	Superficie       int    `json:"superficie"`
}

type ProvinceResponse struct {
	Page        int        `json:"page"`
	TotalPages  int        `json:"total_pages"`
	TotalSource int        `json:"total_source"`
	PageSize    int        `json:"page_size"`
	Columns     []Column   `json:"columns"`
	Source      []Province `json:"source"`
}

// Struct para representar un campo de formulario
type FormField struct {
	Name      string   `json:"name"`
	Value     string   `json:"value"`
	Type      string   `json:"type"`
	ID        string   `json:"id"`
	InputType string   `json:"inputType"`
	Required  bool     `json:"required,omitempty"`
	Options   []Option `json:"options,omitempty"`
}

// Struct para representar las opciones de selección en el campo de formulario
type Option struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Value string `json:"value"`
}
