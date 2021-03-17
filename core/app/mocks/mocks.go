package mocks

import (
	"time"

	"github.com/sk000f/metrix/core/domain"
	"github.com/sk000f/metrix/internal/ci"
)

type RepositoryMock struct {
	data []domain.Deployment
}

func (mR *RepositoryMock) GetAll() ([]domain.Deployment, error) {
	return nil, nil
}

func (mR *RepositoryMock) GetByProject(proj string) ([]domain.Deployment, error) {
	return nil, nil
}

func (mR *RepositoryMock) GetByDateRange(start time.Time, end time.Time) ([]domain.Deployment, error) {
	return nil, nil
}

func (mR *RepositoryMock) GetByProjectAndDateRange(proj string, start time.Time, end time.Time) ([]domain.Deployment, error) {

	// logic to filter data based on params
	var dep []domain.Deployment

	for _, d := range mR.data {
		if d.ProjectName == proj && d.FinishedAt.After(start) && d.FinishedAt.Before(end) && d.EnvironmentName == ci.Production {
			dep = append(dep, d)
		}
	}

	return dep, nil
}

func (mR *RepositoryMock) Update([]domain.Deployment) error {
	return nil
}

func (mR *RepositoryMock) LoadData(d []domain.Deployment) {
	mR.data = d
}

type CIServerMock struct{}

func (ci *CIServerMock) GetAll() ([]domain.Deployment, error) {
	return nil, nil
}

func (ci *CIServerMock) GetByProject(proj string) ([]domain.Deployment, error) {
	return nil, nil
}

func (ci *CIServerMock) GetByDateRange(start time.Time, end time.Time) ([]domain.Deployment, error) {
	return nil, nil
}

func (ci *CIServerMock) GetByProjectAndDateRange(proj string, start time.Time, end time.Time) ([]domain.Deployment, error) {
	return nil, nil
}

func (ci *CIServerMock) Update([]domain.Deployment) error {
	return nil
}
