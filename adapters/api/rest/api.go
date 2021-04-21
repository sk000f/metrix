package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

func DeploymentFrequency(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"value":"123"}`))
}

func LeadTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"value":"123"}`))
}

func ChangeFailRate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"value":"123"}`))
}

func MTTR(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"value":"123"}`))
}

func InitRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/deployment-frequency", DeploymentFrequency).Methods(http.MethodGet)
	r.HandleFunc("/lead-time", LeadTime).Methods(http.MethodGet)
	r.HandleFunc("/change-fail-rate", ChangeFailRate).Methods(http.MethodGet)
	r.HandleFunc("/mttr", MTTR).Methods(http.MethodGet)

	return r
}
