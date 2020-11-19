package docs

import (
	"github.com/OtavioGallego/dsrpt21-container-tracking-devices/models"
)

// Os tipos aqui contidos servem apenas para documentação

// swagger:parameters cadastrarContainer
type parametroContainer struct {
	// Container que será criado. Note que em caso de cadastro o id é ignorado
	// in: body
	// required: true
	Body models.Container
}

// swagger:parameters cadastrarProduto
type parametroProduto struct {
	// Produto que será criado. Note que em caso de cadastro o id é ignorado
	// in: body
	// required: true
	Body models.Produto
}

// swagger:response responseProduto
type responseProduto struct {
	// Produto criado
	// in: body
	Body models.Produto
}

// swagger:response responseContainer
type responseContainer struct {
	// Container criado
	// in: body
	Body models.Container
}

// swagger:response responseLocalizacao
type responseLocalizacao struct {
	// Todas as localizações registradas
	// in: body
	Body models.Localizacao
}

// swagger:response responseLocalizacoes
type responseLocalizacoes struct {
	// Todas as localizações registradas
	// in: body
	Body []models.Localizacao
}

// Mensagem de erro genérica
// swagger:response erroGenerico
type erroGenerico struct {
	// Descrição do erro
	// in: body
	Body struct {
		Mensagem string `json:"mensagem"`
	}
}

// swagger:parameters buscarContainer buscarLocalizacao cadastrarLocalizacao buscarLocalizacoes
type containerID struct {
	// Id do container a ser buscado
	// in: path
	// required: true
	ContainerID int `json:"containerId"`
}
