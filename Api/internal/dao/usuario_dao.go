package dao

import (
	"database/sql"
	"fmt"
)

type Usuario struct {
	ID        int    `json:"id"`
	Nome      string `json:"nome"`
	Email     string `json:"email"`
	SenhaHash string `json:"senha_hash"`
	Tipo      string `json:"tipo"`
}

type UsuarioDAO struct {
	DB *sql.DB
}

func NovoUsuarioDAO(db *sql.DB) *UsuarioDAO {
	return &UsuarioDAO{DB: db}
}

func (u *UsuarioDAO) Listar() ([]Usuario, error) {
	rows, err := u.DB.Query("SELECT id, nome, email, tipo FROM usuarios")
	if err != nil {
		return nil, fmt.Errorf("erro ao listar usuários: %w", err)
	}
	defer rows.Close()

	var usuarios []Usuario
	for rows.Next() {
		var usr Usuario
		if err := rows.Scan(&usr.ID, &usr.Nome, &usr.Email, &usr.Tipo); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usr)
	}

	return usuarios, nil
}

func (u *UsuarioDAO) Criar(nome, email, senhaHash, tipo string) error {
	_, err := u.DB.Exec(
		"INSERT INTO usuarios (nome, email, senha_hash, tipo) VALUES (?, ?, ?, ?)",
		nome, email, senhaHash, tipo,
	)
	return err
}

// ✅ Novo método: Atualizar usuário existente
func (u *UsuarioDAO) Atualizar(id int, nome, email, senhaHash, tipo string) error {
	_, err := u.DB.Exec(
		"UPDATE usuarios SET nome = ?, email = ?, senha_hash = ?, tipo = ? WHERE id = ?",
		nome, email, senhaHash, tipo, id,
	)
	return err
}

// ✅ Novo: atualiza sem mexer na senha
func (u *UsuarioDAO) AtualizarSemSenha(id int, nome, email, tipo string) error {
	_, err := u.DB.Exec(
		"UPDATE usuarios SET nome = ?, email = ?, tipo = ? WHERE id = ?",
		nome, email, tipo, id,
	)
	return err
}

// ✅ Novo método: Deletar usuário
func (u *UsuarioDAO) Deletar(id int) error {
	_, err := u.DB.Exec("DELETE FROM usuarios WHERE id = ?", id)
	return err
}
