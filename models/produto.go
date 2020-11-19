package models

// Produto que será armazenado dentro de um container
// swagger:model
type Produto struct {
	// ID do produto
	//
	// required: false
	// min: 1
	ID int `json:"id"`

	// Nome do produto
	//
	// required: true
	// max length: 50
	Nome string `json:"nome" binding:"required"`

	// Tipo do produto
	//
	// required: true
	// max length: 50
	Tipo string `json:"tipo" binding:"required"`

	// Indica se o produto é nocivo
	//
	// required: false
	Nocivo bool `json:"nocivo"`

	// Preço médio do produto
	//
	// required: true
	// min: 0.01
	PrecoMedio float64 `json:"preco_medio" binding:"required"`

	// Peso médio do produto
	//
	// required: true
	// min: 0.01
	PesoMedio float64 `json:"peso_medio" binding:"required"`
}
