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

type UsuarioHandler struct {
	dao *dao.UsuarioDAO
}

func NovoUsuarioHandler(db *sql.DB) *UsuarioHandler {
	return &UsuarioHandler{dao: dao.NovoUsuarioDAO(db)}
}

// =========================================
// LISTAR USUÁRIOS
// =========================================
func (h *UsuarioHandler) Listar(w http.ResponseWriter, r *http.Request) {
	usuarios, err := h.dao.Listar()
	if err != nil {
		http.Error(w, "Erro ao listar usuários", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarios)
}

// =========================================
// CRIAR NOVO USUÁRIO
// =========================================
func (h *UsuarioHandler) Criar(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Nome  string `json:"nome"`
		Email string `json:"email"`
		Senha string `json:"senha"`
		Tipo  string `json:"tipo,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if input.Nome == "" || input.Email == "" || input.Senha == "" {
		http.Error(w, "Campos obrigatórios ausentes", http.StatusBadRequest)
		return
	}

	hash, err := auth.GerarHashSenha(input.Senha)
	if err != nil {
		http.Error(w, "Erro ao gerar hash da senha", http.StatusInternalServerError)
		return
	}

	if input.Tipo == "" {
		input.Tipo = "cliente"
	}

	if err := h.dao.Criar(input.Nome, input.Email, hash, input.Tipo); err != nil {
		http.Error(w, "Erro ao criar usuário", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"mensagem": "Usuário criado com sucesso",
	})
}

// =========================================
// ATUALIZAR USUÁRIO
// =========================================
func (h *UsuarioHandler) Atualizar(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var input struct {
		Nome  string `json:"nome"`
		Email string `json:"email"`
		Tipo  string `json:"tipo"`
		Senha string `json:"senha,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if input.Nome == "" || input.Email == "" || input.Tipo == "" {
		http.Error(w, "Campos obrigatórios ausentes", http.StatusBadRequest)
		return
	}

	// Atualiza com ou sem senha
	if strings.TrimSpace(input.Senha) != "" {
		hash, err := auth.GerarHashSenha(input.Senha)
		if err != nil {
			http.Error(w, "Erro ao gerar hash da senha", http.StatusInternalServerError)
			return
		}
		err = h.dao.Atualizar(id, input.Nome, input.Email, hash, input.Tipo)
		if err != nil {
			http.Error(w, "Erro ao atualizar usuário", http.StatusInternalServerError)
			return
		}
	} else {
		err = h.dao.AtualizarSemSenha(id, input.Nome, input.Email, input.Tipo)
		if err != nil {
			http.Error(w, "Erro ao atualizar usuário", http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"mensagem": "Usuário atualizado com sucesso"})
}

// =========================================
// DELETAR USUÁRIO
// =========================================
func (h *UsuarioHandler) Deletar(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := h.dao.Deletar(id); err != nil {
		http.Error(w, "Erro ao deletar usuário", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
