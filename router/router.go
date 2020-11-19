package router

import (
	"github.com/OtavioGallego/dsrpt21-container-tracking-devices/handlers"
	"github.com/gin-gonic/gin"
)

// Gerar retorna um router pronto para ser utilizado
func Gerar() *gin.Engine {
	r := gin.Default()

	r.POST("/produtos", handlers.CadastrarProduto)

	r.POST("/containers", handlers.CadastrarContainer)
	r.GET("/containers/:containerId", handlers.BuscarContainer)
	r.GET("/containers/:containerId/localizacoes", handlers.BuscarLocalizacoes)
	r.POST("/containers/:containerId/localizacoes", handlers.CadastrarLocalizacao)

	return r
}
