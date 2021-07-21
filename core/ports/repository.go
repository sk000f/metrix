package ports

import (
	"time"

	"github.com/sk000f/metrix/core/domain"
)

type Repository interface {
	GetAll() ([]domain.Deployment, error)
	GetByProject(proj string) ([]domain.Deployment, error)
	GetByDateRange(start time.Time, end time.Time) ([]domain.Deployment, error)
	GetByProjectAndDateRange(proj int, start time.Time, end time.Time) ([]domain.Deployment, error)
	GetByProjectAndInterval(proj int, days int) ([]domain.Deployment, error)
	Update(d []domain.Deployment) error
}
