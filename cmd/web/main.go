package main

import (
	"net/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sk000f/metrix/adapters/api/rest"
	"github.com/sk000f/metrix/adapters/ciserver/gitlab"
	"github.com/sk000f/metrix/adapters/database/mongodb"
	"github.com/sk000f/metrix/core/app"
	"github.com/sk000f/metrix/core/ports"
)

func main() {
	configureLogging()
	log.Info().Msg("Metrix starting ...")

	ci := setupCIServer()

	db := setupDatabase("")

	srv := setupAppCore(db, ci)

	api := setupAPI(srv)

	log.Fatal().Err(http.ListenAndServe(":8080", api.Router))
}

func setupCIServer() ports.CIServer {
	return gitlab.New("", "", nil)
}

func setupDatabase(conn string) ports.Repository {
	db, err := mongodb.New(conn)
	if err != nil {
		log.Fatal().Err(err)
	}

	return db
}

func setupAppCore(db ports.Repository, ci ports.CIServer) ports.Service {
	return app.New(db, ci)
}

func setupAPI(srv ports.Service) *rest.RestAPI {
	return rest.New(srv)
}

func configureLogging() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}
