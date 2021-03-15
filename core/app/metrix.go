package metrix

import (
	"time"

	"github.com/sk000f/hexarch/core/ports"
)

type metrix struct {
	metricsRepository ports.MetricsRepository
	ciServer          ports.MetricsCIServer
}

func (m *metrix) CalculateDeploymentFrequencyForDateRangeAndProject(start *time.Time, end *time.Time, proj string) (float64, error) {
	return 0.0, nil
}

func (m *metrix) CalculateLeadTimeForDateRangeAndProject(start *time.Time, end *time.Time, proj string) (*time.Time, error) {
	return nil, nil
}

func (m *metrix) CalculateChangeFailRateForDateRangeAndProject(start *time.Time, end *time.Time, proj string) (float64, error) {
	return 0.0, nil
}

func (m *metrix) CalculateMTTRForDateRangeAndProject(start *time.Time, end *time.Time, proj string) (*time.Time, error) {
	return nil, nil
}

func (m *metrix) UpdateDeployments() error { return nil }

func (m *metrix) UpdateDeploymentsForDateRange(start *time.Time, end *time.Time) error { return nil }
