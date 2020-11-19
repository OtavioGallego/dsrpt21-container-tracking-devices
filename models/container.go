package models

// Container que será monitorado
// swagger:model
type Container struct {
	// ID do container
	//
	// required: false
	// min: 1
	ID int `json:"id"`

	// ID do produto armazenado no container
	//
	// required: true
	// min: 1
	ProdutoID int `json:"produto_id" binding:"required"`

	// Quantidade de produtos armazenados
	//
	// required: true
	// min: 1
	Quantidade int `json:"quantidade" binding:"required"`

	// Comprimento externo do container, em metros
	//
	// required: true
	// min: 1
	ComprimentoExterno float64 `json:"comprimento_externo" binding:"required"`

	// Largura externa do container, em metros
	//
	// required: true
	// min: 1
	LarguraExterna float64 `json:"largura_externa" binding:"required"`

	// Altura externa do container, em metros
	//
	// required: true
	// min: 1
	AlturaExterna float64 `json:"altura_externa" binding:"required"`

	// Comprimento interno do container, em metros
	//
	// required: true
	// min: 1
	ComprimentoInterno float64 `json:"comprimento_interno" binding:"required"`

	// Largura interna do container, em metros
	//
	// required: true
	// min: 1
	LarguraInterna float64 `json:"largura_interna" binding:"required"`

	// Altura interna do container, em metros
	//
	// required: true
	// min: 1
	AlturaInterna float64 `json:"altura_interna" binding:"required"`

	// Valor do carregamento, baseado na quantidade de produtos e no preço médio
	//
	// required: true
	// min: 1
	ValorCarregamento float64 `json:"valor_carregamento"`

	// Peso do carregamento, baseado na quantidade de produtos e no peso médio
	//
	// required: true
	// min: 1
	PesoCarregamento float64 `json:"peso_carregamento"`

	// Volume do carregamento, baseado em suas dimensões
	//
	// required: true
	// min: 1
	Volume float64 `json:"volume"`
}
