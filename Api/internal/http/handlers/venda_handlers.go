package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/EnzoBergamaschi/ProjectGo/internal/auth"
	"github.com/EnzoBergamaschi/ProjectGo/internal/dao"
)

type VendaHandler struct {
	dao *dao.VendaDAO
}

func NovaVendaHandler(db *sql.DB) *VendaHandler {
	return &VendaHandler{dao: dao.NovaVendaDAO(db)}
}

// =========================================================
// LISTAR TODAS AS VENDAS
// =========================================================
func (h *VendaHandler) Listar(w http.ResponseWriter, r *http.Request) {
	vendas, err := h.dao.Listar()
	if err != nil {
		http.Error(w, "Erro ao listar vendas: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vendas)
}

// =========================================================
// CRIAR VENDA (total sempre começa em 0)
// =========================================================
func (h *VendaHandler) Criar(w http.ResponseWriter, r *http.Request) {
	var input struct {
		IDUsuario int    `json:"id_usuario"`
		Status    string `json:"status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	input.Status = strings.ToLower(strings.TrimSpace(input.Status))
	validStatuses := map[string]bool{
		"pendente":  true,
		"pago":      true,
		"enviado":   true,
		"cancelado": true,
	}
	if !validStatuses[input.Status] {
		input.Status = "pendente"
	}

	// Se o usuário não vier no body, pega do token JWT
	if input.IDUsuario == 0 {
		if userID, ok := auth.GetUserID(r); ok {
			input.IDUsuario = userID
		} else {
			http.Error(w, "Usuário não autenticado", http.StatusUnauthorized)
			return
		}
	}

	// Cria venda com total = 0
	if err := h.dao.Criar(input.IDUsuario, 0, input.Status); err != nil {
		http.Error(w, "Erro ao criar venda: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Venda criada com sucesso"))
}

// =========================================================
// ATUALIZAR VENDA (não mexe no total manualmente)
// =========================================================
func (h *VendaHandler) Atualizar(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/vendas/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var input struct {
		IDUsuario int    `json:"id_usuario"`
		Status    string `json:"status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	input.Status = strings.ToLower(strings.TrimSpace(input.Status))
	validStatuses := map[string]bool{
		"pendente":  true,
		"pago":      true,
		"enviado":   true,
		"cancelado": true,
	}
	if !validStatuses[input.Status] {
		input.Status = "pendente"
	}

	if err := h.dao.Atualizar(id, input.IDUsuario, 0, input.Status); err != nil {
		http.Error(w, "Erro ao atualizar venda: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Venda atualizada com sucesso"))
}

// =========================================================
// DELETAR VENDA
// =========================================================
func (h *VendaHandler) Deletar(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/vendas/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := h.dao.Deletar(id); err != nil {
		http.Error(w, "Erro ao deletar venda: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Venda deletada com sucesso"))
}
