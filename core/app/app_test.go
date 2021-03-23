package app_test

import (
	"math"
	"reflect"
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

func TestLeadTime(t *testing.T) {
	t.Run("calculate lead time over 30 day period", func(t *testing.T) {
		r := createMockRepository()
		ci := createMockCIServer()

		a := app.New(r, ci)

		got, err := a.LeadTime(
			"test-project",
			time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2021, 3, 31, 0, 0, 0, 0, time.UTC),
		)

		if err != nil {
			t.Fatal(err)
		}

		want := 52

		if got != want {
			t.Errorf("got: %v; want %v", got, want)
		}
	})

	t.Run("calculate lead time over 5 day period", func(t *testing.T) {
		r := createMockRepository()
		ci := createMockCIServer()

		a := app.New(r, ci)

		got, err := a.LeadTime(
			"test-project",
			time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2021, 3, 6, 0, 0, 0, 0, time.UTC),
		)

		if err != nil {
			t.Fatal(err)
		}

		want := 48

		if got != want {
			t.Errorf("got: %v; want %v", got, want)
		}
	})
}

func TestChangeFailRate(t *testing.T) {
	t.Run("calculate change fail rate over 30 day period", func(t *testing.T) {
		r := createMockRepository()
		ci := createMockCIServer()

		a := app.New(r, ci)

		got, err := a.ChangeFailRate(
			"test-project",
			time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2021, 3, 31, 0, 0, 0, 0, time.UTC),
		)

		if err != nil {
			t.Fatal(err)
		}

		want := 14

		if got != want {
			t.Errorf("got: %v; want: %v", got, want)
		}
	})

	t.Run("calculate change fail rate over 5 day period", func(t *testing.T) {
		r := createMockRepository()
		ci := createMockCIServer()

		a := app.New(r, ci)

		got, err := a.ChangeFailRate(
			"test-project",
			time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2021, 3, 6, 0, 0, 0, 0, time.UTC),
		)

		if err != nil {
			t.Fatal(err)
		}

		want := 16

		if got != want {
			t.Errorf("got: %v; want: %v", got, want)
		}
	})
}

func TestUpdateDeployments(t *testing.T) {
	t.Run("update all deployments", func(t *testing.T) {

		r := createMockRepository()
		r.ClearData()
		ci := createMockCIServer()

		a := app.New(r, ci)

		// run update deployments
		err := a.UpdateDeployments()
		if err != nil {
			t.Fatal(err)
		}

		// check the repository contains the right deployments from the mock CI server
		dCI, _ := ci.GetAllDeployments()
		dR, _ := r.GetAll()

		if !reflect.DeepEqual(dCI, dR) {
			t.Errorf("deployments in repository do not match deployments in CI server")
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
	ci.LoadData(data.SampleDeployments)
	return ci
}
