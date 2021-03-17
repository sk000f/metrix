package domain

import (
	"time"

	"github.com/sk000f/metrix/internal/ci"
)

type Deployment struct {
	ID               int                 `json:"id"`
	Status           ci.DeploymentStatus `json:"status"`
	EnvironmentName  ci.EnvironmentName  `json:"environmentName"`
	ProjectID        int                 `json:"projectId"`
	ProjectName      string              `json:"projectName"`
	ProjectPath      string              `json:"projectPath"`
	ProjectNamespace string              `json:"projectNamespace"`
	PipelineID       int                 `json:"pipelineId"`
	FinishedAt       time.Time           `json:"finishedAt"`
	Duration         int                 `json:"duration"`
}
