// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type DeletePacienteResponse struct {
	Success bool   `json:"success"`
	ID      string `json:"id"`
}

type Paciente struct {
	ID        string  `json:"id"`
	Nombre    string  `json:"nombre"`
	Apellido  string  `json:"apellido"`
	Vsemana   int     `json:"Vsemana"`
	Paga      int     `json:"Paga"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
	DeletedAt *string `json:"deletedAt,omitempty"`
}

type PacienteInput struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Vsemana  int    `json:"Vsemana"`
	Paga     int    `json:"Paga"`
}

type UpdatePaciente struct {
	Nombre   *string `json:"nombre,omitempty"`
	Apellido *string `json:"apellido,omitempty"`
	Vsemana  *int    `json:"Vsemana,omitempty"`
	Paga     *int    `json:"Paga,omitempty"`
}

type Terapista struct {
	ID        string  `json:"id"`
	Nombre    string  `json:"nombre"`
	Apellido  string  `json:"apellido"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
	DeletedAt *string `json:"deletedAt,omitempty"`
}
