package repositories

import (
	"database/sql"

	"github.com/OtavioGallego/dsrpt21-container-tracking-devices/models"
)

// Containers representa um repositório de containers
type Containers struct {
	db *sql.DB
}

// NovoRepositorioDeContainers que recebe o banco como parâmetro
func NovoRepositorioDeContainers(db *sql.DB) *Containers {
	return &Containers{db}
}

// Cadastrar insere um container no banco de dados
func (c *Containers) Cadastrar(container models.Container) (int, error) {
	stmt, err := c.db.Prepare(`
	insert into containers (
		produto_id,
		quantidade,
		comprimento_externo,
		largura_externa,
		altura_externa,
		comprimento_interno,
		largura_interna,
		altura_interna,
		valor_carregamento,
		peso_carregamento,
		volume
	) values (
		 ?,
		 ?,
		 ?,
		 ?,
		 ?,
		 ?,
		 ?,
		 ?,
		 ?,
		 ?,
		 ?
	)`)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	resultado, err := stmt.Exec(
		container.ProdutoID,
		container.Quantidade,
		container.ComprimentoExterno,
		container.LarguraExterna,
		container.AlturaExterna,
		container.ComprimentoInterno,
		container.LarguraInterna,
		container.AlturaInterna,
		container.ValorCarregamento,
		container.PesoCarregamento,
		container.Volume,
	)

	if err != nil {
		return 0, err
	}

	idInserido, _ := resultado.LastInsertId()

	return int(idInserido), nil
}

// BuscarPorID retorna um container
func (c *Containers) BuscarPorID(ID int) (models.Container, error) {
	var container models.Container

	linhas, err := c.db.Query("select * from containers where id = ?", ID)
	if err != nil {
		return container, err
	}
	defer linhas.Close()

	if linhas.Next() {
		err = linhas.Scan(
			&container.ID,
			&container.ProdutoID,
			&container.Quantidade,
			&container.ComprimentoExterno,
			&container.LarguraExterna,
			&container.AlturaExterna,
			&container.ComprimentoInterno,
			&container.LarguraInterna,
			&container.AlturaInterna,
			&container.ValorCarregamento,
			&container.PesoCarregamento,
			&container.Volume,
		)

		if err != nil {
			return container, err
		}
	}

	return container, nil
}
