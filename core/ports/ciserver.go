package ports

import (
	"time"

	"github.com/sk000f/metrix/core/domain"
)

type CIServer interface {
	GetAllDeployments() ([]domain.Deployment, error)
	GetDeploymentsByProject(proj string) ([]domain.Deployment, error)
	GetDeploymentsByDateRange(start time.Time, end time.Time) ([]domain.Deployment, error)
	GetDeploymentsByProjectAndDateRange(proj string, start time.Time, end time.Time) ([]domain.Deployment, error)
}
