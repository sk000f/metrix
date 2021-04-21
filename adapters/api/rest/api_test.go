package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetDeploymentFrequency(t *testing.T) {
	r := InitRouter()

	res := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/deployment-frequency", nil)
	if err != nil {
		t.Errorf("failed to create request: %v", err)
	}

	r.ServeHTTP(res, req)

	want := `{"value":"123"}`
	got := res.Body.String()

	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}

func TestGetLeadTime(t *testing.T) {
	r := InitRouter()

	res := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/lead-time", nil)
	if err != nil {
		t.Errorf("failed to create request: %v", err)
	}

	r.ServeHTTP(res, req)

	want := `{"value":"123"}`
	got := res.Body.String()

	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}

func TestGetChangeFailRate(t *testing.T) {
	r := InitRouter()

	res := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/change-fail-rate", nil)
	if err != nil {
		t.Errorf("failed to create request: %v", err)
	}

	r.ServeHTTP(res, req)

	want := `{"value":"123"}`
	got := res.Body.String()

	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}

func TestGetMTTR(t *testing.T) {
	r := InitRouter()

	res := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/mttr", nil)
	if err != nil {
		t.Errorf("failed to create request: %v", err)
	}

	r.ServeHTTP(res, req)

	want := `{"value":"123"}`
	got := res.Body.String()

	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}
