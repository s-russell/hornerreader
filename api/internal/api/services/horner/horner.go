package horner

import (
	"embed"
	"github.com/jmoiron/sqlx"
	"html/template"
	"jerubaal.com/horner/internal/reader"
	"log"
	"net/http"
	"os"
)

var (
	//go:embed templates
	templateFiles embed.FS
)

type HornerService struct {
	Logger    *log.Logger
	templates *template.Template
	db        *sqlx.DB
}

func NewHornerService(db *sqlx.DB) HornerService {
	logger := log.New(os.Stdout, "UserSvc: ", log.LstdFlags|log.Lshortfile)

	templates, err := template.ParseFS(templateFiles, "templates/*.gohtml")
	if err != nil {
		logger.Fatal("Failed to compile Horner service templates", err)
	}

	return HornerService{logger, templates, db}
}

func (h *HornerService) RenderReadingHTML(w http.ResponseWriter, reading *[]reader.HornerReading) {
	err := h.templates.ExecuteTemplate(w, "reading.gohtml", reading)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
