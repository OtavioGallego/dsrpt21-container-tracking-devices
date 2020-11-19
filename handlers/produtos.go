package handlers

import (
	"net/http"

	"github.com/OtavioGallego/dsrpt21-container-tracking-devices/db"
	"github.com/OtavioGallego/dsrpt21-container-tracking-devices/models"
	"github.com/OtavioGallego/dsrpt21-container-tracking-devices/repositories"
	"github.com/gin-gonic/gin"
)

// swagger:route POST /produtos Produtos cadastrarProduto
// Cadastra um novo produto que pode ser armazenado em um container
//
// responses:
//	200: responseProduto
//  400: erroGenerico
//  501: erroGenerico

// CadastrarProduto lida com a requisição para criação de produto
func CadastrarProduto(c *gin.Context) {
	var produto models.Produto

	err := c.ShouldBindJSON(&produto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	repo := repositories.NovoRepositorioDeProdutos(db.Conexao)
	produto.ID, err = repo.Cadastrar(produto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"produto": produto})

}
