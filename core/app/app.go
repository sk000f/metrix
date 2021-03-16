package app

import (
	"time"

	"github.com/sk000f/metrix/core/ports"
)

type app struct {
	repository ports.Repository
	ciServer   ports.CIServer
}

func New(r ports.Repository, ci ports.CIServer) *app {
	return &app{
		repository: r,
		ciServer:   ci,
	}
}

func (a *app) CalculateDeploymentFrequencyForDateRangeAndProject(start time.Time, end time.Time, proj string) (float64, error) {
	return 0.0, nil
}

func (a *app) CalculateLeadTimeForDateRangeAndProject(start time.Time, end time.Time, proj string) (time.Time, error) {
	return time.Now(), nil
}

func (a *app) CalculateChangeFailRateForDateRangeAndProject(start time.Time, end time.Time, proj string) (float64, error) {
	return 0.0, nil
}

func (a *app) CalculateMTTRForDateRangeAndProject(start time.Time, end time.Time, proj string) (time.Time, error) {
	return time.Now(), nil
}

func (a *app) UpdateDeployments() error { return nil }

func (a *app) UpdateDeploymentsForDateRange(start time.Time, end time.Time) error { return nil }
