package rest

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/sk000f/metrix/core/ports"
)

type RestAPI struct {
	srv    ports.Service
	Router *mux.Router
}

func New(s ports.Service) *RestAPI {
	api := &RestAPI{
		srv: s,
	}
	api.InitRouter()
	return api
}

func (api *RestAPI) DeploymentFrequency(w http.ResponseWriter, r *http.Request) {

	res, err := api.srv.DeploymentFrequency(448, 90)
	if err != nil {
		log.Error().Stack().Err(err).
			Msg("api.DeploymentFrequency")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"value":"error"}`))
		return
	}

	o := `{"value":` + fmt.Sprintf("%.2f", res) + `}`

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(o))
}

func (api *RestAPI) LeadTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"value":123}`))
}

func (api *RestAPI) ChangeFailRate(w http.ResponseWriter, r *http.Request) {

	// interval := r.URL.Query().Get("interval")
	// proj := r.URL.Query().Get("project")

	start := time.Date(2020, 7, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2021, 7, 31, 0, 0, 0, 0, time.UTC)

	res, err := api.srv.ChangeFailRate(520, start, end)
	if err != nil {
		log.Error().Stack().Err(err).
			Msg("api.ChangeFailRate")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"value":"error"}`))
		return
	}

	o := `{"value":` + strconv.Itoa(res) + `}`

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(o))
}

func (api *RestAPI) MTTR(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"value":123}`))
}

func (api *RestAPI) Default(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"":""}`))
}

func (api *RestAPI) InitRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/deployment-frequency", api.DeploymentFrequency).Methods(http.MethodGet)
	r.HandleFunc("/lead-time", api.LeadTime).Methods(http.MethodGet)
	r.HandleFunc("/change-fail-rate", api.ChangeFailRate).Methods(http.MethodGet)
	r.HandleFunc("/mttr", api.MTTR).Methods(http.MethodGet)
	r.HandleFunc("/", api.Default).Methods(http.MethodGet)
	api.Router = r
}
