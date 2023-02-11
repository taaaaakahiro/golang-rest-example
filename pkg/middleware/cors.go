package middleware

import (
	"github.com/taaaaakahiro/golang-rest-example/pkg/config"
	"net/http"
)

func CORSHeaderMiddleware(cfg *config.Config) (fn func(http.Handler) http.Handler) { // 引数名を指定してるのでreturnのみでおｋ
	fn = func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			allowOrigin := cfg.Server.AllowCorsOrigin

			w.Header().Set("Access-Control-Allow-Origin", allowOrigin)
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Access-Control-Allow-Credentials", "true")

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}

			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
	return
}
