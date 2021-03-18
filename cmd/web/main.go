package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	configureLogging()
	log.Info().Msg("Metrix starting ...")
}

func configureLogging() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}
