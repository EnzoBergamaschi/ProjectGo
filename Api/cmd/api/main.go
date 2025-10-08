package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/EnzoBergamaschi/ProjectGo/internal/auth"
	"github.com/EnzoBergamaschi/ProjectGo/internal/config"
	"github.com/EnzoBergamaschi/ProjectGo/internal/database"
	"github.com/EnzoBergamaschi/ProjectGo/internal/http/router"
)

func main() {
	cfg := config.Load()
	if cfg.JWTSecret == "" {
		log.Fatal("JWT_SECRET não definido")
	}
	auth.Configure(cfg.JWTSecret, cfg.JWTExpirationHours)

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco: %v", err)
	}
	defer db.Close()

	r := router.New(db)

	addr := fmt.Sprintf(":%s", cfg.AppPort)
	log.Printf("Servidor iniciado na porta %s...\n", cfg.AppPort)

	server := &http.Server{
		Addr: addr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Cabeçalhos de CORS sempre aplicados
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Set("Access-Control-Allow-Credentials", "true")

			// Intercepta absolutamente todos os OPTIONS
			if req.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}

			// Passa para o roteador
			r.ServeHTTP(w, req)
		}),
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
