// Package main API - Container Tracking Devices
//
// Documentação da API - Container Tracking Devices
//
//	Schemes: http
//  host: localhost:5000
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package main

import (
	"github.com/OtavioGallego/dsrpt21-container-tracking-devices/db"
	"github.com/OtavioGallego/dsrpt21-container-tracking-devices/router"
)

func main() {
	db.Iniciar()
	defer db.Conexao.Close()

	router.Gerar().Run(":5000")
}
