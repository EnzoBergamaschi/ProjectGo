package dao

import (
	"database/sql"
	"fmt"
)

type Venda struct {
	ID          int     `json:"id"`
	IDUsuario   int     `json:"id_usuario"`
	UsuarioNome string  `json:"usuario_nome"`
	Total       float64 `json:"total"`
	Status      string  `json:"status"`
}

type VendaDAO struct {
	DB *sql.DB
}

func NovaVendaDAO(db *sql.DB) *VendaDAO {
	return &VendaDAO{DB: db}
}

// ============================================================
// LISTAR todas as vendas com nome do usuário
// ============================================================
func (v *VendaDAO) Listar() ([]Venda, error) {
	rows, err := v.DB.Query(`
		SELECT ve.id, ve.id_usuario, u.nome AS usuario_nome, ve.total, ve.status
		FROM vendas ve
		JOIN usuarios u ON u.id = ve.id_usuario
		ORDER BY ve.id DESC
	`)
	if err != nil {
		return nil, fmt.Errorf("erro ao listar vendas: %w", err)
	}
	defer rows.Close()

	var vendas []Venda
	for rows.Next() {
		var venda Venda
		if err := rows.Scan(&venda.ID, &venda.IDUsuario, &venda.UsuarioNome, &venda.Total, &venda.Status); err != nil {
			return nil, err
		}
		vendas = append(vendas, venda)
	}
	return vendas, nil
}

// ============================================================
// CRIAR nova venda
// ============================================================
func (v *VendaDAO) Criar(idUsuario int, total float64, status string) error {
	if idUsuario <= 0 {
		return fmt.Errorf("id_usuario inválido")
	}

	// Ao criar uma venda, ignora o total passado (deixa o backend recalcular)
	_, err := v.DB.Exec(`
		INSERT INTO vendas (id_usuario, total, status)
		VALUES (?, 0, ?)
	`, idUsuario, status)
	if err != nil {
		return fmt.Errorf("erro ao criar venda: %w", err)
	}
	return nil
}

// ============================================================
// ATUALIZAR venda
// ============================================================
func (v *VendaDAO) Atualizar(id, idUsuario int, total float64, status string) error {
	if id <= 0 {
		return fmt.Errorf("id inválido")
	}

	// Atualiza apenas o status e o usuário. O total é recalculado pelos itens.
	_, err := v.DB.Exec(`
		UPDATE vendas
		SET id_usuario = ?, status = ?
		WHERE id = ?
	`, idUsuario, status, id)

	if err != nil {
		return fmt.Errorf("erro ao atualizar venda: %w", err)
	}

	// Garante que o total fique sincronizado com os itens
	if err := v.AtualizarTotalPelosItens(id); err != nil {
		return fmt.Errorf("erro ao recalcular total: %w", err)
	}

	return nil
}

// ============================================================
// DELETAR venda
// ============================================================
func (v *VendaDAO) Deletar(id int) error {
	if id <= 0 {
		return fmt.Errorf("id inválido")
	}
	_, err := v.DB.Exec("DELETE FROM vendas WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("erro ao deletar venda: %w", err)
	}
	return nil
}

// ============================================================
// REATUALIZAR total com base nos itens da venda
// ============================================================
func (v *VendaDAO) AtualizarTotalPelosItens(idVenda int) error {
	if idVenda <= 0 {
		return fmt.Errorf("id_venda inválido")
	}

	_, err := v.DB.Exec(`
		UPDATE vendas
		SET total = (
			SELECT COALESCE(SUM(iv.quantidade * iv.preco_unitario), 0)
			FROM itens_venda iv
			WHERE iv.id_venda = ?
		)
		WHERE id = ?
	`, idVenda, idVenda)

	if err != nil {
		return fmt.Errorf("erro ao atualizar total da venda: %w", err)
	}

	return nil
}
