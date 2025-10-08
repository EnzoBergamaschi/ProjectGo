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

type ProdutoHandler struct {
	dao *dao.ProdutoDAO
}

func NovoProdutoHandler(db *sql.DB) *ProdutoHandler {
	return &ProdutoHandler{dao: dao.NovoProdutoDAO(db)}
}

func (h *ProdutoHandler) Listar(w http.ResponseWriter, r *http.Request) {
	produtos, err := h.dao.Listar()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produtos)
}

func (h *ProdutoHandler) Criar(w http.ResponseWriter, r *http.Request) {
	userTipo, _ := auth.GetUserTipo(r)
	if userTipo != "admin" {
		http.Error(w, "Apenas administradores podem criar produtos", http.StatusForbidden)
		return
	}

	var input dao.Produto
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	err := h.dao.Criar(input.Nome, input.Descricao, input.Preco, input.Estoque)
	if err != nil {
		http.Error(w, "Erro ao criar produto", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Produto criado com sucesso"))
}

func (h *ProdutoHandler) Atualizar(w http.ResponseWriter, r *http.Request) {
	userTipo, _ := auth.GetUserTipo(r)
	if userTipo != "admin" {
		http.Error(w, "Apenas administradores podem atualizar produtos", http.StatusForbidden)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/produtos/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var input dao.Produto
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	err = h.dao.Atualizar(id, input.Nome, input.Descricao, input.Preco, input.Estoque)
	if err != nil {
		http.Error(w, "Erro ao atualizar produto", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Produto atualizado com sucesso"))
}

func (h *ProdutoHandler) Deletar(w http.ResponseWriter, r *http.Request) {
	userTipo, _ := auth.GetUserTipo(r)
	if userTipo != "admin" {
		http.Error(w, "Apenas administradores podem deletar produtos", http.StatusForbidden)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/produtos/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = h.dao.Deletar(id)
	if err != nil {
		http.Error(w, "Erro ao deletar produto", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Produto deletado com sucesso"))
}
