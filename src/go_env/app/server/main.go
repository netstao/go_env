package main

import (
	_ "expvar"
	"net/http"

	"go_env/strcat"
)

func main() {
	config := LoadConfig()
	handler := strcat.New()
	http.Handle("/strcat", handler)
	http.ListenAndServe(config.Listen, nil)
}
