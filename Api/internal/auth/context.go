package auth

import "net/http"

func GetUserID(r *http.Request) (int, bool) {
	v := r.Context().Value(userIDKey)
	id, ok := v.(int)
	return id, ok
}
func GetUserTipo(r *http.Request) (string, bool) {
	v := r.Context().Value(userTipoKey)
	s, ok := v.(string)
	return s, ok
}
