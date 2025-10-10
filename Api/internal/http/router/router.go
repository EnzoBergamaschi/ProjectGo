package router

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/EnzoBergamaschi/ProjectGo/internal/auth"
	"github.com/EnzoBergamaschi/ProjectGo/internal/http/handlers"
)

func New(db *sql.DB) http.Handler {
	mux := http.NewServeMux()

	corsWrapper := func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Set("Access-Control-Allow-Credentials", "true")

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}

			next(w, r)
		}
	}
	mux.HandleFunc("/", corsWrapper(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "API do ProjectGo está rodando!")
	}))
	mux.HandleFunc("/health", corsWrapper(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	}))
	usuarioHandler := handlers.NovoUsuarioHandler(db)
	mux.HandleFunc("/usuarios", corsWrapper(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			auth.MiddlewareAutenticacao(auth.RequireAdmin(usuarioHandler.Listar))(w, r)
		case http.MethodPost:
			usuarioHandler.Criar(w, r)
		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	}))
	mux.HandleFunc("/usuarios/", corsWrapper(auth.MiddlewareAutenticacao(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			auth.RequireAdmin(usuarioHandler.Atualizar)(w, r)
		case http.MethodDelete:
			auth.RequireAdmin(usuarioHandler.Deletar)(w, r)
		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})))
	produtoHandler := handlers.NovoProdutoHandler(db)
	mux.HandleFunc("/produtos", corsWrapper(auth.MiddlewareAutenticacao(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			produtoHandler.Listar(w, r)
		case http.MethodPost:
			auth.RequireAdmin(produtoHandler.Criar)(w, r)
		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})))
	mux.HandleFunc("/produtos/", corsWrapper(auth.MiddlewareAutenticacao(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			auth.RequireAdmin(produtoHandler.Atualizar)(w, r)
		case http.MethodDelete:
			auth.RequireAdmin(produtoHandler.Deletar)(w, r)
		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})))
	vendaHandler := handlers.NovaVendaHandler(db)
	mux.HandleFunc("/vendas", corsWrapper(auth.MiddlewareAutenticacao(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			vendaHandler.Listar(w, r)
		case http.MethodPost:
			vendaHandler.Criar(w, r)
		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})))
	mux.HandleFunc("/vendas/", corsWrapper(auth.MiddlewareAutenticacao(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			auth.RequireAdmin(vendaHandler.Atualizar)(w, r)
		case http.MethodDelete:
			auth.RequireAdmin(vendaHandler.Deletar)(w, r)
		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})))
	itemVendaHandler := handlers.NovoItemVendaHandler(db)

	mux.HandleFunc("/itens_venda", corsWrapper(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		auth.MiddlewareAutenticacao(func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodPost:
				itemVendaHandler.Criar(w, r)
			default:
				http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			}
		})(w, r)
	}))
	mux.HandleFunc("/itens_venda/", corsWrapper(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		auth.MiddlewareAutenticacao(func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodGet:
				itemVendaHandler.ListarPorVenda(w, r)
			case http.MethodDelete:
				itemVendaHandler.Deletar(w, r)
			default:
				http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			}
		})(w, r)
	}))
	vendaDetalhadaHandler := handlers.NovaVendaDetalhadaHandler(db)
	mux.HandleFunc("/vendas_detalhadas", corsWrapper(auth.MiddlewareAutenticacao(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			vendaDetalhadaHandler.Listar(w, r)
		} else {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})))
	authHandler := handlers.NovoAuthHandler(db)
	mux.HandleFunc("/login", corsWrapper(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			authHandler.Login(w, r)
			return
		}
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}))
	return mux
}
