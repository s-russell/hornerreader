package main

import (
	"jerubaal.com/horner/internal/api"
	"jerubaal.com/horner/internal/web"
	"log"
	"net/http"
)

func main() {

	logger := log.Default()

	db, err := api.GetDB()
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	apiSvc := api.Build(db)

	mux := http.NewServeMux()
	apiSvc.AddRoutes(mux)
	web.AddRoutes(mux)

	logger.Println("Serving api on http://localhost:8888")
	http.ListenAndServe(":8888", mux)
}
