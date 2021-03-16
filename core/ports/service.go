package ports

import "time"

type MetricsService interface {
	CalculateDeploymentFrequencyForDateRangeAndProject(start time.Time, end time.Time, proj string) (float64, error)
	CalculateLeadTimeForDateRangeAndProject(start time.Time, end time.Time, proj string) (time.Time, error)
	CalculateChangeFailRateForDateRangeAndProject(start time.Time, end time.Time, proj string) (float64, error)
	CalculateMTTRForDateRangeAndProject(start time.Time, end time.Time, proj string) (time.Time, error)
	UpdateDeployments() error
	UpdateDeploymentsForDateRange(start time.Time, end time.Time) error
}
