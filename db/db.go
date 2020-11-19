package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" // MySQL Driver
)

// Conexao com o banco de dados
var Conexao *sql.DB

// Iniciar abre a conexão com o banco de dados
func Iniciar() {
	var stringConexao = fmt.Sprintf("root:root@tcp(db:3306)/dsrpt21?parseTime=true")

	db, err := sql.Open("mysql", stringConexao)
	if err != nil {
		panic(err)
	}

	var pingError = db.Ping()

	for pingError != nil {
		log.Println("Esperando docker-compose terminar")
		time.Sleep(time.Second * 5)
		pingError = db.Ping()
	}

	Conexao = db
}
