package ports

import "time"

type Service interface {
	DeploymentFrequency(proj int, days int) (float64, error)
	LeadTime(proj int, start time.Time, end time.Time) (int, error)
	ChangeFailRate(proj int, start time.Time, end time.Time) (int, error)
	MTTR(proj int, start time.Time, end time.Time) (time.Time, error)
	UpdateDeployments() error
	UpdateDeploymentsForDateRange(start time.Time, end time.Time) error
}
