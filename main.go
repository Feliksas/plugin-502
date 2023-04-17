package plugin_502

import (
	"context"
	"log"
	"net/http"
)

func New(ctx context.Context, next http.Handler, _ *Config, _ string) (http.Handler, error) {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hit 502 middleware")
		http.Error(rw, "Bad gateway", http.StatusBadGateway)
	}), nil
}
