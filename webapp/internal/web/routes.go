package web

import (
	"embed"
	"io/fs"
	"net/http"
)

type FrontendConfig struct {
	Title string `json:"title"`
}

var (
	//go:embed static
	staticFS embed.FS
	config   FrontendConfig
)

func AddRoutes(mux *http.ServeMux) {
	web, _ := fs.Sub(staticFS, "static")
	webServer := http.FileServerFS(web)
	mux.HandleFunc("/", webServer.ServeHTTP)
}
