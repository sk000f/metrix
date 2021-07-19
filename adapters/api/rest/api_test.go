package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sk000f/metrix/core/app/mocks"
)

func TestGetDeploymentFrequency(t *testing.T) {
	s := createMockService()
	api := New(s)
	api.InitRouter()

	res := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/deployment-frequency", nil)
	if err != nil {
		t.Errorf("failed to create request: %v", err)
	}

	api.Router.ServeHTTP(res, req)

	want := `{"value":"123"}`
	got := res.Body.String()

	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}

func TestGetLeadTime(t *testing.T) {
	s := createMockService()
	api := New(s)
	api.InitRouter()

	res := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/lead-time", nil)
	if err != nil {
		t.Errorf("failed to create request: %v", err)
	}

	api.Router.ServeHTTP(res, req)

	want := `{"value":"123"}`
	got := res.Body.String()

	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}

func TestGetChangeFailRate(t *testing.T) {
	s := createMockService()
	api := New(s)
	api.InitRouter()

	res := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/change-fail-rate", nil)
	if err != nil {
		t.Errorf("failed to create request: %v", err)
	}

	api.Router.ServeHTTP(res, req)

	want := `{"value":0}`
	got := res.Body.String()

	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}

func TestGetMTTR(t *testing.T) {
	s := createMockService()
	api := New(s)
	api.InitRouter()

	res := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/mttr", nil)
	if err != nil {
		t.Errorf("failed to create request: %v", err)
	}

	api.Router.ServeHTTP(res, req)

	want := `{"value":"123"}`
	got := res.Body.String()

	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}

func createMockService() *mocks.ServiceMock {
	s := &mocks.ServiceMock{}
	return s
}
