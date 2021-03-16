package app

import (
	"time"

	"github.com/sk000f/hexarch/core/ports"
)

type app struct {
	metricsRepository ports.MetricsRepository
	ciServer          ports.MetricsCIServer
}

func (a *app) CalculateDeploymentFrequencyForDateRangeAndProject(start *time.Time, end *time.Time, proj string) (float64, error) {
	return 0.0, nil
}

func (a *app) CalculateLeadTimeForDateRangeAndProject(start *time.Time, end *time.Time, proj string) (*time.Time, error) {
	return nil, nil
}

func (a *app) CalculateChangeFailRateForDateRangeAndProject(start *time.Time, end *time.Time, proj string) (float64, error) {
	return 0.0, nil
}

func (a *app) CalculateMTTRForDateRangeAndProject(start *time.Time, end *time.Time, proj string) (*time.Time, error) {
	return nil, nil
}

func (a *app) UpdateDeployments() error { return nil }

func (a *app) UpdateDeploymentsForDateRange(start *time.Time, end *time.Time) error { return nil }
