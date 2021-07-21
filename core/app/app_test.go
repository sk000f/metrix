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

func TestUnitDeploymentFrequency(t *testing.T) {
	t.Run("calculate deployment frequency over 30 day period", func(t *testing.T) {
		r := createMockRepository()
		ci := createMockCIServer()

		a := app.New(r, ci)

		got, err := a.DeploymentFrequency(1, 90)

		if err != nil {
			t.Fatal(err)
		}

		want := 0.07

		if math.Abs(got) != math.Abs(want) {
			t.Errorf("got %v; want %v", math.Abs(got), math.Abs(want))
		}

	})

	t.Run("calculate deployment frequency over 5 day period", func(t *testing.T) {
		r := createMockRepository()
		ci := createMockCIServer()

		a := app.New(r, ci)

		got, err := a.DeploymentFrequency(1, 90)

		if err != nil {
			t.Fatal(err)
		}

		want := 0.07

		if math.Abs(got) != math.Abs(want) {
			t.Errorf("got %v; want %v", math.Abs(got), math.Abs(want))
		}

	})
}

func TestUnitLeadTime(t *testing.T) {
	t.Run("calculate lead time over 30 day period", func(t *testing.T) {
		r := createMockRepository()
		ci := createMockCIServer()

		a := app.New(r, ci)

		got, err := a.LeadTime(1,
			time.Now().AddDate(0, 0, -30),
			time.Now(),
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
			1,
			time.Now().AddDate(0, 0, -5),
			time.Now(),
		)

		if err != nil {
			t.Fatal(err)
		}

		want := 52

		if got != want {
			t.Errorf("got: %v; want %v", got, want)
		}
	})
}

func TestUnitChangeFailRate(t *testing.T) {
	t.Run("calculate change fail rate over 30 day period", func(t *testing.T) {
		r := createMockRepository()
		ci := createMockCIServer()

		a := app.New(r, ci)

		got, err := a.ChangeFailRate(
			1,
			time.Now().AddDate(0, 0, -30),
			time.Now(),
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
			1,
			time.Now().AddDate(0, 0, -5),
			time.Now(),
		)

		if err != nil {
			t.Fatal(err)
		}

		want := 14

		if got != want {
			t.Errorf("got: %v; want: %v", got, want)
		}
	})
}

func TestUnitUpdateDeployments(t *testing.T) {
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
