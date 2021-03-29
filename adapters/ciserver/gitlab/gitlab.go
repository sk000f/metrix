package gitlab

import (
	"time"

	"github.com/sk000f/metrix/core/domain"
)

type GitLab struct{}

func (g *GitLab) GetAllDeployments() ([]domain.Deployment, error) {
	return nil, nil
}

func (g *GitLab) GetDeploymentsByProject(proj string) ([]domain.Deployment, error) {
	return nil, nil
}

func (g *GitLab) GetDeploymentsByDateRange(start time.Time, end time.Time) ([]domain.Deployment, error) {
	return nil, nil
}

func (g *GitLab) GetDeploymentsByProjectAndDateRange(proj string, start time.Time, end time.Time) ([]domain.Deployment, error) {
	return nil, nil
}
