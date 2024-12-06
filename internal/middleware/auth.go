package middleware

import (
	"log/slog"
	"net/http"

	"ml-orchestrator-module/internal/config"
)

func AuthMiddleware(next http.Handler, cfg *config.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientID := r.Header.Get("Client-ID")
		clientSecret := r.Header.Get("Client-Secret")
		slog.Info("The rec client ID is", string(clientID))
		slog.Info("The save client ID is", string(cfg.ClientID))
		if clientID != cfg.ClientID || clientSecret != cfg.ClientSecret {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
