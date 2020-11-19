package models

import "time"

// Localizacao representa uma localização registrada para um container
// swagger:model
type Localizacao struct {
	// ID da localização
	//
	// required: false
	// min: 1
	ID int `json:"id,omitempty"`

	// ID do container
	//
	// required: true
	// min: 1
	ContainerID int `json:"container_id,omitempty"`

	// Latitude da posição registrada
	//
	// required: true
	// min: 1.0
	Latitude float64 `json:"latitude" binding:"required"`

	// ID do container
	//
	// required: true
	// min: 1.0
	Longitude float64 `json:"longitude" binding:"required"`

	// ID do container
	//
	// required: false
	// format: date-time
	Data time.Time `json:"data"`
}
