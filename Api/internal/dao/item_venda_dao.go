package dao

import (
	"database/sql"
	"fmt"
)

type ItemVenda struct {
	ID            int     `json:"id"`
	IDVenda       int     `json:"id_venda"`
	IDProduto     int     `json:"id_produto"`
	Quantidade    int     `json:"quantidade"`
	PrecoUnitario float64 `json:"preco_unitario"`
}

type ItemVendaDAO struct {
	DB *sql.DB
}

func NovoItemVendaDAO(db *sql.DB) *ItemVendaDAO {
	return &ItemVendaDAO{DB: db}
}
func (d *ItemVendaDAO) Listar() ([]ItemVenda, error) {
	rows, err := d.DB.Query(`
		SELECT id, id_venda, id_produto, quantidade, preco_unitario
		FROM itens_venda
		ORDER BY id DESC
	`)
	if err != nil {
		return nil, fmt.Errorf("erro ao listar itens: %w", err)
	}
	defer rows.Close()

	var itens []ItemVenda
	for rows.Next() {
		var i ItemVenda
		if err := rows.Scan(&i.ID, &i.IDVenda, &i.IDProduto, &i.Quantidade, &i.PrecoUnitario); err != nil {
			return nil, err
		}
		itens = append(itens, i)
	}
	return itens, nil
}
func (d *ItemVendaDAO) ListarPorVenda(idVenda int) ([]ItemVenda, error) {
	if idVenda <= 0 {
		return nil, fmt.Errorf("id_venda inválido")
	}

	rows, err := d.DB.Query(`
		SELECT id, id_venda, id_produto, quantidade, preco_unitario
		FROM itens_venda
		WHERE id_venda = ?
		ORDER BY id ASC
	`, idVenda)
	if err != nil {
		return nil, fmt.Errorf("erro ao listar itens da venda: %w", err)
	}
	defer rows.Close()

	var itens []ItemVenda
	for rows.Next() {
		var i ItemVenda
		if err := rows.Scan(&i.ID, &i.IDVenda, &i.IDProduto, &i.Quantidade, &i.PrecoUnitario); err != nil {
			return nil, err
		}
		itens = append(itens, i)
	}
	return itens, nil
}
func (d *ItemVendaDAO) Criar(idVenda, idProduto, quantidade int, precoUnitario float64) error {
	if idVenda <= 0 || idProduto <= 0 || quantidade <= 0 || precoUnitario <= 0 {
		return fmt.Errorf("dados inválidos ao criar item de venda")
	}

	_, err := d.DB.Exec(`
		INSERT INTO itens_venda (id_venda, id_produto, quantidade, preco_unitario)
		VALUES (?, ?, ?, ?)
	`, idVenda, idProduto, quantidade, precoUnitario)

	if err != nil {
		return fmt.Errorf("erro ao criar item de venda: %w", err)
	}

	return nil
}
func (d *ItemVendaDAO) Atualizar(id, idVenda, idProduto, quantidade int, precoUnitario float64) error {
	if id <= 0 {
		return fmt.Errorf("id inválido para atualização de item")
	}

	_, err := d.DB.Exec(`
		UPDATE itens_venda
		SET id_venda = ?, id_produto = ?, quantidade = ?, preco_unitario = ?
		WHERE id = ?
	`, idVenda, idProduto, quantidade, precoUnitario, id)

	if err != nil {
		return fmt.Errorf("erro ao atualizar item de venda: %w", err)
	}

	return nil
}
func (d *ItemVendaDAO) Deletar(id int) error {
	if id <= 0 {
		return fmt.Errorf("id inválido para exclusão de item")
	}

	_, err := d.DB.Exec("DELETE FROM itens_venda WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("erro ao deletar item: %w", err)
	}

	return nil
}
func (d *ItemVendaDAO) BuscarIDVendaPorItem(idItem int) (int, error) {
	var idVenda int
	err := d.DB.QueryRow(`
		SELECT id_venda
		FROM itens_venda
		WHERE id = ?
	`, idItem).Scan(&idVenda)

	if err != nil {
		return 0, fmt.Errorf("erro ao buscar id_venda: %w", err)
	}
	return idVenda, nil
}
