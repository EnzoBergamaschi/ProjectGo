package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/EnzoBergamaschi/ProjectGo/internal/dao"
)

type ItemVendaHandler struct {
	dao *dao.ItemVendaDAO
	db  *sql.DB
}

func NovoItemVendaHandler(db *sql.DB) *ItemVendaHandler {
	return &ItemVendaHandler{
		dao: dao.NovoItemVendaDAO(db),
		db:  db,
	}
}

// ============================================================
// LISTAR ITENS DE UMA VENDA
// ============================================================
func (h *ItemVendaHandler) ListarPorVenda(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/itens_venda/")
	idVenda, err := strconv.Atoi(idStr)
	if err != nil || idVenda <= 0 {
		http.Error(w, "ID da venda inválido", http.StatusBadRequest)
		return
	}

	itens, err := h.dao.ListarPorVenda(idVenda)
	if err != nil {
		http.Error(w, "Erro ao listar itens: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(itens)
}

// ============================================================
// ADICIONAR ITEM
// ============================================================
func (h *ItemVendaHandler) Criar(w http.ResponseWriter, r *http.Request) {
	var input struct {
		IDVenda       int     `json:"id_venda"`
		IDProduto     int     `json:"id_produto"`
		Quantidade    int     `json:"quantidade"`
		PrecoUnitario float64 `json:"preco_unitario"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if input.IDVenda <= 0 || input.IDProduto <= 0 || input.Quantidade <= 0 || input.PrecoUnitario <= 0 {
		http.Error(w, "Campos obrigatórios ausentes ou inválidos", http.StatusBadRequest)
		return
	}

	if err := h.dao.Criar(input.IDVenda, input.IDProduto, input.Quantidade, input.PrecoUnitario); err != nil {
		http.Error(w, "Erro ao criar item: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 🔄 Atualiza automaticamente o total da venda
	vendaDAO := dao.NovaVendaDAO(h.db)
	if err := vendaDAO.AtualizarTotalPelosItens(input.IDVenda); err != nil {
		http.Error(w, "Erro ao atualizar total da venda: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Item adicionado com sucesso"))
}

// ============================================================
// DELETAR ITEM
// ============================================================
func (h *ItemVendaHandler) Deletar(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/itens_venda/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Recupera o id_venda do item antes de excluir
	idVenda, err := h.dao.BuscarIDVendaPorItem(id)
	if err != nil {
		http.Error(w, "Erro ao buscar venda do item: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.dao.Deletar(id); err != nil {
		http.Error(w, "Erro ao deletar item: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 🔄 Atualiza automaticamente o total da venda
	vendaDAO := dao.NovaVendaDAO(h.db)
	if err := vendaDAO.AtualizarTotalPelosItens(idVenda); err != nil {
		http.Error(w, "Erro ao atualizar total da venda: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Item excluído com sucesso"))
}
