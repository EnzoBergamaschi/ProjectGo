package auth

import (
	"context"
	"net/http"
	"strings"
)

type contextKey string

const (
	userIDKey   contextKey = "usuario_id"
	userTipoKey contextKey = "usuario_tipo"
)

// MiddlewareAutenticacao valida o token JWT e injeta o usuário no contexto
func MiddlewareAutenticacao(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Token não fornecido", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := ValidarJWT(tokenStr)
		if err != nil {
			http.Error(w, "Token inválido ou expirado", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userIDKey, claims.ID)
		ctx = context.WithValue(ctx, userTipoKey, claims.Tipo)
		next(w, r.WithContext(ctx))
	}
}

func RequireAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tipoVal := r.Context().Value(userTipoKey)
		tipo, ok := tipoVal.(string)
		if !ok || tipo != "admin" {
			http.Error(w, "Acesso negado: admin required", http.StatusForbidden)
			return
		}
		next(w, r)
	}
}

func AdminOnly(next http.HandlerFunc) http.HandlerFunc {
	return MiddlewareAutenticacao(RequireAdmin(next))
}
