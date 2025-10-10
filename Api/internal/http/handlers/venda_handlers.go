package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
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
func (h *VendaHandler) Listar(w http.ResponseWriter, r *http.Request) {
	vendas, err := h.dao.Listar()
	if err != nil {
		http.Error(w, "Erro ao listar vendas: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vendas)
}
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
	if input.IDUsuario == 0 {
		if userID, ok := auth.GetUserID(r); ok {
			input.IDUsuario = userID
		} else {
			http.Error(w, "Usuário não autenticado", http.StatusUnauthorized)
			return
		}
	}

	if err := h.dao.Criar(input.IDUsuario, 0, input.Status); err != nil {
		log.Printf("Erro ao criar venda: %v", err)
		http.Error(w, "Erro ao criar venda: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"mensagem": "Venda criada com sucesso",
	})
}
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
		log.Printf("Erro ao atualizar venda: %v", err)
		http.Error(w, "Erro ao atualizar venda: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"mensagem": "Venda atualizada com sucesso",
	})
}
func (h *VendaHandler) Deletar(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/vendas/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := h.dao.Deletar(id); err != nil {
		log.Printf("Erro ao deletar venda: %v", err)
		http.Error(w, "Erro ao deletar venda: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"mensagem": "Venda deletada com sucesso",
	})
}
