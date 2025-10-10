package auth

import "golang.org/x/crypto/bcrypt"

func GerarHashSenha(senha string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	return string(bytes), err
}
func ValidarSenha(hash, senha string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(senha))
	return err == nil
}
