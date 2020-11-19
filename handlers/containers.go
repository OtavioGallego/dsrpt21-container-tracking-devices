package handlers

import (
	"net/http"
	"strconv"

	"github.com/OtavioGallego/dsrpt21-container-tracking-devices/db"
	"github.com/OtavioGallego/dsrpt21-container-tracking-devices/repositories"

	"github.com/OtavioGallego/dsrpt21-container-tracking-devices/models"
	"github.com/gin-gonic/gin"
)

// swagger:route GET /containers/{containerId} Containers buscarContainer
// Returna um container salvo no banco de dados
// responses:
//	200: responseContainer
//  400: erroGenerico
//  500: erroGenerico

// BuscarContainer lida com a requisição para busca de container
func BuscarContainer(c *gin.Context) {

	containerID, err := strconv.Atoi(c.Param("containerId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	repositorioDeContainers := repositories.NovoRepositorioDeContainers(db.Conexao)
	container, err := repositorioDeContainers.BuscarPorID(containerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"container": container})
}

// swagger:route POST /containers Containers cadastrarContainer
// Cadastra um container que comporta um produto
//
// responses:
//	201: responseContainer
//  400: erroGenerico
//  501: erroGenerico

// CadastrarContainer lida com a requisição para criação de container
func CadastrarContainer(c *gin.Context) {
	var (
		container               models.Container
		repositorioDeContainers = repositories.NovoRepositorioDeContainers(db.Conexao)
		repositorioDeProdutos   = repositories.NovoRepositorioDeProdutos(db.Conexao)
	)

	err := c.ShouldBindJSON(&container)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	produto, err := repositorioDeProdutos.BuscarPorID(container.ProdutoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	container.ValorCarregamento = produto.PrecoMedio * float64(container.Quantidade)
	container.PesoCarregamento = produto.PesoMedio * float64(container.Quantidade)
	container.Volume = container.ComprimentoInterno * container.LarguraInterna * container.AlturaInterna

	container.ID, err = repositorioDeContainers.Cadastrar(container)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"container": container})
}
