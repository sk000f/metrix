package app

import (
	"time"

	"github.com/rs/zerolog/log"
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
		log.Error().Stack().Err(err).
			Str("project", proj).
			Time("start", start).
			Time("end", end).
			Msg("app.DeploymentFrequency")
		return 0.0, err
	}

	days := end.Sub(start).Hours() / 24

	var n float64
	for _, d := range dep {
		if d.EnvironmentName == cicd.Production {
			n++
		}
	}

	// deployment frequency (deploys per day)
	// is number of deployments divided by number of days
	df := n / days

	return num.To2dp(df), nil
}

// LeadTime calculates the number of minutes for a deployment to complete from
// code commit to production deployment, for the specified date range and project name
func (a *app) LeadTime(proj string, start time.Time, end time.Time) (int, error) {

	dep, err := a.r.GetByProjectAndDateRange(proj, start, end)
	if err != nil {
		log.Error().Stack().Err(err).
			Str("project", proj).
			Time("start", start).
			Time("end", end).
			Msg("app.LeadTime")
		return 0, err
	}

	var n, t int
	for _, d := range dep {
		if d.EnvironmentName == cicd.Production {
			n++
			t += d.Duration
		}
	}

	// lead time is average number of minutes per deployment
	// which is total deployment duration divided by number of deployments
	lt := t / n

	return lt, nil
}

// ChangeFailRate calculates the percentage of deployments to a production
// environment which are not successful, for the specified date range and project name
func (a *app) ChangeFailRate(proj string, start time.Time, end time.Time) (int, error) {

	dep, err := a.r.GetByProjectAndDateRange(proj, start, end)
	if err != nil {
		log.Error().Stack().Err(err).
			Str("project", proj).
			Time("start", start).
			Time("end", end).
			Msg("app.ChangeFailRate")
		return 0, err
	}

	var n, f float64
	for _, d := range dep {
		if d.EnvironmentName == cicd.Production {
			n++
			if d.Status == cicd.Failure {
				f++
			}
		}
	}

	// change fail rate is number of failed deployments as a percentage
	cfr := int(f / n * 100)

	return cfr, nil
}

func (a *app) MTTR(proj string, start time.Time, end time.Time) (time.Time, error) {
	return time.Now(), nil
}

func (a *app) UpdateDeployments() error {

	d, err := a.ci.GetAllDeployments()
	if err != nil {
		log.Error().Stack().Err(err).
			Msg("app.UpdateDeployments")
		return err
	}

	err = a.r.Update(d)
	if err != nil {
		log.Error().Stack().Err(err).
			Msg("app.UpdateDeployments")
		return err
	}

	return nil
}

func (a *app) UpdateDeploymentsForDateRange(start time.Time, end time.Time) error { return nil }
