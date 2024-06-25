package routes

import (
	"encoding/json"
	"jerubaal.com/horner/internal/api/services/horner"
	"jerubaal.com/horner/internal/reader"
	"net/http"
	"strconv"
)

func Reader(hornerSvc *horner.HornerService) *http.ServeMux {
	readerMux := http.NewServeMux()

	readerMux.HandleFunc("/reading/{readingNumber}", func(writer http.ResponseWriter, request *http.Request) {
		readingNumber, err := strconv.Atoi(request.PathValue("readingNumber"))
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
		}
		reading := reader.GetHornerReadingPlan(readingNumber)

		jsonReading, err := json.Marshal(reading)
		if err != nil {
			// Handle errors if JSON marshaling fails
			writer.WriteHeader(http.StatusInternalServerError)
			hornerSvc.Logger.Printf("Error marshalling horner reading:\n%v\n %v", jsonReading, err)
			writer.Write([]byte("Internal Server Error"))
			return
		}

		// Set Content-Type header and write the JSON response
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		writer.Write(jsonReading)
	})

	readerMux.HandleFunc("/htmx/reading", func(w http.ResponseWriter, request *http.Request) {
		queryParams := request.URL.Query()
		readingNumberParam := queryParams.Get("n")
		if len(readingNumberParam) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}
		readingNumber, err := strconv.Atoi(readingNumberParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		reading := reader.GetHornerReadingPlan(readingNumber)
		w.Header().Set("Content-Type", "text/html")
		hornerSvc.RenderReadingHTML(w, &reading)
	})

	return readerMux
}
