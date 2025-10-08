package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/EnzoBergamaschi/ProjectGo/internal/auth"
	"github.com/EnzoBergamaschi/ProjectGo/internal/dao"
)

type VendaDetalhadaHandler struct {
	dao *dao.VendaDetalhadaDAO
}

func NovaVendaDetalhadaHandler(db *sql.DB) *VendaDetalhadaHandler {
	return &VendaDetalhadaHandler{dao: dao.NovaVendaDetalhadaDAO(db)}
}

func (h *VendaDetalhadaHandler) Listar(w http.ResponseWriter, r *http.Request) {
	userTipo, _ := auth.GetUserTipo(r)
	if userTipo == "cliente" {
		userID, ok := auth.GetUserID(r)
		if !ok {
			http.Error(w, "Usuário não autenticado", http.StatusUnauthorized)
			return
		}
		vendas, err := h.dao.ListarPorUsuario(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(vendas)
		return
	}

	vendas, err := h.dao.Listar()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vendas)
}
