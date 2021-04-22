package main

import (
	"net/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sk000f/metrix/adapters/api/rest"
	"github.com/sk000f/metrix/adapters/ciserver/gitlab"
	"github.com/sk000f/metrix/adapters/database/mongodb"
	"github.com/sk000f/metrix/core/app"
)

func main() {
	configureLogging()
	log.Info().Msg("Metrix starting ...")

	ci := gitlab.New("", "", nil)
	db := mongodb.New()

	srv := app.New(db, ci)

	rest := rest.New(srv)

	log.Fatal().Err(http.ListenAndServe(":8080", rest.Router))
}

func configureLogging() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}
