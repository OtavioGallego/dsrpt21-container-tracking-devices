package repositories

import (
	"database/sql"

	"github.com/OtavioGallego/dsrpt21-container-tracking-devices/models"
)

// Localizacoes representa um repositório de localizações
type Localizacoes struct {
	db *sql.DB
}

// NovoRepositorioDeLocalizacoes que recebe o banco como parâmetro
func NovoRepositorioDeLocalizacoes(db *sql.DB) *Localizacoes {
	return &Localizacoes{db}
}

// Cadastrar insere um produto no banco de dados
func (l *Localizacoes) Cadastrar(localizacao models.Localizacao) (int, error) {
	stmt, err := l.db.Prepare(`
		insert into localizacoes 
		(container_id, latitude, longitude, data)
		values (?, ?, ?, ?)
	`)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	resultado, err := stmt.Exec(
		localizacao.ContainerID,
		localizacao.Latitude,
		localizacao.Longitude,
		localizacao.Data,
	)

	if err != nil {
		return 0, err
	}

	idInserido, _ := resultado.LastInsertId()
	return int(idInserido), nil
}

// Buscar retorna todas as localizações registradas de um container
func (l *Localizacoes) Buscar(containerID int) ([]models.Localizacao, error) {
	linhas, err := l.db.Query(
		`select latitude, longitude, data from localizacoes where container_id = ?`,
		containerID,
	)

	if err != nil {
		return nil, err
	}

	defer linhas.Close()

	var localizacoes []models.Localizacao

	for linhas.Next() {
		var localizacao models.Localizacao

		err = linhas.Scan(
			&localizacao.Latitude,
			&localizacao.Longitude,
			&localizacao.Data,
		)

		if err != nil {
			return nil, err
		}

		localizacoes = append(localizacoes, localizacao)
	}

	return localizacoes, nil
}
