package reading

import (
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"
)

func Serve(db *sqlx.DB) *http.ServeMux {
	readerMux := http.NewServeMux()

	hornerSvc := NewHornerService(db)

	readerMux.HandleFunc("GET /{readingNumber}", func(writer http.ResponseWriter, request *http.Request) {
		readingNumber, err := strconv.Atoi(request.PathValue("readingNumber"))
		if err != nil {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		writer.Header().Set("Content-Type", "text/html")
		hornerSvc.RenderReadingHTML(writer, readingNumber)
	})

	return readerMux
}
