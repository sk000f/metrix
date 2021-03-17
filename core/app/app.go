package app

import (
	"fmt"
	"time"

	"github.com/sk000f/metrix/core/domain"
	"github.com/sk000f/metrix/core/ports"
	"github.com/sk000f/metrix/internal/ci"
	"github.com/sk000f/metrix/internal/num"
)

type app struct {
	r  ports.Repository
	ci ports.CIServer
}

func New(r ports.Repository, ci ports.CIServer) *app {
	return &app{
		r:  r,
		ci: ci,
	}
}

// DeploymentFrequency calculates how many times per day a deployment to a production
// environment occurs, for the specified date range and project name
func (a *app) DeploymentFrequency(proj string, start time.Time, end time.Time) (float64, error) {

	dep, err := a.r.GetByProjectAndDateRange(proj, start, end)

	if err != nil {
		fmt.Print(err.Error())
		return 0.0, err
	}

	// calculate whole days within specified date range
	days := end.Sub(start).Hours() / 24

	// we only want to count deployments to the production environment
	var pDep []domain.Deployment
	for _, d := range dep {
		if d.EnvironmentName == ci.Production {
			pDep = append(pDep, d)
		}
	}

	// count number of deployments and divide by number of days
	df := float64(len(pDep)) / days

	return num.To2dp(df), nil
}

func (a *app) LeadTime(start time.Time, end time.Time, proj string) (time.Time, error) {
	return time.Now(), nil
}

func (a *app) ChangeFailRate(start time.Time, end time.Time, proj string) (float64, error) {
	return 0.0, nil
}

func (a *app) MTTR(start time.Time, end time.Time, proj string) (time.Time, error) {
	return time.Now(), nil
}

func (a *app) UpdateDeployments() error { return nil }

func (a *app) UpdateDeploymentsForDateRange(start time.Time, end time.Time) error { return nil }
