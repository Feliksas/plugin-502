package plugin_5xx

import (
	"context"
	"log"
	"net/http"
)

const defaultError = "Bad Gateway"
const defaultCode = 502

type Config struct {
	ErrorString string
	ErrorCode   int
}

func CreateConfig() *Config {
	return &Config{
		ErrorString: defaultError,
		ErrorCode:   defaultCode,
	}
}

func New(ctx context.Context, next http.Handler, config *Config, _ string) (http.Handler, error) {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hit 5xx middleware")
		http.Error(rw, config.ErrorString, config.ErrorCode)
	}), nil
}
