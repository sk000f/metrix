package mongodb

import (
	"time"

	"github.com/sk000f/metrix/core/domain"
)

// type Config struct {
// 	Username     string
// 	Password     string
// 	DatabaseName string
// 	URI          string
// }

type MongoDB struct {
}

func New() *MongoDB {
	return &MongoDB{}
}

func (m *MongoDB) GetAll() ([]domain.Deployment, error) {
	return nil, nil
}

func (m *MongoDB) GetByProject(proj string) ([]domain.Deployment, error) {
	return nil, nil
}

func (m *MongoDB) GetByDateRange(start time.Time, end time.Time) ([]domain.Deployment, error) {
	return nil, nil
}

func (m *MongoDB) GetByProjectAndDateRange(proj string, start time.Time, end time.Time) ([]domain.Deployment, error) {
	return nil, nil
}

func (m *MongoDB) Update(d []domain.Deployment) error {
	return nil
}
