package ports

import "time"

type MetricsService interface {
	DeploymentFrequency(start time.Time, end time.Time, proj string) (float64, error)
	LeadTime(start time.Time, end time.Time, proj string) (time.Time, error)
	ChangeFailRate(start time.Time, end time.Time, proj string) (float64, error)
	MTTR(start time.Time, end time.Time, proj string) (time.Time, error)
	UpdateDeployments() error
	UpdateDeploymentsForDateRange(start time.Time, end time.Time) error
}
