package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/OtavioGallego/dsrpt21-container-tracking-devices/db"
	"github.com/OtavioGallego/dsrpt21-container-tracking-devices/models"
	"github.com/OtavioGallego/dsrpt21-container-tracking-devices/repositories"
	"github.com/gin-gonic/gin"
)

// swagger:route Post /containers/{containerId}/localizacoes Localizações cadastrarLocalizacao
// Cadastra uma localização para um container
// responses:
//	201: responseLocalizacao
//  400: erroGenerico
//  500: erroGenerico

// CadastrarLocalizacao registra uma localização para um container
func CadastrarLocalizacao(c *gin.Context) {
	var localizacao models.Localizacao

	err := c.ShouldBindJSON(&localizacao)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	containerID, err := strconv.Atoi(c.Param("containerId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	localizacao.ContainerID = containerID
	localizacao.Data = time.Now().UTC()

	repo := repositories.NovoRepositorioDeLocalizacoes(db.Conexao)
	localizacao.ID, err = repo.Cadastrar(localizacao)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"localizacao": localizacao})
}

// swagger:route GET /containers/{containerId}/localizacoes Localizações buscarLocalizacoes
// Returna todas as localizaçṍes registradas para um container
// responses:
//	200: responseLocalizacoes
//  400: erroGenerico
//  500: erroGenerico

// BuscarLocalizacoes busca as localizações de um container
func BuscarLocalizacoes(c *gin.Context) {
	containerID, err := strconv.Atoi(c.Param("containerId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	repositorioDeLocalizacoes := repositories.NovoRepositorioDeLocalizacoes(db.Conexao)
	localizacoes, err := repositorioDeLocalizacoes.Buscar(containerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"localizacoes": localizacoes})
}
