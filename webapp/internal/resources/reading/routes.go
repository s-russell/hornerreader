package reading

import (
	"fmt"
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
		writer.Header().Set("HX-Trigger-After-Swap", fmt.Sprintf("{ \"newReading\": %d }", readingNumber))
		hornerSvc.RenderReadingHTML(writer, readingNumber)
	})

	return readerMux
}
