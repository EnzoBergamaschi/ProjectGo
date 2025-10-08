package auth

import "golang.org/x/crypto/bcrypt"

// Gera o hash da senha
func GerarHashSenha(senha string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	return string(bytes), err
}

// Compara senha com hash
func ValidarSenha(hash, senha string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(senha))
	return err == nil
}
