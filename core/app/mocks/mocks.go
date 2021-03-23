package mocks

import (
	"time"

	"github.com/sk000f/metrix/core/domain"
)

type RepositoryMock struct {
	data []domain.Deployment
}

func (mR *RepositoryMock) GetAll() ([]domain.Deployment, error) {
	return mR.data, nil
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
		if d.ProjectName == proj && d.FinishedAt.After(start) && d.FinishedAt.Before(end) {
			dep = append(dep, d)
		}
	}

	return dep, nil
}

func (mR *RepositoryMock) Update(d []domain.Deployment) error {
	mR.data = d
	return nil
}

func (mR *RepositoryMock) LoadData(d []domain.Deployment) {
	mR.data = d
}

func (mR *RepositoryMock) ClearData() {
	mR.data = nil
}

type CIServerMock struct {
	data []domain.Deployment
}

func (ci *CIServerMock) GetAllDeployments() ([]domain.Deployment, error) {
	return ci.data, nil
}

func (ci *CIServerMock) GetDeploymentsByProject(proj string) ([]domain.Deployment, error) {
	return nil, nil
}

func (ci *CIServerMock) GetDeploymentsByDateRange(start time.Time, end time.Time) ([]domain.Deployment, error) {
	return nil, nil
}

func (ci *CIServerMock) GetDeploymentsByProjectAndDateRange(proj string, start time.Time, end time.Time) ([]domain.Deployment, error) {
	return nil, nil
}

func (ci *CIServerMock) Update(d []domain.Deployment) error {
	return nil
}

func (ci *CIServerMock) LoadData(d []domain.Deployment) {
	ci.data = d
}
