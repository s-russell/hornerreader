package reading

import (
	"embed"
	"github.com/jmoiron/sqlx"
	"html/template"
	"jerubaal.com/horner/internal/resources/reading/horner"
	"log"
	"net/http"
	"os"
	"strconv"
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

type ReaderTemplateData struct {
	PreviousReading    string
	HasPreviousReading bool
	NextReading        string
	Readings           []horner.HornerReading
	ReadingNumber      int
}

func getTemplateData(readingNumber int) ReaderTemplateData {
	reading := horner.GetNumber(readingNumber)
	if readingNumber == 1 {
		return ReaderTemplateData{
			"",
			false,
			"/reading/2",
			reading,
			1,
		}
	} else {
		return ReaderTemplateData{
			"/reading/" + strconv.Itoa(readingNumber-1),
			true,
			"/reading/" + strconv.Itoa(readingNumber+1),
			reading,
			readingNumber,
		}
	}
}

func (h *HornerService) RenderReadingHTML(w http.ResponseWriter, readingNumber int) {

	err := h.templates.ExecuteTemplate(
		w,
		"reading.gohtml",
		getTemplateData(readingNumber),
	)
	if err != nil {
		h.Logger.Printf("Failed to render reading.gohtml template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
