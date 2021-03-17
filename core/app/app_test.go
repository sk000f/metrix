package app_test

import (
	"math"
	"testing"
	"time"

	"github.com/sk000f/metrix/core/app"
	"github.com/sk000f/metrix/core/app/mocks"
	"github.com/sk000f/metrix/resources/data"
)

func TestDeploymentFrequency(t *testing.T) {
	t.Run("calculate deployment frequency over 30 day period", func(t *testing.T) {
		r := createMockRepository()
		ci := createMockCIServer()

		a := app.New(r, ci)

		got, err := a.DeploymentFrequency(
			"test-project",
			time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2021, 3, 31, 0, 0, 0, 0, time.UTC),
		)

		if err != nil {
			t.Fatal(err)
		}

		want := 0.23

		if math.Abs(got) != math.Abs(want) {
			t.Errorf("got %v; want %v", math.Abs(got), math.Abs(want))
		}

	})

	t.Run("calculate deployment frequency over 5 day period", func(t *testing.T) {
		r := createMockRepository()
		ci := createMockCIServer()

		a := app.New(r, ci)

		got, err := a.DeploymentFrequency(
			"test-project",
			time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2021, 3, 6, 0, 0, 0, 0, time.UTC),
		)

		if err != nil {
			t.Fatal(err)
		}

		want := 1.20

		if math.Abs(got) != math.Abs(want) {
			t.Errorf("got %v; want %v", math.Abs(got), math.Abs(want))
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
