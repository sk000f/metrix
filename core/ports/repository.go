package ports

import (
	"time"

	"github.com/sk000f/hexarch/core/domain"
)

type MetricsRepository interface {
	GetAll() ([]domain.Deployment, error)
	GetByProject(proj string) ([]domain.Deployment, error)
	GetByDateRange(start *time.Time, end *time.Time) ([]domain.Deployment, error)
	GetByProjectAndDateRange(proj string, start *time.Time, end *time.Time) ([]domain.Deployment, error)
	Update([]domain.Deployment) error
}
