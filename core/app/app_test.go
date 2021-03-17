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
	t.Run("calculate deployment frequency", func(t *testing.T) {
		r := createMockRepository()
		ci := createMockCIServer()

		a := app.New(r, ci)

		got, err := a.DeploymentFrequency(
			"test-project",
			time.Date(2021, 3, 1, 12, 0, 0, 0, time.UTC),
			time.Date(2021, 3, 31, 12, 0, 0, 0, time.UTC),
		)

		if err != nil {
			t.Fatal(err)
		}

		want := 0.26

		if math.Abs(got) != math.Abs(want) {
			t.Errorf("got %v; want %v", math.Abs(got), math.Abs(want))
		}

	})
}

func TestHelpers(t *testing.T) {
	t.Run("test truncate float down to 2 decimal places", func(t *testing.T) {
		want := 0.25
		got := app.Trunc2dp(0.2546)

		if got != want {
			t.Errorf("got: %v; want %v", got, want)
		}
	})

	t.Run("test truncate float up to 2 decimal places", func(t *testing.T) {
		want := 0.26
		got := app.Trunc2dp(0.2645)

		if got != want {
			t.Errorf("got: %v; want %v", got, want)
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
