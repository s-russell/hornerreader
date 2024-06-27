package main

import (
	"jerubaal.com/horner/internal/bootstrap"
	"log"
	"net/http"
)

func main() {

	logger := log.Default()

	db, err := bootstrap.GetDB()
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	app := bootstrap.NewApplication(db, mux, logger)

	app.Start()
}
