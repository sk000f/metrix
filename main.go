package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sk000f/metrix/adapters/api/rest"
	"github.com/sk000f/metrix/adapters/ciserver/gitlab"
	"github.com/sk000f/metrix/adapters/database/mongodb"
	"github.com/sk000f/metrix/core/app"
	"github.com/sk000f/metrix/core/ports"

	"github.com/joho/godotenv"
)

func main() {

	flags := parseFlags()

	configureLogging()

	cfg := setupConfig()
	log.Info().Msg("Metrix starting ...")

	ci := setupCIServer(cfg)

	db := setupDatabase(cfg)

	srv := setupAppCore(db, ci)

	api := setupAPI(srv)

	if flags["update"] == "y" {
		err := srv.UpdateDeployments()
		if err != nil {
			log.Error().Stack().Err(err).
				Msg("main")
		}
	}

	log.Fatal().Err(http.ListenAndServe(":8080", api.Router))
}

func setupCIServer(cfg *Config) ports.CIServer {
	g, err := gitlab.New(cfg.GitLabToken, cfg.GitLabURL)
	if err != nil {
		log.Fatal().Err(err)
	}
	return g
}

func setupDatabase(cfg *Config) ports.Repository {
	db, err := mongodb.New(cfg.MongoConn)
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

func parseFlags() map[string]string {
	var update string

	flag.StringVar(&update, "update", "n", "update CI data")

	flag.Parse()

	flags := map[string]string{
		"update": update,
	}

	return flags
}

func setupConfig() *Config {
	cfg := new(Config)

	err := godotenv.Load()
	if err != nil {
		log.Error().Stack().Err(err).
			Msg("main.setupConfig")
	}

	cfg.GitLabURL = os.Getenv("METRIX_GITLAB_URL")
	cfg.GitLabToken = os.Getenv("METRIX_GITLAB_TOKEN")
	cfg.MongoConn = os.Getenv("METRIX_MONGO_CONN")

	return cfg
}

type Config struct {
	GitLabURL   string
	GitLabToken string
	MongoConn   string
}
