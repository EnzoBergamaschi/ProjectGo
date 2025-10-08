package auth

import "net/http"

// Retorna o ID do usuário autenticado do contexto
func GetUserID(r *http.Request) (int, bool) {
	v := r.Context().Value(userIDKey)
	id, ok := v.(int)
	return id, ok
}

// Retorna o tipo do usuário autenticado ("admin" ou "cliente")
func GetUserTipo(r *http.Request) (string, bool) {
	v := r.Context().Value(userTipoKey)
	s, ok := v.(string)
	return s, ok
}
