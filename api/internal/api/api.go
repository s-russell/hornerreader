package api

import (
	"github.com/jmoiron/sqlx"
	"jerubaal.com/horner/internal/api/routes"
	"jerubaal.com/horner/internal/api/services"
	"jerubaal.com/horner/internal/api/services/horner"
	"net/http"
)

type API struct {
	UserSvc   *services.UserSvc
	HornerSvc *horner.HornerService
}

func Build(db *sqlx.DB) API {
	userSvc := services.NewUserSvc(db)
	hornerSvc := horner.NewHornerService(db)
	return API{&userSvc, &hornerSvc}
}

func (api *API) AddRoutes(mux *http.ServeMux) {

	mux.Handle("/api/login/*", http.StripPrefix("/api/login", routes.Login(api.UserSvc)))
	mux.Handle("/api/user/*", http.StripPrefix("/api/user", routes.User(api.UserSvc)))
	mux.Handle("/api/reader/*", http.StripPrefix("/api/reader", routes.Reader(api.HornerSvc)))
}
