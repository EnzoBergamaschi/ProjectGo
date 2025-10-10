package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EnzoBergamaschi/ProjectGo/internal/auth"
	"github.com/EnzoBergamaschi/ProjectGo/internal/dao"
)

type AuthHandler struct {
	db *sql.DB
}

func NovoAuthHandler(db *sql.DB) *AuthHandler {
	return &AuthHandler{db: db}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email string `json:"email"`
		Senha string `json:"senha"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	var usuario dao.Usuario
	err := h.db.QueryRow(`
		SELECT id, nome, email, senha_hash, tipo 
		FROM usuarios 
		WHERE email = ?`, input.Email).
		Scan(&usuario.ID, &usuario.Nome, &usuario.Email, &usuario.SenhaHash, &usuario.Tipo)

	if err == sql.ErrNoRows {
		http.Error(w, "Usuário não encontrado", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}
	if !auth.ValidarSenha(usuario.SenhaHash, input.Senha) {
		http.Error(w, "Senha incorreta", http.StatusUnauthorized)
		return
	}
	fmt.Println("DEBUG: login de usuário =", usuario.Email, "| tipo =", usuario.Tipo)

	token, err := auth.GerarJWT(usuario.ID, usuario.Email, usuario.Tipo)
	if err != nil {
		http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": token,
		"user": map[string]interface{}{
			"id":    usuario.ID,
			"nome":  usuario.Nome,
			"email": usuario.Email,
			"tipo":  usuario.Tipo,
		},
	})
}
