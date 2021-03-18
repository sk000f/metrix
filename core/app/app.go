package app

import (
	"fmt"
	"time"

	"github.com/sk000f/metrix/core/domain"
	"github.com/sk000f/metrix/core/ports"
	"github.com/sk000f/metrix/internal/cicd"
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

	days := end.Sub(start).Hours() / 24

	var pDep []domain.Deployment
	for _, d := range dep {
		if d.EnvironmentName == cicd.Production {
			pDep = append(pDep, d)
		}
	}

	// deployment frequency (deploys per day)
	// is number of deployments divided by number of days
	df := float64(len(pDep)) / days

	return num.To2dp(df), nil
}

func (a *app) LeadTime(proj string, start time.Time, end time.Time) (time.Time, error) {
	return time.Now(), nil
}

// ChangeFailRate calculates the percentage of deployments to a production
// environment which are not successful, for the specified date range and project name
func (a *app) ChangeFailRate(proj string, start time.Time, end time.Time) (int, error) {
	dep, err := a.r.GetByProjectAndDateRange(proj, start, end)
	if err != nil {
		fmt.Print(err.Error())
		return 0, err
	}

	var pDep []domain.Deployment
	var f int = 0
	for _, d := range dep {
		if d.EnvironmentName == cicd.Production {
			pDep = append(pDep, d)
			if d.Status == cicd.Failure {
				f++
			}
		}
	}

	// change fail rate is percentage of failed deployments
	cfr := int(float64(f) / float64(len(pDep)) * 100)

	return cfr, nil
}

func (a *app) MTTR(proj string, start time.Time, end time.Time) (time.Time, error) {
	return time.Now(), nil
}

func (a *app) UpdateDeployments() error { return nil }

func (a *app) UpdateDeploymentsForDateRange(start time.Time, end time.Time) error { return nil }
