package repositories

import (
	"database/sql"

	"github.com/OtavioGallego/dsrpt21-container-tracking-devices/models"
)

// Produtos representa um repositório de produtos
type Produtos struct {
	db *sql.DB
}

// NovoRepositorioDeProdutos que recebe o banco como parâmetro
func NovoRepositorioDeProdutos(db *sql.DB) *Produtos {
	return &Produtos{db}
}

// Cadastrar insere um produto no banco de dados
func (p *Produtos) Cadastrar(produto models.Produto) (int, error) {
	stmt, err := p.db.Prepare(`
		insert into produtos (nome, tipo, nocivo, preco_medio, peso_medio)
		values (?, ?, ?, ?, ?)
	`)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	resultado, err := stmt.Exec(
		produto.Nome,
		produto.Tipo,
		produto.Nocivo,
		produto.PrecoMedio,
		produto.PesoMedio,
	)

	if err != nil {
		return 0, err
	}

	idInserido, _ := resultado.LastInsertId()

	return int(idInserido), nil
}

// BuscarPorID retorna um produto
func (p *Produtos) BuscarPorID(ID int) (models.Produto, error) {
	var produto models.Produto

	linhas, err := p.db.Query(`select * from produtos where id = ?`, ID)
	if err != nil {
		return produto, err
	}
	defer linhas.Close()

	if linhas.Next() {
		err = linhas.Scan(
			&produto.ID,
			&produto.Nome,
			&produto.Tipo,
			&produto.Nocivo,
			&produto.PrecoMedio,
			&produto.PesoMedio,
		)

		if err != nil {
			return produto, err
		}
	}

	return produto, nil
}
