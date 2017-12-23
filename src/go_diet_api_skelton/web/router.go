package web

import (
	"net/http"

	"go_diet_api_skelton/handler"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Root)
	return mux
}
