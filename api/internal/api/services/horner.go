package services

import (
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

type HornerService struct {
	Logger *log.Logger
	db     *sqlx.DB
}

func NewHornerService(db *sqlx.DB) HornerService {
	logger := log.New(os.Stdout, "UserSvc: ", log.LstdFlags|log.Lshortfile)
	return HornerService{logger, db}
}
