package domain

import (
	"time"

	"github.com/sk000f/metrix/internal/cicd"
)

type Deployment struct {
	ID               int                   `json:"id"`
	Status           cicd.DeploymentStatus `json:"status"`
	EnvironmentName  cicd.EnvironmentName  `json:"environmentName"`
	ProjectID        int                   `json:"projectId"`
	ProjectName      string                `json:"projectName"`
	ProjectPath      string                `json:"projectPath"`
	ProjectNamespace string                `json:"projectNamespace"`
	PipelineID       int                   `json:"pipelineId"`
	FinishedAt       time.Time             `json:"finishedAt"`
	Duration         int                   `json:"duration"`
}
