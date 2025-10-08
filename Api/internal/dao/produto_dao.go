package dao

import (
	"database/sql"
	"fmt"
)

type Produto struct {
	ID        int     `json:"id"`
	Nome      string  `json:"nome"`
	Descricao string  `json:"descricao"`
	Preco     float64 `json:"preco"`
	Estoque   int     `json:"estoque"`
}

type ProdutoDAO struct {
	DB *sql.DB
}

func NovoProdutoDAO(db *sql.DB) *ProdutoDAO {
	return &ProdutoDAO{DB: db}
}

func (p *ProdutoDAO) Listar() ([]Produto, error) {
	rows, err := p.DB.Query("SELECT id, nome, descricao, preco, estoque FROM produtos")
	if err != nil {
		return nil, fmt.Errorf("erro ao listar produtos: %w", err)
	}
	defer rows.Close()

	var produtos []Produto
	for rows.Next() {
		var prod Produto
		if err := rows.Scan(&prod.ID, &prod.Nome, &prod.Descricao, &prod.Preco, &prod.Estoque); err != nil {
			return nil, err
		}
		produtos = append(produtos, prod)
	}
	return produtos, nil
}

func (p *ProdutoDAO) Criar(nome, descricao string, preco float64, estoque int) error {
	_, err := p.DB.Exec(
		"INSERT INTO produtos (nome, descricao, preco, estoque) VALUES (?, ?, ?, ?)",
		nome, descricao, preco, estoque,
	)
	return err
}

// ✅ Atualizar produto
func (p *ProdutoDAO) Atualizar(id int, nome, descricao string, preco float64, estoque int) error {
	_, err := p.DB.Exec(
		"UPDATE produtos SET nome = ?, descricao = ?, preco = ?, estoque = ? WHERE id = ?",
		nome, descricao, preco, estoque, id,
	)
	return err
}

// ✅ Deletar produto
func (p *ProdutoDAO) Deletar(id int) error {
	_, err := p.DB.Exec("DELETE FROM produtos WHERE id = ?", id)
	return err
}
