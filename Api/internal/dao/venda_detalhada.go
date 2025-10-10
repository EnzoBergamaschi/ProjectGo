package dao

import (
	"database/sql"
	"fmt"
)

type ItemDetalhado struct {
	Produto       string  `json:"produto"`
	Quantidade    int     `json:"quantidade"`
	PrecoUnitario float64 `json:"preco_unitario"`
}

type VendaDetalhada struct {
	IDVenda int             `json:"id_venda"`
	Cliente string          `json:"cliente"`
	Status  string          `json:"status"`
	Total   float64         `json:"total"`
	Itens   []ItemDetalhado `json:"itens"`
}

type VendaDetalhadaDAO struct {
	DB *sql.DB
}

func NovaVendaDetalhadaDAO(db *sql.DB) *VendaDetalhadaDAO {
	return &VendaDetalhadaDAO{DB: db}
}
func (v *VendaDetalhadaDAO) Listar() ([]VendaDetalhada, error) {
	vendaRows, err := v.DB.Query(`
		SELECT v.id, u.nome, v.status, v.total
		FROM vendas v
		JOIN usuarios u ON v.id_usuario = u.id
	`)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar vendas: %w", err)
	}
	defer vendaRows.Close()

	var vendas []VendaDetalhada

	for vendaRows.Next() {
		var vd VendaDetalhada
		if err := vendaRows.Scan(&vd.IDVenda, &vd.Cliente, &vd.Status, &vd.Total); err != nil {
			return nil, err
		}

		itemRows, err := v.DB.Query(`
			SELECT p.nome, i.quantidade, i.preco_unitario
			FROM itens_venda i
			JOIN produtos p ON i.id_produto = p.id
			WHERE i.id_venda = ?
		`, vd.IDVenda)
		if err != nil {
			return nil, err
		}

		var itens []ItemDetalhado
		for itemRows.Next() {
			var it ItemDetalhado
			if err := itemRows.Scan(&it.Produto, &it.Quantidade, &it.PrecoUnitario); err != nil {
				return nil, err
			}
			itens = append(itens, it)
		}
		itemRows.Close()

		vd.Itens = itens
		vendas = append(vendas, vd)
	}

	return vendas, nil
}
func (v *VendaDetalhadaDAO) ListarPorUsuario(userID int) ([]VendaDetalhada, error) {
	vendaRows, err := v.DB.Query(`
		SELECT v.id, u.nome, v.status, v.total
		FROM vendas v
		JOIN usuarios u ON v.id_usuario = u.id
		WHERE v.id_usuario = ?
	`, userID)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar vendas do usuário: %w", err)
	}
	defer vendaRows.Close()

	var vendas []VendaDetalhada

	for vendaRows.Next() {
		var vd VendaDetalhada
		if err := vendaRows.Scan(&vd.IDVenda, &vd.Cliente, &vd.Status, &vd.Total); err != nil {
			return nil, err
		}

		itemRows, err := v.DB.Query(`
			SELECT p.nome, i.quantidade, i.preco_unitario
			FROM itens_venda i
			JOIN produtos p ON i.id_produto = p.id
			WHERE i.id_venda = ?
		`, vd.IDVenda)
		if err != nil {
			return nil, err
		}

		var itens []ItemDetalhado
		for itemRows.Next() {
			var it ItemDetalhado
			if err := itemRows.Scan(&it.Produto, &it.Quantidade, &it.PrecoUnitario); err != nil {
				return nil, err
			}
			itens = append(itens, it)
		}
		itemRows.Close()

		vd.Itens = itens
		vendas = append(vendas, vd)
	}

	return vendas, nil
}
