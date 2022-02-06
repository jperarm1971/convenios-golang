package models

import (
	"time"
)

type Usuario struct {
	ID               int       `json:"id"`
	Nombre           string    `json:"nombre,omitempty"`
	Apellidos        string    `json:"apellidos,omitempty"`
	FechaNacimientos time.Time `json:"fechaNacimiento,omitempty"`
	Email            string    `json:"email"`
	Password         string    `json:"password,omitempty"`
	Avatar           string    `json:"avatar,omitempty"`
	Banner           string    `json:"banner,omitempty"`
	Biografia        string    `json:"biografia,omitempty"`
	Ubicacion        string    `json:"ubicacion,omitempty"`
	SitioWeb         string    `json:"sitioWeb,omitempty"`
}
