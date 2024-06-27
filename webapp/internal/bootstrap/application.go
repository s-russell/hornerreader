package bootstrap

import (
	"github.com/jmoiron/sqlx"
	"jerubaal.com/horner/internal/resources/reading"
	"jerubaal.com/horner/internal/web"
	"log"
	"net/http"
)

type Application struct {
	db     *sqlx.DB
	mux    *http.ServeMux
	logger *log.Logger
}

func NewApplication(db *sqlx.DB, mux *http.ServeMux, logger *log.Logger) *Application {
	return &Application{db, mux, logger}
}

func (app *Application) Start() {
	app.mux.Handle("/reading/*", http.StripPrefix("/reading", reading.Serve(app.db)))

	web.AddRoutes(app.mux)

	app.logger.Println("Serving Horner Reader on http://localhost:8888")
	http.ListenAndServe(":8888", app.mux)
}
