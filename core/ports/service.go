package ports

import "time"

type Service interface {
	DeploymentFrequency(proj string, start time.Time, end time.Time) (float64, error)
	LeadTime(proj string, start time.Time, end time.Time) (int, error)
	ChangeFailRate(proj string, start time.Time, end time.Time) (int, error)
	MTTR(proj string, start time.Time, end time.Time) (time.Time, error)
	UpdateDeployments() error
	UpdateDeploymentsForDateRange(start time.Time, end time.Time) error
}
