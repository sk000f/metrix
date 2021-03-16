package app_test

import (
	"testing"
	"time"

	"github.com/sk000f/metrix/core/app"
	"github.com/sk000f/metrix/core/app/mocks"
	"github.com/sk000f/metrix/resources/data"
)

func TestDeploymentFrequency(t *testing.T) {
	t.Run("calculate deployment frequency", func(t *testing.T) {
		r := createMockRepository()
		ci := createMockCIServer()

		a := app.New(r, ci)

		got, err := a.CalculateDeploymentFrequencyForDateRangeAndProject(
			time.Date(2021, 3, 1, 12, 0, 0, 0, time.UTC),
			time.Date(2021, 3, 1, 12, 0, 0, 0, time.UTC),
			"test-project",
		)

		if err != nil {
			t.Fatal(err)
		}

		want := 0.0

		if got != want {
			t.Errorf("got %v; want %v", got, want)
		}

	})
}

func createMockRepository() *mocks.RepositoryMock {
	r := &mocks.RepositoryMock{}
	r.LoadData(data.SampleDeployments)
	return r
}

func createMockCIServer() *mocks.CIServerMock {
	ci := &mocks.CIServerMock{}
	return ci
}
